package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fu1cr4m/PortSageGUI/internal/core"
	"github.com/Fu1cr4m/PortSageGUI/internal/engine"
)

// LogFunc 定义日志回调函数类型
type LogFunc func(string)

// createDeviceCard 生成单行设备控制卡片
func createDeviceCard(d core.USBDevice, refreshFunc func(), log LogFunc, window fyne.Window) fyne.CanvasObject {
	// 1. 图标与状态颜色
	var iconResource fyne.Resource
	var statusColor color.Color

	if d.IsEnabled {
		switch d.Risk {
		case core.RiskCritical:
			iconResource = theme.WarningIcon()
			statusColor = color.RGBA{R: 200, G: 0, B: 0, A: 255} // Red
		case core.RiskWarning:
			iconResource = theme.StorageIcon()
			statusColor = color.RGBA{R: 200, G: 150, B: 0, A: 255} // Orange
		default:
			iconResource = theme.ConfirmIcon()
			statusColor = color.RGBA{R: 0, G: 150, B: 0, A: 255} // Green
		}
	} else {
		iconResource = theme.CancelIcon() // Disabled icon
		statusColor = color.RGBA{R: 100, G: 100, B: 100, A: 255} // Gray
	}

	icon := canvas.NewImageFromResource(iconResource)
	icon.FillMode = canvas.ImageFillContain
	icon.SetMinSize(fyne.NewSize(32, 32))

	// 2. 文本信息
	nameLabel := widget.NewLabel(d.Name)
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}
	
	detailText := fmt.Sprintf("%s | %s", d.Service, d.InstanceID)
	if !d.IsEnabled {
		detailText += " [DISABLED]"
	}
	detailLabel := widget.NewLabel(detailText)
	detailLabel.TextStyle = fyne.TextStyle{Italic: true}
	
	// 风险标签
	riskLabel := canvas.NewText(d.RiskReason, statusColor)
	riskLabel.TextSize = 12

	textContainer := container.NewVBox(nameLabel, detailLabel, riskLabel)

	// 3. 操作按钮区域
	
	// 开关 (Switch)
	toggleSwitch := widget.NewCheck("Active", func(checked bool) {
		// 避免初始化时触发
		if checked == d.IsEnabled {
			return
		}

		// 执行控制
		action := "Enable"
		if !checked { action = "Disable" }
		
		log(fmt.Sprintf("Attempting to %s device: %s...", action, d.Name))
		
		err := engine.ToggleDeviceState(d.InstanceID, checked)
		if err != nil {
			log(fmt.Sprintf("Error: %v", err))
			dialog := widget.NewModalPopUp(
				widget.NewLabel("Failed to toggle device.\nEnsure you are running as Administrator."),
				window.Canvas(),
			)
			dialog.Show()
			// 刷新列表以恢复 UI 状态
			time.Sleep(time.Second) 
			refreshFunc() 
		} else {
			log(fmt.Sprintf("Success: Device %s.", action+"d"))
			// 延迟一点时间让 Windows 更新状态，然后刷新
			go func() {
				time.Sleep(1500 * time.Millisecond)
				refreshFunc()
			}()
		}
	})
	toggleSwitch.SetChecked(d.IsEnabled)

	// 组装操作栏
	buttons := container.NewHBox(toggleSwitch)
	
	// 如果是存储设备，添加弹出按钮
	if d.CanEject && d.IsEnabled {
		ejectBtn := widget.NewButtonWithIcon("Eject",theme.LogoutIcon(), func() {
			log(fmt.Sprintf("Ejecting storage: %s", d.Name))
			engine.EjectDevice(d.InstanceID)
			refreshFunc()
		})
		buttons.Add(ejectBtn)
	}

	// 4. 最终布局
	return container.NewBorder(
		nil, nil, 
		container.NewCenter(icon), // Left
		buttons,                   // Right
		textContainer,             // Center
	)
}