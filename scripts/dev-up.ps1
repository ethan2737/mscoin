param(
    [switch]$SkipSmoke,
    [switch]$WriteLogs,
    [int]$TimeoutSeconds = 90
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$repoRoot = (Resolve-Path (Join-Path $PSScriptRoot '..')).Path
$modulePath = Join-Path $repoRoot 'scripts\lib\DevStartup.psm1'
Import-Module $modulePath -Force -DisableNameChecking

function Write-Step {
    param([string]$Message)
    Write-Host "==> $Message"
}

try {
    Assert-MscoinCommandAvailable -Names @('docker', 'go', 'pnpm')

    $logDirectory = Resolve-MscoinLogDirectory -RepoRoot $repoRoot -WriteLogs:$WriteLogs
    $baseline = Resolve-MscoinValidationBaseline -RepoRoot $repoRoot
    $infra = Get-MscoinInfrastructureDefinitions -RepoRoot $repoRoot
    $services = Get-MscoinAppServiceDefinitions -RepoRoot $repoRoot

    Write-Step 'Starting Docker infrastructure'
    Start-MscoinDockerInfrastructure -RepoRoot $repoRoot
    $infraStatuses = $infra | ForEach-Object {
        [pscustomobject]@{
            Name = $_.Name
            Status = 'running'
            Port = $_.Port
        }
    }

    Write-Step 'Applying local validation baselines'
    Ensure-MscoinValidationBaseline -RepoRoot $repoRoot -Baseline $baseline

    $portMap = @{}
    foreach ($dependency in $infra) {
        $portMap[$dependency.Name] = $dependency.Port
    }
    foreach ($service in $services) {
        $portMap[$service.Name] = $service.Port
    }

    $serviceStatuses = [System.Collections.Generic.List[object]]::new()

    foreach ($service in $services) {
        foreach ($dependencyName in $service.DependsOn) {
            $dependencyPort = $portMap[$dependencyName]
            if (-not (Wait-MscoinTcpPort -Port $dependencyPort -TimeoutSeconds $TimeoutSeconds)) {
                throw "Dependency '$dependencyName' for service '$($service.Name)' was not ready on port $dependencyPort."
            }
        }

        Write-Step "Ensuring service '$($service.Name)' is ready"
        $status = Ensure-MscoinServiceReady -Service $service -LogDirectory $logDirectory -TimeoutSeconds $TimeoutSeconds
        $serviceStatuses.Add($status)
    }

    if (-not $SkipSmoke) {
        Write-Step 'Running smoke checks'
        Invoke-MscoinSmokeChecks -BaseUrl 'http://127.0.0.1:3000' -Baseline $baseline
    }

    $summary = Format-MscoinStartupSummary `
        -FrontendUrl 'http://127.0.0.1:3000' `
        -LogDirectory $logDirectory `
        -ServiceStatuses ($infraStatuses + $serviceStatuses) `
        -Baseline $baseline

    Write-Host ''
    Write-Host $summary
}
catch {
    Write-Error $_
    exit 1
}
