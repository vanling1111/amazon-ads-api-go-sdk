# Amazon Advertising API Go SDK

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-AGPL--3.0%20%7C%20Commercial-blue.svg)](LICENSE)
[![Status](https://img.shields.io/badge/status-beta-green.svg)](https://github.com/vanling1111/amazon-ads-api-go-sdk)

> **Official Documentation**: https://advertising.amazon.com/API/docs  
> **Project Status**: Beta - All 45 APIs Implemented  
> **Current Version**: v0.2.0



**Amazon Advertising API 的非官方 Go SDK，提供完整的类型安全 API 封装，支持所有广告产品。**

**最新更新**:
- ✅ **所有 45 个 API 已生成**（100% API 覆盖）
- ✅ 自动修复生成代码的已知问题
- ✅ 所有 API 包含 `doc.go` 文档
- ✅ 完整的认证、传输、速率限制模块
- ✅ 核心模块单元测试覆盖率 90%+
- ✅ 一键生成脚本

---

## ✨ 特性

- ✅ **完整的 API 覆盖** - 支持所有主要的 Amazon Ads API
- ✅ **类型安全** - 完整的 Go 类型定义和编译时检查
- ✅ **生产级质量** - 90%+ 测试覆盖率，完善的错误处理
- ✅ **自动认证** - 内置 LWA OAuth 2.0 认证和 Token 刷新
- ✅ **速率限制** - 智能速率限制管理，避免 API 限流
- ✅ **重试机制** - 自动重试和熔断器保护
- ✅ **官方监控** - 自动监控 API 规范变更
- ✅ **丰富示例** - 完整的使用示例和最佳实践

---

## 📦 支持的 API

### ✅ 已实现（45 个 API - 100% 覆盖）

| 类别 | 数量 | 状态 |
|------|------|------|
| 核心广告产品 | 4 | ✅ |
| 账户管理 | 4 | ✅ |
| 报告与分析 | 5 | ✅ |
| DSP | 5 | ✅ |
| 受众与定向 | 3 | ✅ |
| 创意与素材 | 4 | ✅ |
| 产品与目录 | 2 | ✅ |
| 归因与测量 | 2 | ✅ |
| 财务与预算 | 2 | ✅ |
| 数据管理 | 4 | ✅ |
| 其他 (优化/审计/合作伙伴/测试等) | 4 | ✅ |
| 统一 API | 1 | ✅ |
| Retail Ad Service | 2 | ✅ |
| 已弃用/旧版本 | 3 | ✅ |

<details>
<summary>点击查看完整 API 列表（45 个）</summary>

#### 核心广告产品
- ✅ Sponsored Products v3
- ✅ Sponsored Brands v4
- ✅ Sponsored Display v3
- ✅ Sponsored TV

#### 账户管理
- ✅ Profiles v3
- ✅ Portfolios v2
- ✅ Advertising Accounts
- ✅ Manager Accounts

#### 报告与分析
- ✅ Reporting v3
- ✅ Brand Metrics
- ✅ Insights
- ✅ Stores Analytics
- ✅ Marketing Mix Modeling

#### DSP
- ✅ DSP Audiences
- ✅ DSP Conversions
- ✅ DSP Measurement
- ✅ DSP Target KPI
- ✅ DSP Advertisers

#### 受众与定向
- ✅ Audiences Discovery
- ✅ Persona Builder
- ✅ Locations

#### 创意与素材
- ✅ Creative Assets
- ✅ Moderation
- ✅ Pre-Moderation
- ✅ Ad Library

#### 产品与目录
- ✅ Product Metadata
- ✅ Product Eligibility

#### 归因与测量
- ✅ Amazon Attribution
- ✅ Reach Forecasting

#### 财务与预算
- ✅ Billing
- ✅ Account Budgets

#### 数据管理
- ✅ Exports
- ✅ Marketing Stream
- ✅ Hashed Records
- ✅ Data Provider

#### 其他
- ✅ Tactical Recommendations
- ✅ Change History
- ✅ Partner Opportunities
- ✅ Test Accounts
- ✅ Amazon Ads API v1 (Unified)
- ✅ Retail Ad Service (2)
- ✅ Deprecated/Older (3)

完整列表请查看 [docs/API_COVERAGE.md](docs/API_COVERAGE.md)

</details>

---

**已下载规范**: 45 个 OpenAPI 文件（38 JSON + 7 YAML）  
**已实现 API**: 45/45 (100% ✅)  
**生成文件数**: 8,449 个  
**代码生成**: swagger-codegen-cli-3.0.62.jar  
**自动修复**: Creative Assets, Billing, Amazon Ads v1  
**编译状态**: ✅ 全部通过

---

## 🚀 快速开始

### 安装

```bash
go get github.com/vanling1111/amazon-ads-api-go-sdk
```

### 基础使用

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi"
    sp "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi/sponsored-products-v3"
)

func main() {
    // 1. 创建客户端配置
    cfg := &adsapi.Config{
        ClientID:     "amzn1.application-oa2-client.xxx",
        ClientSecret: "your-client-secret",
        RefreshToken: "Atzr|xxx",
        ProfileID:    "123456789",
        Region:       "NA",
    }

    // 2. 创建主客户端
    client, err := adsapi.NewClient(cfg)
    if err != nil {
        log.Fatal(err)
    }

    // 3. 创建 Sponsored Products API 客户端配置
    spConfig := sp.NewConfiguration()
    spConfig.HTTPClient = client.GetHTTPClient()
    spConfig.BasePath = "https://advertising-api.amazon.com"

    // 4. 创建 Sponsored Products API 客户端
    spClient := sp.NewAPIClient(spConfig)

    // 5. 调用 API（示例：列出广告活动）
    // 注意：实际使用时需要根据生成的 API 方法调用
    // campaigns, response, err := spClient.CampaignsApi.ListCampaigns(context.Background(), ...)
    // if err != nil {
    //     log.Fatal(err)
    // }

    fmt.Println("Client initialized successfully!")
}
```

**注意**: 
- 实际 API 调用方法请参考各 API 包的 `doc.go` 文件
- swagger-codegen 生成的客户端使用方式与手动封装略有不同
- 查看 `pkg/adsapi/sponsored-products-v3/doc.go` 获取完整示例

---

## 📚 文档

- **[架构设计](docs/ARCHITECTURE.md)** - 项目架构和设计原则
- **[API 覆盖](docs/API_COVERAGE.md)** - 完整的 API 覆盖清单
- **[开发指南](docs/DEVELOPMENT.md)** - 贡献代码指南

### API 包文档

所有 API 的详细文档都在各自的 `doc.go` 文件中：

- **Profiles API v3**: `pkg/adsapi/profiles-v3/doc.go`
- **Portfolios API v2**: `pkg/adsapi/portfolios-v2/doc.go`
- **Sponsored Products API v3**: `pkg/adsapi/sponsored-products-v3/doc.go`
- **Sponsored Brands API v4**: `pkg/adsapi/sponsored-brands-v4/doc.go`
- **Sponsored Display API v3**: `pkg/adsapi/sponsored-display-v3/doc.go`
- **Reporting API v3**: `pkg/adsapi/reporting-v3/doc.go`

### 核心模块文档

- **认证**: `internal/auth/` - LWA OAuth 2.0
- **传输**: `internal/transport/` - HTTP 客户端、重试
- **速率限制**: `internal/ratelimit/` - 令牌桶算法
- **指标**: `internal/metrics/` - Prometheus 监控

---

## 🎯 使用场景

### ERP 系统集成

详细示例请参考各 API 包的 `doc.go` 文件：

- **Profiles API**: `pkg/adsapi/profiles-v3/doc.go`
- **Sponsored Products API**: `pkg/adsapi/sponsored-products-v3/doc.go`
- **Reporting API**: `pkg/adsapi/reporting-v3/doc.go`

### 生成的代码结构

所有 API 客户端都使用 `swagger-codegen` 从官方 OpenAPI 规范生成，包含：

- **APIClient**: 主客户端结构
- **Configuration**: 配置管理
- **API Services**: 各功能模块的 API 服务（`api_*.go`）
- **Models**: 数据模型定义（`model_*.go`）

### 代码生成

如需重新生成代码：

```powershell
# 批量生成所有 API
.\scripts\generate-all-apis-batch.ps1

# 单个 API 生成示例
java -jar swagger-codegen-cli-3.0.62.jar generate `
  -i specs/sponsored-products-v3.json `
  -l go `
  -o pkg/adsapi/sponsored-products-v3 `
  --additional-properties packageName=sponsoredproductsv3
```

---

## 🔧 高级功能

### 认证管理

```go
// 内置 LWA OAuth 2.0 自动认证
// Token 自动刷新，无需手动管理
cfg := &adsapi.Config{
    ClientID:     "your-client-id",
    ClientSecret: "your-client-secret",
    RefreshToken: "your-refresh-token",
}
```

### 速率限制

```go
// 使用令牌桶算法自动管理速率限制
// 避免触发 API 限流
import "github.com/vanling1111/amazon-ads-api-go-sdk/internal/ratelimit"

limiter := ratelimit.NewManager(10.0, 100) // 10 请求/秒，桶容量 100
```

### 指标监控

```go
// Prometheus 指标收集
import "github.com/vanling1111/amazon-ads-api-go-sdk/internal/metrics"

metrics := metrics.NewPrometheusMetrics()
metrics.RecordRequest("sponsored-products", "ListCampaigns", 200, 0.5)
```

### HTTP 客户端自定义

```go
// 自定义 HTTP 客户端（重试、超时等）
import "github.com/vanling1111/amazon-ads-api-go-sdk/internal/transport"

httpClient := transport.NewHTTPClient(&transport.Config{
    MaxRetries:    3,
    Timeout:       30 * time.Second,
    EnableMetrics: true,
})
```

---

## 🔍 API 监控系统

SDK 内置自动监控工具，实时跟踪 Amazon Ads API 的变更。

### 监控功能

- ✅ **全面覆盖** - 监控所有 45 个官方 API
- ✅ **自动检测** - API 新增、删除、版本更新
- ✅ **SHA256 校验** - 精确检测规范文件变更
- ✅ **GitHub 集成** - 自动创建 Issue 通知
- ✅ **定时运行** - 每周自动 + 手动触发

### 本地运行

```bash
cd tools/api-monitor
go run main.go
```

### GitHub Actions 自动监控

- **自动运行**: 每周一 00:00 UTC
- **手动触发**: Actions 页面手动运行
- **变更通知**: 自动创建 Issue

### 监控报告示例

```
🔍 Starting Amazon Ads API Monitor...
📊 Monitoring 45 APIs in total

🆕 New APIs:
  • New Feature API (v1.0)

🔄 Updated APIs:
  • Sponsored Products v3: 3.0 → 3.1
  • Reporting v3: 2.5 → 3.0
```

详细文档：[tools/api-monitor/README.md](tools/api-monitor/README.md)

---

## 🔗 相关项目

- **[Amazon SP-API Go SDK](https://github.com/vanling1111/amazon-sp-api-go-sdk)** - 本项目的姊妹项目，提供 Selling Partner API 封装

---

## 📄 许可证

**双许可证模式**:

### 1️⃣ AGPL-3.0 (开源项目)

免费用于：
- ✅ 个人项目
- ✅ 教育用途
- ✅ 开源项目（必须同样开源）

### 2️⃣ Commercial License (商业项目)

需购买商业许可证：
- 🏢 企业内部使用
- 🔒 闭源商业产品
- 💼 SaaS 服务

**购买商业许可证**: 联系 vanling1111@gmail.com

---

## 🤝 贡献

欢迎贡献！请查看 [CONTRIBUTING.md](CONTRIBUTING.md)

**贡献者要求**:
- ✅ 遵循 Go 官方代码规范
- ✅ 提供单元测试（覆盖率 90%+）
- ✅ 使用 Google 风格注释（中文）
- ✅ 通过所有 CI 检查

---

## 📞 支持

- **文档**: https://github.com/vanling1111/amazon-ads-api-go-sdk/docs
- **Issues**: https://github.com/vanling1111/amazon-ads-api-go-sdk/issues
- **邮箱**: vanling1111@gmail.com

---

## 🙏 致谢

- 参考了 [Amazon SP-API Go SDK](https://github.com/vanling1111/amazon-sp-api-go-sdk) 的架构设计
- 感谢 Amazon Advertising API 官方文档

---

**项目状态**: ✅ Alpha 版本  
**当前版本**: v0.2.0 (2025-10-07)  
**API 覆盖**: 45/45 (100%) ✅  
**已下载规范**: 45 个 OpenAPI 文件 (38 JSON + 7 YAML)  
**生成文件数**: 8,449 个  
**核心模块测试覆盖**: 90%+  
**维护者**: [@vanling1111](https://github.com/vanling1111)

