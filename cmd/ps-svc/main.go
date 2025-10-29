package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/YOUR_ORG/portsage/pkg/policy"
    "github.com/YOUR_ORG/portsage/pkg/storage"
)

var (
    policyPath = flag.String("policy", "assets/policy/sample_office_travel.yaml", "policy file path")
    addr       = flag.String("addr", "127.0.0.1:7777", "control endpoint")
)

func main() {
    flag.Parse()

    pol, err := policy.LoadFromFile(*policyPath)
    if err != nil {
        log.Fatalf("load policy: %v", err)
    }
    log.Printf("PortSage svc starting with policy v%d (%d rules)", pol.Version, len(pol.Rules))

    db, err := storage.Open("portsage_audit.db")
    if err != nil {
        log.Fatalf("open audit db: %v", err)
    }
    defer db.Close()

    mux := http.NewServeMux()
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "ok") })
    mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "policy_version=%d rules=%d time=%s\n", pol.Version, len(pol.Rules), time.Now().Format(time.RFC3339))
    })
    log.Printf("control endpoint: http://%s", *addr)
    log.Fatal(http.ListenAndServe(*addr, mux))
}
