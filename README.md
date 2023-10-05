> **Warning**
> Work in progress..

# difii
A CLI tool to diff 2 directories.

## Usage
```console
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

### diff summary
`--summary` flag shows diff summary.
```console
$ difii --compare ../random-b --summary
----- Compare -----
I'll show you any additions or deletions in [.] when compared to [../random-b].

----- Summary -----
-5 +4 diffs in main.md

```

### diff detail
`--inspect` flag shows diff detail.
```console
$ difii --compare ../random-b --inspect
----- Compare -----
I'll show you any additions or deletions in [.] when compared to [../random-b].

----- Inspect -----
main.md:2	- KLMNOPQRST  // diff
main.md:5	- QRSTUVWXYZA
main.md:4	+ QRSTUVWXYY   // diff
main.md:7	- MNOPQRSTUVWXYZ
main.md:8	- ABCDEFGHIJ
main.md:6	+ MNOOPQRSTUVWXYZ  // diff
main.md:7	+ ABCDEFGHJI  // diff
main.md:10	- UVWXYZABCD
main.md:9	+ UVWXYZABCC  // diff

```
