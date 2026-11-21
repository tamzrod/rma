# RMA — Agent Entry Point

Repository: `tamzrod/rma`
Development branch: `development`

## Routes
- **operation cwal**: read `operation cwal.md`, the first marker blocks in `workflow/active_work/README.md` and `handoff.md`, and the authorized CURRENT task only. Run `bash scripts/cwal-check.sh` before work. The owner-approved OpenHands resume for blocked RMA-001 is specified in `operation cwal.md`; do not assume the exception applies to other tasks.
- **PLAN**: read-only design, task proposal, or risk evaluation; no implementation or promotion.
- **CHECK / pre-PR audit**: independent read-only assessment of actual diff, tests, source contracts, and evidence. Findings do not authorize fixes.

## Non-negotiable RMA contract
- RMA core owns bounded raw memory, consistent atomic operations, and (if approved) raw committed-change capture. It must never interpret protocol/application/AI semantics.
- No Modbus, DNP3, IEC 61850, OPC UA, MQTT, AI agent, authentication, policy, or network logic inside `internal/core/memory`.
- Protocol adapters own mapping and delivery. Authorization remains in external wrappers. Do not silently change the published raw TCP interface.
- Maintain the fixed-layout sealing contract; prevent unbounded allocations, races, panics on malformed input, and silent change loss.
- No hardware, deployment, production listener changes, force-push, direct main push, PR creation, or merge without separate owner authorization.

## Engineering standards
- Read the smallest relevant source and tests; do not bulk-import previous CWAL repositories or copy their project-specific governance.
- One coherent deliverable at a time. Prefer Go standard library; avoid new dependencies unless justified.
- Add negative, boundary, concurrency, and race tests when affected. Report actual PASS / FAIL / NOT_RUN and exact command.
- Verify remote SHA after any authorized checkpoint; stop on conflicting refs or edits.
- Current work is selected only by matching markers, not by this file or by a roadmap.
