# OPERATION CWAL — RMA (OpenHands runner)

Repository: `tamzrod/rma`  
Execution branch: `development`  
CURRENT: read the matching markers in `workflow/active_work/README.md` and `handoff.md`.

**Owner instruction (2026-11-21):** When the owner launches OpenHands with `operation cwal`, resume the already authorized `RMA-001` verification/correction mission even though its current marker is `TASK_BLOCKED`. The prior blocker was lack of a Go 1.25 execution environment, not a rejected engineering design. This authorization allows only RMA-001's existing scope and its normal acceptance/archival steps. It does not approve RBE, TCP refactors, new protocols, deployment, PR creation, merges, or main writes. Do not interpret this as blanket permission to resume other blocked tasks.

## OpenHands startup — execute, don't just plan

1. Open the local checkout. Read `AGENTS.md`, this file, and the first six marker lines in `workflow/active_work/README.md` and `handoff.md`.
2. Verify origin URL identifies `tamzrod/rma`, fetch remote refs safely, and check branch `development`. Do **not** reset, force-checkout, stash away unknown work, or overwrite changes. On an empty workspace, clone the repo and check out `development`. If local changes exist, preserve and reconcile safely or stop.
3. Read only `workflow/active_work/RMA-001.md`, `docs/verification/RMA-001.md`, and directly relevant memory core/test files. Existing remote changes are an unverified checkpoint: inspect, correct and test; do not assume PASS.
4. Run `bash scripts/cwal-check.sh`. If it reports the known `TASK_BLOCKED` state but verifies matching scope/branch, continue by this **explicit owner-approved resume**. If branch, repository, markers, ownership or local/remote refs conflict, stop and report.
5. Verify `go version` (Go >= 1.25). Use an already available toolchain or install Go 1.25+ **only in the isolated OpenHands development environment**, from official trusted distribution and with integrity verification. No changes to devices, production hosts or external services. If provisioning is unavailable, document NOT_RUN and stop.

## RMA-001 work loop

1. Examine current diff relative to `main` and previous checkpoint. Scope strictly to the approved core files/tests; maintain exported public API and wire compatibility.
2. Run focused tests first, reproduce any failures, and fix only actual in-scope defects. Review bit-order and non-byte-aligned copy behavior, read/write atomicity, integer arithmetic, missing-bank errors, sealing concurrency, allocation caps and atomic batch construction.
3. Add relevant negative/race tests when coverage gaps are discovered. Avoid introducing RBE, application semantics, auth, network behavior, protocol adapters, dependencies, persistence or new architecture in RMA-001.
4. Run the gates below and record command, environment and exact result. Fix failures inside RMA-001's boundaries, then rerun affected gates. A command not executed is `NOT_RUN`, never PASS.
5. Maintain `docs/verification/RMA-001.md` with factual evidence (no fabricated results). If any required gate cannot pass, retain CURRENT and record the precise blocker in both current-state files. Do not archive.
6. When all acceptance criteria actually pass, archive the task to `archive/active-work/RMA-001.md`, preserve its requirements, and set both marker blocks to `Repository: tamzrod/rma`, `Package: RMA-CORE-HARDENING`, `Branch: development`, `CURRENT: NONE`, `QUEUED: NONE`, `Status: WORK_PACKAGE_READY`. Update handoff and verification report.
7. Commit coherent changes, ordinary push to `origin/development` only. Fetch/read back remote SHA and marker/file state; only then claim completed remote checkpoint. If unable to push/read back, state local-only or unverified, not complete.

## Verification gates (on OpenHands workspace)

```bash
git status --short
git branch --show-current
git remote -v
git fetch origin development main
bash scripts/cwal-check.sh
go version
gofmt -l internal/core/memory/*.go
go test ./internal/core/memory -count=1
go test -race ./internal/core/memory -count=1
go test ./... -count=1
go vet ./...
go build ./...
git diff --check
```

Use `gofmt -w` on modified Go source as necessary, then repeat formatting check. If the environment cannot run race detection, label it `NOT_RUN` with an explicit reason and keep RMA-001 unaccepted until required verification is possible. If a test involves networking, hardware, external services or mutation beyond approved files, ask for separate authorization rather than running it.

## Hard boundaries

- Preserve RMA as a deterministic raw memory engine: no protocol or AI semantics inside `internal/core/memory`.
- Keep policies, identities, authentication and protocol event delivery outside the core.
- No direct pushes to `main`; no PR creation, merge, release, production deployment or EG628 changes.
- No force pushes, destructive cleanups, unverifiable Git checkpoints or hidden scope expansion.
- Do not auto-start RMA-002 or another mission after RMA-001 closes.

## Required final report

State: CURRENT, status, exact remote SHA (or local-only), changed files, defects corrected, tests PASS/FAIL/NOT_RUN with commands, limits/unknowns and next action. Keep it concise.

**Trigger:** `operation cwal` → verify checkpoint → resume RMA-001 → fix/test → evidence → remote checkpoint → STOP.
