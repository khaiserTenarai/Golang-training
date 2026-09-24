Write-Host "1. RUN"
go run ./cmd/app

Write-Host ""
Write-Host "2. BUILD"
go build ./...

Write-Host ""
Write-Host "3. TEST"
go test ./...

Write-Host ""
Write-Host "4. FORMAT"
go fmt ./...

Write-Host ""
Write-Host "5. VET"
go vet ./...