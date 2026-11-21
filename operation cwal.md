# OPERATION CWAL — RMA

Purpose: autonomous, bounded development of the RMA raw-memory foundation. This is a **small workflow**, not a governance framework.

## Startup
1. Read `AGENTS.md`, this file, and only the top six state markers of `workflow/active_work/README.md` and `handoff.md`.
2. Run `bash scripts/cwal-check.sh`. Confirm `tamzrod/rma`, the designated `development` branch, matching markers and safe local/remote state.
3. If CURRENT is NONE, report `NO_CURRENT` and stop. Only an explicitly owner-approved CURRENT task authorizes implementation. A backlog, audit, proposed task or old checkpoint does not.
4. Read the CURRENT task and directly relevant code/contracts/tests. Expand context one dependency at a time. Do not read queued task bodies until their predecessor has completed.

## Execution
1. Work autonomously within CURRENT's exact scope and edit boundaries. Fix ordinary in-scope implementation/test failures without seeking repeated approval.
2. Apply test-first where useful; run task-specific gates plus relevant `gofmt`, `go test ./...`, `go vet ./...`, `go build ./...`. Run `go test -race ./...` for concurrency changes when supported. Each result is PASS, FAIL or NOT_RUN with reason. An unexecuted check is never PASS.
3. Check runtime and network exposure separately: no hardware actuation, external deployment, live device write, or security-sensitive listener enablement is implied by code authorization.
4. On longer tasks, checkpoint verified, meaningful slices to the designated development branch with CURRENT unchanged and status TASK_IN_PROGRESS. Never describe intermediate work as completed.
5. On full acceptance, preserve the completed task under `archive/active-work/<ID>.md`, add concise `docs/verification/<ID>.md` with commands/results/limitations/SHA, update both marker blocks consistently, and commit/push normally to `development`.
6. Read back remote SHA and both markers before claiming TASK_CHECKPOINTED. Continue only to the next *already owner-approved* QUEUED task; otherwise stop.
7. On failure, conflict, safety uncertainty or expanding scope, keep CURRENT, record blocker and stop. Never reset, force-push, overwrite another agent's work, manufacture evidence, silently promote tasks or change the branch.

## Final gate
- `CURRENT: NONE` after accepted tasks means WORK_PACKAGE_READY only when a package was actually approved and completed; it does **not** authorize PR, merge, release or deployment.
- `CHECK` / pre-PR audit is read-only and independent from the development agent. Corrections require an authorized task, not silent self-approval.
- Verify raw-memory correctness, invalid IDs, arithmetic overflow, atomic concurrent access, TCP framing, startup error propagation, lifecycle and RBE loss/ordering **when affected by CURRENT**.
- RBE design must separate atomic memory-change capture from subscriber delivery and protocol-specific event queues. No implicit persistence or semantic interpretation.

**Loop:** matching state → authorized CURRENT → focused implementation → actual tests → remote checkpoint → next approved task or stop.
