# 11. Measure test coverage

"Coverage" means what percentage of your code actually gets run by your
tests. High coverage doesn't automatically mean bug-free code, but low
coverage definitely means large chunks of code have never been checked by
anything.

## Commands

```
go test -cover
```

This prints a summary line like:

```
ok      11-test-coverage   0.002s  coverage: 75.0% of statements
```

For a more detailed, file-by-file, line-by-line view:

```
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

The second command opens an HTML report in your browser where covered
lines are shown in green and uncovered lines in red.

## Why this example is at ~75%, not 100%

`grade_test.go` only tests 3 of the 4 branches inside `Grade()` - the A
case, the B case, and the default F case. The C case (marks >= 50 and
< 75) is never exercised by any test, so that line shows up as "not
covered" in the report.

To bring this up to 100%, you'd add one more test:

```go
func TestGradeC(t *testing.T) {
	if got := Grade(60); got != "C" {
		t.Errorf("Grade(60) = %s, want C", got)
	}
}
```

I don't have a Go toolchain in this sandbox to actually generate the real
percentage, but leaving one branch deliberately untested here is meant to
make the concept concrete - coverage isn't just an abstract number, it
points at exactly which lines of real logic have never been run by a test.
