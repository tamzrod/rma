package config

// Path: internal/config/errors.go

import "errors"

var (
ErrMissingMemory        = errors.New("missing memory section")
ErrNoMemoryBanks        = errors.New("memory.banks must not be empty")
ErrDuplicateMemoryID   = errors.New("duplicate memory bank id")
ErrInvalidUnitWidth    = errors.New("unit_width_bits must be > 0")
ErrInvalidUnits        = errors.New("units must be > 0")
ErrInvalidInitBlock    = errors.New("init.allow_cidr must not be empty")
ErrInvalidAccessBlock  = errors.New("access.allow_cidr must not be empty")
)
