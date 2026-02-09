# STATE SEALING

## 1. Purpose
- Define memory lifecycle
- Separate lifecycle from security and protocol
- Guarantee deterministic runtime behavior

## 2. Scope
- Applies per memory bank
- Applies only during process lifecycle
- Independent of listeners, protocols, and data access

## 3. Lifecycle Phases
- Bootstrap phase
- INIT phase
- Runtime phase

## 4. Sealing Rules
- Presence of init: implies sealed memory
- Sealing is implicit (no explicit flag)
- Sealing is one-way for process lifetime
- No runtime unseal

## 5. INIT Semantics
- INIT is controlled initial write
- INIT happens only before runtime
- INIT intent is explicit via MemoryID flag
- INIT is not a protocol operation

## 6. Runtime Semantics
- Runtime allows data access only
- Runtime forbids lifecycle mutation
- No reseal, resize, or unlock

## 7. MemoryID Flags
- INIT flag declares intent
- DATA flag is default
- Flags do not grant authority
- Flags do not change memory layout

## 8. Forbidden Operations
- Unlock flag
- Control bits in memory data
- Special addresses or magic regions
- Runtime lifecycle commands

## 9. Authority Boundary
- State sealing is not an authority decision
- Authority is consulted externally
- RMA core never evaluates identity

## 10. Invariants
- Lifecycle ends before listeners start
- Restart equals new lifecycle
- Memory core is lifecycle-aware but policy-blind
