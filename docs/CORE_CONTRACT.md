# Core Contract

The core of RMA is responsible ONLY for:
- managing memory areas
- enforcing bounds and access rules
- providing deterministic read/write behavior

The core MUST NOT:
- know about protocols (Modbus, OPC, REST, MQTT, etc.)
- perform serialization or parsing
- open sockets or handle networking
- perform logging beyond fatal invariants
- depend on adapters

Core code must be:
- testable in isolation
- deterministic under concurrency
- bounded in memory usage
