# Adapter Rules

Adapters exist to translate external protocols into RMA core operations.

Adapters MAY:
- handle network I/O
- parse protocol frames
- serialize and deserialize data
- perform logging and metrics
- manage protocol-specific state

Adapters MUST:
- treat core as a black box
- use only exposed core contracts
- never bypass bounds checks
- never mutate memory outside core APIs

Adapters MUST NOT:
- embed memory logic
- assume memory layout beyond contracts
- introduce non-determinism into core behavior
