package memory

// Path: internal/core/memory/build_from_specs.go

// BuildFromSpecs builds all memory banks from declarative specs.
// YAML is the sole authority for memory bank creation.
// Must be called before sealing.
func (c *Core) BuildFromSpecs(specs []BankSpec) error {
if c.sealed {
return ErrSealed
}

for _, spec := range specs {
if spec.UnitWidthBits == 0 {
return ErrInvalidUnitWidth
}
if spec.Units == 0 {
return ErrInvalidUnits
}

id := MemoryID(spec.ID)

if _, exists := c.banks[id]; exists {
return ErrDuplicateBank
}

b, err := newBank(spec)
if err != nil {
return err
}

c.banks[id] = b
}

return nil
}
