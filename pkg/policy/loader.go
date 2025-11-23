package policy

import (
    "os"
    "gopkg.in/yaml.v3"
)

func LoadFromFile(path string) (*Policy, error) {
    b, err := os.ReadFile(path)
    if err != nil { return nil, err }
    var p Policy
    if err := yaml.Unmarshal(b, &p); err != nil { return nil, err }
    if p.Version == 0 { p.Version = 1 }
    return &p, nil
}
