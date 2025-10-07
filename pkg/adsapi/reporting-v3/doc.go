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

// Package reportingv3 提供 Amazon Advertising API Reporting v3 客户端。
//
// # 概述
//
// Reporting v3 API 用于生成和下载广告数据报告。
// 支持各种维度的数据分析，包括广告活动、关键词、商品等级别的性能数据。
//
// 主要功能包括：
//   - 报告生成 (Report Creation) - 创建异步报告任务
//   - 报告下载 (Report Download) - 下载生成的报告数据
//   - 报告状态查询 (Report Status) - 查询报告生成进度
//   - 多种报告类型 - 支持 Sponsored Products、Sponsored Brands、Sponsored Display 报告
//
// # 官方文档
//
// API 文档: https://advertising.amazon.com/API/docs/en-us/reporting/v3/overview
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
//	    "time"
//
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi"
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi/reporting-v3"
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
//	    // 创建 Reporting API 配置
//	    cfg := reportingv3.NewConfiguration()
//	    cfg.BasePath = models.RegionNorthAmerica.Endpoint
//
//	    // 创建 API 客户端
//	    apiClient := reportingv3.NewAPIClient(cfg)
//
//	    // 创建报告请求
//	    ctx := context.Background()
//	    reportRequest := reportingv3.ReportRequest{
//	        ReportDate: "20250101",
//	        Metrics:    []string{"impressions", "clicks", "cost"},
//	    }
//
//	    report, resp, err := apiClient.ReportsApi.CreateReport(ctx, reportRequest, nil)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    fmt.Printf("报告 ID: %s\n", report.ReportId)
//	}
//
// # 报告生成流程
//
// 1. 创建报告请求 - 指定报告类型、日期范围、指标等
// 2. 等待报告生成 - 报告异步生成，通常需要几分钟
// 3. 查询报告状态 - 使用 reportId 查询生成进度
// 4. 下载报告数据 - 报告生成完成后下载 JSON/CSV 文件
//
// # 报告类型
//
// 支持以下报告类型：
//   - Campaign Report - 广告活动级别报告
//   - Ad Group Report - 广告组级别报告
//   - Keyword Report - 关键词级别报告
//   - Product Ad Report - 商品广告级别报告
//   - Search Term Report - 搜索词报告
//
// # 指标说明
//
// 常用指标包括：
//   - impressions - 展示次数
//   - clicks - 点击次数
//   - cost - 花费金额
//   - sales - 销售额
//   - conversions - 转化次数
//   - acos - 广告成本销售比
//
// # 认证与授权
//
// Reporting API 使用 Amazon LWA OAuth 2.0 进行认证。
// 所有请求都需要 Profile ID。
//
// # 速率限制
//
// Reporting API 有以下速率限制：
//   - 报告创建: 0.1 请求/秒
//   - 报告查询: 2 请求/秒
//
// # 最佳实践
//
// 1. **异步处理** - 报告生成是异步的，使用轮询或 webhook 获取结果
// 2. **批量下载** - 尽可能批量下载报告，减少 API 调用
// 3. **缓存结果** - 缓存已下载的报告，避免重复请求
// 4. **错误重试** - 报告生成可能失败，需要实现重试机制
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
//   - v0.1.0 (2025-10-07) - 初始版本，基于 Reporting API v3
package reportingv3

