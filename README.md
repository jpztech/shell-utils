# shell-utils
Shell utilities for easing development

## ugit

### Commit and push
```bash
ugit commitpush "commit message"
```
1. git add .
2. git commit -m "commit message"
3. git push origin $currentbranch

### Force going back to main branch
```bash
ugit forcemain
```
1. git checkout main
2. git pull origin main
3. git branch -D $prevbranch
