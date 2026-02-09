package memory

// internal/core/memory/errors.go

import "errors"

var (
ErrSealed              = errors.New("core sealed: mutation forbidden")
ErrBankExists          = errors.New("memory bank already exists")
ErrBankNotFound        = errors.New("memory bank not found")
ErrInvalidUnitWidth    = errors.New("invalid unit width")
ErrInvalidUnits        = errors.New("invalid unit count")
ErrInvalidCount        = errors.New("invalid count")
ErrOutOfBounds         = errors.New("out of bounds")
ErrPayloadSizeMismatch = errors.New("payload size mismatch")
)
