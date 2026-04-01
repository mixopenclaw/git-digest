# Task Queue (Autopilot-managed)

This file is the canonical task queue for autonomous agents working on the git-digest project.
Agents and humans should use the sections below to coordinate work.

Task format (example):
- ID: GDT-001
- Title: Short descriptive title
- Priority: P0/P1/P2 (P0 highest)
- Estimate: 30m | 2h | 1d
- Assignee: @mix or unassigned
- Repo/Path: git-digest/ (optional)
- Notes: Short notes, blockers, links

When an autonomous runner works a task it MUST:
1. Move the task from Ready -> In Progress and add a timestamp + runner name.
2. Create a branch named: task/GDT-001-<short-title>
3. Work iteratively with tests. Commit frequently with clear messages.
4. If "allowPush: true" is configured, push the branch and open a PR when ready.
   Otherwise, create a patch under tasks/patches/ and mark the task Done locally.
5. If blocked, move the task to Blocked with a short reason and a next-action.
6. Update the task entry with final status, commit hashes (if any), and PR URL (if opened).

---

## Ready

- (Add new tasks here. Use the format above. Autopilot will pick highest-priority Ready task.)

## In Progress

- ID: OPENCODE-BATCH-01
  Title: Apply batch-01 plan (10 items)
  Priority: P0
  Estimate: 34h
  Assignee: @mix
  Runner: Kodex
  Started: 2026-04-01 01:36 Asia/Singapore
  Branch: opencode/batch-01-apply
  Commits:
    - c907bf4: api/scan: add POST /api/scan endpoint
    - 9486b7c: api/results: add GET /api/scan/:id/results
    - 9f0ffbd: api/metrics: integrate Prometheus metrics
    - 93d0dde: api/logging: structured server logs
    - a2fdc82: api/logging: structured server logs (use backend/logx)
    - dd70782: api/db: add migrations and simple store layer
    - a24738c: api/trace: add basic tracing hooks
    - 5c99dba: api/cleanup: add retention worker for old artifacts
    - e8656b9: api/health-checks: add DB and storage checks to readiness
    - 1db7d53: api/docs: add OpenAPI spec skeleton
    - 20eeb07: scanner/init: add scanner package skeleton
  Patches:
    - tasks/patches/0001-api-results-add-GET-api-scan-id-results.patch
    - tasks/patches/0002-api-metrics-integrate-Prometheus-metrics.patch
    - tasks/patches/0003-api-logging-structured-server-logs.patch
    - tasks/patches/0004-api-logging-structured-server-logs-use-backend-logx.patch
    - tasks/patches/0005-api-db-add-migrations-and-simple-store-layer.patch
    - tasks/patches/0006-api-trace-add-basic-tracing-hooks.patch
    - tasks/patches/0007-api-cleanup-add-retention-worker-for-old-artifacts.patch
    - tasks/patches/0008-api-health-checks-add-DB-and-storage-checks-to-readi.patch
    - tasks/patches/0009-api-docs-add-OpenAPI-spec-skeleton.patch
    - tasks/patches/0010-scanner-init-add-scanner-package-skeleton.patch
    - tasks/patches/0001-api-scan-add-POST-api-scan-endpoint.patch
  allowPush: true
  RemoteBranch: opencode/batch-01-apply-on-main
  PR: https://github.com/mixopenclaw/git-digest/pull/12
  Notes: Created scaffolding and patch files for the 10-item batch plan. No credentials were used and no remote pushes were made. To push branches and open PRs, set allowPush: true on the task or reply with `ALLOW_PUSH`.


## Blocked

- (Tasks blocked by external factors; include reason and suggested next step.)

## Done

- (Completed tasks. Add commit hash or PR URL where applicable.)

---

## Maintainer notes

- Patches created by autonomous runs will be stored in: tasks/patches/
- For safety, autonomous runners will NOT push or open PRs unless explicitly allowed by the human owner (Mix). To grant permission, reply to the agent with: `ALLOW_PUSH` or configure the queue entry `allowPush: true` on a per-task basis.
- If a task requires human approval (credentials, destructive changes, or merges), the runner will stop and add a clear request in this file and in the run log.

---

## Global Autopilot Settings

- globalAllowPush: true  # granted by Mix in-channel on 2026-04-01 00:58 (Asia/Singapore)
- automergeEnabled: true # PRs will be merged automatically when CI passes unless interrupted
- models:
  - planner: github-copilot/gpt-4o
  - builder: github-copilot/gpt-4.1
  - reviewer: github-copilot/gpt-5-mini
  - fallback_chain: ["openrouter/auto", "mistral/mistral-large-latest"]

Notes: These global settings were set from the #dev channel per user instruction. Per-task `allowPush` entries are still honored; globalAllowPush acts as an override for agents when present.
