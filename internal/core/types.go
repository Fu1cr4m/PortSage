package core

type RiskLevel int

const (
	RiskSafe RiskLevel = iota
	RiskWarning
	RiskCritical
)

func (r RiskLevel) String() string {
	switch r {
	case RiskSafe: return "SAFE"
	case RiskWarning: return "WARNING"
	case RiskCritical: return "CRITICAL"
	default: return "UNKNOWN"
	}
}

type USBDevice struct {
	InstanceID  string    // Windows PnP 实例 ID (用于控制)
	Name        string    // 友好名称
	Description string    // 描述
	Service     string    // 驱动服务
	Status      string    // Windows 状态 (OK, Error, Degraded)
	IsEnabled   bool      // 逻辑状态
	Risk        RiskLevel // 风险等级
	RiskReason  string    // 风险原因
	CanEject    bool      // 是否为存储设备 (可弹出)
}