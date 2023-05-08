## Command Interface
```bash
difii <source-dir> <destination-dir> \
    --only <filename> \
    --overwrite
```
### source-dir
取り込み元のディレクトリ  
### destination-dir
取り込み先のディレクトリ

### 標準出力
```bash
## Summary
|name|diff|
|<filename>|+2 -1|
|<dirname>|<dirname> is directory. skipped.|

## Detail
<filename> has +2 -1 diffs.
+ aaa
+ bbb
- ccc

Do you overwrite ? [Y/n] 
```
