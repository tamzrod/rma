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

func (b *Builder) Build(cfg *config.Config) error {
	specs := make([]memory.BankSpec, 0, len(cfg.Memory.Banks))
	for _, bank := range cfg.Memory.Banks {
		specs = append(specs, memory.BankSpec{
			ID:            memory.MemoryID(bank.ID),
			UnitWidthBits: bank.UnitWidthBits,
			Units:         bank.Units,
			Name:          bank.Name,
		})
	}

	core := memory.NewCore()
	if err := core.BuildFromSpecs(specs); err != nil {
		return err
	}

	b.core = core
	return nil
}

func (b *Builder) Init(cfg *config.Config) error {
	return nil
}

func (b *Builder) Seal() error {
	b.core.Seal()
	return nil
}

func (b *Builder) Core() *memory.Core {
	return b.core
}
