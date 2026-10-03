1. Unit Testing:
    I created tests to check whether employee details are valid or invalid.
    I used table-driven tests so I could test different employee inputs.
    I ran go test -v to check the test results.

2. Code Formatting
    I used go fmt ./... to format the Go files properly.

3. Static Analysis
    I ran go vet ./... to check the code for possible issues.

4. Test Coverage
    I used go test -cover to check how much of the code was covered by tests.
    I created a coverage file using go test -coverprofile="coverage.out".
    I used go tool cover -func="coverage.out" to see the coverage of each function.

5. Race Detection:
    I added a test with concurrent goroutines accessing the same employee data.
    This was done to demonstrate a race condition caused by unsynchronized
    access to shared data.
    I ran go test -race to detect the data race.
    The Go race detector identified the race condition during the test.