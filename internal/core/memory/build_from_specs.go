package memory

// BuildFromSpecs validates a complete batch before publishing it. On error,
// no bank from that batch is made visible. The layout lock serializes Seal.
func (c *Core) BuildFromSpecs(specs []BankSpec) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.sealed {
		return ErrSealed
	}
	pending := make(map[MemoryID]*bank, len(specs))
	for _, spec := range specs {
		id := spec.ID
		if _, exists := c.banks[id]; exists {
			return ErrDuplicateBank
		}
		if _, exists := pending[id]; exists {
			return ErrDuplicateBank
		}
		b, err := newBank(spec)
		if err != nil {
			return err
		}
		pending[id] = b
	}
	for id, b := range pending {
		c.banks[id] = b
	}
	return nil
}
