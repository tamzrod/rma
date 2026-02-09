package rawtcp

// Path: internal/listener/rawtcp/starter.go
//
// Raw TCP listener starter.
// Starts listeners defined in config.
// No logging here  actual listening is logged by Listener.Run().

import (
	"context"

	"github.com/tamzrod/rma/internal/config"
	"github.com/tamzrod/rma/internal/core/memory"
)

type Starter struct{}

func NewStarter() *Starter {
	return &Starter{}
}

func (s *Starter) Start(cfg *config.Config, core *memory.Core) error {
	for _, l := range cfg.Listeners {
		addr := l.Bind

		listener := New(addr, core)

		go func() {
			_ = listener.Run(context.Background())
		}()
	}
	return nil
}
