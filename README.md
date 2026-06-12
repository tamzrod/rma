# RMA — Real-Time Memory Appliance

RMA is a deterministic, protocol-agnostic real-time memory core designed for industrial telemetry and infrastructure systems.

It exposes raw memory as a bounded, controllable resource while keeping policy, authority, and lifecycle enforcement strictly outside the core.

RMA is architecture-first.  
Documentation defines behavior. Code implements contracts.

---

## Design Intent

RMA exists to provide a stable, deterministic truth layer for telemetry systems.

Control decisions depend on visible, reliable state.  
RMA guarantees that state integrity is preserved inside explicit bounds.

It is not a feature framework.  
It is not a policy engine.  
It is a memory core.

---

## Core Principles

- Memory is the product
- Determinism over convenience
- Explicit lifecycle, no implicit mutation
- Strict bounds enforcement
- Protocols are adapters, never core
- Authority is external, never embedded
- No hidden behavior, no runtime magic

---

## What RMA Is

- A raw memory core with strict bounds and deterministic behavior
- A foundation for telemetry adapters (Modbus, MQTT, REST, etc.)
- A reusable infrastructure component
- Designed for industrial, embedded, and distributed systems

---

## What RMA Is NOT

- Not an authentication or identity system
- Not a protocol implementation framework
- Not a policy or authorization engine
- Not dynamically reconfigurable at runtime
- Not self-authorizing or self-mutating

RMA enforces memory integrity.  
Everything else is external.

---

## Architectural Boundaries

The core must remain:

- Policy-blind
- Authority-blind
- Identity-blind
- Transport-agnostic

Adapters may interpret.  
Core never does.

---

## Docker

RMA ships with a multi-stage `Dockerfile` that produces a minimal runtime image with a preloaded default configuration.

### Default configuration

The image bundles `config/default.yaml`:

```yaml
memory:
  banks:
    - id: 0
      name: default
      unit_width_bits: 16
      units: 1024

listeners:
  - name: rawtcp
    bind: ":9000"
```

### Build

```sh
docker build -t rma .
```

### Run

```sh
docker run -p 9000:9000 rma
```

The container starts the raw TCP listener on port `9000` using the preloaded configuration.

### Override configuration

Mount a custom YAML file to replace the default configuration at runtime:

```sh
docker run -p 9000:9000 -v $(pwd)/my-config.yaml:/app/config/default.yaml rma
```

The override must satisfy the same schema and validation rules as any other RMA configuration file.

---

## Repository Layout
