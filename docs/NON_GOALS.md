# Non-Goals

RMA is NOT:
- a database
- a historian
- a protocol gateway
- a message broker
- a cache with eviction policies
- a distributed system (by default)

RMA will NOT:
- auto-scale memory
- persist data to disk
- interpret register semantics
- perform unit conversions
- manage user authentication

These concerns belong to adapters or external systems.
