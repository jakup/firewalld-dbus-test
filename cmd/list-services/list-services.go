package main

import (
    "fmt"
    "github.com/godbus/dbus"
    "os"
)

func main() {
    conn, err := dbus.SystemBus()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to connect to system bus:", err)
        os.Exit(1)
    }

    obj := conn.Object("org.fedoraproject.FirewallD1", dbus.ObjectPath("/org/fedoraproject/FirewallD1"))

    var services []string
    err = obj.Call("org.fedoraproject.FirewallD1.listServices", 0).Store(&services)
    if err != nil {
        fmt.Fprintln(os.Stderr, "org.fedoraproject.FirewallD1.listServices error:", err)
        os.Exit(1)
    }

    for _, service := range services {
        fmt.Println(service)
    }
}
