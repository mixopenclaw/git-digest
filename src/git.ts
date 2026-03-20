import { execaCommand } from 'execa';

export async function gitLog(count: number) {
  try {
    const { stdout } = await execaCommand(`git log -n ${count} --pretty=format:%h %s -- %an`);
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
    const { stdout } = await execaCommand('git config --get remote.origin.url');
    const url = stdout.trim();
    const repo = url.replace(/^git@/, 'https://').replace(':', '/').replace(/\.git$/, '');
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

export function scanTextForTodos(text: string){
  return text.split('\n').filter(l=>l.includes('TODO'))
}
