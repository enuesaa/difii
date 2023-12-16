# difii
A CLI tool to diff 2 folders.

> [!Note]
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
A CLI tool to inspect diffs interactively.

Usage:
  difii <dir1> <dir2> [flags]

Flags:
      --inspect        Inspect diffs.
  -i, --interactive    Use interactive prompt.
      --only strings   Specify filename to compare.
```

### summary
```console
$ difii ./testdata/random-a ./testdata/random-b
----- Summary -----
Any additions or deletions to [./testdata/random-a] are shown.

-5 +4 diffs in main.md

```

### diff detail
`--inspect` flag shows diff detail.
```console
$ difii ./testdata/random-a ./testdata/random-b --inspect
----- Summary -----
Any additions or deletions to [./testdata/random-a] are shown.

-5 +4 diffs in main.md

----- Inspect -----
main.md:2  - KLMNOPQRST
main.md:5  - QRSTUVWXYZA
main.md:4  + QRSTUVWXYY
main.md:7  - MNOPQRSTUVWXYZ
main.md:8  - ABCDEFGHIJ
main.md:6  + MNOOPQRSTUVWXYZ
main.md:7  + ABCDEFGHJI
main.md:10 - UVWXYZABCD
main.md:9  + UVWXYZABCC

```
