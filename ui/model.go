package ui

import (
	"fmt"
	"time"

	"github.com/Fu1cr4m/PortSage/internal/core"
	"github.com/Fu1cr4m/PortSage/internal/engine"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	table         table.Model
	devices       []core.USBDevice
	storageBlocked bool
	err           error
	loading       bool
}

// 消息类型
type scanResultMsg struct {
	devices []core.USBDevice
	err     error
}

type blockStatusMsg struct {
	blocked bool
	err     error
}

func NewModel() Model {
	// 表格配置
	columns := []table.Column{
		{Title: "Risk", Width: 8},
		{Title: "Device Name", Width: 30},
		{Title: "Type/Description", Width: 35},
		{Title: "Status", Width: 15},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.BorderStyle(lipgloss.NormalBorder()).BorderBottom(true).Bold(true)
	s.Selected = s.Selected.Foreground(lipgloss.Color("229")).Background(lipgloss.Color("57")).Bold(false)
	t.SetStyles(s)

	return Model{
		table:   t,
		loading: true,
	}
}

func (m Model) Init() tea.Cmd {
	// 初始化时同时执行扫描和检查屏蔽状态
	return tea.Batch(scanCmd, checkBlockStatusCmd)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r": // 刷新
			m.loading = true
			return m, scanCmd
		case " ": // 空格键切换屏蔽状态
			newStatus := !m.storageBlocked
			// 立即触发更改
			err := engine.SetStorageBlockState(newStatus)
			if err != nil {
				m.err = err
			} else {
				m.storageBlocked = newStatus
			}
			return m, checkBlockStatusCmd // 重新检查确认
		}

	case scanResultMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.devices = msg.devices
			m.updateTable()
		}

	case blockStatusMsg:
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.storageBlocked = msg.blocked
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n\nPress 'q' to quit", m.err)
	}

	// 头部标题
	header := titleStyle.Render("PORTSAGE - USB SECURITY MANAGER")

	// 防火墙/屏蔽状态显示
	statusIcon := "🔓 UNLOCKED"
	statusColor := lipgloss.Color("#50FA7B") // Green
	if m.storageBlocked {
		statusIcon = "🔒 USB STORAGE BLOCKED"
		statusColor = lipgloss.Color("#FF5555") // Red
	}
	firewallStatus := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(statusColor).
		Padding(0, 1).
		Bold(true).
		Render(" GLOBAL POLICY: " + statusIcon + " ")

	// 表格内容
	tableView := m.table.View()
	if m.loading {
		tableView = "Scanning hardware bus..."
	}

	// 底部提示
	help := statusStyle.Render("Controls: [Space] Toggle Storage Block • [R] Refresh • [Q] Quit")

	return appStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			header,
			firewallStatus,
			"\n",
			tableView,
			"\n",
			help,
		),
	)
}

func (m *Model) updateTable() {
	rows := []table.Row{}
	for _, d := range m.devices {
		// 格式化 Risk 列
		riskStr := d.Risk.String()
		
		// 格式化 Reason
		reason := d.RiskReason
		if reason == "" {
			reason = "Normal"
		}

		rows = append(rows, table.Row{
			riskStr,
			d.Name,
			d.Description,
			reason,
		})
	}
	m.table.SetRows(rows)
}

// --- Commands ---

func scanCmd() tea.Msg {
	// 模拟一点延迟让 UI 看起来在工作
	time.Sleep(500 * time.Millisecond)
	devs, err := engine.ScanUSBDevices()
	return scanResultMsg{devices: devs, err: err}
}

func checkBlockStatusCmd() tea.Msg {
	blocked, err := engine.IsStorageBlocked()
	return blockStatusMsg{blocked: blocked, err: err}
}