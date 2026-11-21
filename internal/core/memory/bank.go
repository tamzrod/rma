package memory

import (
	"math"
	"sync"
)

// maxBankBytes bounds resource use from user-supplied configuration.
// Larger memory pools require an explicit capacity contract revision.
const maxBankBytes uint64 = 64 << 20

type bank struct {
	mu     sync.RWMutex
	spec   BankSpec
	buf    []byte
	bitLen uint64
}

func newBank(spec BankSpec) (*bank, error) {
	if spec.UnitWidthBits == 0 {
		return nil, ErrInvalidUnitWidth
	}
	if spec.Units == 0 {
		return nil, ErrInvalidUnits
	}
	w := uint64(spec.UnitWidthBits)
	if spec.Units > math.MaxUint64/w {
		return nil, ErrMemoryTooLarge
	}
	totalBits := w * spec.Units
	totalBytes := bitsToBytes(totalBits)
	if totalBytes > maxBankBytes || totalBytes > uint64(int(^uint(0)>>1)) {
		return nil, ErrMemoryTooLarge
	}
	return &bank{
		spec: spec,
		buf: make([]byte, int(totalBytes)),
		bitLen: totalBits,
	}, nil
}

// span validates in units before multiplication, preventing wraparound.
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
