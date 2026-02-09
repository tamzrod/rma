# RMA  Rod Memory Appliance

## Purpose
RMA is a deterministic, protocol-agnostic, in-memory register engine.

It exists to provide:
- fast, bounded, predictable memory access
- strict separation between memory and protocols
- a reusable core for industrial and non-industrial systems

## Core Rules
- Memory is the product
- No protocol logic in core
- No network code in core
- No serialization in core
- Deterministic behavior over feature richness
- All protocol support is adapter-based
