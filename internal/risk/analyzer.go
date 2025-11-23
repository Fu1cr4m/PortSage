package risk

import (
	"strings"
	"github.com/Fu1cr4m/PortSage/internal/core"
)

// Assess 计算设备的风险等级
func Assess(device *core.USBDevice) {
	// 默认安全
	device.Risk = core.RiskSafe
	device.RiskReason = "Verified Driver"

	name := strings.ToLower(device.Name)
	desc := strings.ToLower(device.Description)
	service := strings.ToLower(device.Service)

	// 规则 1: 存储设备风险 (数据泄露/病毒传播)
	if strings.Contains(service, "usbstor") || strings.Contains(desc, "mass storage") {
		device.Risk = core.RiskLow
		device.RiskReason = "Data Storage Device"
	}

	// 规则 2: 疑似 BadUSB (HID 注入风险)
	// 通常键盘/鼠标由 kbdhid 或 mouhid 服务管理
	// 如果描述含糊且包含 Input 但服务不是标准的，标记高危
	if strings.Contains(desc, "input") || strings.Contains(name, "keyboard") {
		if !strings.Contains(service, "kbd") && !strings.Contains(service, "hid") {
			device.Risk = core.RiskHigh
			device.RiskReason = "Suspicious Input Device"
		}
	}

	// 规则 3: 无法识别的设备
	if device.Service == "" {
		device.Risk = core.RiskLow
		device.RiskReason = "No Driver Service"
	}
}