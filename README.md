# git-digest

git-digest is a small developer CLI that summarizes recent git commits, extracts inline TODOs from commit messages, and can open the repository's pull requests page.

Install (local dev):

```bash
npm install
npm run build
node dist/index.js --count 10
```

Example:

```bash
git-digest --count 15 --open-pr
```
