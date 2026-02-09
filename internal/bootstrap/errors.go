package bootstrap

// Path: internal/bootstrap/errors.go

import "errors"

var (
ErrConfigLoadFailed   = errors.New("config load failed")
ErrMemoryBuildFailed  = errors.New("memory build failed")
ErrInitFailed         = errors.New("memory init failed")
ErrSealFailed         = errors.New("memory seal failed")
ErrListenerFailed     = errors.New("listener startup failed")
)
