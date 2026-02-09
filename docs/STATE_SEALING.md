# STATE SEALING

## 1. Purpose
- Define memory lifecycle boundaries
- Separate lifecycle from security, policy, and protocol
- Guarantee deterministic runtime behavior
- Prevent runtime structural mutation

---

## 2. Scope
- Applies per memory bank
- Applies only within a single process lifecycle
- Independent of listeners, protocols, and data access mechanisms
- Independent of authority and identity evaluation

---

## 3. Lifecycle Phases
- Bootstrap phase
- INIT phase
- Runtime phase

Lifecycle phases are strictly ordered and non-overlapping.

---

## 4. Sealing Rules (Implicit)
- Presence of an init block implies the memory is SEALED
- No explicit sealed flag exists
- Sealing is implicit and declarative
- Sealing is one-way for the lifetime of the process
- No runtime unseal is permitted

Absence of init implies the memory is OPEN.

---

## 5. INIT Semantics
- INIT is a controlled memory-shaping operation
- INIT occurs only during bootstrap, before runtime
- INIT intent is explicit via MemoryID flags
- INIT is not a protocol command
- INIT is not a runtime operation

INIT defines who may SHAPE memory, not who may USE it.

---

## 6. Runtime Semantics
- Runtime permits data-plane access only
- Runtime forbids:
  - lifecycle mutation
  - resealing
  - resizing
  - unlocking
- Runtime behavior is deterministic and bounded

Any lifecycle change requires process restart.

---

## 7. MemoryID Intent Flags
- INIT flag declares INIT intent
- DATA flag is the default runtime intent
- Flags declare intent only
- Flags do not grant authority
- Flags do not alter memory layout or structure

Intent evaluation is external to the memory core.

---

## 8. Forbidden Operations
- Unlock flags
- Reseal flags
- Control bits embedded in memory data
- Special addresses or magic regions
- Runtime lifecycle commands
- Data-plane driven lifecycle changes

Any of the above constitutes an architectural violation.

---

## 9. Authority Boundary
- State sealing is not an authority decision
- Authority is evaluated externally (wrapper, gateway, sidecar)
- Authority may allow or deny INIT and ACCESS
- RMA core never evaluates identity, CIDR, or permissions

INIT authority implies ACCESS authority.
ACCESS authority never implies INIT authority.

---

## 10. Invariants
- init implies sealed memory
- INIT authority implies ACCESS authority
- Lifecycle ends before listeners start
- Unlock requires full process restart
- MemoryID carries intent, not authority
- RMA core is lifecycle-aware but policy-blind
