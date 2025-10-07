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

// Package sponsoreddisplayv3 提供 Amazon Advertising API Sponsored Display v3 客户端。
//
// # 概述
//
// Sponsored Display v3 API 用于管理展示型推广广告。
// 展示型推广广告可以在 Amazon 站内和站外展示，帮助触达更广泛的受众。
//
// 主要功能包括：
//   - 广告活动管理 (Campaigns) - 创建、更新、列表、删除展示型推广广告活动
//   - 广告组管理 (Ad Groups) - 管理广告活动下的广告组
//   - 定向管理 (Targeting) - 受众定向、商品定向、兴趣定向
//   - 否定定向 (Negative Targeting) - 排除特定受众或商品
//   - 创意管理 (Creatives) - 管理广告创意素材
//   - 预算管理 (Budget Rules) - 自动预算调整规则
//   - 优化规则 (Optimization Rules) - 自动化广告优化
//   - 推荐服务 (Recommendations) - 受众、出价推荐
//
// # 官方文档
//
// API 文档: https://advertising.amazon.com/API/docs/en-us/sponsored-display/3-0/openapi
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
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi/sponsored-display-v3"
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/internal/models"
//	)
//
//	func main() {
//	    // 创建基础客户端
//	    baseClient, err := adsapi.NewClient(
//	        adsapi.WithRegion(models.RegionNorthAmerica),
//	        adsapi.WithCredentials("client-id", "client-secret", "refresh-token"),
//	        adsapi.WithProfileID(1234567890),
//	    )
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    // 创建 Sponsored Display API 配置
//	    cfg := sponsoreddisplayv3.NewConfiguration()
//	    cfg.BasePath = models.RegionNorthAmerica.Endpoint
//
//	    // 创建 API 客户端
//	    apiClient := sponsoreddisplayv3.NewAPIClient(cfg)
//
//	    // 列出广告活动
//	    ctx := context.Background()
//	    campaigns, resp, err := apiClient.CampaignsApi.ListCampaigns(ctx, nil)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    fmt.Printf("找到 %d 个展示型推广广告活动\n", len(campaigns))
//	}
//
// # 定向策略
//
// Sponsored Display 支持多种定向方式：
//   - 受众定向 (Audience Targeting) - 基于购物行为、兴趣等
//   - 商品定向 (Product Targeting) - 基于 ASIN、分类等
//   - 再营销 (Remarketing) - 触达浏览过你商品的用户
//
// # 广告位置
//
// 展示型推广广告可以在以下位置展示：
//   - Amazon 商品详情页
//   - Amazon 搜索结果页
//   - 第三方网站和应用（站外投放）
//
// # 认证与授权
//
// Sponsored Display API 使用 Amazon LWA OAuth 2.0 进行认证。
// 所有请求都需要 Profile ID。
//
// # 速率限制
//
// Sponsored Display API 有以下速率限制：
//   - 大多数端点: 2 请求/秒
//   - 批量操作: 0.5 请求/秒
//
// # 代码生成说明
//
// 本包中除 doc.go 外的所有代码均由 swagger-codegen 自动生成。
// 不要手动修改自动生成的文件。
//
// # 相关包
//
//   - github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi - 基础客户端
//   - github.com/vanling1111/amazon-ads-api-go-sdk/internal/auth - LWA 认证
//
// # 版本历史
//
//   - v0.1.0 (2025-10-07) - 初始版本，基于 Sponsored Display API v3
package sponsoreddisplayv3

