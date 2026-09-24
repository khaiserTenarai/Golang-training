# Task 08 — Go Documentation

## GoDoc rules
- A comment directly above an exported identifier documents it.
- It starts with the identifier's name: `// NewStore returns ...`.
- Package docs go in `doc.go` starting with `// Package employee ...`.
- Indent code blocks with a tab; link symbols with `[NewStore]`.
- `Example...` functions in `_test.go` files become runnable docs.

## Generate documentation
```bash
# In the terminal
go doc ./employee              # package summary
go doc ./employee Store        # one type
go doc -all ./employee         # everything

# As a local website (like pkg.go.dev)
go install golang.org/x/pkgsite/cmd/pkgsite@latest
pkgsite -open .                # opens http://localhost:8080

# Save to a file
go doc -all ./employee > employee-docs.txt
```
`employee-docs.txt` in this folder is the generated output.
