Commands:
go test -cover -> Checked overall test coverage 
go test -coverprofile="coverage.out"  ->  Generates a coverage report file
go tool cover -func="coverage.out"  ->  To view coverage for each function