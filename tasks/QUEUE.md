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

- (Tasks currently being worked on.)

## Blocked

- (Tasks blocked by external factors; include reason and suggested next step.)

## Done

- ID: OPENCODE-BATCH-01
  Title: Apply batch-01 plan (10 items)
  Priority: P0
  Assignee: @mix
  Runner: Kodex
  Started: 2026-04-01 01:36 Asia/Singapore
  Completed: 2026-04-01 03:01:05 UTC
  Branch: opencode/batch-01-apply-on-main
  Commits:
    - 208e97f: fix(batch-01): resolve compilation errors, add unit tests for all backend/scanner packages
    - (previous scaffolding commits retained in history)
  PR: https://github.com/mixopenclaw/git-digest/pull/12
  MergeCommit: a3a01012e40989b276f479fc6d7455031afebab4
  Notes: Merged after CI passed. Changes included scaffolding for API endpoints, Prometheus metrics, structured logging, DB migrations/store layer, tracing hooks, cleanup worker, health checks, OpenAPI skeleton, scanner package, and multiple unit tests. Local and CI tests passed.

---

## Maintainer notes

- Patches created by autonomous runs are stored in: tasks/patches/
- For safety, autonomous runners STOP and request explicit human approval if a task requires credentials, destructive changes, or merges unless per-task permissions are set. In this case OPENCODE-BATCH-01 had allowPush: true and automergeEnabled: true, and was merged automatically when CI passed.

---

## Global Autopilot Settings

- globalAllowPush: true  # granted by Mix in-channel on 2026-04-01 00:58 (Asia/Singapore)
- automergeEnabled: true # PRs will be merged automatically when CI passes unless interrupted
- models:
  - planner: github-copilot/gpt-4o
  - builder: github-copilot/gpt-4.1
  - reviewer: github-copilot/gpt-5-mini
  - fallback_chain: ["openrouter/auto", "mistral/mistral-large-latest"]
