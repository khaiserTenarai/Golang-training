# Code Review — employee_flawed.go

10 issues found, in the order they appear in the file.

## 1. `Employee` struct has no doc comment

Exported types (capitalized, so usable from other packages) are expected
to have a comment starting with their own name, explaining what they're
for. Right now `Employee` has none.

**Fix:**
```go
// Employee represents one employee's basic details.
type Employee struct {
```

## 2. `Emp_list` uses the wrong naming style, and doesn't need to be exported

Go convention is `camelCase`, not `Snake_Case` - `Emp_list` should be
`empList`. It's also capitalized (exported) even though nothing outside
this file needs direct access to it; that just makes it easier for other
code to accidentally modify the list unsafely from anywhere.

**Fix:** rename to `empList` (lowercase, unexported).

## 3. `AddEmployee` silently ignores a real error

```go
salary, _ := strconv.ParseFloat(salaryText, 64)
```

If `salaryText` isn't a valid number (like `"abc"` in `main()`), this
error is thrown away with `_`, and `salary` silently becomes `0` instead.
The program has no way of knowing the input was actually bad.

**Fix:**
```go
salary, err := strconv.ParseFloat(salaryText, 64)
if err != nil {
    fmt.Println("invalid salary:", err)
    return
}
```

## 4. Magic number `0.1` for tax rate

```go
tax := emp.Salary * 0.1
```

What is `0.1`? Someone reading this later has to guess it's a 10% tax
rate. Numbers like this should be named constants.

**Fix:**
```go
const taxRate = 0.10
...
tax := emp.Salary * taxRate
```

## 5. `AddEmployee` does too many unrelated things

In one function, it: parses a string, creates an employee, appends it to
a global list, calculates tax, AND prints output. If any one of those
needs to change (e.g. tax logic), you have to touch a function that's
also responsible for three other things - and it's hard to test any one
part in isolation.

**Fix:** split into smaller functions, e.g. `parseSalary`, `calculateTax`,
and keep `AddEmployee` focused on just adding the employee.

## 6. `GetEmployee` has no bounds checking — this is a real, live bug

```go
func GetEmployee(index int) *Employee {
	return Emp_list[index]
}
```

`main()` calls `PrintSalary(5)` after only adding 2 employees (indexes 0
and 1). `Emp_list[5]` doesn't exist, so this will actually panic at
runtime with "index out of range" - this isn't a style nitpick, it's a
genuine crash waiting to happen any time the caller passes an out-of-range
index.

**Fix:**
```go
func GetEmployee(index int) (*Employee, error) {
	if index < 0 || index >= len(empList) {
		return nil, fmt.Errorf("no employee at index %d", index)
	}
	return empList[index], nil
}
```

## 7. `PrintAllNames` builds a string inefficiently

```go
result := ""
for _, emp := range Emp_list {
	result = result + emp.Name + ", "
}
```

Every `+` here creates a brand new string and copies everything into it -
fine for 2 employees, wasteful for a few thousand. `strings.Builder` (or
`strings.Join`) avoids the repeated copying.

**Fix:**
```go
var b strings.Builder
for _, emp := range empList {
	b.WriteString(emp.Name)
	b.WriteString(", ")
}
return b.String()
```

## 8. `Emp_list` is a shared global with no protection

If two goroutines ever called `AddEmployee` at the same time, both
appending to the same slice with no lock, that's a data race (the same
kind of bug covered in question 8). Nothing here causes it today because
`main()` only calls things one after another, but the moment this code is
used concurrently, it breaks.

**Fix:** either avoid a package-level mutable slice (pass the list around
explicitly), or protect it with a `sync.Mutex` if it truly needs to be
shared.

## 9. Trailing comma left in the joined name string

`PrintAllNames` always leaves a trailing `", "` after the last name (e.g.
`"Anita, Ravi, "` instead of `"Anita, Ravi"`). Small, but it's the kind of
thing that looks sloppy in real output.

**Fix:** use `strings.Join(names, ", ")` instead of manual concatenation -
it never adds a trailing separator.

## 10. No tests exist for any of this

There is no `_test.go` file anywhere alongside `employee_flawed.go`. None
of `AddEmployee`, `GetEmployee`, or `PrintAllNames` have ever actually been
verified by an automated test - issue 6 (the out-of-range panic) is
exactly the kind of bug a simple unit test would have caught before it
ever reached `main()`.

**Fix:** add `employee_test.go` with cases covering a normal add, an
invalid salary string, an out-of-range `GetEmployee` call, and
`PrintAllNames` with 0/1/many employees - similar to the test files in
questions 9, 10, and 12 of this assignment.
