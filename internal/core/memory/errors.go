package memory

import "errors"

var (
	ErrSealed              = errors.New("memory is sealed")
	ErrDuplicateBank       = errors.New("duplicate memory bank id")
	ErrBankNotFound        = errors.New("memory bank not found")
	ErrInvalidUnitWidth    = errors.New("invalid unit width")
	ErrInvalidUnits        = errors.New("invalid units count")
	ErrInvalidCount        = errors.New("invalid count")
	ErrOutOfBounds         = errors.New("out of bounds access")
	ErrPayloadSizeMismatch = errors.New("payload size mismatch")
	ErrMemoryTooLarge      = errors.New("memory allocation exceeds supported bounds")
)
