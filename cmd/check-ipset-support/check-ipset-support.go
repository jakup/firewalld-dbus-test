package main

import (
    "fmt"
    "github.com/godbus/dbus"
    "os"
)

func checkError(err error, message string) {
    if err == nil {
        return
    }

    var errMsg string = err.Error()
    if errMsg == "" {
        errMsg = "insufficient permissions?"
    }

    fmt.Fprintf(os.Stderr, "%s: %s\n", message, errMsg)
    os.Exit(1)
}

func main() {
    conn, err := dbus.SystemBus()
    checkError(err, "Failed to connect to system bus")

    obj := conn.Object("org.fedoraproject.FirewallD1", dbus.ObjectPath("/org/fedoraproject/FirewallD1"))

    p, err := obj.GetProperty("org.fedoraproject.FirewallD1.IPSet")
    checkError(err, "Unable to read IPSet property")

    hasIPSetSupport := p.Value().(bool)
    fmt.Println("firewall has IPSet support:", hasIPSetSupport)
    if ! hasIPSetSupport {
        os.Exit(0)
    }

    p, err = obj.GetProperty("org.fedoraproject.FirewallD1.IPSetTypes")
    checkError(err, "Unable to read IPSetTypes property")

    supportedTypes := p.Value().([]string)
    fmt.Println("supported IPSet types:")
    for _, t := range supportedTypes {
        fmt.Printf("  - %s\n", t)
    }
}
