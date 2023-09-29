# designdoc
## Usage
```bash
difii --compare <dir> --apply # Overwrite working files with comparison.
difii --compare <dir> --report --report-file report.html --auto-approve
```

## dev command
```bash
go run . --inspect --compare testdata/tourism-a --workdir testdata/tourism-b --summary
```
