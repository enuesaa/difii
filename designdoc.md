# designdoc

## Usage
```bash
difii 
    --source <source-dir>
    --dest <dest-dir> \
    --only <filename> \
    --include-git \
    --no-interactive \
    --overwrite
```

### `--source <source-dir>`
取り込み元のディレクトリ  
### `--dest <dest-dir>`
取り込み先のディレクトリ
### `--only <filename>`
diff対象のファイルを指定する. 複数指定可能
### `--include-git`
By default, .git directory is ignored. If you pass this option, you can also diff git directory.
### `--no-interactive`
Disable interactive prompt.
### `--overwrite`
Overwrite all files. This behavior is same as cp command.

### 標準出力
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
