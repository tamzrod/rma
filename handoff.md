Repository: tamzrod/rma
Package: RMA-CORE-HARDENING
Branch: development
CURRENT: RMA-001
QUEUED: NONE
Status: TASK_BLOCKED

# RMA — CWAL Handoff

## Installed workflow
- Lightweight CWAL adapted from Atlas-PPC, DNP3, PowerPlantMonitor and EG628; RMA-specific architecture stays authoritative.
- Execution branch is `development`; no automatic development task or merge.
- No production changes, tests, or deployments authorized by this installation.

## Next owner decision
Promote a bounded first CURRENT task to harden `internal/core/memory`: concurrency, missing-bank errors, arithmetic/allocation bounds, and race/negative tests. Keep the raw protocol-neutral contract. Alternatively request PLAN for RBE first.

## Checkpoint convention
On work, record exact commit SHA, tests PASS/FAIL/NOT_RUN, remaining risks and the next safe action. Keep this handoff short; archived evidence belongs under `docs/verification/`.

## RMA-001 checkpoint — tests not executed
- Implemented memory locks, checked bounds/allocations, safe missing-bank errors and focused tests on development.
- Required Go 1.25+ test/race/vet/build gates are NOT_RUN; environment has Go 1.23.2 and cannot access GitHub through local git.
- Status TASK_BLOCKED pending real build/test gate. See `docs/verification/RMA-001.md`.
- Do not archive or promote automatically. No PR, merge, deploy or main changes.
