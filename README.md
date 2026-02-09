# RMA — Raw Memory Appliance

RMA is a deterministic, protocol-agnostic memory appliance.
It exposes raw memory as a controllable resource while keeping
policy, authority, and lifecycle strictly outside the core.

This repository is architecture-first.
Code follows contracts, not the other way around.

## Core Principles

- Memory is the product
- Protocols are adapters
- Lifecycle is finite and explicit
- Authority is external
- Runtime behavior is deterministic
- No hidden magic, no implicit control

## What RMA Is

- A raw memory core with strict bounds
- A foundation for external adapters
- A reusable system component
- Designed for industrial and embedded use

## What RMA Is NOT

- Not an authentication system
- Not a protocol framework
- Not a policy engine
- Not dynamically reconfigurable at runtime
- Not self-authorizing

## Repository Layout

.
├─ adapter/      # External-facing adapters (protocols, ingress)
├─ contract/     # Architectural contracts and invariants
├─ core/         # Core logic (policy-blind)
├─ docs/         # Authoritative documentation
├─ errors/       # Shared error definitions
├─ internal/     # Internal helpers and plumbing
├─ memory/       # Memory core implementation
├─ go.mod
└─ README.md

## Authoritative Documentation

All behavior is defined by documentation in `docs/`.
These documents are contracts, not commentary.

- STATE_SEALING.md  
  - Memory lifecycle and sealing rules  
  - INIT phase vs runtime  
  - No runtime unlocks  

- YAML_CONFIGURATION.md  
  - Declarative memory and listener configuration  
  - No authority or lifecycle execution  

- AUTHORITY_EXTERNAL.md  
  - CIDR-based INIT and ACCESS enforcement  
  - Explicitly outside the RMA core  

## Lifecycle Model (High-Level)

1. Load YAML configuration
2. Initialize memory banks
3. Perform INIT where defined
4. Seal memory state
5. Start listeners
6. Runtime read/write only

Unlocking memory requires a process restart with new configuration.

## Authority Model

- RMA core has no knowledge of identity, CIDR, or permissions
- Authority is enforced externally (wrapper, gateway, sidecar)
- INIT authority implies ACCESS authority
- No authority present means open access

## Development Rules

- Architecture changes require doc updates first
- Core must remain policy-blind
- No runtime lifecycle mutation
- Full-file changes only
- PowerShell is used for file operations

## Status

Architecture locked.
Implementation proceeds incrementally under contract control.
