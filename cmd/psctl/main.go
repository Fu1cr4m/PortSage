package main

import (
    "flag"
    "fmt"
    "io"
    "net/http"
)

var (
    cmd  = flag.String("cmd", "status", "status|enforce|audit|ctx|temp-allow|rollback")
    addr = flag.String("addr", "127.0.0.1:7777", "svc endpoint")
)

func main() {
    flag.Parse()
    switch *cmd {
    case "status":
        resp, err := http.Get("http://" + *addr + "/status")
        if err != nil { panic(err) }
        defer resp.Body.Close()
        b, _ := io.ReadAll(resp.Body)
        fmt.Print(string(b))
    default:
        fmt.Println("TODO:", *cmd)
    }
}
