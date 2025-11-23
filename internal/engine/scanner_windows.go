package engine

import (
	"strings"
	"github.com/Fu1cr4m/PortSage/internal/core"
	"github.com/Fu1cr4m/PortSage/internal/risk"
	"github.com/yusufpapurcu/wmi"
)

type win32PnPEntity struct {
	Name        string
	DeviceID    string
	Description string
	Service     string
}

// ScanUSBDevices 通过 WMI 获取所有连接的 USB 设备
func ScanUSBDevices() ([]core.USBDevice, error) {
	var dst []win32PnPEntity
	// 查询所有 PnP 设备中 DeviceID 以 USB 开头的
	query := "SELECT Name, DeviceID, Description, Service FROM Win32_PnPEntity WHERE DeviceID LIKE 'USB%'"
	
	err := wmi.Query(query, &dst)
	if err != nil {
		return nil, err
	}

	var devices []core.USBDevice
	for _, d := range dst {
		// 过滤掉一些 USB 根集线器，只保留实际外设
		// 这是一个简化的过滤逻辑，实际可以根据需求调整
		if strings.Contains(d.Name, "Root Hub") || strings.Contains(d.Name, "Controller") {
			continue
		}

		dev := core.USBDevice{
			DeviceID:    d.DeviceID,
			Name:        d.Name,
			Description: d.Description,
			Service:     d.Service,
		}
		
		// 进行风险评估
		risk.Assess(&dev)
		
		devices = append(devices, dev)
	}
	return devices, nil
}