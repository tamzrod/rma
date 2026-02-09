# Concurrency Model

RMA core is concurrent by design.

Principles:
- concurrency safety is enforced inside the core
- adapters are not responsible for locking
- read/write operations must be race-safe
- deterministic behavior is required under contention

The exact synchronization primitives are an implementation detail,
but the observable behavior must remain consistent regardless of:
- adapter type
- caller concurrency
- execution order
