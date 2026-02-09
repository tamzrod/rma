package bootstrap

// Path: internal/bootstrap/bootstrap.go

import (
"github.com/tamzrod/rma/internal/config"
)

// Interfaces keep bootstrap policy-blind and replaceable

type MemoryBuilder interface {
Build(cfg *config.Config) error
Init(cfg *config.Config) error
Seal() error
}

type ListenerStarter interface {
Start(cfg *config.Config) error
}

// Run executes the full RMA bootstrap sequence.
// Order is strict and must not be altered.
func Run(
configPath string,
mem MemoryBuilder,
listeners ListenerStarter,
) error {

// 1. Load + validate configuration
cfg, err := config.Load(configPath)
if err != nil {
return ErrConfigLoadFailed
}

// 2. Build memory layout
if err := mem.Build(cfg); err != nil {
return ErrMemoryBuildFailed
}

// 3. Perform INIT (if defined)
if err := mem.Init(cfg); err != nil {
return ErrInitFailed
}

// 4. Seal memory
if err := mem.Seal(); err != nil {
return ErrSealFailed
}

// 5. Start listeners (runtime begins)
if err := listeners.Start(cfg); err != nil {
return ErrListenerFailed
}

return nil
}
