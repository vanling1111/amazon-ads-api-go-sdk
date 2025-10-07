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

// Package sponsoredbrandsv4 提供 Amazon Advertising API Sponsored Brands v4 客户端。
//
// # 概述
//
// Sponsored Brands v4 API 用于管理品牌推广广告。
// 品牌推广广告可以展示品牌 Logo、自定义标题和多个商品，帮助提升品牌知名度。
//
// 主要功能包括：
//   - 广告活动管理 (Campaigns) - 创建、更新、列表、删除品牌推广广告活动
//   - 广告组管理 (Ad Groups) - 管理广告活动下的广告组
//   - 关键词管理 (Keywords) - 添加和管理广告关键词
//   - 否定关键词管理 (Negative Keywords) - 排除不相关的搜索词
//   - 创意管理 (Creatives) - 管理广告创意内容（图片、视频、标题等）
//   - 预算管理 (Budget Rules) - 自动预算调整规则
//   - 推荐服务 (Recommendations) - 关键词、出价、创意推荐
//
// # 官方文档
//
// API 文档: https://advertising.amazon.com/API/docs/en-us/sponsored-brands/4-0/openapi
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
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi/sponsored-brands-v4"
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
//	    // 创建 Sponsored Brands API 配置
//	    cfg := sponsoredbrandsv4.NewConfiguration()
//	    cfg.BasePath = models.RegionNorthAmerica.Endpoint
//
//	    // 创建 API 客户端
//	    apiClient := sponsoredbrandsv4.NewAPIClient(cfg)
//
//	    // 列出广告活动
//	    ctx := context.Background()
//	    campaigns, resp, err := apiClient.CampaignsApi.ListCampaigns(ctx, nil)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    fmt.Printf("找到 %d 个品牌推广广告活动\n", len(campaigns))
//	}
//
// # 品牌推广广告类型
//
// Sponsored Brands 支持多种广告格式：
//   - Product Collection - 商品集合广告（展示品牌 Logo + 多个商品）
//   - Store Spotlight - 店铺焦点广告（引导到品牌旗舰店）
//   - Video - 视频广告（展示品牌视频）
//
// # 认证与授权
//
// Sponsored Brands API 使用 Amazon LWA OAuth 2.0 进行认证。
// 所有请求都需要 Profile ID。
//
// # 速率限制
//
// Sponsored Brands API 有以下速率限制：
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
//   - v0.1.0 (2025-10-07) - 初始版本，基于 Sponsored Brands API v4
package sponsoredbrandsv4
