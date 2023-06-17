**Work in progress...**
# difii
A cli tool to inspect diffs interactively.  

## Usage
```bash
$ difii --source ./appp-aaa/styles --dest ./app-bbb/styles

Diffs Summary
 -4  +0 diffs in emotion.d.ts
-20  +5 diffs in global.ts
-47 +29 diffs in theme.ts
 -0  +0 diffs in use.ts

To inspect diffs:
  difii --source ./appp-aaa/styles --dest ./app-bbb/styles inspect
```

## Commands
```bash
difii         # show summary
difii inspect # show diffs 
difii apply   # import diffs. work in progress..
```
