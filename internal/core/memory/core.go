package memory

// Path: internal/core/memory/core.go

type Core struct {
banks  map[MemoryID]*bank
sealed bool
}

func NewCore() *Core {
return &Core{
banks: make(map[MemoryID]*bank),
}
}

func (c *Core) Seal() {
c.sealed = true
}
