package engine

import (
	"strings"
	"github.com/Fu1cr4m/PortSageGUI/internal/core"
	"github.com/Fu1cr4m/PortSageGUI/internal/risk"
	"github.com/yusufpapurcu/wmi"
)

// Win32_PnPEntity 对应 Windows WMI 类
type win32PnPEntity struct {
	Name                 string
	DeviceID             string
	Description          string
	Service              string
	Status               string // "OK", "Error", "Degraded", "Unknown"
	ConfigManagerErrorCode uint32 // 0=正常, 22=已禁用
}

func ScanUSBDevices() ([]core.USBDevice, error) {
	var dst []win32PnPEntity
	// 查询所有 USB 设备
	query := "SELECT Name, DeviceID, Description, Service, Status, ConfigManagerErrorCode FROM Win32_PnPEntity WHERE DeviceID LIKE 'USB%'"
	if err := wmi.Query(query, &dst); err != nil {
		return nil, err
	}

	var devices []core.USBDevice
	for _, d := range dst {
		// 过滤根集线器和复合设备容器，只显示终端设备
		if strings.Contains(d.Name, "Root Hub") || strings.Contains(d.Name, "Controller") || strings.Contains(d.Name, "Composite Device") {
			continue
		}

		// 判断是否启用: ErrorCode 22 代表被禁用
		isEnabled := d.ConfigManagerErrorCode != 22
		
		isStorage := strings.Contains(d.Service, "USBSTOR") || strings.Contains(strings.ToLower(d.Description), "mass storage")

		dev := core.USBDevice{
			InstanceID:  d.DeviceID,
			Name:        d.Name,
			Description: d.Description,
			Service:     d.Service,
			Status:      d.Status,
			IsEnabled:   isEnabled,
			CanEject:    isStorage,
		}
		
		risk.Assess(&dev)
		devices = append(devices, dev)
	}
	return devices, nil
}