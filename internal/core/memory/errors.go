package memory

// Path: internal/core/memory/errors.go

import "errors"

var (
// Core / lifecycle
ErrSealed        = errors.New("memory is sealed")
ErrDuplicateBank = errors.New("duplicate memory bank id")

// Bank validation
ErrInvalidUnitWidth = errors.New("invalid unit width")
ErrInvalidUnits     = errors.New("invalid units count")
ErrInvalidCount     = errors.New("invalid count")
ErrOutOfBounds      = errors.New("out of bounds access")
ErrPayloadSizeMismatch = errors.New("payload size mismatch")
)
