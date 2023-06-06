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
difii plan   # dry run import
difii import # run import
```

### difii (root command)
- summaryを出力する
- 左に source, 右にdestを表示する

### difii plan
import処理の実行計画を表示する
```bash
difii plan

<filename> has +2 -1 diffs.
+ aaa
+ bbb
- ccc
```

### difii import
import処理をする
```bash
difii import --auto-approve
```
