#!/usr/bin/env node
import { program } from 'commander';
import { summarize } from './git';
import pkg from '../package.json';

program
  .name('git-digest')
  .description('Summarize recent git commits, extract TODOs, and open relevant PRs.\n\nGit-Digest helps teams track work and review outstanding TODOs directly from their commit history. It scans for TODOs, links relevant PRs, and can open them in your browser.')
  .version(pkg.version)
  .usage('[options]')
  .option('-n, --count <number>', 'number of recent commits to summarize', '20')
  .option('--open-pr', 'open the most relevant PR in the default browser')
  .addHelpText('after', '\n\nExample:\n  $ git-digest --count 10\n  $ git-digest --open-pr')
  .showHelpAfterError('(add --help for additional information)')
  .showSuggestionAfterError(true)
  .action(async (opts) => {
    const count = parseInt(opts.count, 10) || 20;
    const result = await summarize({ count });
    console.log(result.summary);
    if (opts.openPr || opts.openPR || opts.open_pr) {
      if (result.prUrl) {
        const open = await import('open');
        await open.default(result.prUrl);
      } else {
        console.log('No PR URL found to open.');
      }
    }
  });

program.parse(process.argv);
