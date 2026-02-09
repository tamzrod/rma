package main

// Path: cmd/rma/main.go

import (
"fmt"
"os"

"github.com/tamzrod/rma/internal/bootstrap"
"github.com/tamzrod/rma/internal/config"
"github.com/tamzrod/rma/internal/memorybuilder"
)

// noopListenerStarter satisfies bootstrap.ListenerStarter.
type noopListenerStarter struct{}

func (n *noopListenerStarter) Start(_ *config.Config) error {
// No listeners yet.
return nil
}

func main() {
if len(os.Args) < 2 {
fmt.Fprintln(os.Stderr, "usage: rma <config.yaml>")
os.Exit(1)
}

configPath := os.Args[1]

mem := memorybuilder.New()
listeners := &noopListenerStarter{}

if err := bootstrap.Run(configPath, mem, listeners); err != nil {
fmt.Fprintln(os.Stderr, "startup failed:", err)
os.Exit(1)
}

fmt.Println("RMA started successfully (memory sealed, runtime active)")
}
