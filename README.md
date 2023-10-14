# difii
A CLI tool to diff 2 directories.

> **Note**
> Work in progress.. `difii` is currently under development.

## Install
```bash
git clone https://github.com/enuesaa/difii.git --depth 1
cd difii
go install
```

## Usage
```console
$ difii --help
Usage:
  difii <compare-dir> [flags]

Flags:
      --inspect   Inspect diffs.
      --summary   Show diffs summary.

Global Flags:
      --workdir string   Working dir. Default value is current dir.
      --only strings     Filename to compare
  -i, --interactive      Start interactive prompt.
      --help             help
      --version          version
```

### diff summary
`--summary` flag shows diff summary.
```console
$ difii ./testdata/random-b --summary --workdir ./testdata/random-a
----- Compare -----
I'll show you any additions or deletions of [./testdata/random-a] when compared to [./testdata/random-b].

----- Summary -----
-4 +5 diffs in main.md

```

### diff detail
`--inspect` flag shows diff detail.
```console
$ difii ./testdata/random-b --inspect --workdir ./testdata/random-a
----- Compare -----
I'll show you any additions or deletions of [./testdata/random-a] when compared to [./testdata/random-b].

----- Inspect -----
main.md:2  + KLMNOPQRST
main.md:4  - QRSTUVWXYY
main.md:5  + QRSTUVWXYZA
main.md:6  - MNOOPQRSTUVWXYZ
main.md:7  - ABCDEFGHJI
main.md:7  + MNOPQRSTUVWXYZ
main.md:8  + ABCDEFGHIJ
main.md:9  - UVWXYZABCC
main.md:10 + UVWXYZABCD

```
