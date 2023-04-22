# difii
A cli tool to import diffs interactively.  

## Command Interface
```bash
difii <from> <to> [--limit 1] [--only <filename>] [--no-interactive] [--overwrite]
```
### from
取り込み元のディレクトリ
### to
取り込み先のディレクトリ
### 標準出力
差分

## Example
### Normal
```bash
difii ../aaa . --no-interactive

README.md has +1 -1 diff.
+ aa
- bb
```
### Interactive
```bash
difii 

from: ../aaa
to: .
only: README.md
only: NOTEXISTS.md

README.md has +1 -1 diff.
+ aa
- bb

Do you overwrite ?
```
