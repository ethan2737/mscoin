Set-StrictMode -Version Latest

function Resolve-MscoinValidationBaseline {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot
    )

    [pscustomobject]@{
        AuthSqlPath     = Join-Path $RepoRoot 'scripts\sql\local-auth-baseline.sql'
        MarketSqlPath   = Join-Path $RepoRoot 'scripts\sql\local-market-baseline.sql'
        ExchangeSqlPath = Join-Path $RepoRoot 'scripts\sql\local-exchange-baseline.sql'
        LoginPhone      = '13800000000'
        Password        = '123456'
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
            Name = 'exchange'
            Port = 8083
            WorkingDirectory = $backendRoot
            Command = 'go run exchange/main.go -f exchange/etc/conf.yaml'
            DependsOn = @('ucenter', 'market', 'mysql', 'redis', 'etcd', 'kafka')
        }
        [pscustomobject]@{
            Name = 'jobcenter'
            Port = 0
            WorkingDirectory = $backendRoot
            Command = 'go run jobcenter/main.go -f jobcenter/etc/conf.yaml'
            DependsOn = @('ucenter', 'market', 'mongo', 'redis', 'kafka')
            ReuseMatch = 'jobcenter/main.go -f jobcenter/etc/conf.yaml'
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
            Name = 'exchange-api'
            Port = 8890
            WorkingDirectory = $backendRoot
            Command = 'go run exchange-api/main.go -f exchange-api/etc/conf.yaml'
            DependsOn = @('exchange')
        }
        [pscustomobject]@{
            Name = 'swap-api'
            Port = 8086
            WorkingDirectory = $backendRoot
            Command = 'go run swap-api/main.go -f swap-api/etc/conf.yaml'
            DependsOn = @('market', 'mysql', 'redis', 'etcd', 'mongo')
        }
        [pscustomobject]@{
            Name = 'frontend'
            Port = 3000
            WorkingDirectory = $frontendRoot
            Command = 'pnpm run dev'
            DependsOn = @('ucenter-api', 'market-api', 'exchange-api', 'swap-api')
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
        if ($service.Port -and $service.Port -gt 0) {
            $lines.Add("- $($service.Name): $($service.Status) (port $($service.Port))")
        }
        else {
            $lines.Add("- $($service.Name): $($service.Status) (worker)")
        }
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

function Test-MscoinProcessRunning {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Match
    )

    $processes = Get-CimInstance Win32_Process -ErrorAction SilentlyContinue | Where-Object {
        $_.CommandLine -and $_.CommandLine -like "*$Match*"
    }

    return ($processes | Measure-Object).Count -gt 0
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
    docker exec -e MYSQL_PWD=root mysql8 mysql -uroot -e "CREATE DATABASE IF NOT EXISTS exchange CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;" | Out-Null
    if ($LASTEXITCODE -ne 0) {
        throw 'Failed to ensure exchange database exists.'
    }

    Invoke-MscoinSqlFile -SqlFilePath $Baseline.AuthSqlPath -Database 'mscoin'
    Invoke-MscoinSqlFile -SqlFilePath $Baseline.MarketSqlPath -Database 'market'
    Invoke-MscoinSqlFile -SqlFilePath $Baseline.ExchangeSqlPath -Database 'exchange'
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

    if ($Service.Port -and $Service.Port -gt 0 -and (Test-MscoinTcpPortOpen -Port $Service.Port)) {
        return [pscustomobject]@{
            Name = $Service.Name
            Port = $Service.Port
            Status = 'reused'
            StdOut = $null
            StdErr = $null
        }
    }

    if ((-not $Service.Port -or $Service.Port -le 0) -and $Service.PSObject.Properties.Name -contains 'ReuseMatch' -and (Test-MscoinProcessRunning -Match $Service.ReuseMatch)) {
        return [pscustomobject]@{
            Name = $Service.Name
            Port = $Service.Port
            Status = 'reused'
            StdOut = $null
            StdErr = $null
        }
    }

    $processInfo = Start-MscoinManagedProcess -Service $Service -LogDirectory $LogDirectory
    if (-not $Service.Port -or $Service.Port -le 0) {
        Start-Sleep -Seconds 5
        if ($processInfo.Process.HasExited) {
            $diagnostic = ''
            if ($processInfo.StdErr -and (Test-Path $processInfo.StdErr)) {
                $diagnostic = Get-Content -Path $processInfo.StdErr -Tail 40 | Out-String
            }
            else {
                $diagnostic = "Process exited with code $($processInfo.Process.ExitCode). Re-run with -WriteLogs for persisted stderr."
            }

            throw "Worker '$($Service.Name)' exited during startup.`n$diagnostic"
        }

        return [pscustomobject]@{
            Name = $Service.Name
            Port = $Service.Port
            Status = 'started'
            StdOut = $processInfo.StdOut
            StdErr = $processInfo.StdErr
        }
    }

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

function Test-MscoinThumbTrendReady {
    param(
        [Parameter(Mandatory = $true)]
        [string]$BaseUrl
    )

    try {
        $response = Invoke-RestMethod -Uri "$BaseUrl/market/symbol-thumb-trend" -Method Post -ContentType 'application/json' -Body '{}' -TimeoutSec 15
        if ($response.code -ne 0 -or $null -eq $response.data) {
            return $false
        }

        $readyItems = @($response.data | Where-Object {
            [double]($_.close) -gt 0 -or [double]($_.high) -gt 0 -or [double]($_.low) -gt 0
        })
        return $readyItems.Count -gt 0
    }
    catch {
        return $false
    }
}

function Invoke-MscoinLocalLogin {
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

    Invoke-RestMethod -Uri "$BaseUrl/uc/login" -Method Post -ContentType 'application/json' -Body $body -TimeoutSec 15
}

function Test-MscoinExchangePlateReady {
    param(
        [Parameter(Mandatory = $true)]
        [string]$BaseUrl
    )

    try {
        $response = Invoke-RestMethod -Uri "$BaseUrl/market/exchange-plate-mini" -Method Post -ContentType 'application/json' -Body '{"symbol":"BTC/USDT"}' -TimeoutSec 15
        if ($response.code -ne 0 -or $null -eq $response.data) {
            return $false
        }
        $askCount = @($response.data.ask.items).Count
        $bidCount = @($response.data.bid.items).Count
        return ($askCount -gt 0 -and $bidCount -gt 0)
    }
    catch {
        return $false
    }
}

function Test-MscoinLatestTradeReady {
    param(
        [Parameter(Mandatory = $true)]
        [string]$BaseUrl
    )

    try {
        $response = Invoke-RestMethod -Uri "$BaseUrl/market/latest-trade" -Method Post -ContentType 'application/json' -Body '{"symbol":"BTC/USDT","size":5}' -TimeoutSec 15
        return ($response.code -eq 0 -and @($response.data).Count -gt 0)
    }
    catch {
        return $false
    }
}

function Test-MscoinExchangeOrderReady {
    param(
        [Parameter(Mandatory = $true)]
        [string]$BaseUrl,
        [Parameter(Mandatory = $true)]
        [pscustomobject]$Baseline
    )

    try {
        $loginResponse = Invoke-MscoinLocalLogin -BaseUrl $BaseUrl -Baseline $Baseline
        if ($loginResponse.code -ne 0 -or $null -eq $loginResponse.data.token) {
            return $false
        }
        $headers = @{ 'x-auth-token' = $loginResponse.data.token }
        $response = Invoke-RestMethod -Uri "$BaseUrl/exchange/order/current" -Method Post -Headers $headers -ContentType 'application/json' -Body '{"symbol":"BTC/USDT","pageNo":1,"pageSize":10}' -TimeoutSec 15
        return ($response.code -eq 0 -and $null -ne $response.data)
    }
    catch {
        return $false
    }
}

function Invoke-MscoinSmokeChecks {
    param(
        [Parameter(Mandatory = $true)]
        [string]$BaseUrl,
        [Parameter(Mandatory = $true)]
        [pscustomobject]$Baseline,
        [int]$TimeoutSeconds = 90
    )

    $deadline = (Get-Date).AddSeconds($TimeoutSeconds)
    while ((Get-Date) -lt $deadline) {
        if (
            (Test-MscoinFrontendReady -Url $BaseUrl) -and
            (Test-MscoinLoginReady -BaseUrl $BaseUrl -Baseline $Baseline) -and
            (Test-MscoinThumbTrendReady -BaseUrl $BaseUrl) -and
            (Test-MscoinExchangePlateReady -BaseUrl $BaseUrl) -and
            (Test-MscoinLatestTradeReady -BaseUrl $BaseUrl) -and
            (Test-MscoinExchangeOrderReady -BaseUrl $BaseUrl -Baseline $Baseline)
        ) {
            return
        }
        Start-Sleep -Seconds 2
    }

    if (-not (Test-MscoinFrontendReady -Url $BaseUrl)) {
        throw "Frontend smoke check failed for $BaseUrl."
    }

    if (-not (Test-MscoinLoginReady -BaseUrl $BaseUrl -Baseline $Baseline)) {
        throw "Local login smoke check failed for $BaseUrl/uc/login."
    }
    if (-not (Test-MscoinThumbTrendReady -BaseUrl $BaseUrl)) {
        throw "Market thumb smoke check failed for $BaseUrl/market/symbol-thumb-trend."
    }
    if (-not (Test-MscoinExchangePlateReady -BaseUrl $BaseUrl)) {
        throw "Exchange plate smoke check failed for $BaseUrl/market/exchange-plate-mini."
    }
    if (-not (Test-MscoinLatestTradeReady -BaseUrl $BaseUrl)) {
        throw "Latest trade smoke check failed for $BaseUrl/market/latest-trade."
    }
    throw "Exchange order smoke check failed for $BaseUrl/exchange/order/current."
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
    'Test-MscoinProcessRunning',
    'Ensure-MscoinLogDirectory',
    'Start-MscoinDockerInfrastructure',
    'Ensure-MscoinValidationBaseline',
    'Ensure-MscoinServiceReady',
    'Test-MscoinThumbTrendReady',
    'Test-MscoinExchangePlateReady',
    'Test-MscoinLatestTradeReady',
    'Test-MscoinExchangeOrderReady',
    'Invoke-MscoinSmokeChecks'
)
