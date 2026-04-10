Set-StrictMode -Version Latest

function Resolve-MscoinValidationBaseline {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot
    )

    [pscustomobject]@{
        AuthSqlPath  = Join-Path $RepoRoot 'scripts\sql\local-auth-baseline.sql'
        MarketSqlPath = Join-Path $RepoRoot 'scripts\sql\local-market-baseline.sql'
        LoginPhone   = '13800000000'
        Password     = '123456'
    }
}

function Get-MscoinInfrastructureDefinitions {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot
    )

    @(
        [pscustomobject]@{ Name = 'mysql'; DockerService = 'mysql'; Port = 33306 }
        [pscustomobject]@{ Name = 'redis'; DockerService = 'Redis'; Port = 33679 }
        [pscustomobject]@{ Name = 'etcd'; DockerService = 'Etcd'; Port = 32379 }
        [pscustomobject]@{ Name = 'mongo'; DockerService = 'mongo'; Port = 37018 }
        [pscustomobject]@{ Name = 'kafka'; DockerService = 'kafka'; Port = 39092 }
        [pscustomobject]@{ Name = 'kafdrop'; DockerService = 'kafdrop'; Port = 39000 }
    )
}

function Get-MscoinAppServiceDefinitions {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot
    )

    $backendRoot = Join-Path $RepoRoot 'mscoin-backend'
    $frontendRoot = Join-Path $RepoRoot 'mscoin-frontend'

    @(
        [pscustomobject]@{
            Name = 'ucenter'
            Port = 8081
            WorkingDirectory = $backendRoot
            Command = 'go run ucenter/main.go -f ucenter/etc/conf.yaml'
            DependsOn = @('mysql', 'redis', 'etcd', 'kafka')
        }
        [pscustomobject]@{
            Name = 'market'
            Port = 8082
            WorkingDirectory = $backendRoot
            Command = 'go run market/main.go -f market/etc/conf.yaml'
            DependsOn = @('mysql', 'redis', 'etcd', 'mongo', 'kafka')
        }
        [pscustomobject]@{
            Name = 'ucenter-api'
            Port = 8888
            WorkingDirectory = $backendRoot
            Command = 'go run ucenter-api/main.go -f ucenter-api/etc/conf.yaml'
            DependsOn = @('ucenter', 'market')
        }
        [pscustomobject]@{
            Name = 'market-api'
            Port = 8889
            WorkingDirectory = $backendRoot
            Command = 'go run market-api/main.go -f market-api/etc/conf.yaml'
            DependsOn = @('market', 'kafka')
        }
        [pscustomobject]@{
            Name = 'frontend'
            Port = 3000
            WorkingDirectory = $frontendRoot
            Command = 'pnpm run dev'
            DependsOn = @('ucenter-api', 'market-api')
        }
    )
}

function Format-MscoinStartupSummary {
    param(
        [Parameter(Mandatory = $true)]
        [string]$FrontendUrl,
        [AllowNull()]
        [string]$LogDirectory,
        [Parameter(Mandatory = $true)]
        [object[]]$ServiceStatuses,
        [Parameter(Mandatory = $true)]
        [pscustomobject]$Baseline
    )

    $lines = [System.Collections.Generic.List[string]]::new()
    $lines.Add('MSCOIN local dev stack is ready.')
    $lines.Add('')
    $lines.Add("Frontend: $FrontendUrl")
    if ([string]::IsNullOrWhiteSpace($LogDirectory)) {
        $lines.Add('Logs: disabled (re-run with -WriteLogs to persist startup logs)')
    }
    else {
        $lines.Add("Logs: $LogDirectory (only newly started services write files)")
    }
    $lines.Add('')
    $lines.Add('Service status:')

    foreach ($service in $ServiceStatuses) {
        $lines.Add("- $($service.Name): $($service.Status) (port $($service.Port))")
    }

    $lines.Add('')
    $lines.Add('Local validation account:')
    $lines.Add("- phone: $($Baseline.LoginPhone)")
    $lines.Add("- password: $($Baseline.Password)")

    ($lines -join [Environment]::NewLine)
}

function Test-MscoinFrontendReady {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Url
    )

    try {
        $response = Invoke-WebRequest -Uri $Url -UseBasicParsing -TimeoutSec 10
        return ($response.StatusCode -eq 200 -and $response.Content -match '<!DOCTYPE html>')
    }
    catch {
        return $false
    }
}

function Test-MscoinLoginReady {
    param(
        [Parameter(Mandatory = $true)]
        [string]$BaseUrl,
        [Parameter(Mandatory = $true)]
        [pscustomobject]$Baseline
    )

    $body = @{
        username = $Baseline.LoginPhone
        password = $Baseline.Password
        captcha  = @{
            mode   = 'fallback'
            passed = $true
        }
    } | ConvertTo-Json -Depth 4

    try {
        $response = Invoke-RestMethod -Uri "$BaseUrl/uc/login" -Method Post -ContentType 'application/json' -Body $body -TimeoutSec 15
        return ($response.code -eq 0 -and $null -ne $response.data.token)
    }
    catch {
        return $false
    }
}

function Assert-MscoinCommandAvailable {
    param(
        [Parameter(Mandatory = $true)]
        [string[]]$Names
    )

    foreach ($name in $Names) {
        if (-not (Get-Command $name -ErrorAction SilentlyContinue)) {
            throw "Required command '$name' is not available."
        }
    }
}

function Test-MscoinTcpPortOpen {
    param(
        [Parameter(Mandatory = $true)]
        [int]$Port
    )

    try {
        $client = [System.Net.Sockets.TcpClient]::new()
        $async = $client.BeginConnect('127.0.0.1', $Port, $null, $null)
        if (-not $async.AsyncWaitHandle.WaitOne(1000, $false)) {
            $client.Close()
            return $false
        }

        $client.EndConnect($async)
        $client.Close()
        return $true
    }
    catch {
        return $false
    }
}

function Wait-MscoinTcpPort {
    param(
        [Parameter(Mandatory = $true)]
        [int]$Port,
        [int]$TimeoutSeconds = 60,
        [int]$PollIntervalMilliseconds = 1000
    )

    $deadline = (Get-Date).AddSeconds($TimeoutSeconds)
    while ((Get-Date) -lt $deadline) {
        if (Test-MscoinTcpPortOpen -Port $Port) {
            return $true
        }
        Start-Sleep -Milliseconds $PollIntervalMilliseconds
    }

    return $false
}

function Resolve-MscoinLogDirectory {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot,
        [switch]$WriteLogs
    )

    if (-not $WriteLogs) {
        return $null
    }

    Join-Path $RepoRoot '.logs\dev-up'
}

function Ensure-MscoinLogDirectory {
    param(
        [AllowNull()]
        [string]$LogDirectory
    )

    if ([string]::IsNullOrWhiteSpace($LogDirectory)) {
        return $null
    }

    if (-not (Test-Path $LogDirectory)) {
        $null = New-Item -ItemType Directory -Path $LogDirectory -Force
    }

    return $LogDirectory
}

function Start-MscoinDockerInfrastructure {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot
    )

    $backendRoot = Join-Path $RepoRoot 'mscoin-backend'
    $infra = Get-MscoinInfrastructureDefinitions -RepoRoot $RepoRoot
    $serviceNames = $infra | Select-Object -ExpandProperty DockerService -Unique
    $serviceArgs = $serviceNames -join ' '
    Push-Location $backendRoot
    try {
        & docker compose up -d $serviceNames | Out-Null
        if ($LASTEXITCODE -ne 0) {
            throw 'Failed to start Docker infrastructure via docker compose.'
        }
    }
    finally {
        Pop-Location
    }

    foreach ($dependency in $infra) {
        if (-not (Wait-MscoinTcpPort -Port $dependency.Port -TimeoutSeconds 90)) {
            throw "Infrastructure dependency '$($dependency.Name)' did not become ready on port $($dependency.Port)."
        }
    }
}

function Invoke-MscoinSqlFile {
    param(
        [Parameter(Mandatory = $true)]
        [string]$SqlFilePath,
        [Parameter(Mandatory = $true)]
        [string]$Database
    )

    $sql = Get-Content -Raw $SqlFilePath
    $sql | docker exec -e MYSQL_PWD=root -i mysql8 mysql -uroot $Database | Out-Null
    if ($LASTEXITCODE -ne 0) {
        throw "Failed to apply SQL file '$SqlFilePath' to database '$Database'."
    }
}

function Ensure-MscoinValidationBaseline {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot,
        [Parameter(Mandatory = $true)]
        [pscustomobject]$Baseline
    )

    docker exec -e MYSQL_PWD=root mysql8 mysql -uroot -e "CREATE DATABASE IF NOT EXISTS market CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;" | Out-Null
    if ($LASTEXITCODE -ne 0) {
        throw 'Failed to ensure market database exists.'
    }

    Invoke-MscoinSqlFile -SqlFilePath $Baseline.AuthSqlPath -Database 'mscoin'
    Invoke-MscoinSqlFile -SqlFilePath $Baseline.MarketSqlPath -Database 'market'
}

function Start-MscoinManagedProcess {
    param(
        [Parameter(Mandatory = $true)]
        [pscustomobject]$Service,
        [AllowNull()]
        [string]$LogDirectory
    )

    $stdout = $null
    $stderr = $null

    if ([string]::IsNullOrWhiteSpace($LogDirectory)) {
        $process = Start-Process `
            -FilePath 'powershell' `
            -ArgumentList @('-NoLogo', '-NoProfile', '-Command', "& { $($Service.Command) *> `$null }") `
            -WorkingDirectory $Service.WorkingDirectory `
            -WindowStyle Hidden `
            -PassThru
    }
    else {
        $null = Ensure-MscoinLogDirectory -LogDirectory $LogDirectory
        $stdout = Join-Path $LogDirectory "$($Service.Name).out.log"
        $stderr = Join-Path $LogDirectory "$($Service.Name).err.log"

        if (Test-Path $stdout) {
            Remove-Item $stdout -Force
        }

        if (Test-Path $stderr) {
            Remove-Item $stderr -Force
        }

        $process = Start-Process `
            -FilePath 'powershell' `
            -ArgumentList @('-NoLogo', '-NoProfile', '-Command', $Service.Command) `
            -WorkingDirectory $Service.WorkingDirectory `
            -RedirectStandardOutput $stdout `
            -RedirectStandardError $stderr `
            -WindowStyle Hidden `
            -PassThru
    }

    [pscustomobject]@{
        Process = $process
        StdOut  = $stdout
        StdErr  = $stderr
    }
}

function Ensure-MscoinServiceReady {
    param(
        [Parameter(Mandatory = $true)]
        [pscustomobject]$Service,
        [AllowNull()]
        [string]$LogDirectory,
        [int]$TimeoutSeconds = 90
    )

    if (Test-MscoinTcpPortOpen -Port $Service.Port) {
        return [pscustomobject]@{
            Name = $Service.Name
            Port = $Service.Port
            Status = 'reused'
            StdOut = $null
            StdErr = $null
        }
    }

    $processInfo = Start-MscoinManagedProcess -Service $Service -LogDirectory $LogDirectory
    if (-not (Wait-MscoinTcpPort -Port $Service.Port -TimeoutSeconds $TimeoutSeconds)) {
        $diagnostic = ''
        if ($processInfo.StdErr -and (Test-Path $processInfo.StdErr)) {
            $diagnostic = Get-Content -Path $processInfo.StdErr -Tail 40 | Out-String
        }
        elseif ($processInfo.Process.HasExited) {
            $diagnostic = "Process exited with code $($processInfo.Process.ExitCode). Re-run with -WriteLogs for persisted stderr."
        }
        else {
            $diagnostic = 'No persisted startup logs were captured. Re-run with -WriteLogs for service stderr.'
        }

        throw "Service '$($Service.Name)' did not become ready on port $($Service.Port).`n$diagnostic"
    }

    [pscustomobject]@{
        Name = $Service.Name
        Port = $Service.Port
        Status = 'started'
        StdOut = $processInfo.StdOut
        StdErr = $processInfo.StdErr
    }
}

function Invoke-MscoinSmokeChecks {
    param(
        [Parameter(Mandatory = $true)]
        [string]$BaseUrl,
        [Parameter(Mandatory = $true)]
        [pscustomobject]$Baseline
    )

    if (-not (Test-MscoinFrontendReady -Url $BaseUrl)) {
        throw "Frontend smoke check failed for $BaseUrl."
    }

    if (-not (Test-MscoinLoginReady -BaseUrl $BaseUrl -Baseline $Baseline)) {
        throw "Local login smoke check failed for $BaseUrl/uc/login."
    }
}

Export-ModuleMember -Function @(
    'Resolve-MscoinValidationBaseline',
    'Resolve-MscoinLogDirectory',
    'Get-MscoinInfrastructureDefinitions',
    'Get-MscoinAppServiceDefinitions',
    'Format-MscoinStartupSummary',
    'Test-MscoinFrontendReady',
    'Test-MscoinLoginReady',
    'Assert-MscoinCommandAvailable',
    'Test-MscoinTcpPortOpen',
    'Wait-MscoinTcpPort',
    'Ensure-MscoinLogDirectory',
    'Start-MscoinDockerInfrastructure',
    'Ensure-MscoinValidationBaseline',
    'Ensure-MscoinServiceReady',
    'Invoke-MscoinSmokeChecks'
)
