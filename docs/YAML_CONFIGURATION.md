# YAML CONFIGURATION

## 1. Purpose
- Declarative description of memory topology and listeners
- Defines WHAT exists, not WHO or WHEN
- Zero runtime logic
- Zero policy execution
- Zero lifecycle control

---

## 2. Design Principles
- Explicit over implicit (except where explicitly defined)
- Absence implies openness
- Declarative only, no execution semantics
- No speculative or future-facing fields
- No protocol coupling
- No authority enforcement inside RMA

---

## 3. Top-Level Structure
- memory
- listeners

No other top-level sections are defined by RMA.

---

## 4. Memory Banks
Each memory bank defines a fixed, bounded memory region.

Required:
- Unique numeric id
- Fixed unit_width_bits
- Fixed units

Optional:
- init
- access

Memory banks do not define protocol behavior, identity, or lifecycle actions.

---

## 5. init Block (Implicit Sealing)
- Presence of init implicitly marks the memory as SEALED
- No explicit sealed flag exists
- init controls who may perform INIT
- INIT is a lifecycle action, not runtime access
- Evaluated only during bootstrap
- CIDR-based declaration only
- Absence of init means the memory is OPEN

Example intent:
- init defines who may SHAPE memory
- init does not define runtime behavior directly

---

## 6. access Block (Runtime Access Declaration)
- Restricts runtime access to memory
- CIDR-based declaration only
- No operation semantics (read/write) are defined in YAML
- Evaluated only during runtime
- Absence of access means OPEN access

---

## 7. Inheritance Rules
- INIT authority IMPLIES ACCESS authority
- ACCESS authority NEVER implies INIT authority
- Effective runtime access is the UNION of:
  - init.allow_cidr
  - access.allow_cidr

No duplication of CIDRs is required.

---

## 8. Default Behavior
- No init and no access:
  - Memory is OPEN
  - No sealing
  - No access restriction

- init only:
  - Memory is SEALED
  - INIT restricted by init.allow_cidr
  - Runtime access restricted to init.allow_cidr

- access only:
  - Memory is OPEN
  - Runtime access restricted to access.allow_cidr

- init and access:
  - Memory is SEALED
  - INIT restricted by init.allow_cidr
  - Runtime access restricted to (init  access)

---

## 9. Lifecycle and Intent (Non-Declarative)
- YAML does not execute INIT
- YAML does not seal memory
- YAML does not unlock memory
- YAML does not define lifecycle transitions

INIT intent is expressed externally (e.g. via MemoryID flags) and enforced by wrappers.

---

## 10. Listeners
- Define raw data-plane endpoints
- No protocol type enumeration
- No authority role
- No lifecycle role
- No identity or CIDR evaluation

Listeners only expose entry points.

---

## 11. Explicit Non-Goals
- No auth engine
- No authority enforcement
- No explicit sealing flags
- No unlock or reseal mechanisms
- No protocol definitions
- No lifecycle control flags
- No control-plane commands

All policy and lifecycle enforcement occurs OUTSIDE RMA.

---

## 12. Invariants
- MemoryID carries intent, not authority
- init implies sealed memory
- init authority implies access authority
- Unlock requires process restart
- RMA core remains policy-blind
