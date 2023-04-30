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


## Memo
git diff のように行頭に+-が表示されるのがいいと思う.   
ハンクごとにdiffが表示するか.     
加えて Github みたいに違う文字列が太字になっている方が良いと思う.   
ベストは取り込んだ後に編集がいらないこと. 
