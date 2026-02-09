package memory

// internal/core/memory/bank.go

type bank struct {
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

totalBits := uint64(spec.UnitWidthBits) * spec.Units
totalBytes := bitsToBytes(totalBits)

b := &bank{
spec:   spec,
buf:    make([]byte, totalBytes),
bitLen: totalBits,
}
return b, nil
}

func (b *bank) span(start uint64, count uint32) (uint64, uint64, error) {
if count == 0 {
return 0, 0, ErrInvalidCount
}
w := uint64(b.spec.UnitWidthBits)
off := start * w
len := uint64(count) * w

if off+len > b.bitLen {
return 0, 0, ErrOutOfBounds
}
return off, len, nil
}

func (b *bank) read(start uint64, count uint32) ([]byte, error) {
off, bits, err := b.span(start, count)
if err != nil {
return nil, err
}
out := make([]byte, bitsToBytes(bits))
bitCopy(out, 0, b.buf, off, bits)
return out, nil
}

func (b *bank) write(start uint64, count uint32, data []byte) error {
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
