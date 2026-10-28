
# register-services.ps1

$ErrorActionPreference = "Stop"
$ConsulAddress = "http://127.0.0.1:8500"

$Services = @(
    @{
        ID = "employee-service"
        Name = "employee-service"
        Port = 8081
        HealthCheck = $true
    },
    @{
        ID = "department-service"
        Name = "department-service"
        Port = 8082
        HealthCheck = $true
    },
    @{
        ID = "auth-service"
        Name = "auth-service"
        Port = 8083
        HealthCheck = $true
    },
    @{
        ID = "event-consumer"
        Name = "event-consumer"
        Port = 8084
        HealthCheck = $false
    }
)

# Verify Consul is running
try {
    $leader = Invoke-RestMethod `
        -Uri "$ConsulAddress/v1/status/leader" `
        -TimeoutSec 5

    if ([string]::IsNullOrWhiteSpace([string]$leader)) {
        throw "Consul has no leader."
    }

    Write-Host "Consul is running." -ForegroundColor Green
}
catch {
    Write-Host "ERROR: Consul is unavailable." -ForegroundColor Red
    Write-Host $_.Exception.Message
    exit 1
}

$Failed = $false

foreach ($Service in $Services) {
    $Payload = @{
        ID      = $Service.ID
        Name    = $Service.Name
        Address = "127.0.0.1"
        Port    = $Service.Port
        Tags    = @("go", "microservice")
    }

    if ($Service.HealthCheck) {
        $Payload.Check = @{
            HTTP                          = "http://127.0.0.1:$($Service.Port)/health"
            Interval                      = "10s"
            Timeout                       = "3s"
            DeregisterCriticalServiceAfter = "1m"
        }
    }

    $Json = $Payload | ConvertTo-Json -Depth 10

    try {
        Invoke-RestMethod `
            -Uri "$ConsulAddress/v1/agent/service/register" `
            -Method Put `
            -ContentType "application/json" `
            -Body $Json `
            -TimeoutSec 10 | Out-Null

        Write-Host "Registered $($Service.Name) on port $($Service.Port)" `
            -ForegroundColor Green
    }
    catch {
        $Failed = $true
        Write-Host "Failed: $($Service.Name)" -ForegroundColor Red
        Write-Host $_.Exception.Message
    }
}

Write-Host "`nServices registered with Consul:" -ForegroundColor Cyan

try {
    $Registered = Invoke-RestMethod `
        -Uri "$ConsulAddress/v1/agent/services" `
        -TimeoutSec 5

    $Registered.PSObject.Properties |
        ForEach-Object {
            [PSCustomObject]@{
                ID      = $_.Value.ID
                Name    = $_.Value.Service
                Address = $_.Value.Address
                Port    = $_.Value.Port
            }
        } |
        Format-Table -AutoSize
}
catch {
    Write-Host "Could not list registered services." -ForegroundColor Yellow
}

if ($Failed) {
    Write-Host "One or more registrations failed." -ForegroundColor Red
    exit 1
}
