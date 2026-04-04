# Git Digest — Frontend

React + Vite single-page application for Git Digest.

## Prerequisites

- Node.js >= 18
- npm >= 9

## Getting Started

```bash
# Install dependencies
npm install

# Start the dev server (http://localhost:5173)
npm run dev

# Build for production
npm run build

# Preview the production build locally
npx vite preview
```

## Available Scripts

| Script          | Description                              |
| --------------- | ---------------------------------------- |
| `npm run dev`   | Start Vite dev server with HMR           |
| `npm run build` | Create optimised production build in `dist/` |
| `npm test`      | Run unit tests (Jest / React Testing Library) |

## Project Structure

```
frontend/
├── e2e/                  # Playwright end-to-end tests
│   ├── playwright.config.js
│   └── tests/
├── src/
│   ├── api/              # API helpers (auth, fetch wrappers)
│   ├── components/       # Reusable UI components
│   ├── i18n/             # Internationalisation scaffolding
│   ├── icons/            # SVG icon system
│   ├── locales/          # Translation files (en.json, …)
│   ├── pages/            # Route-level page components
│   ├── store/            # Lightweight client-side state
│   ├── styles/           # CSS variables, utilities, layout
│   ├── App.jsx           # Root app component
│   ├── main.jsx          # Vite entry point
│   └── routes.js         # Route definitions (lazy-loaded)
├── index.html            # HTML shell
├── vite.config.ts        # Vite configuration
└── package.json
```

## Linting & Formatting

```bash
# Lint (if ESLint is configured at repo root)
npx eslint src/

# Format with Prettier (if installed)
npx prettier --write "src/**/*.{js,jsx,css,json}"
```

## End-to-End Tests

```bash
# Install Playwright browsers (first time)
npx playwright install --with-deps chromium

# Run E2E tests
npx playwright test --config e2e/playwright.config.js
```
