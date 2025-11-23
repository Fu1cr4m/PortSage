package policy

import "time"

type ContextSnapshot struct {
    NetworkCategory string
    SessionLocked   bool
    SSID            string
    Now             time.Time
}

type DeviceEvent struct {
    Class string
    CompatibleIDs []string
    When time.Time
}

type Decision struct {
    Action string
    Reason string
}

// Match: minimal placeholder – always allow
func Match(p *Policy, ctx ContextSnapshot, ev DeviceEvent, runtime map[string]any) Decision {
    return Decision{Action: "allow", Reason: "placeholder"}
}

func Periodic(p *Policy, ctx ContextSnapshot, runtime map[string]any) []struct{Decision Decision; Event DeviceEvent} {
    return nil
}
