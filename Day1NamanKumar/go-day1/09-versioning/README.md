# Task 09 — Versioning (Semantic Versioning)

## Create the tags
```bash
./create-tags.sh        # or manually:
git tag -a v1.0.0 -m "first stable release"
git tag -a v1.1.0 -m "new feature"
git tag -a v1.1.1 -m "bug fix"
git tag                 # list
git push origin --tags  # publish (makes them usable via go get)
```

## MAJOR.MINOR.PATCH
| Part | Bump when… | Example | Go consumers |
|------|-----------|---------|--------------|
| **MAJOR** (`v1 → v2.0.0`) | You make **breaking** changes: remove/rename an exported function, change a signature. | `GetEmployee(id int)` → `GetEmployee(ctx, id string)` | Must change import path to `module/v2` |
| **MINOR** (`v1.0.0 → v1.1.0`) | You **add** functionality in a backward-compatible way. | New `SearchByDept()` function | Safe to upgrade |
| **PATCH** (`v1.1.0 → v1.1.1`) | You make backward-compatible **bug fixes** only. | Fix wrong salary rounding | Safe to upgrade |

Rules: bumping MINOR resets PATCH to 0; bumping MAJOR resets both. `v0.x.y` means unstable — anything may change.

```bash
go get example.com/employeeapp@v1.1.1   # exact version
go get example.com/employeeapp@latest   # newest tag
go list -m -versions example.com/employeeapp
```

## Embed the version in the binary
```bash
go build -ldflags "-X main.version=$(git describe --tags)" -o app .
./app   # versioning demo, version: v1.1.1
```
