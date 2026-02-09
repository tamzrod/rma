package memory

// internal/core/memory/bitops.go

func bitsToBytes(bits uint64) uint64 {
if bits == 0 {
return 0
}
return (bits + 7) / 8
}

func getBit(buf []byte, bit uint64) byte {
b := bit / 8
o := bit % 8
if b >= uint64(len(buf)) {
return 0
}
return (buf[b] >> o) & 1
}

func setBit(buf []byte, bit uint64, v byte) {
b := bit / 8
o := bit % 8
if b >= uint64(len(buf)) {
return
}
if v&1 == 1 {
buf[b] |= (1 << o)
} else {
buf[b] &^= (1 << o)
}
}

func bitCopy(dst []byte, dstOff uint64, src []byte, srcOff uint64, bits uint64) {
for i := uint64(0); i < bits; i++ {
setBit(dst, dstOff+i, getBit(src, srcOff+i))
}
}
