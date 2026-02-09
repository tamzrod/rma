# RMA YAML Settings  Memory Banks

## Status
Documentation-only specification.
Architecture-first. No implementation implied.

---

## 1) Purpose of YAML
YAML is declarative configuration.
It describes what memory banks exist and their fixed properties.

YAML MUST NOT describe protocol semantics, runtime behavior, or transport rules.

---

## 2) Top-Level Shape (Conceptual)
memory:
  banks:
    - <bank>
    - <bank>
    - ...

---

## 3) Bank Fields

### 3.1 Required Fields
Each bank MUST define:

- id
  - Type: unsigned integer conceptually up to 64-bit
  - Must be unique within the configuration
  - Must remain stable for the lifetime of the process (after seal)

- name
  - Type: string
  - Human/debug label only
  - Not used for addressing

- unit_width_bits
  - Type: positive integer
  - Defines the fixed logical unit width of this bank in bits
  - Width is implied by MemoryID; requests do not carry width
  - Custom widths are allowed (e.g., 3, 5, 13) as long as they are explicit

- units
  - Type: positive integer
  - Number of logical units in the bank
  - Capacity is bounded: total_bits = unit_width_bits * units

---

### 3.2 Optional Fields (Human/Policy Only)
Optional fields MAY include:

- notes: string
- tags: list of strings

Optional behavior flags (future, not required now):
- read_only: boolean

---

## 4) Example Configuration (Illustrative)
memory:
  banks:
    - id: 10
      name: bit_main
      unit_width_bits: 1
      units: 1024

    - id: 20
      name: byte_main
      unit_width_bits: 8
      units: 4096

    - id: 30
      name: reg16_main
      unit_width_bits: 16
      units: 10000

    - id: 40
      name: blk128_fast
      unit_width_bits: 128
      units: 512

    - id: 55
      name: enum5_test
      unit_width_bits: 5
      units: 100

---

## 5) Validation Rules (Hard)
The configuration MUST be rejected if:

1. Duplicate bank id values exist
2. unit_width_bits is missing or <= 0
3. units is missing or <= 0
4. total_bits = unit_width_bits * units overflows supported capacity
5. Any field attempts to define protocol semantics or offsets

---

## 6) Guidelines (Soft)
These are not hard failures, but recommended:

- If unit_width_bits % 8 != 0, keep units small and intentional
- Prefer byte-aligned widths for bulk storage
- Use narrow widths (15 bits) for flags/enums, not payload storage

---

## 7) Sealing Relationship
All banks must be declared before sealing.
After sealing:
- no bank additions
- no width changes
- no unit count changes

---

## 8) Non-Goals
This YAML spec does not define:
- Modbus mappings
- address offsets
- scaling, units, engineering values
- transport configuration

Those belong to adapters and boundary layers.

---
