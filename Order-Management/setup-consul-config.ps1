# Order Management System - Store and Verify Configuration in Consul KV
# Consul must be running first. This script does not start or stop services.
# If script execution is blocked, run this once in the same PowerShell window:
# Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass

$ErrorActionPreference = "Stop"
$ConsulAddress = "http://127.0.0.1:8500"
$ConfigKey = "$ConsulAddress/v1/kv/order-management/config"

Write-Host "`n========================================" -ForegroundColor Cyan
Write-Host " Order Management - Consul Configuration" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan

Write-Host "`n[1/3] Checking Consul..." -ForegroundColor Yellow
try {
    $Leader = Invoke-RestMethod -Uri "$ConsulAddress/v1/status/leader" -TimeoutSec 3
    if ([string]::IsNullOrWhiteSpace([string]$Leader)) { throw "Consul has no elected leader." }
    Write-Host "Consul is running. Leader: $Leader" -ForegroundColor Green
}
catch {
    Write-Host "ERROR: Consul is not available at $ConsulAddress. Run start-all.ps1 first or start Consul manually." -ForegroundColor Red
    exit 1
}

Write-Host "`n[2/3] Storing configuration in Consul KV..." -ForegroundColor Yellow
# Ports match this project's .env: Gateway 8080, Product 8081, Order 8082, Auth 8084.
$config = @{
    api_gateway_url = "http://localhost:8080"
    product_service_url = "http://localhost:8081"
    order_service_url = "http://localhost:8082"
    auth_service_url = "http://localhost:8084"
    http_client_timeout_seconds = 5
    http_client_max_retries = 3
    circuit_breaker_failure_threshold = 5
    circuit_breaker_reset_timeout_seconds = 30
} | ConvertTo-Json -Depth 5

try {
    $result = Invoke-RestMethod -Method Put -Uri $ConfigKey -ContentType "application/json" -Body $config -TimeoutSec 5
    if ($result -ne $true) { throw "Consul did not confirm the KV write." }
    Write-Host "Configuration stored successfully." -ForegroundColor Green
}
catch {
    Write-Host "ERROR: Could not store configuration: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

Write-Host "`n[3/3] Reading back and verifying configuration..." -ForegroundColor Yellow
try {
    $Response = Invoke-WebRequest -Uri "$($ConfigKey)?raw" -Method Get -UseBasicParsing -TimeoutSec 5
    $SavedConfig = $Response.Content | ConvertFrom-Json
    $SavedConfig | Format-List
    Write-Host "Configuration verified successfully." -ForegroundColor Green
}
catch {
    Write-Host "ERROR: Could not verify saved configuration: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

Write-Host "`nDone. Configuration is saved at Consul KV key: order-management/config" -ForegroundColor Cyan
Write-Host "Note: Consul KV values affect runtime behavior only if the Go application explicitly reads this KV key." -ForegroundColor Yellow
