#!/usr/bin/env node
import { program } from 'commander';
import { summarize } from './git';
import pkg from '../package.json';
import chalk from 'chalk';

program
  .name('git-digest')
  .option('--json', 'output results as JSON')
  .option('--format <type>', 'output format: text|json', 'text')
  .description(
    chalk.cyanBright('Summarize recent git commits, extract TODOs, and open relevant PRs.') +
      '\n\nGit-Digest helps teams track work and review outstanding TODOs directly from their commit history.' +
      ' It scans for TODOs, links relevant PRs, and can open them in your browser.'
  )
  .version(pkg.version)
  .usage('[options]')
  .option('-n, --count <number>', 'number of recent commits to summarize', (v) => {
    const parsed = parseInt(v, 10);
    if (isNaN(parsed) || parsed < 1 || parsed > 1000) {
      throw new Error('Count must be a number between 1 and 1000');
    }
    return parsed;
  }, 20)
  .option('--open-pr', 'open the most relevant PR in the default browser')
  .option('-v, --verbose', 'enable verbose logging')
  .addHelpText('afterAll', '\n' + chalk.green('Examples:') +
    '\n  $ git-digest --count 10' +
    '\n  $ git-digest --open-pr' +
    '\n  $ git-digest -n 5 -v')
  .showHelpAfterError('(add --help for additional information)')
  .showSuggestionAfterError(true)
  .action(async (opts) => {
    try {
      const count = opts.count;
      const result = await summarize({ count });

      // Determine output format
      const outputFormat = opts.json ? 'json' : opts.format;
      if (outputFormat === 'json') {
        console.log(JSON.stringify(result, null, 2));
      } else {
        // Colorized output
        console.log(chalk.cyanBright(result.summary));
      }

      if (opts.openPr) {
        if (result.prUrl) {
          const open = await import('open');
          await open.default(result.prUrl);
           if (opts.verbose) console.log(chalk.gray('Opened PR URL in default browser.'));
        } else if (outputFormat !== 'json') {
           console.log(chalk.yellow('No PR URL found to open.'));
        } else {
          console.log(chalk.yellow('No PR URL found to open.'));
          process.exitCode = 2;
        }
      }
    } catch (err: any) {
if (opts.verbose) {
         console.error(chalk.redBright('Error:'), err && err.stack ? err.stack : err);
       } else {
         const outputFormat = opts.json ? 'json' : opts.format;
         if (opts.json || outputFormat === 'json') {
           console.error(JSON.stringify({ error: err && err.message ? err.message : err }, null, 2));
         } else {
           console.error(chalk.redBright('Fatal:'), err && err.message ? err.message : err);
         }
       }
      process.exitCode = 1;
    }
  });

program.parse(process.argv);
