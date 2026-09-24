# Task 14 — Troubleshooting

`broken/` contains the intentionally broken code, `fixed/` the corrected version.
Reproduce with `cd broken && go build ./...`.

---

## Error 1 — Configuration error: module path mismatch

**Symptom**
```
main.go:6:2: package employeeapp/employee is not in std (/usr/lib/go-1.22/src/employeeapp/employee)
```

**Cause**
`go.mod` declared `module employeapp` (typo — missing an "e"), but `main.go`
imports `employeeapp/employee`. Because the import path doesn't start with the
module name, Go assumes it must be a standard-library package, can't find it in
GOROOT, and reports "not in std".

**How I identified it**
The message "not in std" for a package I wrote myself means Go doesn't recognise
the path as belonging to my module. I compared `head -1 go.mod` with the import
line and spotted the spelling difference.

**Fix**
```diff
- module employeapp
+ module employeeapp
```
(Alternatively, change the import. The rule: **import path = module name + folder path**.)

---

## Error 2 — Build error: type mismatch and unused import

After fixing Error 1, the build moved on and failed in the `employee` package:

**Symptom**
```
employee/employee.go:3:8: "fmt" imported and not used
employee/employee.go:13:22: cannot use "101" (untyped string constant) as int value in struct literal
```

**Cause**
1. `fmt` was imported but never used — Go treats unused imports as **compile errors**, not warnings.
2. `ID` is declared as `int`, but the code assigned the string `"101"`. Go is statically typed and does no implicit conversion. It also ignored the `id` parameter.

**How I identified it**
The compiler gives the exact `file:line:column`. Line 3 col 8 is the import;
line 13 col 22 is the `ID:` field value.

**Fix**
```diff
-import "fmt"
-
 ...
-	return Employee{ID: "101", Name: "Asha"}
+	return Employee{ID: id, Name: "Asha"}
```
(If a string really came from input, convert it: `id, err := strconv.Atoi(s)`.)

**Verification**
```
$ cd fixed && go vet ./... && go run .
{101 Asha}
```

---

## Lessons / quick checklist
| Symptom | Usual cause | Fix |
|---------|-------------|-----|
| `package X is not in std` | Import path ≠ module name in go.mod | Match them |
| `no required module provides package` | External dependency not added | `go get <pkg>` / `go mod tidy` |
| `imported and not used` / `declared and not used` | Leftover code | Remove it, or `_ =` while debugging |
| `cannot use X (type A) as type B` | Type mismatch | Convert explicitly (`strconv`, `float64(x)`) |
| `go: cannot find main module` | Running outside a folder with go.mod | `cd` into module or `go mod init` |
| Errors appear in "waves" | Compiler stops at the first failing package | Fix, rebuild, repeat |
