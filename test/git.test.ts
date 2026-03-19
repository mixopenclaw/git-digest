import { describe, it, expect } from 'vitest';
import { extractTodos, summarizeCommits } from '../src/git';

describe('git helpers', () => {
  it('extracts todos from commit lines', () => {
    const commits = ['a1b2c Fix bug', 'd3e4f TODO: refactor parsing', 'z9y8x feat: add tests TODO add more'];
    const todos = extractTodos(commits);
    expect(todos).toEqual(['refactor parsing', 'add more']);
  });

  it('summarizes commits', () => {
    const commits = ['a1b2c First', 'b2c3d Second'];
    const s = summarizeCommits(commits);
    expect(s).toContain('Recent commits:');
  });
});
