package memory

// internal/core/memory/core.go

import "sync"

type Core struct {
mu     sync.RWMutex
sealed bool
banks  map[MemoryID]*bank
}

func NewCore() *Core {
return &Core{
banks: make(map[MemoryID]*bank),
}
}

func (c *Core) Seal() {
c.mu.Lock()
defer c.mu.Unlock()
c.sealed = true
}

func (c *Core) AddBank(spec BankSpec) error {
c.mu.Lock()
defer c.mu.Unlock()

if c.sealed {
return ErrSealed
}
if _, ok := c.banks[spec.ID]; ok {
return ErrBankExists
}
b, err := newBank(spec)
if err != nil {
return err
}
c.banks[spec.ID] = b
return nil
}

func (c *Core) Read(req ReadReq) ([]byte, error) {
c.mu.RLock()
b, ok := c.banks[req.ID]
c.mu.RUnlock()
if !ok {
return nil, ErrBankNotFound
}
return b.read(req.Start, req.Count)
}

func (c *Core) Write(req WriteReq) error {
c.mu.RLock()
b, ok := c.banks[req.ID]
c.mu.RUnlock()
if !ok {
return ErrBankNotFound
}
return b.write(req.Start, req.Count, req.Data)
}
