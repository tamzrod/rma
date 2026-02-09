# AUTHORITY (EXTERNAL)

## 1. Purpose
- Decide who may INIT or ACCESS
- Enforce policy without polluting RMA
- Remain replaceable and optional

## 2. Scope
- External wrapper, gateway, or sidecar
- Not part of RMA core
- Not embedded in memory logic

## 3. Authority Subjects
- Source IP or CIDR
- Listener context (optional)
- Process phase (INIT versus runtime)

## 4. INIT Authorization
- Evaluated only if init exists
- CIDR allow-list based
- INIT flag required
- One-time per lifecycle

## 5. Access Authorization
- Evaluated during runtime
- CIDR allow-list based
- Effective access includes INIT CIDRs

## 6. Phase Enforcement
- INIT allowed only during bootstrap
- Runtime forbids INIT
- Authority cannot change lifecycle

## 7. Decision Order
- Identify source
- Determine phase
- Evaluate INIT versus ACCESS
- Allow or reject request

## 8. Defaults
- No authority present means open access
- Sealed memory without INIT authorization causes startup failure

## 9. Forbidden Capabilities
- Unlocking memory
- Resealing memory
- Resizing memory
- Mutating memory topology

## 10. Invariants
- Authority never alters memory layout
- Authority never lives in RMA core
- Policy changes do not require core changes
