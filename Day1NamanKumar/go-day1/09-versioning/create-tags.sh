#!/usr/bin/env bash
# Creates a demo Git history with three semantic-version tags.
set -euo pipefail

git init -q
git add go.mod main.go && git commit -qm "Initial release"
git tag -a v1.0.0 -m "v1.0.0: first stable release"

echo "// Added: greeting feature" >> main.go
git commit -qam "feat: add greeting feature"
git tag -a v1.1.0 -m "v1.1.0: new backward-compatible feature"

echo "// Fixed: typo in output" >> main.go
git commit -qam "fix: correct typo in output"
git tag -a v1.1.1 -m "v1.1.1: bug fix"

git tag -n              # list tags with messages
git log --oneline --decorate
