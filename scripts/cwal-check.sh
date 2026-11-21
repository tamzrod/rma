#!/usr/bin/env bash
# Read-only state check. Does not establish task authorization or remote freshness.
set -euo pipefail
cd "$(dirname "$0")/.."
fail() { echo "CWAL CHECK FAIL: $*" >&2; exit 1; }
for f in "AGENTS.md" "operation cwal.md" "workflow/active_work/README.md" "handoff.md"; do
  [[ -f "$f" ]] || fail "missing $f"
done
markers() {
  local key value
  for key in Repository Package Branch CURRENT QUEUED Status; do
    value="$(sed -n "s/^$key: //p" "$1" | head -n 1)"
    [[ -n "$value" ]] || fail "missing $key in $1"
    printf '%s\n' "$value"
  done
}
mapfile -t a < <(markers workflow/active_work/README.md)
mapfile -t b < <(markers handoff.md)
for i in "${!a[@]}"; do [[ "${a[i]}" == "${b[i]}" ]] || fail "marker mismatch at $i"; done
[[ "${a[0]}" == "tamzrod/rma" ]] || fail "wrong repository"
[[ "${a[2]}" == "development" ]] || fail "wrong branch marker"
[[ "$(git branch --show-current)" == "development" ]] || fail "wrong checkout branch"
[[ -z "$(git status --porcelain)" ]] || fail "dirty worktree; preserve/reconcile first"
if [[ "${a[3]}" == NONE ]]; then
  [[ "${a[4]}" == NONE && "${a[5]}" == NO_CURRENT ]] || fail "inconsistent empty state"
  echo "CWAL CHECK OK: NO_CURRENT (no engineering authority)"; exit 0
fi
[[ "${a[1]}" != NONE ]] || fail "missing package"
[[ "${a[3]}" =~ ^[A-Za-z0-9._-]+$ ]] || fail "invalid task id"
[[ -f "workflow/active_work/${a[3]}.md" ]] || fail "CURRENT task file missing"
case "${a[5]}" in TASK_READY|TASK_IN_PROGRESS|TASK_BLOCKED) ;; *) fail "invalid status";; esac
echo "CWAL CHECK OK: CURRENT=${a[3]} (state only; verify actual remote ref and owner-approved scope)"
