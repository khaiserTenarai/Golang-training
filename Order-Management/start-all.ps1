# Order Management System - Start Consul and All Go Services
# Run from the project root, or right-click and Run with PowerShell.
# If script execution is blocked, run this once in the same PowerShell window:
# Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass

$ErrorActionPreference = "Stop"
$ProjectRoot = $PSScriptRoot
$ConsulAddress = "http://127.0.0.1:8500"
# Change this path if consul.exe is not in PATH and is stored elsewhere.
$ConsulExe = "C:\Users\M.Chanikya\Downloads\consul_2.1.0_windows_amd64\consul.exe"

$Services = @(
    @{ Name = "product-service"; Port = 8081; Path = "cmd/product-service" },
    @{ Name = "order-service";   Port = 8082; Path = "cmd/order-service" },
    @{ Name = "auth-service";    Port = 8084; Path = "cmd/auth-service" },
    @{ Name = "api-gateway";     Port = 8080; Path = "cmd/api-gateway" }
)

function Test-TcpPort {
    param([string]$HostName = "127.0.0.1", [int]$Port)
    $Client = New-Object System.Net.Sockets.TcpClient
    try {
        $Result = $Client.BeginConnect($HostName, $Port, $null, $null)
        if (-not $Result.AsyncWaitHandle.WaitOne(1000, $false)) { return $false }
        $Client.EndConnect($Result)
        return $true
    }
    catch { return $false }
    finally { $Client.Close() }
}

function Wait-ForPort {
    param([int]$Port, [int]$TimeoutSeconds = 90)
    $Deadline = (Get-Date).AddSeconds($TimeoutSeconds)
    while ((Get-Date) -lt $Deadline) {
        if (Test-TcpPort -HostName "127.0.0.1" -Port $Port) { return $true }
        Start-Sleep -Seconds 2
    }
    return $false
}

Write-Host "`n=== ORDER MANAGEMENT STARTUP ===" -ForegroundColor Cyan

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    throw "Go is not installed or is not available in PATH. Install Go and reopen PowerShell."
}

$EnvFile = Join-Path $ProjectRoot ".env"
if (-not (Test-Path $EnvFile)) {
    throw ".env file not found at $EnvFile. Create .env from .env.example and configure PostgreSQL."
}

foreach ($Service in $Services) {
    $MainFile = Join-Path $ProjectRoot (Join-Path $Service.Path "main.go")
    if (-not (Test-Path $MainFile)) { throw "Service main.go not found: $MainFile" }
}

# Prefer consul.exe found in PATH; otherwise use the configured path above.
$ConsulCommand = Get-Command consul -ErrorAction SilentlyContinue
if ($ConsulCommand) { $ConsulExe = $ConsulCommand.Source }

if (-not (Test-TcpPort -HostName "127.0.0.1" -Port 8500)) {
    if (-not (Test-Path $ConsulExe)) {
        throw "Consul is not running and consul.exe was not found at '$ConsulExe'. Edit `$ConsulExe in start-all.ps1 or add consul.exe to PATH."
    }
    Write-Host "Starting Consul in local development mode..." -ForegroundColor Yellow
    Start-Process -FilePath $ConsulExe -ArgumentList @("agent", "-dev", "-ui", "-client=127.0.0.1") -WorkingDirectory (Split-Path $ConsulExe)
    $ConsulDeadline = (Get-Date).AddSeconds(30)
    while ((Get-Date) -lt $ConsulDeadline -and -not (Test-TcpPort -HostName "127.0.0.1" -Port 8500)) {
        Start-Sleep -Seconds 1
    }
}

if (-not (Test-TcpPort -HostName "127.0.0.1" -Port 8500)) { throw "Consul did not start on port 8500." }
try {
    $Leader = Invoke-RestMethod -Uri "$ConsulAddress/v1/status/leader" -TimeoutSec 3
    if ([string]::IsNullOrWhiteSpace([string]$Leader)) { throw "No Consul leader elected." }
    Write-Host "Consul is ready. Leader: $Leader" -ForegroundColor Green
}
catch { throw "Consul is listening but its API is not ready: $($_.Exception.Message)" }

Write-Host "`nBuilding all Go packages..." -ForegroundColor Yellow
Push-Location $ProjectRoot
try {
    go build ./...
    if ($LASTEXITCODE -ne 0) { throw "Go build failed. Fix the build errors before starting services." }
}
finally { Pop-Location }
Write-Host "Build successful." -ForegroundColor Green

foreach ($Service in $Services) {
    $Port = [int]$Service.Port
    Write-Host "`nStarting $($Service.Name) on port $Port..." -ForegroundColor Cyan

    if (Test-TcpPort -HostName "127.0.0.1" -Port $Port) {
        Write-Host "Port $Port is already listening; not starting a duplicate process." -ForegroundColor Yellow
        continue
    }

    # Start from project root so the service can load the root .env file.
    $EscapedRoot = $ProjectRoot.Replace("'", "''")
    $Command = "Set-Location -LiteralPath '$EscapedRoot'; go run ./$($Service.Path)"
    Start-Process -FilePath "powershell.exe" -WorkingDirectory $ProjectRoot -ArgumentList @("-NoExit", "-ExecutionPolicy", "Bypass", "-Command", $Command)

    Write-Host "Waiting for $($Service.Name)..." -ForegroundColor Yellow
    if (-not (Wait-ForPort -Port $Port -TimeoutSeconds 90)) {
        Write-Warning "$($Service.Name) did not open port $Port. Check that service's PowerShell window and PostgreSQL settings."
        continue
    }

    try {
        $Health = Invoke-WebRequest -Uri "http://127.0.0.1:$Port/health/ready" -TimeoutSec 5 -UseBasicParsing
        if ($Health.StatusCode -ge 200 -and $Health.StatusCode -lt 300) {
            Write-Host "$($Service.Name) is listening and ready." -ForegroundColor Green
        }
        else {
            Write-Warning "$($Service.Name) is listening, but readiness returned HTTP $($Health.StatusCode). Check PostgreSQL and service logs."
        }
    }
    catch {
        Write-Warning "$($Service.Name) is listening, but /health/ready failed. Check the service window and PostgreSQL. Consul may show it as Critical until readiness succeeds."
    }
}

Write-Host "`n=== STARTUP SCRIPT FINISHED ===" -ForegroundColor Cyan
Write-Host "Consul UI:   http://localhost:8500/ui/"
Write-Host "API Gateway: http://localhost:8080"
Write-Host "Health:      http://localhost:8080/health/live"
Write-Host "Services self-register with Consul when they start. Check each service window for errors." -ForegroundColor Yellow
