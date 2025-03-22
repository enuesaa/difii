# difii
A CLI tool to diff 2 folders interactively.

## Install
```bash
go install github.com/enuesaa/difii@v0.0.14
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

## Commands
```bash
# By default, `difii` counts the number of changed lines for each file. 
difii <dir-a> <dir-b>

# Use the `--inspect` flag to look up more details.
difii <dir-a> <dir-b> --inspect
```

### Example
```console
$ difii ./testdata/random-a ./testdata/random-b
-5 +4 diffs in main.md

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
