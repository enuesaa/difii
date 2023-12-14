# refactor
Need re-consider command interface because some flags are misleading.

- `--inspect` is an operational flag that changes command behavior while `--workdir` is a setting flag to configure command.
- relation of first argument and `--workdir` flag is not kind to user. which is a, which is b.
- remove import command because unit of import may be huge hunk or mini hunk and this is meaningless.

## Command
```console
$ difii <dir-a> <dir-b> // this shows summary
$ difii <dir-a> <dir-b> --inspect // this shows detail and also summary.
```
