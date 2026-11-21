package memory

import (
	"math"
	"sync"
)

// Bank and total capacity limits prevent configuration-induced OOM.
// A larger footprint requires an explicit capacity-contract revision.
const (
	maxBankBytes  uint64 = 64 << 20
	maxTotalBytes uint64 = 256 << 20
)

type bank struct {
	mu   sync.RWMutex
	spec BankSpec
	buf  []byte
}

func validatedSize(spec BankSpec) (uint64, error) {
	if spec.UnitWidthBits == 0 {
		return 0, ErrInvalidUnitWidth
	}
	if spec.Units == 0 {
		return 0, ErrInvalidUnits
	}
	w := uint64(spec.UnitWidthBits)
	if spec.Units > math.MaxUint64/w {
		return 0, ErrMemoryTooLarge
	}
	totalBytes := bitsToBytes(w * spec.Units)
	if totalBytes > maxBankBytes || totalBytes > uint64(^uint(0)>>1) {
		return 0, ErrMemoryTooLarge
	}
	return totalBytes, nil
}

func newBank(spec BankSpec) (*bank, error) {
	totalBytes, err := validatedSize(spec)
	if err != nil {
		return nil, err
	}
	return &bank{spec: spec, buf: make([]byte, int(totalBytes))}, nil
}

// span checks logical unit bounds before offset arithmetic.
func (b *bank) span(start uint64, count uint32) (uint64, uint64, error) {
	if count == 0 {
		return 0, 0, ErrInvalidCount
	}
	if start >= b.spec.Units || uint64(count) > b.spec.Units-start {
		return 0, 0, ErrOutOfBounds
	}
	w := uint64(b.spec.UnitWidthBits)
	return start * w, uint64(count) * w, nil
}

func (b *bank) read(start uint64, count uint32) ([]byte, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	off, bits, err := b.span(start, count)
	if err != nil {
		return nil, err
	}
	out := make([]byte, int(bitsToBytes(bits)))
	bitCopy(out, 0, b.buf, off, bits)
	return out, nil
}

func (b *bank) write(start uint64, count uint32, data []byte) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	off, bits, err := b.span(start, count)
	if err != nil {
		return err
	}
	if uint64(len(data)) != bitsToBytes(bits) {
		return ErrPayloadSizeMismatch
	}
	bitCopy(b.buf, off, data, 0, bits)
	return nil
}
