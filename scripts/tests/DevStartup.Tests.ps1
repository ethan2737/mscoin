$repoRoot = (Resolve-Path (Join-Path $PSScriptRoot '..\..')).Path
$modulePath = Join-Path $repoRoot 'scripts\lib\DevStartup.psm1'

Import-Module $modulePath -Force -DisableNameChecking

Describe 'Get-MscoinAppServiceDefinitions' {
    $services = Get-MscoinAppServiceDefinitions -RepoRoot $repoRoot

    It 'returns the minimum local stack in startup order' {
        (($services | Select-Object -ExpandProperty Name) -join ',') | Should Be 'ucenter,market,exchange,jobcenter,ucenter-api,market-api,exchange-api,frontend'
    }

    It 'defines the frontend dev service on port 3000' {
        $frontend = $services | Where-Object { $_.Name -eq 'frontend' }

        $frontend.Port | Should Be 3000
        $frontend.Command | Should Be 'pnpm run dev'
    }

    It 'includes the jobcenter worker in the local stack' {
        $jobcenter = $services | Where-Object { $_.Name -eq 'jobcenter' }

        $jobcenter | Should Not BeNullOrEmpty
        $jobcenter.Command | Should Be 'go run jobcenter/main.go -f jobcenter/etc/conf.yaml'
    }

    It 'includes exchange services in the local stack' {
        $exchange = $services | Where-Object { $_.Name -eq 'exchange' }
        $exchangeApi = $services | Where-Object { $_.Name -eq 'exchange-api' }

        $exchange.Port | Should Be 8083
        $exchangeApi.Port | Should Be 8890
    }
}

Describe 'Resolve-MscoinValidationBaseline' {
    $baseline = Resolve-MscoinValidationBaseline -RepoRoot $repoRoot
    $marketSql = Get-Content -Raw $baseline.MarketSqlPath
    $authSql = Get-Content -Raw $baseline.AuthSqlPath

    It 'returns the local auth baseline sql path' {
        $baseline.AuthSqlPath | Should Exist
    }

    It 'returns the local market baseline sql path' {
        $baseline.MarketSqlPath | Should Exist
    }

    It 'returns the local exchange baseline sql path' {
        $baseline.ExchangeSqlPath | Should Exist
    }

    It 'returns the documented local login account' {
        $baseline.LoginPhone | Should Be '13800000000'
        $baseline.Password | Should Be '123456'
    }

    It 'includes 1INCH coin metadata in the local market baseline' {
        $marketSql | Should Match "'1INCH'"
    }

    It 'includes a 1INCH wallet for the local validation account' {
        $authSql | Should Match "'1INCH', 'LOCAL-1INCH-ADDRESS-10001'"
    }
}

Describe 'Resolve-MscoinLogDirectory' {
    It 'returns null when persistent logs are disabled' {
        $logDirectory = Resolve-MscoinLogDirectory -RepoRoot $repoRoot

        $logDirectory | Should Be $null
    }

    It 'returns the default log directory when persistent logs are enabled' {
        $logDirectory = Resolve-MscoinLogDirectory -RepoRoot $repoRoot -WriteLogs

        $logDirectory | Should Be (Join-Path $repoRoot '.logs\dev-up')
    }
}

Describe 'Format-MscoinStartupSummary' {
    It 'renders urls, log root and login credentials in the summary' {
        $summary = Format-MscoinStartupSummary `
            -FrontendUrl 'http://127.0.0.1:3000' `
            -LogDirectory 'E:\Project\web3\mscoin\.logs\dev-up' `
            -ServiceStatuses @(
                [pscustomobject]@{ Name = 'ucenter'; Status = 'running'; Port = 8081 },
                [pscustomobject]@{ Name = 'frontend'; Status = 'running'; Port = 3000 }
            ) `
            -Baseline ([pscustomobject]@{
                LoginPhone = '13800000000'
                Password = '123456'
            })

        $summary | Should Match 'http://127.0.0.1:3000'
        $summary | Should Match '13800000000'
        $summary | Should Match '123456'
        $summary | Should Match '\.logs\\dev-up'
    }

    It 'renders disabled log mode when persistent logs are off' {
        $summary = Format-MscoinStartupSummary `
            -FrontendUrl 'http://127.0.0.1:3000' `
            -LogDirectory $null `
            -ServiceStatuses @(
                [pscustomobject]@{ Name = 'ucenter'; Status = 'running'; Port = 8081 }
            ) `
            -Baseline ([pscustomobject]@{
                LoginPhone = '13800000000'
                Password = '123456'
            })

        $summary | Should Match 'Logs: disabled'
    }
}

Describe 'Startup smoke checks' {
    It 'treats a 200 frontend response as ready' {
        Mock Invoke-WebRequest -ModuleName DevStartup { [pscustomobject]@{ StatusCode = 200; Content = '<!DOCTYPE html>' } }

        (Test-MscoinFrontendReady -Url 'http://127.0.0.1:3000') | Should Be $true
    }

    It 'treats a successful local login response as ready' {
        Mock Invoke-RestMethod -ModuleName DevStartup {
            @{
                code = 0
                message = 'success'
                data = @{
                    token = 'local-token'
                }
            }
        }

        $baseline = [pscustomobject]@{
            LoginPhone = '13800000000'
            Password = '123456'
        }

        (Test-MscoinLoginReady -BaseUrl 'http://127.0.0.1:3000' -Baseline $baseline) | Should Be $true
    }
}
