import execa from 'execa';

export async function gitLog(count: number) {
  try {
    const { stdout } = await execa('git', ['log', `-n`, String(count), `--pretty=format:%h %s -- %an`]);
    return stdout.split('\n').filter(Boolean);
  } catch (e) {
    return [];
  }
}

export function extractTodos(commits: string[]) {
  const todos: string[] = [];
  for (const line of commits) {
    const m = line.match(/TODO[:]?\s*(.*)/i);
    if (m) todos.push(m[1].trim());
  }
  return todos;
}

export function summarizeCommits(commits: string[]) {
  if (commits.length === 0) return 'No commits found.';
  const lines = commits.slice(0, 5).map((l) => `- ${l}`);
  return `Recent commits:\n${lines.join('\n')}`;
}

export async function findPrUrl() {
  try {
    // Try to find PR URL from git remote or hub
    const { stdout } = await execa('git', ['config', '--get', 'remote.origin.url']);
    const url = stdout.trim();
    // naive transform from git@github.com:user/repo.git or https://github.com/user/repo.git to https://github.com/user/repo
    const repo = url.replace(/^git@/, 'https://').replace(':', '/').replace(/\.git$/, '');
    // open PRs page
    return `${repo}/pulls`;
  } catch (e) {
    return null;
  }
}

export async function summarize({ count }: { count: number }) {
  const commits = await gitLog(count);
  const todos = extractTodos(commits);
  const summary = `${summarizeCommits(commits)}\n\nFound ${todos.length} TODO(s).\n${todos.slice(0,5).map((t,i)=>`${i+1}. ${t}`).join('\n')}`;
  const prUrl = await findPrUrl();
  return { summary, prUrl };
}
