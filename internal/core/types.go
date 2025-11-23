package core

// RiskLevel 定义风险等级
type RiskLevel int

const (
	RiskSafe RiskLevel = iota
	RiskLow
	RiskHigh
)

func (r RiskLevel) String() string {
	switch r {
	case RiskSafe:
		return "SAFE"
	case RiskLow:
		return "LOW"
	case RiskHigh:
		return "HIGH"
	default:
		return "UNKNOWN"
	}
}

// USBDevice 描述一个物理连接的 USB 设备
type USBDevice struct {
	DeviceID    string    // 硬件唯一 ID (如 USB\VID_xxxx&PID_xxxx\...)
	Name        string    // 友好名称
	Description string    // 设备描述
	Service     string    // 使用的驱动服务
	Risk        RiskLevel // 风险等级
	RiskReason  string    // 风险原因
}