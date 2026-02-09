package memorybuilder

// Path: internal/memorybuilder/errors.go

import "errors"

var (
ErrNilConfig      = errors.New("nil config")
ErrMemoryCreate   = errors.New("memory creation failed")
ErrInitPhase      = errors.New("init phase failed")
ErrAlreadySealed  = errors.New("memory already sealed")
)
