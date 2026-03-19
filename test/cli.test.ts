import { describe, it, expect } from 'vitest';
import { execa } from 'execa';
import path from 'path';

const CLI_PATH = path.resolve(__dirname, '../dist/src/index.js');

describe('CLI --json output', () => {
  it('outputs valid JSON and includes keys', async () => {
    const { stdout } = await execa('node', [CLI_PATH, '--json', '--count', '2']);
    const output = JSON.parse(stdout);
    expect(output).toHaveProperty('summary');
    expect(output).toHaveProperty('prUrl');
    expect(typeof output.summary).toBe('string');
  });
  it('outputs valid JSON with --format json', async () => {
    const { stdout } = await execa('node', [CLI_PATH, '--format', 'json', '--count', '2']);
    const output = JSON.parse(stdout);
    expect(output).toHaveProperty('summary');
    expect(output).toHaveProperty('prUrl');
  });
});
