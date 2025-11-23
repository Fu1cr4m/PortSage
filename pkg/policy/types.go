package policy

type Policy struct {
    Version  int      `yaml:"version"`
    Defaults Defaults `yaml:"defaults"`
    Rules    []Rule   `yaml:"rules"`
}

type Defaults struct {
    When struct {
        NetworkCategory []string `yaml:"network_category"`
        SessionLocked   *bool    `yaml:"session_locked"`
    } `yaml:"when"`
    Action  string `yaml:"action"`
    Mount   struct{ Readonly bool } `yaml:"mount"`
    Runtime struct {
        Scan string `yaml:"scan"`
        KeyInjectionGuard struct {
            AfterNewKeyboardMs int `yaml:"after_new_keyboard_ms"`
            RateThresholdHz    int `yaml:"rate_threshold_hz"`
            JitterMsP95        int `yaml:"jitter_ms_p95"`
        } `yaml:"key_injection_guard"`
    } `yaml:"runtime"`
}

type Rule struct {
    Name   string `yaml:"name"`
    Match  Match  `yaml:"match"`
    Action string `yaml:"action"`
    Mount  *struct{ Readonly bool } `yaml:"mount,omitempty"`
    Runtime *struct{
        Scan string `yaml:"scan,omitempty"`
        FirewallProfile string `yaml:"firewall_profile,omitempty"`
    } `yaml:"runtime,omitempty"`
    Notify bool `yaml:"notify,omitempty"`
}

type Match struct {
    Class                string   `yaml:"class,omitempty"`
    UsbCompatibleIDs     []string `yaml:"usb_compatible_ids,omitempty"`
    WlanSSID             []string `yaml:"wlan_ssid,omitempty"`
    TimeRange            string   `yaml:"time_range,omitempty"`
    ArrivalWithinMs      int      `yaml:"arrival_within_ms,omitempty"`
    RuntimeKeyInjection  *struct {
        RateHz      string `yaml:"rate_hz,omitempty"`
        JitterMsP95 string `yaml:"jitter_ms_p95,omitempty"`
    } `yaml:"runtime.key_injection,omitempty"`
}
