# RMA Core Contract (Derived from MMA2.0 Salvage)

## Status
Derived from MMA2.0 forensic salvage.
This document refines and extends the existing CORE_CONTRACT.md.
Architecture-first. No implementation implied.

---

## 1. Core Purpose
RMA core is a deterministic, bounded, protocol-agnostic memory engine.

The core provides:
- Memory areas and layouts
- Deterministic read/write operations
- Bounds enforcement
- Internal concurrency safety
- State lifecycle gating (state sealing)

The core does NOT provide:
- Protocol parsing or serialization
- Network I/O
- Persistence or caching
- Semantic meaning, scaling, or units
- Adapter orchestration

---

## 2. Ownership & Boundary Rules
The core is the sole owner of memory state.

External components:
- MUST NOT mutate memory except through core contracts
- MUST NOT assume internal layout
- MUST treat the core as a black box

Protocol concepts MUST NOT appear in core.

---

## 3. Memory Model Contract
### 3.1 Primitives
- Registers: 16-bit words
- Bits: single-bit addressing

### 3.2 Memory Areas
Each area:
- Has a fixed, bounded size
- Has a stable identifier
- Has a defined layout

Areas are immutable after sealing.

### 3.3 Addressing
- Address interpretation is core-owned
- Adapters submit address requests
- Internal addressing details are not part of the contract

### 3.4 Bounds Enforcement
All operations enforce:
- Area existence
- Address validity
- Range validity
- Domain correctness

Failures return explicit errors.

---

## 4. Read / Write Operations
### 4.1 Determinism
Given the same configuration and operation order,
observable behavior MUST be deterministic.

### 4.2 Atomicity
Each operation is atomic at the request level unless explicitly stated otherwise.

### 4.3 Concurrency
The core is internally concurrency-safe.
Adapters must not provide locks.

---

## 5. State Lifecycle & Sealing
### 5.1 States
- Unsealed: configuration allowed
- Sealed: runtime-only operations

### 5.2 Seal Rules
- Seal is irreversible
- No configuration mutation post-seal

### 5.3 Authority of Sealing
Only the core performs the seal transition.
External requests may ask but not enforce sealing.

---

## 6. Authority Contract (Boundary Layer)
Authority is outside the core.

Authority:
- Decides if an operation may be attempted
- Never mutates memory
- Never bypasses bounds or sealing rules

Placement:
Adapter  Authority  Core

---

## 7. Adapter Contract
Adapters:
- Parse protocol frames
- Translate requests into core operations
- Map core errors to protocol responses

Adapters MUST NOT:
- Duplicate memory logic
- Maintain shadow layouts
- Bypass authority or core checks

---

## 8. Error Model
Core exposes explicit error categories:
- AreaNotFound
- AddressOutOfRange
- LengthOutOfRange
- DomainMismatch
- SealedViolation
- InternalInvariantViolation

---

## 9. Hard Invariants (STOP Conditions)
Violation of any is fatal to RMA compliance:
1. Protocol logic inside core
2. Adapter bypass of core bounds
3. Observable nondeterminism
4. Unbounded memory growth
5. Reversible or bypassable sealing
6. Authority mutating memory

---

## 10. Non-Goals
The core will not:
- Persist data
- Scale horizontally by itself
- Perform semantic interpretation
- Implement protocol schemas

---
