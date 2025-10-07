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

// Package profilesv3 提供 Amazon Advertising API Profiles v3 客户端。
//
// # 概述
//
// Profiles v3 API 用于管理广告主配置文件（Advertising Profiles）。
// Profile 是 Amazon Ads API 中的核心概念，代表一个广告主账户在特定市场的配置。
//
// 主要功能包括：
//   - Profile 管理 - 列出、获取、更新广告主配置文件
//   - 市场信息 - 获取可用的市场和区域信息
//   - 账户关联 - 管理 Profile 与广告账户的关联
//
// # 官方文档
//
// API 文档: https://advertising.amazon.com/API/docs/en-us/guides/get-started/profiles
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
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi/profiles-v3"
//	    "github.com/vanling1111/amazon-ads-api-go-sdk/internal/models"
//	)
//
//	func main() {
//	    // 创建基础客户端
//	    baseClient, err := adsapi.NewClient(
//	        adsapi.WithRegion(models.RegionNorthAmerica),
//	        adsapi.WithCredentials("client-id", "client-secret", "refresh-token"),
//	    )
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    // 创建 Profiles API 配置
//	    cfg := profilesv3.NewConfiguration()
//	    cfg.BasePath = models.RegionNorthAmerica.Endpoint
//
//	    // 创建 API 客户端
//	    apiClient := profilesv3.NewAPIClient(cfg)
//
//	    // 列出所有 Profiles
//	    // 注意：Profiles API 通常不需要 ProfileID header
//	    ctx := context.Background()
//	    profiles, resp, err := apiClient.ProfilesApi.ListProfiles(ctx, nil)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    fmt.Printf("找到 %d 个 Profiles\n", len(profiles))
//	}
//
// # 认证与授权
//
// Profiles API 使用 Amazon LWA OAuth 2.0 进行认证。
// 注意：Profiles API 是特殊的，因为它用于获取 Profile ID，
// 所以大多数 Profiles API 调用不需要 Amazon-Advertising-API-Scope header。
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
//   - v0.1.0 (2025-10-07) - 初始版本，基于 Profiles API v3
package profilesv3
