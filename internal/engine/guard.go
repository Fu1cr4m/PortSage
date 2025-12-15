package engine

import (
	"golang.org/x/sys/windows/registry"
)

const (
	usbStorPath = `SYSTEM\CurrentControlSet\Services\USBSTOR`
	startKey    = "Start"
)

// CheckStorageStatus: true = Blocked (Disabled), false = Allowed
func CheckStorageStatus() (bool, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, usbStorPath, registry.QUERY_VALUE)
	if err != nil {
		return false, nil // 默认认为没被禁用
	}
	defer k.Close()

	val, _, err := k.GetIntegerValue(startKey)
	if err != nil {
		return false, err
	}
	return val == 4, nil // 4 表示禁用
}

func SetStorageBlock(block bool) error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, usbStorPath, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	var val uint32 = 3 // 3 = Enabled
	if block {
		val = 4 // 4 = Disabled
	}
	return k.SetDWordValue(startKey, val)
}