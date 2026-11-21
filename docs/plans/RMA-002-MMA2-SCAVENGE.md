# RMA-002 — MMA2 Memory Mechanism Scavenging Plan

Status: DOCUMENTED_ONLY / NOT_PROMOTED
Repository: tamzrod/rma
Target branch: development
Planning baseline: 9be07629bf7632434a7bbc0da626ea93e5169112 (reported RMA-001 checkpoint)
CURRENT: NONE
QUEUED: NONE

## Purpose

Reuse selected, verified concepts from tamzrod/mma2 to strengthen RMA as a deterministic protocol-agnostic memory appliance. RMA is usable by general applications, including game engines; MMA2 protocol features must not be imported into the RMA memory core.

This document is a candidate inventory and planning record. It is NOT implementation authority, a finalized API, a new architecture revision, or a claim that MMA2 packages can be copied unchanged.

## Scavenge candidates — focused memory behavior only

| Priority | Mechanism | MMA2 evidence | RMA adaptation |
| --- | --- | --- | --- |
| P1 | Atomic observed writes | `internal/memorycore/observed_write.go` | Compare previous and incoming logical bits under the same bank write lock. Do not expose MMA2-specific bit/register areas or seal probes in the core. |
| P1 | Actual-value change detection | `internal/rbe/engine.go` (`emitChanged`) | Suppress identical writes and ignore unused padding; support arbitrary RMA unit widths, not just bits and 16-bit registers. |
| P1 | Selective RBE range monitoring | `internal/rbe/engine.go` (`Rule`, `intersects`) | Only configured ranges trigger RBE. Multiple overlapping selections remain independent. Rules and application policy stay outside memory core. |
| P1 | Compact RBE identifiers | `internal/rbe/engine.go`, `internal/rbe/tcp.go` | Design an optional 16-bit, big-endian RBE rule ID: 1..65535; 0 reserved. Two-byte output is a proposed adapter contract, not an internal core limitation. |
| P1 | Multiple independent subscribers | `internal/rbe/tcp.go` | Reuse independent per-subscriber delivery semantics in an optional notification layer; never let slow subscribers block memory writes. |
| P1 | Bounded notifications and overflow handling | `internal/rbe/tcp.go` | Bounded queues with explicit reconnect/resynchronization behavior; two-byte rule IDs alone do not guarantee delivery or carry revisions. |
| P1 | Selective State Sealing | `internal/memorycore/observed_write.go`, `internal/rbe/engine.go` | Separate follow-up design for optional per-bank/range protection in an authority/policy layer; retain pre-write observation semantics without embedding authorization in the core. |
| P1 | Pre-write state observation | `internal/memorycore/observed_write.go` (`BitProbe`) | Evaluate state consistently with writes; any sealing design must explicitly specify lock ordering, write permissions, and RBE suppression independently. |
| P1 | Committed-write observation | `internal/memorycore/observer.go` | Consider a lightweight internal committed-change record; do not allow arbitrary blocking/reentrant callbacks under bank locks. |
| P2 | Dirty-range coalescing | `internal/persistence/observer.go` (`coalesce`, `DirtyRange`) | Reuse range math only, preserving ownership boundaries; do not import persistence or clear-on-read dirty snapshots as a shared event cursor. |
| P2 | Generation/revision concepts | `internal/persistence/observer.go` (`dirtyGen`) | Prefer bounded, per-bank monotonic revisions with explicit wraparound and history-gap contracts, if justified. |
| P1 | Differential and concurrent tests | `internal/rbe/engine_test.go` | Adapt identical-write, padding, overlap, randomized reference-model, and concurrency cases for RMA widths 1, 3, 5, 8, 16, 32, 64, 128 bits. |

## RMA-002 proposed bounded implementation slice (requires separate promotion)

1. Atomic generic change detection within existing bank writes.
2. Selective logical-unit RBE rule matching independent of protocols.
3. Proposed 16-bit rule identifiers (0 reserved) for an optional 2-byte trigger adapter.
4. Bounded, nonblocking subscriber/notification behavior; define explicit loss/reconciliation contract.
5. Tests for unchanged writes, arbitrary widths, padding, intersecting ranges, invalid selections, rejected writes, contention and subscriber overflow.

Do not silently combine revision-journal design, sealing implementation, or persistence into RMA-002. Evaluate revisions separately if needed for multi-consumer loss detection and coherent snapshot reads.

## State Sealing — important, selective, but separately bounded

State Sealing is a desired capability, not an automatic global state and not a protocol feature. Define which banks or ranges are sealable, what operation unseals them, which writes are rejected or permitted, and whether notifications are suppressed while sealed. Preserve the distinction between memory changes and permission to publish notifications. Under the RMA architecture contract, authority/policy must run outside the raw memory core.

## Explicit exclusions

- Modbus, Raw Ingest, DNP3, protocol-specific address spaces, function codes, packet parsing, and protocol adapters.
- MMA2 one-byte RBE wire compatibility as a requirement; the new 16-bit format is intentionally a proposed different contract.
- Persistence, snapshots, restore, disk I/O, access-event reporting, IP filtering, user authentication, and unrelated MMA2 features.
- Game-engine domain semantics, entity models, inventory rules, and multi-bank transactions inside core memory.
- Any claim that a two-byte RBE trigger contains value, address, revision, delivery acknowledgment, or replay data.

## HARD design questions before code

- Verify precise existing RMA bit packing and padding semantics against actual `bank.go` and bit-copy implementation.
- Specify whether RBE dispatch occurs after the bank lock is released; preserve record ordering across concurrent writers.
- Specify how selected ranges are validated without narrowing RMA's uint64 address space.
- Specify how a consumer reconciles after a lost two-byte notification.
- Determine if revision/history tracking is essential in RMA-002 or warrants a later package.
- Specify State Sealing policy and change-notification interaction independently.
- Verify whether any proposal changes core responsibility, memory ownership, adapter boundaries or concurrency model; if so, obtain explicit architecture revision approval first.

## Acceptance tests (planning inventory, NOT_RUN)

- Identical write produces no RBE trigger.
- Single-bit and 3-bit unit comparison ignores padding and preserves neighbors.
- Selected and nonselected writes behave independently.
- Overlapping selections produce their independent trigger IDs once per changed write.
- Rejected writes produce no notification and no mutation.
- 16-bit ID boundaries: 0 rejected, 1 and 65535 accepted, duplicates rejected.
- Two-byte big-endian serialization is tested at the adapter boundary.
- Concurrent writers cannot produce torn reads or falsely ordered change reports.
- Slow subscriber cannot block memory writes; overflow has testable reconciliation behavior.
- Randomized reference-model tests over diverse unit widths.
- RMA-001 tests and standard formatting, race, test, vet, build, diff-check and CWAL gates continue passing.

## CWAL governance

Documenting this plan does not promote RMA-002, update `workflow/active_work` or `handoff.md`, authorize implementation, merge, deployment, or modify `main`.

CURRENT: NONE
QUEUED: NONE
Next safe action: owner explicitly authorizes a bounded implementation work package after resolving HARD questions.
