package bootstrap

// Path: internal/bootstrap/bootstrap.go
//
// Bootstrap orchestrates startup lifecycle:
//   Load config -> Build memory -> Init -> Seal -> Start listeners

import (
	"github.com/tamzrod/rma/internal/config"
	"github.com/tamzrod/rma/internal/core/memory"
	"github.com/tamzrod/rma/internal/memorybuilder"
)

// ListenerStarter defines how runtime listeners are started.
type ListenerStarter interface {
	Start(*config.Config, *memory.Core) error
}

func Run(
	configPath string,
	mem *memorybuilder.Builder,
	listeners ListenerStarter,
) error {

	// 1) Load config
	cfg, err := config.Load(configPath)
	if err != nil {
		return err
	}

	// 2) Build memory
	if err := mem.Build(cfg); err != nil {
		return err
	}

	// 3) Init (noop for now)
	if err := mem.Init(cfg); err != nil {
		return err
	}

	// 4) Seal
	if err := mem.Seal(); err != nil {
		return err
	}

	core := mem.Core()

	// 5) Start listeners
	if listeners != nil {
		if err := listeners.Start(cfg, core); err != nil {
			return err
		}
	}

	return nil
}
