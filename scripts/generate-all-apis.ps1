# Amazon Ads API 一键生成脚本
# 功能：生成所有 API + 自动修复已知问题

$ErrorActionPreference = "Stop"

# ========== 配置 ==========
$projectRoot = "C:\Users\Administrator\amazon-ads-api-go-sdk"
$specsDir = "$projectRoot\specs"
$pkgDir = "$projectRoot\pkg\adsapi"
$swaggerJar = "$projectRoot\swagger-codegen-cli-3.0.62.jar"

# ========== 所有 API 配置 ==========
$apiConfigs = @(
    # 核心广告产品
    @{ Name = "Sponsored Products v3"; SpecFile = "sponsored-products-v3.json"; PackageName = "sponsoredproductsv3"; OutputDir = "sponsored-products-v3"; Status = "completed" },
    @{ Name = "Sponsored Brands v4"; SpecFile = "sponsored-brands-v4.json"; PackageName = "sponsoredbrandsv4"; OutputDir = "sponsored-brands-v4"; Status = "completed" },
    @{ Name = "Sponsored Display v3"; SpecFile = "sponsored-display-v3.yaml"; PackageName = "sponsoreddisplayv3"; OutputDir = "sponsored-display-v3"; Status = "completed" },
    @{ Name = "Sponsored TV"; SpecFile = "sponsored-tv.json"; PackageName = "sponsoredtv"; OutputDir = "sponsored-tv"; Status = "pending" },
    
    # 核心账户管理
    @{ Name = "Profiles v3"; SpecFile = "profiles-v3.yaml"; PackageName = "profilesv3"; OutputDir = "profiles-v3"; Status = "completed" },
    @{ Name = "Portfolios v2"; SpecFile = "portfolios-v2.json"; PackageName = "portfoliosv2"; OutputDir = "portfolios-v2"; Status = "completed" },
    @{ Name = "Advertising Accounts"; SpecFile = "advertising-accounts.json"; PackageName = "advertisingaccounts"; OutputDir = "advertising-accounts"; Status = "pending" },
    @{ Name = "Manager Accounts"; SpecFile = "manager-accounts.json"; PackageName = "manageraccounts"; OutputDir = "manager-accounts"; Status = "pending" },
    
    # 报告和分析
    @{ Name = "Reporting v3"; SpecFile = "reporting-v3.json"; PackageName = "reportingv3"; OutputDir = "reporting-v3"; Status = "completed" },
    @{ Name = "Brand Metrics"; SpecFile = "brand-metrics.json"; PackageName = "brandmetrics"; OutputDir = "brand-metrics"; Status = "pending" },
    @{ Name = "Insights"; SpecFile = "insights.json"; PackageName = "insights"; OutputDir = "insights"; Status = "pending" },
    @{ Name = "Stores Analytics"; SpecFile = "stores-analytics.json"; PackageName = "storesanalytics"; OutputDir = "stores-analytics"; Status = "pending" },
    @{ Name = "Marketing Mix Modeling"; SpecFile = "marketing-mix-modeling.json"; PackageName = "marketingmixmodeling"; OutputDir = "marketing-mix-modeling"; Status = "pending" },
    
    # DSP
    @{ Name = "DSP Audiences"; SpecFile = "dsp-audiences.json"; PackageName = "dspaudiences"; OutputDir = "dsp-audiences"; Status = "pending" },
    @{ Name = "DSP Conversions"; SpecFile = "dsp-conversions.json"; PackageName = "dspconversions"; OutputDir = "dsp-conversions"; Status = "pending" },
    @{ Name = "DSP Measurement"; SpecFile = "dsp-measurement.json"; PackageName = "dspmeasurement"; OutputDir = "dsp-measurement"; Status = "pending" },
    @{ Name = "DSP Target KPI"; SpecFile = "dsp-target-kpi.json"; PackageName = "dsptargetkpi"; OutputDir = "dsp-target-kpi"; Status = "pending" },
    @{ Name = "DSP Advertisers"; SpecFile = "dsp-advertisers.yaml"; PackageName = "dspadvertisers"; OutputDir = "dsp-advertisers"; Status = "pending" },
    
    # 受众和定位
    @{ Name = "Audiences Discovery"; SpecFile = "audiences-discovery.json"; PackageName = "audiencesdiscovery"; OutputDir = "audiences-discovery"; Status = "pending" },
    @{ Name = "Persona Builder"; SpecFile = "persona-builder.json"; PackageName = "personabuilder"; OutputDir = "persona-builder"; Status = "pending" },
    @{ Name = "Locations"; SpecFile = "locations.json"; PackageName = "locations"; OutputDir = "locations"; Status = "pending" },
    
    # 创意和素材
    @{ Name = "Creative Assets"; SpecFile = "creative-assets.yaml"; PackageName = "creativeassets"; OutputDir = "creative-assets"; Status = "pending" },
    @{ Name = "Moderation"; SpecFile = "moderation.json"; PackageName = "moderation"; OutputDir = "moderation"; Status = "pending" },
    @{ Name = "Pre-Moderation"; SpecFile = "pre-moderation.json"; PackageName = "premoderation"; OutputDir = "pre-moderation"; Status = "pending" },
    @{ Name = "Ad Library"; SpecFile = "ad-library.json"; PackageName = "adlibrary"; OutputDir = "ad-library"; Status = "pending" },
    
    # 产品和目录
    @{ Name = "Product Metadata"; SpecFile = "product-metadata.json"; PackageName = "productmetadata"; OutputDir = "product-metadata"; Status = "pending" },
    @{ Name = "Product Eligibility"; SpecFile = "product-eligibility.json"; PackageName = "producteligibility"; OutputDir = "product-eligibility"; Status = "pending" },
    
    # 归因和测量
    @{ Name = "Amazon Attribution"; SpecFile = "amazon-attribution.json"; PackageName = "amazonattribution"; OutputDir = "amazon-attribution"; Status = "pending" },
    @{ Name = "Reach Forecasting"; SpecFile = "reach-forecasting.json"; PackageName = "reachforecasting"; OutputDir = "reach-forecasting"; Status = "pending" },
    
    # 财务和预算
    @{ Name = "Billing"; SpecFile = "billing.json"; PackageName = "billing"; OutputDir = "billing"; Status = "pending" },
    @{ Name = "Account Budgets"; SpecFile = "account-budgets.json"; PackageName = "accountbudgets"; OutputDir = "account-budgets"; Status = "pending" },
    
    # 数据管理
    @{ Name = "Exports"; SpecFile = "exports.json"; PackageName = "exports"; OutputDir = "exports"; Status = "pending" },
    @{ Name = "Marketing Stream"; SpecFile = "marketing-stream.json"; PackageName = "marketingstream"; OutputDir = "marketing-stream"; Status = "pending" },
    @{ Name = "Hashed Records"; SpecFile = "hashed-records.json"; PackageName = "hashedrecords"; OutputDir = "hashed-records"; Status = "pending" },
    @{ Name = "Data Provider"; SpecFile = "data-provider.yaml"; PackageName = "dataprovider"; OutputDir = "data-provider"; Status = "pending" },
    
    # 其他
    @{ Name = "Tactical Recommendations"; SpecFile = "tactical-recommendations.json"; PackageName = "tacticalrecommendations"; OutputDir = "tactical-recommendations"; Status = "pending" },
    @{ Name = "Change History"; SpecFile = "change-history.json"; PackageName = "changehistory"; OutputDir = "change-history"; Status = "pending" },
    @{ Name = "Partner Opportunities"; SpecFile = "partner-opportunities.json"; PackageName = "partneropportunities"; OutputDir = "partner-opportunities"; Status = "pending" },
    @{ Name = "Test Accounts"; SpecFile = "test-accounts.json"; PackageName = "testaccounts"; OutputDir = "test-accounts"; Status = "pending" },
    
    # 统一 API
    @{ Name = "Amazon Ads API v1"; SpecFile = "amazon-ads-v1-campaign-management.json"; PackageName = "amazonadsv1"; OutputDir = "amazon-ads-v1"; Status = "pending" },
    
    # Retail Ad Service
    @{ Name = "Retail Ad Service"; SpecFile = "retail-ad-service.json"; PackageName = "retailadservice"; OutputDir = "retail-ad-service"; Status = "pending" },
    @{ Name = "Retail Ad Service Identity"; SpecFile = "retail-ad-service-retailer-identity.json"; PackageName = "retailadserviceidentity"; OutputDir = "retail-ad-service-identity"; Status = "pending" },
    
    # 已弃用/旧版本
    @{ Name = "Posts (Deprecated)"; SpecFile = "posts.json"; PackageName = "posts"; OutputDir = "posts"; Status = "pending" },
    @{ Name = "Sponsored Products v2"; SpecFile = "sponsored-products-v2.yaml"; PackageName = "sponsoredproductsv2"; OutputDir = "sponsored-products-v2"; Status = "pending" },
    @{ Name = "Sponsored Brands v3"; SpecFile = "sponsored-brands-v3.yaml"; PackageName = "sponsoredbrandsv3"; OutputDir = "sponsored-brands-v3"; Status = "pending" }
)

# ========== 自动修复函数（集成在脚本内） ==========
function Apply-AutoFix {
    param(
        [string]$ApiDir,
        [string]$PackageName,
        [string]$ApiName
    )
    
    $fixCount = 0
    
    # 修复 1: Creative Assets 联合类型
    if ($ApiDir -match "creative-assets") {
        $file1 = Join-Path $ApiDir "model_ca_value_range_filter_options.go"
        if (Test-Path $file1) {
            $content = Get-Content $file1 -Raw
            if ($content -match "type CaValueRangeFilterOptions struct") {
                $newContent = $content -replace "type CaValueRangeFilterOptions struct \{[^}]+\}", "// CaValueRangeFilterOptions can be SIZE or DATE_UPLOADED`ntype CaValueRangeFilterOptions interface{}"
                Set-Content -Path $file1 -Value $newContent -NoNewline
                Write-Host "    ✓ Fixed CaValueRangeFilterOptions" -ForegroundColor Cyan
                $fixCount++
            }
        }
        
        $file2 = Join-Path $ApiDir "model_ca_value_filter_options.go"
        if (Test-Path $file2) {
            $content = Get-Content $file2 -Raw
            if ($content -match "type CaValueFilterOptions struct") {
                $newContent = $content -replace "type CaValueFilterOptions struct \{[^}]+\}", "// CaValueFilterOptions can be one of multiple filter types`ntype CaValueFilterOptions interface{}"
                Set-Content -Path $file2 -Value $newContent -NoNewline
                Write-Host "    ✓ Fixed CaValueFilterOptions" -ForegroundColor Cyan
                $fixCount++
            }
        }
    }
    
    # 修复 2: Billing 缺失 CountryCode
    if ($ApiDir -match "billing") {
        $countryCodeFile = Join-Path $ApiDir "model_country_code.go"
        if (-not (Test-Path $countryCodeFile)) {
            $countryCodeContent = @"
/*
 * Advertising Billing
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package $PackageName

// CountryCode represents country/region codes
type CountryCode string

// CountryCode enum values
const (
    CountryCodeAE CountryCode = "AE"
    CountryCodeAU CountryCode = "AU"
    CountryCodeBE CountryCode = "BE"
    CountryCodeBR CountryCode = "BR"
    CountryCodeCA CountryCode = "CA"
    CountryCodeCL CountryCode = "CL"
    CountryCodeCO CountryCode = "CO"
    CountryCodeDE CountryCode = "DE"
    CountryCodeEG CountryCode = "EG"
    CountryCodeES CountryCode = "ES"
    CountryCodeFR CountryCode = "FR"
    CountryCodeIE CountryCode = "IE"
    CountryCodeIN CountryCode = "IN"
    CountryCodeIT CountryCode = "IT"
    CountryCodeJP CountryCode = "JP"
    CountryCodeMX CountryCode = "MX"
    CountryCodeNG CountryCode = "NG"
    CountryCodeNL CountryCode = "NL"
    CountryCodePL CountryCode = "PL"
    CountryCodeSA CountryCode = "SA"
    CountryCodeSE CountryCode = "SE"
    CountryCodeSG CountryCode = "SG"
    CountryCodeTR CountryCode = "TR"
    CountryCodeUK CountryCode = "UK"
    CountryCodeUS CountryCode = "US"
    CountryCodeZA CountryCode = "ZA"
)
"@
            Set-Content -Path $countryCodeFile -Value $countryCodeContent
            Write-Host "    ✓ Created CountryCode type" -ForegroundColor Cyan
            $fixCount++
        }
    }
    
    # 修复 3: Amazon Ads v1 非法标识符
    if ($ApiDir -match "amazon-ads-v1") {
        $lookbackFile = Join-Path $ApiDir "model_lookback.go"
        if (Test-Path $lookbackFile) {
            $content = Get-Content $lookbackFile -Raw
            # 检查是否有非法的以数字开头的常量名
            if ($content -match "\s+\d+__") {
                $newContent = $content -replace "const \([^)]+\)", @"
const (
    // Lookback period constants (days)
    LOOKBACK_DAYS_7   Lookback = "DAYS_7"
    LOOKBACK_DAYS_14  Lookback = "DAYS_14"
    LOOKBACK_DAYS_30  Lookback = "DAYS_30"
    LOOKBACK_DAYS_60  Lookback = "DAYS_60"
    LOOKBACK_DAYS_90  Lookback = "DAYS_90"
    LOOKBACK_DAYS_180 Lookback = "DAYS_180"
    LOOKBACK_DAYS_365 Lookback = "DAYS_365"
)
"@
                Set-Content -Path $lookbackFile -Value $newContent -NoNewline
                Write-Host "    ✓ Fixed Lookback constants" -ForegroundColor Cyan
                $fixCount++
            }
        }
    }
    
    if ($fixCount -gt 0) {
        Write-Host "    Applied $fixCount fixes" -ForegroundColor Green
    }
}

# ========== 主流程 ==========
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Amazon Ads API Code Generator" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Total APIs: $($apiConfigs.Count)" -ForegroundColor Yellow
Write-Host "  Completed: $($apiConfigs | Where-Object { $_.Status -eq 'completed' } | Measure-Object | Select-Object -ExpandProperty Count)" -ForegroundColor Green
Write-Host "  Pending: $($apiConfigs | Where-Object { $_.Status -eq 'pending' } | Measure-Object | Select-Object -ExpandProperty Count)" -ForegroundColor Yellow
Write-Host ""

$successCount = 0
$failCount = 0

foreach ($api in $apiConfigs | Where-Object { $_.Status -eq "pending" }) {
    $specFile = Join-Path $specsDir $api.SpecFile
    $tempOutputDir = "$projectRoot\temp-$($api.OutputDir)"
    $finalOutputDir = Join-Path $pkgDir $api.OutputDir
    
    Write-Host "[$($successCount + $failCount + 1)/$($apiConfigs | Where-Object { $_.Status -eq 'pending' } | Measure-Object | Select-Object -ExpandProperty Count)] $($api.Name)" -ForegroundColor Green
    
    if (-not (Test-Path $specFile)) {
        Write-Host "  ✗ Spec file not found" -ForegroundColor Red
        $failCount++
        continue
    }
    
    # 清理
    if (Test-Path $tempOutputDir) { Remove-Item -Recurse -Force $tempOutputDir }
    if (Test-Path $finalOutputDir) { Remove-Item -Recurse -Force $finalOutputDir }
    
    # 生成
    Write-Host "  Generating..." -ForegroundColor Gray
    java -jar "$swaggerJar" generate -i "$specFile" -l go -o "$tempOutputDir" --additional-properties packageName="$($api.PackageName)" 2>&1 | Out-Null
    
    if ($LASTEXITCODE -ne 0 -or -not (Test-Path $tempOutputDir)) {
        Write-Host "  ✗ Generation failed" -ForegroundColor Red
        $failCount++
        continue
    }
    
    # 移动文件
    New-Item -ItemType Directory -Path $finalOutputDir -Force | Out-Null
    Move-Item -Path "$tempOutputDir\*" -Destination $finalOutputDir -Force
    Remove-Item -Recurse -Force $tempOutputDir
    
    # 创建 model_object.go
    $modelObjectFile = Join-Path $finalOutputDir "model_object.go"
    @"
/*
 * $($api.Name)
 *
 * Amazon Advertising API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package $($api.PackageName)

// Object is a generic map for unstructured data.
type Object map[string]interface{}
"@ | Set-Content -Path $modelObjectFile
    
    # 应用自动修复
    Apply-AutoFix -ApiDir $finalOutputDir -PackageName $api.PackageName -ApiName $api.Name
    
    Write-Host "  ✓ Complete" -ForegroundColor Green
    $successCount++
}

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Generation Complete!" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Success: $successCount" -ForegroundColor Green
Write-Host "  Failed: $failCount" -ForegroundColor Red
Write-Host ""
Write-Host "Next steps:" -ForegroundColor Yellow
Write-Host "  1. Run: go build ./..." -ForegroundColor Gray
Write-Host "  2. Run: go test ./..." -ForegroundColor Gray
Write-Host ""

