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
difii        # show diffs
difii import # run import
```

### difii (root command)
- summaryを出力する
- 左に source, 右にdestを表示する

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
