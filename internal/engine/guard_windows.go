package engine

import (
	"golang.org/x/sys/windows/registry"
)

const (
	usbStorPath = `SYSTEM\CurrentControlSet\Services\USBSTOR`
	startKey    = "Start"
)

// IsStorageBlocked 检查当前是否禁用了 USB 存储
func IsStorageBlocked() (bool, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, usbStorPath, registry.QUERY_VALUE)
	if err != nil {
		// 如果找不到键值，可能驱动未加载，默认视为未禁用
		return false, nil
	}
	defer k.Close()

	val, _, err := k.GetIntegerValue(startKey)
	if err != nil {
		return false, err
	}

	// 4 = Disabled, 3 = Enabled (Manual)
	return val == 4, nil
}

// SetStorageBlockState 开启或关闭 USB 存储屏蔽
// enableBlock = true -> 禁用 USB (Start=4)
// enableBlock = false -> 启用 USB (Start=3)
func SetStorageBlockState(enableBlock bool) error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, usbStorPath, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	var val uint32 = 3 // 默认启用
	if enableBlock {
		val = 4 // 禁用
	}

	return k.SetDWordValue(startKey, val)
}