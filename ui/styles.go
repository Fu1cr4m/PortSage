package ui

import "github.com/charmbracelet/lipgloss"

var (
	// 颜色定义
	colorPurple = lipgloss.Color("#BD93F9")
	colorGreen  = lipgloss.Color("#50FA7B")
	colorRed    = lipgloss.Color("#FF5555")
	colorGray   = lipgloss.Color("#6272A4")
	colorBg     = lipgloss.Color("#282A36")

	// 基础容器
	appStyle = lipgloss.NewStyle().
		Padding(1, 2).
		Background(colorBg)

	// 标题样式
	titleStyle = lipgloss.NewStyle().
		Foreground(colorBg).
		Background(colorPurple).
		Padding(0, 1).
		Bold(true).
		MarginBottom(1)

	// 状态栏
	statusStyle = lipgloss.NewStyle().
		Foreground(colorGreen).
		MarginTop(1)

	// 风险标签
	riskSafe = lipgloss.NewStyle().Foreground(colorGreen).Bold(true)
	riskLow  = lipgloss.NewStyle().Foreground(lipgloss.Color("#F1FA8C")) // Yellow
	riskHigh = lipgloss.NewStyle().Foreground(colorRed).Bold(true).Blink(true)
)