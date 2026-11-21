Repository: tamzrod/rma
Package: NONE
Branch: development
CURRENT: NONE
QUEUED: NONE
Status: NO_CURRENT

# RMA — CWAL Handoff

## Installed workflow
- Lightweight CWAL adapted from Atlas-PPC, DNP3, PowerPlantMonitor and EG628; RMA-specific architecture stays authoritative.
- Execution branch is `development`; no automatic development task or merge.
- No production changes, tests, or deployments authorized by this installation.

## Next owner decision
Promote a bounded first CURRENT task to harden `internal/core/memory`: concurrency, missing-bank errors, arithmetic/allocation bounds, and race/negative tests. Keep the raw protocol-neutral contract. Alternatively request PLAN for RBE first.

## Checkpoint convention
On work, record exact commit SHA, tests PASS/FAIL/NOT_RUN, remaining risks and the next safe action. Keep this handoff short; archived evidence belongs under `docs/verification/`.
