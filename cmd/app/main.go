package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fu1cr4m/PortSageGUI/internal/engine"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("PortSage v2.0 - Advanced USB Manager")
	myWindow.Resize(fyne.NewSize(900, 650))

	// --- 1. 日志系统 ---
	logData := []string{"System initialized. Scanning devices..."}
	logList := widget.NewList(
		func() int { return len(logData) },
		func() fyne.CanvasObject { return widget.NewLabel("Log Entry") },
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), logData[len(logData)-1-i])) // 倒序显示
		},
	)
	
	// 添加日志的辅助函数
	addLog := func(msg string) {
		logData = append(logData, msg)
		logList.Refresh()
	}

	// --- 2. 设备管理 Tab ---
	deviceListContainer := container.NewVBox()
	scrollDeviceContainer := container.NewVScroll(deviceListContainer)

	refreshDevices := func() {
		deviceListContainer.Objects = nil // 清空
		addLog("Refreshing device list via WMI...")
		
		devices, err := engine.ScanUSBDevices()
		if err != nil {
			deviceListContainer.Add(widget.NewLabel(fmt.Sprintf("Scan Error: %v", err)))
			addLog(fmt.Sprintf("Error scanning: %v", err))
		} else {
			if len(devices) == 0 {
				deviceListContainer.Add(widget.NewLabel("No USB peripherals detected."))
			}
			for _, d := range devices {
				card := createDeviceCard(d, func() {
					// 刷新回调（递归调用需要注意，但在事件中是安全的）
					// 这里我们只能通过外部触发刷新，无法直接递归，所以稍后绑定
				}, addLog, myWindow)
				deviceListContainer.Add(card)
				deviceListContainer.Add(widget.NewSeparator())
			}
		}
		deviceListContainer.Refresh()
		addLog(fmt.Sprintf("Scan complete. Found %d devices.", len(devices)))
	}

	// 由于闭包问题，我们需要重新包装 createDeviceCard 的回调
	// 这里采用简化的方式，重新定义刷新逻辑
	refreshWrapper := func() { refreshDevices() }

	// 修改上面的循环逻辑 (为了能调用 refreshWrapper)
	refreshDevices = func() {
		deviceListContainer.Objects = nil
		addLog("Refreshing device list...")
		devices, err := engine.ScanUSBDevices()
		if err != nil {
			deviceListContainer.Add(widget.NewLabel(fmt.Sprintf("Error: %v", err)))
			return
		}
		for _, d := range devices {
			card := createDeviceCard(d, refreshWrapper, addLog, myWindow)
			deviceListContainer.Add(card)
			deviceListContainer.Add(widget.NewSeparator())
		}
		deviceListContainer.Refresh()
	}

	// 顶部工具栏
	refreshBtn := widget.NewButtonWithIcon("Scan Hardware", theme.ViewRefreshIcon(), func() {
		refreshDevices()
	})
	toolbar := container.NewHBox(refreshBtn)
	
	deviceTabContent := container.NewBorder(toolbar, nil, nil, nil, scrollDeviceContainer)

	// --- 3. 审计日志 Tab ---
	// 简单的清空日志按钮
	clearLogBtn := widget.NewButton("Clear Logs", func() {
		logData = []string{"Logs cleared."}
		logList.Refresh()
	})
	logTabContent := container.NewBorder(nil, clearLogBtn, nil, nil, logList)

	// --- 4. 组装 Tabs ---
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Devices Manager", theme.ComputerIcon(), deviceTabContent),
		container.NewTabItemWithIcon("Audit Logs", theme.DocumentIcon(), logTabContent),
	)

	myWindow.SetContent(tabs)

	// 启动时初次扫描
	refreshDevices()
	
	myWindow.ShowAndRun()
}