**Work in progress...**
# difii
A CLI tool to diff 2 directories.

## Usage
```bash
$ difii --help
Usage:
  difii [flags]

Flags:
      --inspect   Inspect diffs.
      --summary   Show diffs summary.

Global Flags:
      --compare string   Compare dir.
      --workdir string   Working dir. Default value is current dir.
      --only strings     Filename to compare
  -i, --interactive      Start interactive prompt.
      --help             help
      --version          version
```

### `--summary` flag shows diffs summary.
```bash
$ difii --compare ../another-styles --summary

Comparing..

- .
- ../another-styles

I'll show you any additions or deletions in [.] when compared to [../another-styles].

-----------

Summary

-4 +0 diffs in emotion.d.ts
-2 +5 diffs in global.ts
-4 +2 diffs in theme.ts
-0 +0 diffs in use.ts
```
