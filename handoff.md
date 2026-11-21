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

## Current owner-approved action
Launch OpenHands with `operation cwal` to resume RMA-001 verification and in-scope corrections, following the explicit resume authorization in `operation cwal.md`. Do not promote a new task.

## Checkpoint convention
On work, record exact commit SHA, tests PASS/FAIL/NOT_RUN, remaining risks and the next safe action. Keep this handoff short; archived evidence belongs under `docs/verification/`.

## RMA-001 checkpoint — tests not executed
- Implemented memory locks, checked bounds/allocations, safe missing-bank errors and focused tests on development.
- Required Go 1.25+ test/race/vet/build gates are NOT_RUN; environment has Go 1.23.2 and cannot access GitHub through local git.
- Status TASK_BLOCKED pending real build/test gate. See `docs/verification/RMA-001.md`.
- Do not archive or promote automatically. No PR, merge, deploy or main changes.

## OpenHands runner instruction — 2026-11-21
- Owner authorized OpenHands to resume RMA-001 despite TASK_BLOCKED due to unavailable Go verification in the earlier environment.
- Verify existing code and tests with Go 1.25+, correct in scope, and record real PASS/FAIL/NOT_RUN.
- Task markers remain TASK_BLOCKED until OpenHands actually verifies or updates them. Do not infer success from this instruction update.
- Stop after RMA-001; no RBE work, PR, merge, main push, hardware or deployment.
