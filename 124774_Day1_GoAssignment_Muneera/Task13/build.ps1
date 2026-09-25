Write-Host "Running..."
go run Task13.go

Write-Host "Building..."
go build Task13.go

Write-Host "Testing..."
go test ./...

Write-Host "Formatting..."
go fmt ./...

Write-Host "Checking..."
go vet ./...