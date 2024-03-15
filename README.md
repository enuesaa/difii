# difii
A CLI tool to diff 2 folders interactively.

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
A CLI tool to diff 2 folders interactively.

Usage:
  difii <dir1> <dir2> [flags]

Flags:
      --help           Show help messages.
      --inspect        Inspect diffs.
  -i, --interactive    Use interactive prompt.
      --only strings   Specify filename to compare.
      --version        Show version information.
```

### Commands
```console
$ difii <dir-a> <dir-b> # this prints diffs per file
$ difii <dir-a> <dir-b> --inspect # this prints diffs detail
```

### Print diffs per file
```console
$ difii ./testdata/random-a ./testdata/random-b
-5 +4 diffs in main.md
```

### Print diffs detail
`--inspect` flag shows diff detail.
```console
$ difii ./testdata/random-a ./testdata/random-b --inspect
main.md:2   - KLMNOPQRST
main.md:5   - QRSTUVWXYZA
main.md:4   + QRSTUVWXYY
main.md:7   - MNOPQRSTUVWXYZ
main.md:8   - ABCDEFGHIJ
main.md:6   + MNOOPQRSTUVWXYZ
main.md:7   + ABCDEFGHJI
main.md:10  - UVWXYZABCD
main.md:9   + UVWXYZABCC
```
