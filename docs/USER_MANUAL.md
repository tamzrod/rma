# RMA User Manual

**Repository:** [tamzrod/rma](https://github.com/tamzrod/rma)

RMA (Real-Time Memory Appliance) is a deterministic, protocol-agnostic real-time memory core designed for industrial telemetry and infrastructure systems.

---

## Table of Contents

1. [Overview](#1-overview)
2. [Prerequisites](#2-prerequisites)
3. [Installation](#3-installation)
4. [Configuration](#4-configuration)
5. [Running RMA](#5-running-rma)
6. [Running with Docker](#6-running-with-docker)
7. [Memory Bank Concepts](#7-memory-bank-concepts)
8. [Access Control](#8-access-control)
9. [Listeners](#9-listeners)
10. [Troubleshooting](#10-troubleshooting)

---

## 1. Overview

RMA provides a stable, deterministic truth layer for telemetry systems. It manages bounded memory regions and enforces strict access rules, while remaining policy-blind and transport-agnostic.

Key properties:

- **Deterministic** — memory behavior is predictable under concurrency
- **Bounded** — every memory bank has a fixed, declared size
- **Protocol-agnostic** — adapters connect to RMA; RMA does not implement protocols
- **Authority-external** — access policy is declared in configuration, not embedded in core

---

## 2. Prerequisites

- [Go 1.25+](https://go.dev/dl/) (to build from source)
- [Docker](https://docs.docker.com/get-docker/) (to run the container image)
- A YAML configuration file describing your memory topology

---

## 3. Installation

### Build from source

```sh
git clone https://github.com/tamzrod/rma.git
cd rma
go build -o rma ./cmd/rma
```

The resulting `rma` binary is a self-contained executable.

---

## 4. Configuration

RMA is configured with a single YAML file passed at startup. The file has two top-level sections: `memory` and `listeners`.

### Minimal example

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

### Memory bank fields

| Field | Required | Description |
|---|---|---|
| `id` | Yes | Unique numeric identifier for the bank |
| `name` | Yes | Human-readable label |
| `unit_width_bits` | Yes | Width of each memory unit in bits (e.g. 1, 8, 16, 32, 128) |
| `units` | Yes | Number of units in the bank |
| `init` | No | Seals the bank and restricts who may initialise it |
| `access` | No | Restricts runtime read/write access |

### Listener fields

| Field | Required | Description |
|---|---|---|
| `name` | Yes | Human-readable label for the listener |
| `bind` | Yes | Address and port to bind (e.g. `":9000"` or `"0.0.0.0:9001"`) |

For the full configuration reference and validation rules see [YAML_CONFIGURATION.md](YAML_CONFIGURATION.md).

---

## 5. Running RMA

```sh
rma <config.yaml>
```

**Example:**

```sh
rma config/default.yaml
```

On successful startup you will see:

```
RMA started successfully (memory sealed, runtime active)
```

RMA then runs until the process is terminated. There is no interactive shell or control plane.

---

## 6. Running with Docker

### Build the image

```sh
docker build -t rma .
```

### Start with the bundled default configuration

```sh
docker run -p 9000:9000 rma
```

### Override the configuration at runtime

Mount your own YAML file over the bundled default:

```sh
docker run -p 9000:9000 -v $(pwd)/my-config.yaml:/app/config/default.yaml rma
```

The override file must satisfy the same schema and validation rules as any other RMA configuration file.

---

## 7. Memory Bank Concepts

### Open memory (default)

A bank with no `init` or `access` blocks is fully open — any caller may read or write it.

```yaml
- id: 10
  name: bit_main
  unit_width_bits: 1
  units: 1024
```

### Sealed memory

Adding an `init` block seals the bank. Only callers whose IP address falls within the declared CIDRs may perform the INIT lifecycle action. Callers permitted to INIT are also permitted runtime access.

```yaml
- id: 20
  name: byte_main
  unit_width_bits: 8
  units: 4096
  init:
    allow_cidr:
      - "127.0.0.1/32"
      - "192.168.1.0/24"
```

### Open memory with restricted runtime access

Adding only an `access` block leaves the bank unsealed but limits who may read or write at runtime.

```yaml
- id: 30
  name: reg16_main
  unit_width_bits: 16
  units: 10000
  access:
    allow_cidr:
      - "10.0.0.0/8"
```

### Sealed memory with expanded runtime access

Combining `init` and `access` seals the bank and grants runtime access to the union of both CIDR lists.

```yaml
- id: 40
  name: blk128_fast
  unit_width_bits: 128
  units: 512
  init:
    allow_cidr:
      - "127.0.0.1/32"
  access:
    allow_cidr:
      - "192.168.1.0/24"
```

---

## 8. Access Control

Access control in RMA is CIDR-based and declarative. The rules are:

- **INIT authority implies ACCESS authority** — a caller allowed to INIT is also allowed runtime access.
- **ACCESS authority does not imply INIT authority** — a caller listed only in `access` cannot INIT a sealed bank.
- **Effective runtime access** is the union of `init.allow_cidr` and `access.allow_cidr`.
- **Absence of both blocks** means the bank is fully open.
- **Unlock requires a process restart** — there is no runtime mechanism to unseal a sealed bank.

RMA enforces memory integrity. Authentication, authorisation, and identity management are the responsibility of external systems.

---

## 9. Listeners

Listeners expose raw data-plane endpoints. They carry no protocol, authority, or lifecycle role.

```yaml
listeners:
  - name: tcp_9000
    bind: "0.0.0.0:9000"

  - name: tcp_9001
    bind: "0.0.0.0:9001"
```

Multiple listeners can be declared. Each binds independently. Adapter code connects to a listener to provide protocol-specific behaviour (Modbus, MQTT, REST, etc.).

---

## 10. Troubleshooting

### `usage: rma <config.yaml>`

RMA requires exactly one argument — the path to your configuration file.

```sh
rma path/to/config.yaml
```

### `startup failed: ...`

The error message describes which configuration value failed validation. Check your YAML against the field tables in [Section 4](#4-configuration) and the detailed rules in [YAML_CONFIGURATION.md](YAML_CONFIGURATION.md).

### Port already in use

Ensure no other process is bound to the port declared in your `listeners` section, or change the `bind` address in your configuration file.

### Memory not accessible after restart

If a bank is sealed (`init` block present), sealed state is restored from configuration on every startup. Callers must re-perform the INIT action after each process restart.

---

*For architecture and design documentation see the other files in the [docs/](.) directory.*
