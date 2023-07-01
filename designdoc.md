# designdoc

## Global Options
```bash
difii
    --compare <compare-dir> \ # 取り込み元のディレクトリ
    --workdir <work-dir> \    # 取り込み先のディレクトリ. default value is current dir.
    --only <filename>       # diff対象のファイルを絞り込む. 複数指定可能
```

### difii apply
This command apply diffs to dest directory.
```bash
difii apply --auto-approve
```

## Development Plans
### v0.1
### v0.2
- use bubbletea.
- implements interavive prompt.
### v0.3
- `difii open` ... start web server and diffs are shown on browser.
