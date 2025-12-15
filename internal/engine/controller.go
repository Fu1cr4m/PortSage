package engine

import (
	"fmt"
	"os/exec"
	"syscall"
)

// ToggleDeviceState 切换设备状态
// enable = true (Enable-PnpDevice), enable = false (Disable-PnpDevice)
func ToggleDeviceState(instanceID string, enable bool) error {
	cmdStr := "Disable-PnpDevice"
	if enable {
		cmdStr = "Enable-PnpDevice"
	}

	// 构造 PowerShell 命令: -Confirm:$false 用于跳过确认提示
	// 必须以管理员权限运行主程序，否则此处会失败
	psScript := fmt.Sprintf("%s -InstanceId '%s' -Confirm:$false", cmdStr, instanceID)

	cmd := exec.Command("powershell", "-NoProfile", "-WindowStyle", "Hidden", "-Command", psScript)
	
	// 隐藏执行窗口 (Windows 特有)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("operation failed: %v, output: %s", err, string(output))
	}
	return nil
}

// EjectDevice 尝试安全弹出设备 (简化实现，实际弹出需要复杂的 API，这里先做逻辑占位)
// 在 V2.0 中，我们暂时通过禁用操作来模拟“断开连接”
func EjectDevice(instanceID string) error {
	return ToggleDeviceState(instanceID, false)
}