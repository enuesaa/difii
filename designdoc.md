# designdoc

## Usage
```bash
difii 
    --source <source-dir>
    --dest <dest-dir> \
    --only <filename> \
    --include-git
```
差分を出力するのみ。

### `--source <source-dir>`
取り込み元のディレクトリ  
### `--dest <dest-dir>`
取り込み先のディレクトリ. Current dir is default.
### `--only <filename>`
diff対象のファイルを指定する. 複数指定可能
### `--include-git`
By default, .git directory is ignored. If you pass this option, you can also diff git directory.

### 標準出力
左に source, 右にdestを表示する
```bash
<filename> has +2 -1 diffs.
+ aaa
+ bbb
- ccc

Do you overwrite ? [Y/n] 
```

### 標準出力 (Summary)
```bash
## Summary
|name|diff|
|<filename>|+2 -1|
|<dirname>|<dirname> is directory. skipped.|
```

## Usage of Plan command
```bash
difii plan
    --source <source-dir>
    --dest <dest-dir> \
    --only <filename> \
    --include-git
```
import処理の実行計画を表示する

## Usage of Apply command
```bash
difii apply
    --source <source-dir>
    --dest <dest-dir> \
    --only <filename> \
    --include-git \
    --auto-approve
```
import処理をする
