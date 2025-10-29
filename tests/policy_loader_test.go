package tests

import (
    "testing"
    "github.com/YOUR_ORG/portsage/pkg/policy"
)

func TestLoadPolicy(t *testing.T){
    p, err := policy.LoadFromFile("../assets/policy/sample_office_travel.yaml")
    if err != nil { t.Fatal(err) }
    if p.Version < 1 || len(p.Rules) == 0 { t.Fatalf("unexpected policy: %+v", p) }
}
