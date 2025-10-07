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

// Package portfoliosv2 提供 Amazon Advertising API Portfolios v2 客户端。
//
// # 概述
//
// Portfolios v2 API 用于管理广告组合（Portfolios）。
// Portfolio 是一组广告活动的集合，可以统一管理预算和查看性能数据。
//
// 主要功能包括：
//   - Portfolio 管理 - 创建、更新、列表、删除广告组合
//   - 预算管理 - 设置和管理组合级别的预算
//   - 广告活动关联 - 将广告活动添加到或从组合中移除
//   - 性能查看 - 查看组合级别的性能数据
//
// # 官方文档
//
// API 文档: https://advertising.amazon.com/API/docs/en-us/reference/portfolios/2-0/overview
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
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi/portfolios-v2"
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
//	    // 创建 Portfolios API 配置
//	    cfg := portfoliosv2.NewConfiguration()
//	    cfg.BasePath = models.RegionNorthAmerica.Endpoint
//
//	    // 创建 API 客户端
//	    apiClient := portfoliosv2.NewAPIClient(cfg)
//
//	    // 列出所有 Portfolios
//	    ctx := context.Background()
//	    portfolios, resp, err := apiClient.PortfoliosApi.ListPortfolios(ctx, nil)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    fmt.Printf("找到 %d 个广告组合\n", len(portfolios))
//	}
//
// # Portfolio 的作用
//
// 使用 Portfolio 可以：
//   - 按产品线、品牌或季节组织广告活动
//   - 统一管理多个广告活动的预算
//   - 快速查看组合级别的性能数据
//   - 简化广告活动的批量管理
//
// # 预算管理
//
// Portfolio 支持两种预算模式：
//   - Daily Budget - 每日预算限制
//   - Lifetime Budget - 总预算限制
//
// # 认证与授权
//
// Portfolios API 使用 Amazon LWA OAuth 2.0 进行认证。
// 所有请求都需要 Profile ID。
//
// # 速率限制
//
// Portfolios API 有以下速率限制：
//   - 大多数端点: 2 请求/秒
//
// # 最佳实践
//
// 1. **合理分组** - 按业务逻辑组织 Portfolio，便于管理
// 2. **预算监控** - 定期检查 Portfolio 的预算使用情况
// 3. **性能分析** - 利用 Portfolio 级别的数据进行性能对比
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
//   - v0.1.0 (2025-10-07) - 初始版本，基于 Portfolios API v2
package portfoliosv2
