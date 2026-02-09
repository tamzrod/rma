package config

// Path: internal/config/validate.go

func Validate(cfg *Config) error {
if cfg.Memory.Banks == nil {
return ErrMissingMemory
}
if len(cfg.Memory.Banks) == 0 {
return ErrNoMemoryBanks
}

seen := make(map[uint64]struct{})

for _, b := range cfg.Memory.Banks {
if b.UnitWidthBits == 0 {
return ErrInvalidUnitWidth
}
if b.Units == 0 {
return ErrInvalidUnits
}
if _, ok := seen[b.ID]; ok {
return ErrDuplicateMemoryID
}
seen[b.ID] = struct{}{}

if b.Init != nil && len(b.Init.AllowCIDR) == 0 {
return ErrInvalidInitBlock
}
if b.Access != nil && len(b.Access.AllowCIDR) == 0 {
return ErrInvalidAccessBlock
}
}

return nil
}
