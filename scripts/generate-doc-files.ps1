# 批量生成所有 API 的 doc.go 文件

$ErrorActionPreference = "Stop"

$projectRoot = "C:\Users\Administrator\amazon-ads-api-go-sdk"
$pkgDir = "$projectRoot\pkg\adsapi"

# API 元数据配置
$apiMetadata = @{
    "sponsored-products-v3" = @{
        Name = "Sponsored Products v3"
        Package = "sponsoredproductsv3"
        Description = "Sponsored Products v3 是 Amazon Advertising API 中最核心的 API，提供了完整的商品推广广告管理功能。"
        Features = @(
            "广告活动管理 (Campaigns)",
            "广告组管理 (Ad Groups)",
            "关键词管理 (Keywords)",
            "否定关键词管理 (Negative Keywords)",
            "商品广告管理 (Product Ads)",
            "商品定向管理 (Product Targeting)",
            "定向条款管理 (Targeting Clauses)",
            "预算管理 (Budget Rules)",
            "优化规则 (Optimization Rules)",
            "推荐服务 (Recommendations)"
        )
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/sponsored-products/3-0/openapi"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-products/openapi.json"
        Status = "completed"
    }
    "sponsored-brands-v4" = @{
        Name = "Sponsored Brands v4"
        Package = "sponsoredbrandsv4"
        Description = "Sponsored Brands v4 提供品牌推广广告管理功能，帮助提升品牌知名度。"
        Features = @(
            "品牌广告活动管理",
            "品牌广告组管理",
            "品牌关键词管理",
            "品牌创意管理",
            "视频广告管理",
            "商店聚光灯广告"
        )
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/sponsored-brands/4-0/openapi"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-brands/openapi.json"
        Status = "completed"
    }
    "sponsored-display-v3" = @{
        Name = "Sponsored Display v3"
        Package = "sponsoreddisplayv3"
        Description = "Sponsored Display v3 提供展示广告管理功能，通过程序化展示广告触达目标受众。"
        Features = @(
            "展示广告活动管理",
            "展示广告组管理",
            "受众定向管理",
            "商品定向管理",
            "创意管理",
            "预算管理"
        )
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/sponsored-display/3-0/openapi"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-display/openapi.yaml"
        Status = "completed"
    }
    "profiles-v3" = @{
        Name = "Profiles v3"
        Package = "profilesv3"
        Description = "Profiles API 提供账户配置文件管理功能，用于管理广告账户和市场配置。"
        Features = @(
            "配置文件列表查询",
            "配置文件详情获取",
            "市场配置管理",
            "账户信息管理"
        )
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/account-management/profiles"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/profiles/openapi.yaml"
        Status = "completed"
    }
    "portfolios-v2" = @{
        Name = "Portfolios v2"
        Package = "portfoliosv2"
        Description = "Portfolios API 提供广告组合管理功能，帮助组织和管理多个广告活动。"
        Features = @(
            "组合创建和管理",
            "组合预算管理",
            "广告活动分组",
            "组合报告"
        )
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/campaign-management/portfolios"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/portfolios/openapi.json"
        Status = "completed"
    }
    "reporting-v3" = @{
        Name = "Reporting v3"
        Package = "reportingv3"
        Description = "Reporting v3 提供统一的广告报告生成和下载功能。"
        Features = @(
            "报告创建",
            "报告状态查询",
            "报告下载",
            "自定义指标",
            "多维度数据分析"
        )
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/reporting/v3/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/reporting/openapi.json"
        Status = "completed"
    }
    # 新生成的 API
    "sponsored-tv" = @{
        Name = "Sponsored TV"
        Package = "sponsoredtv"
        Description = "Sponsored TV API 提供电视流媒体广告管理功能。"
        Features = @("广告活动管理", "广告组管理", "创意管理", "定向管理", "预测服务")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/sponsored-tv/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-tv/openapi.json"
    }
    "advertising-accounts" = @{
        Name = "Advertising Accounts"
        Package = "advertisingaccounts"
        Description = "Advertising Accounts API 提供广告账户管理功能。"
        Features = @("账户查询", "账户详情", "账户配置")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/account-management/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/advertising-accounts/openapi.json"
    }
    "manager-accounts" = @{
        Name = "Manager Accounts"
        Package = "manageraccounts"
        Description = "Manager Accounts API 提供管理账户功能。"
        Features = @("管理账户查询", "子账户管理", "权限管理")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/account-management/manager-accounts"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/manager-accounts/openapi.json"
    }
    "brand-metrics" = @{
        Name = "Brand Metrics"
        Package = "brandmetrics"
        Description = "Brand Metrics API 提供品牌指标分析功能。"
        Features = @("品牌知名度指标", "品牌搜索指标", "市场份额分析")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/reporting/brand-metrics"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/brand-metrics/openapi.json"
    }
    "insights" = @{
        Name = "Insights"
        Package = "insights"
        Description = "Insights API 提供广告洞察和分析功能。"
        Features = @("搜索词洞察", "竞品分析", "市场趋势")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/insights/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/insights/openapi.json"
    }
    "stores-analytics" = @{
        Name = "Stores Analytics"
        Package = "storesanalytics"
        Description = "Stores Analytics API 提供品牌旗舰店分析功能。"
        Features = @("店铺流量分析", "销售数据", "访客行为")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/stores/analytics"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/stores-analytics/openapi.json"
    }
    "marketing-mix-modeling" = @{
        Name = "Marketing Mix Modeling"
        Package = "marketingmixmodeling"
        Description = "Marketing Mix Modeling API 提供营销组合建模和分析功能。"
        Features = @("归因分析", "ROI 计算", "渠道效果评估")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/amazon-marketing-cloud/marketing-mix-modeling"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/marketing-mix-modeling/openapi.json"
    }
    "dsp-audiences" = @{
        Name = "DSP Audiences"
        Package = "dspaudiences"
        Description = "DSP Audiences API 提供 DSP 受众管理功能。"
        Features = @("受众创建", "受众分段", "受众分析")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/dsp/audiences"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/dsp-audiences/openapi.json"
    }
    "dsp-conversions" = @{
        Name = "DSP Conversions"
        Package = "dspconversions"
        Description = "DSP Conversions API 提供 DSP 转化跟踪功能。"
        Features = @("转化事件上报", "像素管理", "转化分析")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/dsp/conversions"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/dsp-conversions/openapi.json"
    }
    "dsp-measurement" = @{
        Name = "DSP Measurement"
        Package = "dspmeasurement"
        Description = "DSP Measurement API 提供 DSP 广告效果测量功能。"
        Features = @("品牌提升研究", "归因分析", "效果测量")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/dsp/measurement"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/dsp-measurement/openapi.json"
    }
    "dsp-target-kpi" = @{
        Name = "DSP Target KPI"
        Package = "dsptargetkpi"
        Description = "DSP Target KPI API 提供 DSP 目标 KPI 管理功能。"
        Features = @("KPI 设置", "目标优化", "性能监控")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/dsp/target-kpi"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/dsp-target-kpi/openapi.json"
    }
    "dsp-advertisers" = @{
        Name = "DSP Advertisers"
        Package = "dspadvertisers"
        Description = "DSP Advertisers API 提供 DSP 广告主管理功能。"
        Features = @("广告主账户管理", "订单管理", "广告系列管理")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/dsp/advertisers"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/dsp-advertisers/openapi.yaml"
    }
    "audiences-discovery" = @{
        Name = "Audiences Discovery"
        Package = "audiencesdiscovery"
        Description = "Audiences Discovery API 提供受众发现和推荐功能。"
        Features = @("受众推荐", "相似受众", "受众洞察")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/audiences/discovery"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/audiences-discovery/openapi.json"
    }
    "persona-builder" = @{
        Name = "Persona Builder"
        Package = "personabuilder"
        Description = "Persona Builder API 提供用户画像构建功能。"
        Features = @("画像创建", "行为分析", "兴趣标签")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/audiences/persona-builder"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/persona-builder/openapi.json"
    }
    "locations" = @{
        Name = "Locations"
        Package = "locations"
        Description = "Locations API 提供地理位置定向功能。"
        Features = @("地理位置查询", "区域定向", "位置分析")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/targeting/locations"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/locations/openapi.json"
    }
    "creative-assets" = @{
        Name = "Creative Assets"
        Package = "creativeassets"
        Description = "Creative Assets API 提供创意素材管理功能。"
        Features = @("素材上传", "素材管理", "素材审核状态")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/creative-asset/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/creative-assets/openapi.yaml"
    }
    "moderation" = @{
        Name = "Moderation"
        Package = "moderation"
        Description = "Moderation API 提供广告审核功能。"
        Features = @("审核状态查询", "政策违规检查", "审核历史")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/moderation/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/moderation/openapi.json"
    }
    "pre-moderation" = @{
        Name = "Pre-Moderation"
        Package = "premoderation"
        Description = "Pre-Moderation API 提供预审核功能。"
        Features = @("预审核检查", "合规性验证", "政策预检")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/moderation/pre-moderation"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/pre-moderation/openapi.json"
    }
    "ad-library" = @{
        Name = "Ad Library"
        Package = "adlibrary"
        Description = "Ad Library API 提供广告库管理功能。"
        Features = @("广告素材库", "历史广告查询", "广告创意管理")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/ad-library/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/ad-library/openapi.json"
    }
    "product-metadata" = @{
        Name = "Product Metadata"
        Package = "productmetadata"
        Description = "Product Metadata API 提供商品元数据查询功能。"
        Features = @("商品信息查询", "类目数据", "商品属性")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/product-metadata/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/product-metadata/openapi.json"
    }
    "product-eligibility" = @{
        Name = "Product Eligibility"
        Package = "producteligibility"
        Description = "Product Eligibility API 提供商品广告资格检查功能。"
        Features = @("资格检查", "限制查询", "合规性验证")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/product-eligibility/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/product-eligibility/openapi.json"
    }
    "amazon-attribution" = @{
        Name = "Amazon Attribution"
        Package = "amazonattribution"
        Description = "Amazon Attribution API 提供跨渠道归因分析功能。"
        Features = @("归因标签管理", "转化跟踪", "渠道效果分析")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/amazon-attribution-api"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/amazon-attribution/openapi.json"
    }
    "reach-forecasting" = @{
        Name = "Reach Forecasting"
        Package = "reachforecasting"
        Description = "Reach Forecasting API 提供覆盖预测功能。"
        Features = @("覆盖预测", "频次预估", "受众规模预测")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/reach-forecasting/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/reach-forecasting/openapi.json"
    }
    "billing" = @{
        Name = "Billing"
        Package = "billing"
        Description = "Billing API 提供账单和支付管理功能。"
        Features = @("账单查询", "发票下载", "支付方式管理", "消费报表")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/billing/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/billing/openapi.json"
    }
    "account-budgets" = @{
        Name = "Account Budgets"
        Package = "accountbudgets"
        Description = "Account Budgets API 提供账户预算管理功能。"
        Features = @("预算设置", "预算监控", "预算使用情况")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/budget-usage/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/account-budgets/openapi.json"
    }
    "exports" = @{
        Name = "Exports"
        Package = "exports"
        Description = "Exports API 提供数据导出功能。"
        Features = @("数据导出", "批量下载", "自定义导出")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/exports/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/exports/openapi.json"
    }
    "marketing-stream" = @{
        Name = "Marketing Stream"
        Package = "marketingstream"
        Description = "Marketing Stream API 提供实时营销数据流功能。"
        Features = @("实时数据流", "事件订阅", "数据推送")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/marketing-stream/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/marketing-stream/openapi.json"
    }
    "hashed-records" = @{
        Name = "Hashed Records"
        Package = "hashedrecords"
        Description = "Hashed Records API 提供哈希记录管理功能。"
        Features = @("记录上传", "数据加密", "隐私保护")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/hashed-records/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/hashed-records/openapi.json"
    }
    "data-provider" = @{
        Name = "Data Provider"
        Package = "dataprovider"
        Description = "Data Provider API 提供数据提供商集成功能。"
        Features = @("数据源管理", "数据同步", "第三方集成")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/data-provider/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/data-provider/openapi.yaml"
    }
    "tactical-recommendations" = @{
        Name = "Tactical Recommendations"
        Package = "tacticalrecommendations"
        Description = "Tactical Recommendations API 提供战术推荐功能。"
        Features = @("优化建议", "出价推荐", "预算推荐", "关键词推荐")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/rules/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/tactical-recommendations/openapi.json"
    }
    "change-history" = @{
        Name = "Change History"
        Package = "changehistory"
        Description = "Change History API 提供变更历史记录功能。"
        Features = @("变更记录", "审计日志", "操作历史")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/change-history/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/change-history/openapi.json"
    }
    "partner-opportunities" = @{
        Name = "Partner Opportunities"
        Package = "partneropportunities"
        Description = "Partner Opportunities API 提供合作伙伴机会管理功能。"
        Features = @("机会查询", "合作管理", "收益分析")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/partner-opportunities/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/partner-opportunities/openapi.json"
    }
    "test-accounts" = @{
        Name = "Test Accounts"
        Package = "testaccounts"
        Description = "Test Accounts API 提供测试账户管理功能。"
        Features = @("测试账户创建", "沙箱环境", "测试数据")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/testing/test-accounts"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/test-accounts/openapi.json"
    }
    "amazon-ads-v1" = @{
        Name = "Amazon Ads API v1"
        Package = "amazonadsv1"
        Description = "Amazon Ads API v1 是新的统一 API，提供跨广告产品的统一接口。"
        Features = @("统一广告活动管理", "跨产品广告组管理", "统一定向管理")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/betas/unified-api/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/amazon-ads-v1/openapi.json"
    }
    "retail-ad-service" = @{
        Name = "Retail Ad Service"
        Package = "retailadservice"
        Description = "Retail Ad Service API 提供零售广告服务功能。"
        Features = @("零售广告管理", "商品推广", "店内广告")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/retail-ad-service/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/retail-ad-service/openapi.json"
    }
    "retail-ad-service-identity" = @{
        Name = "Retail Ad Service Identity"
        Package = "retailadserviceidentity"
        Description = "Retail Ad Service Identity API 提供零售广告身份管理功能。"
        Features = @("身份验证", "权限管理", "零售商配置")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/retail-ad-service/identity"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/retail-ad-service-retailer-identity/openapi.json"
    }
    "posts" = @{
        Name = "Posts (Deprecated)"
        Package = "posts"
        Description = "Posts API 提供帖子管理功能（已弃用，将于 2025-07-31 关闭）。"
        Features = @("帖子创建", "帖子管理", "内容发布")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/guides/posts/overview"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/posts/openapi.json"
    }
    "sponsored-products-v2" = @{
        Name = "Sponsored Products v2"
        Package = "sponsoredproductsv2"
        Description = "Sponsored Products v2 是旧版商品推广 API（建议使用 v3）。"
        Features = @("广告活动管理（v2）", "关键词管理（v2）", "报告（v2）")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/sponsored-products/2-0/openapi"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-products-v2/openapi.yaml"
    }
    "sponsored-brands-v3" = @{
        Name = "Sponsored Brands v3"
        Package = "sponsoredbrandsv3"
        Description = "Sponsored Brands v3 是旧版品牌推广 API（建议使用 v4）。"
        Features = @("品牌广告管理（v3）", "视频广告（v3）")
        DocUrl = "https://advertising.amazon.com/API/docs/en-us/sponsored-brands/3-0/openapi"
        SpecUrl = "https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-brands-v3/openapi.yaml"
    }
}

# 生成 doc.go 内容的函数
function Generate-DocFile {
    param(
        [string]$Dir,
        [hashtable]$Metadata
    )
    
    $packageName = $Metadata.Package
    $apiName = $Metadata.Name
    $description = $Metadata.Description
    $features = $Metadata.Features -join "`n//   - "
    $docUrl = $Metadata.DocUrl
    $specUrl = $Metadata.SpecUrl
    
    $content = @"
// Copyright 2025 Amazon Ads API Go SDK Authors.
//
// This file is part of Amazon Ads API Go SDK.
//
// Amazon Ads API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//    - Free for personal, educational, and open source projects
//    - Your project must also be open sourced under AGPL-3.0
//    - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//    - Required for any commercial, enterprise, or proprietary use
//    - Allows closed source distribution
//    - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license.

// Package $packageName 提供 Amazon Advertising API $apiName 客户端。
//
// # 概述
//
// $description
// 本包是基于官方 OpenAPI 规范通过 swagger-codegen 自动生成的 Go 客户端。
//
// 主要功能包括：
//   - $features
//
// # 官方文档
//
// API 文档: $docUrl
//
// OpenAPI 规范: $specUrl
//
// # 快速开始
//
// 基本用法示例：
//
//	package main
//
//	import (
//	    "context"
//	    "log"
//
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi"
//	    "$packageName "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi/$($Dir -replace '\\','/')"
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/internal/models"
//	)
//
//	func main() {
//	    // 1. 创建基础客户端
//	    baseClient, err := adsapi.NewClient(
//	        adsapi.WithRegion(models.RegionNA),
//	        adsapi.WithCredentials("client-id", "client-secret", "refresh-token"),
//	        adsapi.WithProfileID(1234567890),
//	    )
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    // 2. 创建 API 配置
//	    cfg := $packageName.NewConfiguration()
//	    cfg.HTTPClient = baseClient.HTTPClient()
//	    cfg.Host = models.RegionNA.Endpoint
//
//	    // 3. 创建 API 客户端
//	    apiClient := $packageName.NewAPIClient(cfg)
//
//	    // 4. 调用 API
//	    ctx := context.Background()
//	    // 具体的 API 调用方法请参考官方文档
//	}
//
// # 认证与授权
//
// $apiName 使用 Amazon LWA (Login with Amazon) OAuth 2.0 进行认证。
// 所有请求都需要在 Header 中提供：
//   - Authorization: Bearer {access_token}
//   - Amazon-Advertising-API-ClientId: {client_id}
//   - Amazon-Advertising-API-Scope: {profile_id}
//
// 基础客户端（adsapi.Client）会自动处理认证和 Token 刷新。
//
// # 错误处理
//
// API 可能返回以下错误类型：
//   - 400 Bad Request - 请求参数无效
//   - 401 Unauthorized - 认证失败或 Token 过期
//   - 403 Forbidden - 没有权限访问该资源
//   - 404 Not Found - 资源不存在
//   - 429 Too Many Requests - 超过速率限制
//   - 500 Internal Server Error - 服务器内部错误
//
// # 代码生成说明
//
// 本包中除 doc.go 外的所有代码均由 swagger-codegen 自动生成。
// 不要手动修改自动生成的文件，所有修改都应该通过更新 OpenAPI 规范并重新生成来完成。
//
// # 相关包
//
//   - github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi - 基础客户端和配置
//   - github.com/vanling1111/amazon-ads-api-go-sdk/internal/auth - LWA 认证
//   - github.com/vanling1111/amazon-ads-api-go-sdk/internal/ratelimit - 速率限制
//
// # 版本历史
//
//   - v0.1.0 (2025-10-07) - 初始版本
//
// # 许可证
//
// 本包使用双许可证模式：
//   - AGPL-3.0 用于开源项目
//   - Commercial License 用于商业项目
//
// 详见项目根目录的 LICENSE 文件。
package $packageName

"@
    
    return $content
}

# 主流程
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Generating doc.go Files" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

$totalCount = 0
$successCount = 0

Get-ChildItem -Path $pkgDir -Directory | ForEach-Object {
    $dirName = $_.Name
    $docFile = Join-Path $_.FullName "doc.go"
    
    if ($apiMetadata.ContainsKey($dirName)) {
        $totalCount++
        
        # 跳过已有 doc.go 且状态为 completed 的
        if ((Test-Path $docFile) -and $apiMetadata[$dirName].Status -eq "completed") {
            Write-Host "[$totalCount] $dirName - Skipped (already exists)" -ForegroundColor Gray
            $successCount++
            return
        }
        
        $content = Generate-DocFile -Dir $dirName -Metadata $apiMetadata[$dirName]
        Set-Content -Path $docFile -Value $content -Encoding UTF8
        
        Write-Host "[$totalCount] $dirName - Generated" -ForegroundColor Green
        $successCount++
    }
    else {
        Write-Host "[?] $dirName - No metadata (skipped)" -ForegroundColor Yellow
    }
}

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Complete!" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Generated: $successCount / $totalCount" -ForegroundColor Green
Write-Host ""

