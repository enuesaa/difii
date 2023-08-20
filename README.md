**Work in progress...**
# difii
A CLI tool to compare two directories.  

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

### Show diffs summary
```bash
$ difii --compare ./sample-app/styles --summary

Diffs Summary
 -4  +0 diffs in emotion.d.ts
-20  +5 diffs in global.ts
-47 +29 diffs in theme.ts
 -0  +0 diffs in use.ts
```
