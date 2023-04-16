# Refactor
差分を取り込むツールにしたい  

## Command Interface
```bash
difii <from> <to> [--no-interactive]
```
### from
存在するファイルもしくはディレクトリ
### to
存在しなくてもよい. 存在しない場合にはファイルもしくはディレクトリを作成する
### 標準出力
差分

## Example
### Normal
```bash
difii ../../aaa/README.md ./README.md --no-interactive
+ aa
- bb
```
### Interactive
```bash
difii 

from: ../../aaa/README.md
to: ./README.md

Do you overwrite ?
+ aa
- bb
```

## TODO
- cp との違いはなんなのか