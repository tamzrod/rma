package memory

// internal/core/memory/doc.go
//
// RMA Memory Core (protocol-agnostic)
//
// - Multiple memory banks identified by 64-bit MemoryID
// - Fixed unit width per bank (in bits)
// - Requests are: MemoryID + Start (0-based) + Count
// - Width is implied by MemoryID
// - Deterministic, bounded, protocol-agnostic
// - Layout is immutable after sealing
