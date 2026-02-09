package memory

// Path: internal/core/memory/core.go
//
// Core owns banks and exposes minimal RAM-like doors.
// No semantics. No policy. No safety net.

type Core struct {
	banks  map[MemoryID]*bank
	sealed bool
}

func NewCore() *Core {
	return &Core{
		banks: make(map[MemoryID]*bank),
	}
}

// Seal locks the layout. After this, memory behaves like RAM.
func (c *Core) Seal() {
	c.sealed = true
}

// WriteUnits writes raw data into memory using unit addressing.
func (c *Core) WriteUnits(id MemoryID, start uint64, count uint32, data []byte) error {
	b := c.banks[id]
	return b.write(start, count, data)
}

// ReadUnits reads raw data from memory using unit addressing.
func (c *Core) ReadUnits(id MemoryID, start uint64, count uint32) ([]byte, error) {
	b := c.banks[id]
	return b.read(start, count)
}
