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

Global Flags:
      --workdir string   Working dir. Default value is current dir.
      --only strings     Filename to compare
  -i, --interactive      Start interactive prompt.
      --help             help
      --version          version
```

### summary
```console
$ difii ./testdata/random-b --workdir ./testdata/random-a
I'll show you any additions or deletions of [./testdata/random-a] when compared to [./testdata/random-b].

----- Summary -----
-4 +5 diffs in main.md

```

### diff detail
`--inspect` flag shows diff detail.
```console
$ difii ./testdata/random-b --inspect --workdir ./testdata/random-a
I'll show you any additions or deletions of [./testdata/random-a] when compared to [./testdata/random-b].

----- Summary -----
-4 +5 diffs in main.md

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
