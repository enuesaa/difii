## Command Interface
```bash
difii 
    --source <source-dir>
    --destination <destination-dir> \
    --only <filename> \
    --overwrite \
    --no-interactive
```
### source-dir
取り込み元のディレクトリ  
### destination-dir
取り込み先のディレクトリ
デフォルトは ./

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
