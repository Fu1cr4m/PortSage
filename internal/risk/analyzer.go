package risk

import (
	"strings"
	"github.com/Fu1cr4m/PortSageGUI/internal/core"
)

func Assess(device *core.USBDevice) {
	device.Risk = core.RiskSafe
	device.RiskReason = "Safe"

	desc := strings.ToLower(device.Description)
	service := strings.ToLower(device.Service)

	if strings.Contains(service, "usbstor") {
		device.Risk = core.RiskWarning
		device.RiskReason = "Mass Storage"
	}

	// 键盘鼠标如果是 HID 则视为安全，但如果描述可疑则警告
	if (strings.Contains(desc, "input") || strings.Contains(desc, "keyboard")) && 
	   (!strings.Contains(service, "hid") && !strings.Contains(service, "kbd")) {
		device.Risk = core.RiskCritical
		device.RiskReason = "Suspicious HID"
	}
}