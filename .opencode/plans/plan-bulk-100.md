# Plan: 100 Actionable Commits (Bulk)

This plan lists 100 ordered, actionable commits grouped by area so an automated build/CI can apply them. Each entry includes:
- Sequential number (1..100)
- Commit message (<=80 chars)
- Files to change (paths)
- 1–3 line patch summary (specific edits)
- Tests: yes/no and brief test intent

Group distribution: CLI (10), Frontend (15), Backend (20), TODO Scanner (10), CI (10), Tests (15), Docs (10), Examples (10).

---

CLI (1..10)

1. Commit: `cli/init: add CLI entrypoint and help`  
Files: `cli/main.go`, `cli/README.md`  
Patch: Add `main.go` with a basic cobra/flag-based entrypoint, implement `--help` and `--version` flags; add README stub.  
Tests: yes — unit test for help output parsing

2. Commit: `cli/flags: add global flags and config loading`  
Files: `cli/config.go`, `cli/main.go`  
Patch: Implement config struct, load from `--config` path and env vars; validate required fields.  
Tests: yes — verify config precedence and error when missing required fields

3. Commit: `cli/commands: add `init` command skeleton`  
Files: `cli/cmd/init.go`, `cli/main.go`  
Patch: Add `init` subcommand with flags `--force` and `--template`; wire into root command.  
Tests: no — integration covered later

4. Commit: `cli/commands: add `scan` command skeleton`  
Files: `cli/cmd/scan.go`, `cli/main.go`  
Patch: Add `scan` command with `--path` and `--format` flags; stubbed runner that delegates to scanner package.  
Tests: yes — unit test ensures correct flags are passed to runner

5. Commit: `cli/logging: add structured logging and verbosity`  
Files: `cli/log.go`, `cli/main.go`  
Patch: Add log package with levels; wire a `--verbose` flag to increase log level; replace prints in main.  
Tests: yes — ensure log level toggles and output contains level prefix

6. Commit: `cli/output: add JSON and human output formats`  
Files: `cli/output.go`, `cli/cmd/scan.go`  
Patch: Implement `RenderJSON()` and `RenderHuman()` helpers; `scan` command chooses format via `--format`.  
Tests: yes — assert JSON output validity and human-readable string presence

7. Commit: `cli/errors: unify error handling and exit codes`  
Files: `cli/errors.go`, `cli/main.go`  
Patch: Introduce error types with mapped exit codes; ensure `main` uses consistent `os.Exit` codes on failure.  
Tests: yes — unit tests for error-to-exit-code mapping

8. Commit: `cli/plugins: add plugin discovery interface`  
Files: `cli/plugin.go`, `plugins/README.md`  
Patch: Add plugin interface, discovery from configured `plugins` dir; add README describing plugin contract.  
Tests: no — integration later

9. Commit: `cli/completion: add shell completion scripts`  
Files: `cli/completion.go`, `cli/README.md`  
Patch: Generate bash/zsh completion functions; document how to install them in README.  
Tests: no — manual verification recommended

10. Commit: `cli/auth: add auth hooks and token caching`  
Files: `cli/auth.go`, `cli/cache.go`  
Patch: Add token store (file-based), auth hook interface, and cache TTL; stubbed token refresh logic.  
Tests: yes — unit tests for cache read/write and TTL expiration

Frontend (11..25)

11. Commit: `web/init: add React + Vite app skeleton`  
Files: `frontend/package.json`, `frontend/vite.config.ts`, `frontend/src/main.jsx`, `frontend/src/index.css`  
Patch: Initialize frontend scaffold with Vite, React entrypoint, basic CSS import, and package scripts.  
Tests: no — build/test in CI later

12. Commit: `web/layout: add app shell and responsive nav`  
Files: `frontend/src/App.jsx`, `frontend/src/components/Nav.jsx`, `frontend/src/styles/layout.css`  
Patch: Implement top nav, responsive drawer for mobile, and app shell that renders children.  
Tests: yes — unit test snapshots for nav render at different widths

13. Commit: `web/theme: define CSS variables and base typography`  
Files: `frontend/src/styles/variables.css`, `frontend/src/index.css`  
Patch: Add color tokens, font stacks (non-default choice), and base typographic scale.  
Tests: no — visual check

14. Commit: `web/home: add landing page with hero`  
Files: `frontend/src/pages/Home.jsx`, `frontend/src/styles/home.css`  
Patch: Create hero section, brief product value props, CTA buttons wired to routes.  
Tests: yes — render test asserting CTA presence

15. Commit: `web/scan-ui: add scanner form and results list`  
Files: `frontend/src/pages/Scan.jsx`, `frontend/src/components/ScanForm.jsx`, `frontend/src/components/ResultsList.jsx`  
Patch: Add form to select path/format, async submit to backend, and results list component showing findings.  
Tests: yes — component test for form emits correct request payload

16. Commit: `web/state: add lightweight client store`  
Files: `frontend/src/store/store.js`, `frontend/src/store/hooks.js`  
Patch: Implement small context-based store for UI state (current scan, results, user session).  
Tests: yes — unit tests for store reducers and hooks

17. Commit: `web/auth: add login flow and token handling`  
Files: `frontend/src/pages/Login.jsx`, `frontend/src/api/auth.js`  
Patch: Add login form, perform POST to `/api/login`, store token in memory/localStorage via auth helper.  
Tests: yes — mock API test ensuring token stored after login

18. Commit: `web/accessibility: add ARIA roles and keyboard nav`  
Files: `frontend/src/components/*` (multiple), `frontend/src/styles/accessibility.css`  
Patch: Add ARIA attributes to interactive components, improve focus styles and keyboard shortcuts for actions.  
Tests: no — accessibility audit in CI later

19. Commit: `web/perf: implement lazy loading and code splitting`  
Files: `frontend/src/routes.js`, `frontend/src/pages/Scan.jsx`  
Patch: Lazy-load heavy route components with React.lazy and Suspense; update route config.  
Tests: no — performance measurement in CI later

20. Commit: `web/icons: add SVG icon system and icon set`  
Files: `frontend/src/icons/index.js`, `frontend/src/icons/*.svg`  
Patch: Add icons helper to inline SVGs and an initial set (logo, search, close).  
Tests: no — visual verification

21. Commit: `web/i18n: add i18n scaffolding`  
Files: `frontend/src/i18n/index.js`, `frontend/src/locales/en.json`  
Patch: Integrate i18n library, add English locale and translate strings used on landing and scan form.  
Tests: yes — ensure components render translated strings when locale set

22. Commit: `web/error: add global error boundary and toast system`  
Files: `frontend/src/components/ErrorBoundary.jsx`, `frontend/src/components/Toast.jsx`  
Patch: Add error boundary around app and toast notification system for async errors.  
Tests: yes — component test for ErrorBoundary capturing child error

23. Commit: `web/styles: centralize variables and utility classes`  
Files: `frontend/src/styles/utils.css`, `frontend/src/styles/variables.css`  
Patch: Add utility classes (.sr-only, spacing helpers) and import them into index.css.  
Tests: no — style-only change

24. Commit: `web/e2e: add Playwright scaffolding`  
Files: `frontend/e2e/playwright.config.js`, `frontend/e2e/tests/home.spec.js`  
Patch: Add Playwright config and a smoke test that loads the home page and checks the hero CTA.  
Tests: yes — E2E smoke test run in CI

25. Commit: `web/docs: document frontend dev and build commands`  
Files: `frontend/README.md`  
Patch: Add instructions for `npm install`, `npm run dev`, `npm run build`, and lint/test commands.  
Tests: no

Backend (26..45)

26. Commit: `api/init: add server bootstrap and routing`  
Files: `backend/server.go`, `backend/router.go`, `backend/go.mod`  
Patch: Add HTTP server entrypoint, router setup, and graceful shutdown logic.  
Tests: yes — unit test for health endpoint

27. Commit: `api/health: add /health and readiness endpoints`  
Files: `backend/handlers/health.go`, `backend/router.go`  
Patch: Implement `/health` and `/ready` with JSON status and uptime metrics.  
Tests: yes — assert 200 and JSON schema

28. Commit: `api/auth: add JWT auth middleware`  
Files: `backend/middleware/auth.go`, `backend/handlers/login.go`  
Patch: Implement login handler issuing JWT, middleware to protect endpoints, and token parsing.  
Tests: yes — unit tests for token issuance and middleware rejecting invalid token

29. Commit: `api/scan: add POST /api/scan endpoint`  
Files: `backend/handlers/scan.go`, `backend/services/scan_service.go`  
Patch: Add API endpoint accepting scan parameters, queue scan job, and return job ID.  
Tests: yes — unit test for request validation and job creation

30. Commit: `api/jobs: add background job runner and store`  
Files: `backend/jobs/runner.go`, `backend/jobs/store.go`  
Patch: Implement in-process job runner with concurrency limit and persistent job state store (file/sql stub).  
Tests: yes — test job runner executes jobs and updates status

31. Commit: `api/results: add GET /api/scan/:id/results`  
Files: `backend/handlers/results.go`, `backend/services/result_store.go`  
Patch: Implement results retrieval endpoint with pagination and format param.  
Tests: yes — test pagination and format switching

32. Commit: `api/storage: add artifacts storage abstraction`  
Files: `backend/storage/storage.go`, `backend/storage/local.go`  
Patch: Add storage interface and local filesystem implementation for storing scan artifacts.  
Tests: yes — file write/read roundtrip tests

33. Commit: `api/metrics: integrate Prometheus metrics`  
Files: `backend/metrics/metrics.go`, `backend/server.go`  
Patch: Expose `/metrics`, add counters for scans, errors, and job durations.  
Tests: no — metrics smoke verified in integration

34. Commit: `api/config: centralize backend config and secrets`  
Files: `backend/config.go`, `backend/env.example`  
Patch: Add config loader for DB, storage, job concurrency, and JWT secret; example env file.  
Tests: yes — config parsing and validation tests

35. Commit: `api/logging: structured server logs`  
Files: `backend/log.go`, `backend/server.go`  
Patch: Replace fmt logs with structured logger and include request IDs.  
Tests: yes — ensure request logs include request ID header

36. Commit: `api/db: add migrations and simple store layer`  
Files: `backend/db/migrations/0001_init.sql`, `backend/db/store.go`  
Patch: Add initial SQL migration for jobs/results, and store helpers to read/write jobs.  
Tests: yes — in-memory DB migration and store tests

37. Commit: `api/trace: add basic tracing hooks`  
Files: `backend/trace/trace.go`, `backend/server.go`  
Patch: Add OpenTelemetry skeleton, create spans for requests and jobs.  
Tests: no — observability integration later

38. Commit: `api/cleanup: add retention worker for old artifacts`  
Files: `backend/jobs/cleanup.go`, `backend/config.go`  
Patch: Implement scheduled cleanup job removing artifacts older than configured retention.  
Tests: yes — test cleanup removes expected files by age

39. Commit: `api/health-checks: add DB and storage checks to readiness`  
Files: `backend/handlers/health.go`, `backend/db/store.go`  
Patch: Add deeper readiness checks calling DB ping and storage list.  
Tests: yes — readiness fails when underlying checks fail (mocked)

40. Commit: `api/docs: add OpenAPI spec skeleton`  
Files: `backend/api/openapi.yaml`, `backend/README.md`  
Patch: Add initial OpenAPI file documenting `/api/scan`, `/api/scan/{id}/results`, `/login`, and `/health`.  
Tests: no — spec review

TODO Scanner (46..55)

41. Commit: `scanner/init: add scanner package skeleton`  
Files: `scanner/scanner.go`, `scanner/README.md`  
Patch: Create scanner interface `Scan(path, opts) -> Findings` and README describing signature.  
Tests: no

42. Commit: `scanner/parsers: add file walker and basic parsers`  
Files: `scanner/walker.go`, `scanner/parsers/text_parser.go`  
Patch: Implement recursive file walker with ignore patterns and a simple text parser to detect TODO/FIXME.  
Tests: yes — walker unit tests and parser detection tests

43. Commit: `scanner/rules: add regex rules config`  
Files: `scanner/rules/rules.go`, `scanner/rules/default.json`  
Patch: Implement rules loader for regex-based detectors and add default rules for TODO/FIXME/NOTE.  
Tests: yes — rule matching tests with example snippets

44. Commit: `scanner/output: produce structured findings`  
Files: `scanner/output.go`, `scanner/models.go`  
Patch: Define `Finding` model with file, line, tag, message; implement JSON output writer.  
Tests: yes — serialization roundtrip tests

45. Commit: `scanner/ignore: support .scannerignore`  
Files: `scanner/ignore.go`, `.scannerignore.example`  
Patch: Implement ignore file parsing and integrate into file walker, add example ignore file.  
Tests: yes — ensure ignored paths are excluded

46. Commit: `scanner/threading: add concurrent scanning`  
Files: `scanner/worker_pool.go`, `scanner/walker.go`  
Patch: Add worker pool to parse files concurrently with configurable concurrency limit.  
Tests: yes — concurrency test ensures no data races (run with -race in CI)

47. Commit: `scanner/formatters: add SARIF and plain text outputs`  
Files: `scanner/format/sarif.go`, `scanner/format/plain.go`  
Patch: Implement SARIF output for findings and a compact plain text formatter.  
Tests: yes — validate SARIF JSON schema and sample plain output

48. Commit: `scanner/integration: wire scanner into backend job runner`  
Files: `backend/services/scan_service.go`, `scanner/scanner.go`  
Patch: Hook the scanner interface into the job runner so scan jobs execute using scanner package.  
Tests: yes — integration test that job runner invokes scanner and stores results (mock scanner)

49. Commit: `scanner/alerts: annotate findings with severity`  
Files: `scanner/models.go`, `scanner/rules/rules.go`  
Patch: Add severity levels to rules and propagate severity to findings for downstream filtering.  
Tests: yes — rule-to-severity mapping tests

50. Commit: `scanner/perf: add incremental scan state`  
Files: `scanner/state/store.go`, `scanner/state/hash.go`  
Patch: Implement simple state store to skip unchanged files using content hashing; expose opt-in flag.  
Tests: yes — state persistence and skip logic tests

CI (56..65)

51. Commit: `ci/init: add GitHub Actions workflow skeleton`  
Files: `.github/workflows/ci.yml`  
Patch: Add CI workflow to run `build`, `lint`, and `unit` jobs matrix for `linux` and `node` steps.  
Tests: no — this defines CI pipeline

52. Commit: `ci/lint: add linters and formatting checks`  
Files: `.github/workflows/lint.yml`, `frontend/.eslintrc.js`, `backend/.golangci.yml`  
Patch: Add separate lint workflow that runs ESLint and golangci-lint and fails on issues.  
Tests: no

53. Commit: `ci/tests: run unit tests and coverage upload`  
Files: `.github/workflows/tests.yml`, `ci/upload-coverage.sh`  
Patch: Add test workflow for frontend/backend unit tests and upload coverage artifact to actions.  
Tests: no — workflow itself runs tests

54. Commit: `ci/e2e: add Playwright E2E job`  
Files: `.github/workflows/e2e.yml`, `frontend/e2e/action-setup.sh`  
Patch: Configure E2E job to build app, start server, and run Playwright tests.  
Tests: no

55. Commit: `ci/security: add Dependabot config and SCA`  
Files: `.github/dependabot.yml`, `.github/workflows/snyk.yml`  
Patch: Enable Dependabot for npm/go and add security scanning job placeholder.  
Tests: no

56. Commit: `ci/cache: add caching for dependencies`  
Files: `.github/workflows/ci.yml`  
Patch: Update CI to cache `~/.cache` for npm and Go modules between runs.  
Tests: no

57. Commit: `ci/docs: add job to build and publish docs artifact`  
Files: `.github/workflows/docs.yml`, `ci/publish-docs.sh`  
Patch: Add workflow to build docs (mkdocs or similar) and publish as artifact on CI runs.  
Tests: no

58. Commit: `ci/containers: add Docker build and scan steps`  
Files: `.github/workflows/containers.yml`, `docker/Dockerfile`  
Patch: Add container build job and placeholder security scan (trivy/snyk).  
Tests: no

59. Commit: `ci/notifications: add Slack/GitHub status integration`  
Files: `.github/workflows/ci.yml`  
Patch: Add job-level notifications using Actions step that can call webhook on failure/success (config via secrets).  
Tests: no

60. Commit: `ci/branch-protection: add PR checks requirement`  
Files: `docs/PR_PROCESS.md`, `.github/workflows/ci.yml`  
Patch: Document required checks and add status checks in workflow for PR gating.  
Tests: no

Tests (61..75)

61. Commit: `tests/unit: add backend unit test harness`  
Files: `backend/tests/unit/README.md`, `backend/tests/unit/example_test.go`  
Patch: Add Go test helpers and an example unit test demonstrating mocking DB and HTTP handlers.  
Tests: yes — running `go test ./...` should include example_test

62. Commit: `tests/unit: add frontend Jest setup and sample test`  
Files: `frontend/jest.config.js`, `frontend/src/components/__tests__/Nav.test.jsx`  
Patch: Configure Jest + React Testing Library and add a sample nav render test.  
Tests: yes — jest runs sample test

63. Commit: `tests/integration: add API integration harness`  
Files: `backend/tests/integration/setup_test.go`, `backend/tests/integration/scan_flow_test.go`  
Patch: Add test harness to spin up in-memory server and DB, and integration test that runs end-to-end scan flow.  
Tests: yes — integration test included

64. Commit: `tests/e2e: add Playwright CI scripts`  
Files: `frontend/e2e/playwright.config.js`, `frontend/e2e/tests/scan.spec.js`  
Patch: Add e2e tests that exercise scanning UI and results rendering; integrate with CI workflow earlier.  
Tests: yes — E2E test run expected in CI

65. Commit: `tests/perf: add basic perf harness for scanning`  
Files: `scanner/tests/perf/bench.go`, `scanner/tests/data/large_repo_sample`  
Patch: Add micro-benchmark to measure scan speed on a sample dataset; record baseline.  
Tests: no — benchmark not run in unit CI

66. Commit: `tests/security: add auth and token tests`  
Files: `backend/tests/unit/auth_test.go`  
Patch: Add tests verifying JWT expiry, refresh, and middleware rejection.  
Tests: yes

67. Commit: `tests/db: add transactional tests and cleanup helpers`  
Files: `backend/tests/helpers/db_helper.go`, `backend/tests/unit/store_test.go`  
Patch: Implement DB helper to run tests in transactions and rollback to keep DB clean.  
Tests: yes

68. Commit: `tests/ci: add smoke test step`  
Files: `.github/workflows/ci.yml`  
Patch: Insert a `ci:smoke` job that verifies server boots and basic endpoints respond before other jobs.  
Tests: no

69. Commit: `tests/race: enable go race detector in CI`  
Files: `.github/workflows/tests.yml`  
Patch: Add `-race` flag to `go test` in CI for unit/integration suites where applicable.  
Tests: no

70. Commit: `tests/coverage: unify coverage reporting`  
Files: `ci/upload-coverage.sh`, `frontend/package.json`  
Patch: Standardize coverage threshold and upload process for frontend and backend.  
Tests: no

Docs (76..85)

71. Commit: `docs/overview: add project overview and architecture`  
Files: `docs/ARCHITECTURE.md`, `README.md`  
Patch: Create architecture diagram notes, flow of scan job, and update root README with links.  
Tests: no

72. Commit: `docs/development: add local dev setup`  
Files: `docs/DEVELOPMENT.md`  
Patch: Document how to run frontend, backend, scanner locally, and how to run tests and linters.  
Tests: no

73. Commit: `docs/cli: document CLI commands and examples`  
Files: `docs/CLI.md`, `cli/README.md`  
Patch: Add examples for `scan`, `init`, and config formats; include sample command outputs.  
Tests: no

74. Commit: `docs/api: document REST API and examples`  
Files: `docs/API.md`, `backend/api/openapi.yaml`  
Patch: Expand OpenAPI documentation with request/response examples and curl snippets.  
Tests: no

75. Commit: `docs/security: add auth and secrets handling`  
Files: `docs/SECURITY.md`  
Patch: Document how secrets are stored, rotated, and how to configure JWT secret and CI secrets.  
Tests: no

76. Commit: `docs/contributing: add PR and code style guide`  
Files: `CONTRIBUTING.md`  
Patch: Add PR process, branch naming, commit message conventions, and code style rules.  
Tests: no

77. Commit: `docs/releases: add changelog and release process`  
Files: `docs/RELEASES.md`, `CHANGELOG.md`  
Patch: Add release numbering strategy, changelog template, and GitHub release steps.  
Tests: no

78. Commit: `docs/ops: runbook for production incidents`  
Files: `docs/OPERATIONS.md`  
Patch: Add runbook steps for incident response, restoring jobs, and common recovery commands.  
Tests: no

79. Commit: `docs/benchmarks: document scanner performance baselines`  
Files: `docs/BENCHMARKS.md`, `scanner/tests/perf/README.md`  
Patch: Record baseline scan times and instructions to run benchmarks.  
Tests: no

80. Commit: `docs/examples: include common configs and .scannerignore`  
Files: `docs/EXAMPLE_CONFIGS.md`, `.scannerignore.example`  
Patch: Add curated example configs for CI, local dev, and scanner ignore files.  
Tests: no

Examples (86..95)

81. Commit: `examples/basic: add simple repo example`  
Files: `examples/basic/README.md`, `examples/basic/src/main.py`  
Patch: Add small repo with few files and some TODO comments for scanner to detect.  
Tests: yes — scanner test against example repo

82. Commit: `examples/monorepo: add multi-package example`  
Files: `examples/monorepo/packages/pkg-a/index.js`, `examples/monorepo/packages/pkg-b/index.js`  
Patch: Create monorepo layout including nested TODOs and a `package.json` at root.  
Tests: yes — scanner verifies traversal across packages

83. Commit: `examples/large: add larger synthetic repo sample`  
Files: `examples/large/README.md`, `examples/large/*`  
Patch: Add scripted generator or sample files to simulate a larger codebase for perf testing.  
Tests: no — used by perf harness

84. Commit: `examples/docker: add example Dockerized server`  
Files: `examples/docker/Dockerfile`, `examples/docker/README.md`  
Patch: Provide Dockerfile to run backend API and instructions to mount artifacts and run scanner.  
Tests: no

85. Commit: `examples/ci-pipeline: add sample GitHub Actions`  
Files: `examples/ci-pipeline/.github/workflows/sample.yml`  
Patch: Add sample workflow demonstrating how to call scanner in CI and upload findings as artifact.  
Tests: no

86. Commit: `examples/plugin: add sample plugin implementing interface`  
Files: `examples/plugin/plugin-example/README.md`, `examples/plugin/plugin-example/main.go`  
Patch: Provide minimal plugin demonstrating plugin contract and sample output.  
Tests: yes — plugin unit test to exercise contract

87. Commit: `examples/frontend-integration: add example connecting UI to API`  
Files: `examples/frontend-integration/README.md`, `examples/frontend-integration/proxy.conf`  
Patch: Document and provide sample proxy config to run frontend against local backend for demos.  
Tests: no

88. Commit: `examples/ignore-sets: add sample .scannerignore configs`  
Files: `examples/ignore-sets/.scannerignore-node`, `examples/ignore-sets/.scannerignore-python`  
Patch: Curated ignore patterns per language ecosystem and README explaining rationale.  
Tests: yes — run scanner to confirm ignores applied

89. Commit: `examples/reporting: add sample SARIF and JSON outputs`  
Files: `examples/reporting/sample.sarif.json`, `examples/reporting/sample.json`  
Patch: Add sample outputs for reference and test verification.  
Tests: yes — validator ensures SARIF file conforms to schema

90. Commit: `examples/workflow: add complete demo with steps`  
Files: `examples/workflow/README.md`, `examples/workflow/run.sh`  
Patch: Provide complete demo script that runs scanner, serves API, and opens frontend for manual demo.  
Tests: no

Final Commits / Wrap-up (96..100)

91. Commit: `chore/license: add project license file`  
Files: `LICENSE`  
Patch: Add MIT license file and update README with license badge.  
Tests: no

92. Commit: `chore/gitignore: add common ignores and IDE files`  
Files: `.gitignore`  
Patch: Add standard entries for node_modules, build artifacts, .vscode, .idea, and OS files.  
Tests: no

93. Commit: `chore/ci-config: canonicalize workflows and badges`  
Files: `README.md`, `.github/workflows/ci.yml`  
Patch: Update README badges and ensure CI workflow names match documented checks.  
Tests: no

94. Commit: `chore/cleanup: remove placeholder files and tidy modules`  
Files: Various placeholders (list in commit), `go.mod`, `package.json`  
Patch: Remove deprecated placeholder files, run `go mod tidy` and clean npm package metadata.  
Tests: no

95. Commit: `chore/release: add release workflow and tagging script`  
Files: `.github/workflows/release.yml`, `scripts/release.sh`  
Patch: Add automated release workflow that tags main on merge and creates GitHub release with changelog.  
Tests: no

96. Commit: `ci/verify: add pre-commit hooks and local checks`  
Files: `.husky/pre-commit`, `.pre-commit-config.yaml`  
Patch: Add pre-commit hooks for formatting, lint, and running quick unit tests locally.  
Tests: no

97. Commit: `meta/license-check: add license checker to CI`  
Files: `.github/workflows/license.yml`, `ci/license-check.sh`  
Patch: Add job to verify third-party license compliance for dependencies.  
Tests: no

98. Commit: `meta/branding: add logo and brand assets`  
Files: `assets/logo.svg`, `frontend/src/assets/logo.svg`, `README.md`  
Patch: Add logo SVG and reference it in frontend and README.  
Tests: no

99. Commit: `meta/version: set initial semantic version`  
Files: `VERSION`  
Patch: Add `0.1.0` as default project version and document version bump process in RELEASES.md.  
Tests: no

100. Commit: `meta/checks: add final checklist and CI gating`  
Files: `docs/RELEASES.md`, `docs/CHECKLIST.md`  
Patch: Add final release checklist to verify tests, lint, docs, examples, and CI passing before cut.  
Tests: no

---

File saved to: `.opencode/plans/plan-bulk-100.md`

Plan notes: commits are intentionally small and focused so an automated build or apply tool can implement them sequentially. Many edits are scaffolding and will require follow-up implementation details; tests flagged YES should be added to CI to validate each change.
