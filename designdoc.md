# designdoc
## Global Options
```bash
difii
    --source <source-dir> \ # 取り込み元のディレクトリ
    --dest <dest-dir> \     # 取り込み先のディレクトリ. default value is current dir.
    --only <filename>       # diff対象のファイルを絞り込む. 複数指定可能
```

## Commands
```bash
difii         # show summary
difii inspect # show diffs 
difii import  # import diffs
```

### difii (root command)
This command shows diff summary.

### difii inspect
This command shows diff details.

### difii import
import処理をする
```bash
difii import --auto-approve
```

## Development Plans
### v0.1
### v0.2
- use bubbletea.
- implements interavive prompt.
### v0.3
- `difii open` ... start web server and diffs are shown on browser.
