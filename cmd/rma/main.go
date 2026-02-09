package main

// Path: cmd/rma/main.go

import (
	"fmt"
	"os"

	"github.com/tamzrod/rma/internal/bootstrap"
	"github.com/tamzrod/rma/internal/listener/rawtcp"
	"github.com/tamzrod/rma/internal/memorybuilder"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: rma <config.yaml>")
		os.Exit(1)
	}

	configPath := os.Args[1]

	mem := memorybuilder.New()
	listeners := rawtcp.NewStarter()

	if err := bootstrap.Run(configPath, mem, listeners); err != nil {
		fmt.Fprintln(os.Stderr, "startup failed:", err)
		os.Exit(1)
	}

	fmt.Println("RMA started successfully (memory sealed, runtime active)")
	select {}
}
