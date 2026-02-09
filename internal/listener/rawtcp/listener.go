package rawtcp

// Path: internal/listener/rawtcp/listener.go
//
// Raw TCP listener.
// Accepts connections and hands them to the raw adapter.
// No parsing, no policy, no memory logic here.

import (
	"context"
	"log"
	"net"

	rawadapter "github.com/tamzrod/rma/internal/adapter/raw"
	"github.com/tamzrod/rma/internal/core/memory"
)

type Listener struct {
	Addr string
	Core *memory.Core
}

func New(addr string, core *memory.Core) *Listener {
	return &Listener{
		Addr: addr,
		Core: core,
	}
}

func (l *Listener) Run(ctx context.Context) error {
	ln, err := net.Listen("tcp", l.Addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	log.Printf("[rawtcp] listening on %s", l.Addr)

	// Shutdown watcher
	go func() {
		<-ctx.Done()
		log.Printf("[rawtcp] shutdown signal received")
		ln.Close()
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return nil
			default:
				return err
			}
		}

		log.Printf("[rawtcp] connection accepted: %s", conn.RemoteAddr())

		adapter := rawadapter.New(l.Core)

		go func() {
			adapter.Handle(conn)
			log.Printf("[rawtcp] connection closed: %s", conn.RemoteAddr())
		}()
	}
}
