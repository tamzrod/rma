package memory

// internal/core/memory/types.go

type MemoryID uint64

// BankSpec defines a fixed memory bank.
type BankSpec struct {
ID            MemoryID
UnitWidthBits uint32 // may be any positive integer (1,3,5,8,16,32,128,...)
Units         uint64 // number of logical units
Name          string // human/debug only
}

type ReadReq struct {
ID    MemoryID
Start uint64
Count uint32
}

type WriteReq struct {
ID    MemoryID
Start uint64
Count uint32
Data  []byte
}
