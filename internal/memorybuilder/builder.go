package memorybuilder

// Path: internal/memorybuilder/builder.go

import (
"github.com/tamzrod/rma/internal/config"
"github.com/tamzrod/rma/internal/core/memory"
)

type Builder struct {
core *memory.Core
}

func New() *Builder {
return &Builder{}
}

// Build prepares the memory core.
// Actual bank creation is delegated to the memory core contract
// and will be wired once the bank-creation API is finalized.
func (b *Builder) Build(cfg *config.Config) error {
if cfg == nil {
return ErrNilConfig
}

b.core = memory.NewCore()
return nil
}

// Init performs INIT-phase actions.
// Actual shaping logic will be added once memory INIT semantics
// are formally wired.
func (b *Builder) Init(cfg *config.Config) error {
if b.core == nil {
return ErrInitPhase
}
return nil
}

// Seal finalizes the memory layout.
func (b *Builder) Seal() error {
if b.core == nil {
return ErrMemoryCreate
}

b.core.Seal()
return nil
}
