package memory

import "sync"

// Core owns the memory-bank registry. Its lock protects layout changes;
// each bank protects its own read/write contents independently.
type Core struct {
	mu     sync.RWMutex
	banks  map[MemoryID]*bank
	sealed bool
}

func NewCore() *Core {
	return &Core{banks: make(map[MemoryID]*bank)}
}

// Seal prevents new banks from being created. It does not prohibit data writes.
func (c *Core) Seal() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sealed = true
}

// WriteUnits atomically writes a bounded range into one bank.
func (c *Core) WriteUnits(id MemoryID, start uint64, count uint32, data []byte) error {
	c.mu.RLock()
	b, ok := c.banks[id]
	c.mu.RUnlock()
	if !ok {
		return ErrBankNotFound
	}
	return b.write(start, count, data)
}

// ReadUnits returns an independent, consistent copy of one bank range.
func (c *Core) ReadUnits(id MemoryID, start uint64, count uint32) ([]byte, error) {
	c.mu.RLock()
	b, ok := c.banks[id]
	c.mu.RUnlock()
	if !ok {
		return nil, ErrBankNotFound
	}
	return b.read(start, count)
}
