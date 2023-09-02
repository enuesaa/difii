# designdoc
## Usage
```bash
difii -i
difii --compare <dir>
difii --compare <dir> --summary
difii --compare <dir> --inspect
difii --compare <dir> --apply
difii --compare <dir> --report --report-file report.html
```

## dev command
```bash
go run . --inspect --compare testdata/tourism-a --workdir testdata/tourism-filename-changed --summary
```

## Development Plan
### Inspect
- とりあえず差分がわかれば良い
- 前後2行くらい出力する
- 背景色変えても良い?
