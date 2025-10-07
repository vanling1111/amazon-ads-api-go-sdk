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

// Package sponsoredproductsv3 提供 Amazon Advertising API Sponsored Products v3 客户端。
//
// # 概述
//
// Sponsored Products v3 是 Amazon Advertising API 中最核心的 API，提供了完整的商品推广广告管理功能。
// 本包是基于官方 OpenAPI 规范通过 swagger-codegen 自动生成的 Go 客户端。
//
// 主要功能包括：
//   - 广告活动管理 (Campaigns) - 创建、更新、列表、删除广告活动
//   - 广告组管理 (Ad Groups) - 管理广告活动下的广告组
//   - 关键词管理 (Keywords) - 添加和管理广告关键词
//   - 否定关键词管理 (Negative Keywords) - 排除不相关的搜索词
//   - 商品广告管理 (Product Ads) - 管理推广的商品
//   - 商品定向管理 (Product Targeting) - 基于 ASIN 的定向
//   - 定向条款管理 (Targeting Clauses) - 自定义定向规则
//   - 预算管理 (Budget Rules) - 自动预算调整规则
//   - 优化规则 (Optimization Rules) - 自动化广告优化
//   - 推荐服务 (Recommendations) - 关键词、预算、出价推荐
//
// # 官方文档
//
// API 文档: https://advertising.amazon.com/API/docs/en-us/sponsored-products/3-0/openapi
//
// OpenAPI 规范: https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-products/openapi.json
//
// # 快速开始
//
// 基本用法示例：
//
//	package main
//
//	import (
//	    "context"
//	    "fmt"
//	    "log"
//
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi"
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi/sponsored-products-v3"
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/internal/models"
//	)
//
//	func main() {
//	    // 1. 创建基础客户端
//	    baseClient, err := adsapi.NewClient(
//	        adsapi.WithRegion(models.RegionNA),
//	        adsapi.WithCredentials("your-client-id", "your-client-secret", "your-refresh-token"),
//	        adsapi.WithProfileID(1234567890),
//	    )
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    // 2. 创建 Sponsored Products API 配置
//	    cfg := sponsoredproductsv3.NewConfiguration()
//	    cfg.HTTPClient = baseClient.HTTPClient() // 使用基础客户端的 HTTP 客户端
//	    cfg.Host = models.RegionNA.Endpoint      // 设置 API 端点
//
//	    // 3. 创建 API 客户端
//	    apiClient := sponsoredproductsv3.NewAPIClient(cfg)
//
//	    // 4. 调用 API - 列出广告活动
//	    ctx := context.Background()
//	    request := sponsoredproductsv3.SponsoredProductsListCampaignsRequestContent{
//	        MaxResults: 10,
//	    }
//
//	    campaigns, resp, err := apiClient.CampaignsApi.ListSponsoredProductsCampaigns(
//	        ctx,
//	        request,
//	        "your-client-id",
//	        "1234567890", // Profile ID
//	        nil,
//	    )
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    fmt.Printf("找到 %d 个广告活动\n", len(campaigns.Campaigns))
//	}
//
// # API 结构
//
// 本包提供以下 API 服务（所有服务都可以通过 APIClient 访问）：
//
//   - AdGroupsApi - 广告组管理
//   - CampaignsApi - 广告活动管理
//   - KeywordsApi - 关键词管理
//   - NegativeKeywordsApi - 否定关键词管理
//   - ProductAdsApi - 商品广告管理
//   - ProductTargetingApi - 商品定向管理
//   - TargetingClausesApi - 定向条款管理
//   - BudgetRulesApi - 预算规则管理
//   - OptimizationRulesApi - 优化规则管理
//   - ConsolidatedRecommendationsApi - 综合推荐服务
//   - 以及其他 14 个专门的 API 服务
//
// # 认证与授权
//
// Sponsored Products API v3 使用 Amazon LWA (Login with Amazon) OAuth 2.0 进行认证。
// 所有请求都需要在 Header 中提供：
//   - Authorization: Bearer {access_token}
//   - Amazon-Advertising-API-ClientId: {client_id}
//   - Amazon-Advertising-API-Scope: {profile_id}
//
// 基础客户端（adsapi.Client）会自动处理认证和 Token 刷新。
//
// # 速率限制
//
// Sponsored Products API v3 有以下速率限制：
//   - 大多数端点: 2 请求/秒
//   - 批量操作: 0.5 请求/秒
//   - 报告相关: 0.1 请求/秒
//
// 基础客户端会自动处理速率限制和重试逻辑。
//
// # 错误处理
//
// API 可能返回以下错误类型：
//
//   - 400 Bad Request - 请求参数无效
//   - 401 Unauthorized - 认证失败或 Token 过期
//   - 403 Forbidden - 没有权限访问该资源
//   - 404 Not Found - 资源不存在
//   - 429 Too Many Requests - 超过速率限制
//   - 500 Internal Server Error - 服务器内部错误
//
// 错误响应示例：
//
//	{
//	  "code": "INVALID_ARGUMENT",
//	  "details": "campaignId is required",
//	  "requestId": "ABC123XYZ"
//	}
//
// # 最佳实践
//
// 1. **批量操作** - 尽可能使用批量 API（如 CreateSponsoredProductsCampaigns）而不是逐个调用
//
// 2. **分页处理** - 使用 nextToken 进行分页，避免一次性获取大量数据：
//
//	var allCampaigns []Campaign
//	var nextToken *string
//	for {
//	    request := SponsoredProductsListCampaignsRequestContent{
//	        MaxResults: 100,
//	        NextToken: nextToken,
//	    }
//	    resp, _, err := client.CampaignsApi.ListSponsoredProductsCampaigns(ctx, request, clientID, profileID, nil)
//	    if err != nil {
//	        return err
//	    }
//	    allCampaigns = append(allCampaigns, resp.Campaigns...)
//	    if resp.NextToken == nil {
//	        break
//	    }
//	    nextToken = resp.NextToken
//	}
//
// 3. **错误重试** - 对于临时错误（429, 5xx），使用指数退避重试
//
// 4. **上下文超时** - 为所有请求设置合理的超时时间：
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
// 5. **日志记录** - 记录所有 API 调用以便于调试和监控
//
// # 代码生成说明
//
// 本包中除 doc.go 外的所有代码均由 swagger-codegen 3.0.64 自动生成。
// 生成命令：
//
//	swagger-codegen generate \
//	  -i specs/sponsored-products-v3.json \
//	  -l go \
//	  -o pkg/adsapi/sponsored-products-v3 \
//	  -DpackageName=sponsoredproductsv3
//
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
//   - v0.1.0 (2025-10-07) - 初始版本，基于 Sponsored Products API v3.0
//
// # 许可证
//
// 本包使用双许可证模式：
//   - AGPL-3.0 用于开源项目
//   - Commercial License 用于商业项目
//
// 详见项目根目录的 LICENSE 文件。
package sponsoredproductsv3

