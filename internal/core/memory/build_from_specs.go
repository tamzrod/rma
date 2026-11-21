package memory

// BuildFromSpecs validates the whole batch and aggregate capacity before
// allocating or publishing any bank. Seal is serialized with this operation.
func (c *Core) BuildFromSpecs(specs []BankSpec) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.sealed {
		return ErrSealed
	}
	var total uint64
	for _, b := range c.banks {
		total += uint64(len(b.buf))
	}
	seen := make(map[MemoryID]struct{}, len(specs))
	for _, spec := range specs {
		if _, exists := c.banks[spec.ID]; exists {
			return ErrDuplicateBank
		}
		if _, exists := seen[spec.ID]; exists {
			return ErrDuplicateBank
		}
		seen[spec.ID] = struct{}{}
		size, err := validatedSize(spec)
		if err != nil {
			return err
		}
		if size > maxTotalBytes-total {
			return ErrMemoryTooLarge
		}
		total += size
	}
	pending := make(map[MemoryID]*bank, len(specs))
	for _, spec := range specs {
		b, err := newBank(spec)
		if err != nil {
			return err
		}
		pending[spec.ID] = b
	}
	for id, b := range pending {
		c.banks[id] = b
	}
	return nil
}
