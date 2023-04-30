## Command Interface
```bash
difii <fromfile> <tofile> [--no-interactive] [--overwrite]
```
### fromfile
取り込み元のファイル
### tofile
取り込み先のファイル
### 標準出力
差分

## Example
### Normal
```bash
difii ../aaa.txt bb.txt --no-interactive
+ aa
bb
- cc
```
### Interactive
```bash
difii 

from: ../aaa.txt
to: bb.txt

+ aa
Do you overwrite ?

bb
- cc
Do you overwrite ?
```

## Memo
git diff のように行頭に+-が表示されるのがいいと思う.   
ハンクごとにdiffが表示するか.     
加えて Github みたいに違う文字列が太字になっている方が良いと思う.   
ベストは取り込んだ後に編集がいらないこと. 
