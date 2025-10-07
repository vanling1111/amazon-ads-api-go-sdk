# 开发指南

> **版本**: v0.1.0  
> **最后更新**: 2025-10-07

本文档描述 Amazon Ads API Go SDK 的开发规范和流程。

---

## 📋 目录

1. [开发环境](#开发环境)
2. [项目结构](#项目结构)
3. [代码规范](#代码规范)
4. [测试规范](#测试规范)
5. [文档规范](#文档规范)
6. [提交规范](#提交规范)
7. [发布流程](#发布流程)

---

## 🛠️ 开发环境

### 必需工具

| 工具             | 版本要求     | 用途                     |
|------------------|--------------|--------------------------|
| Go               | 1.21+        | 编程语言                 |
| Git              | 2.30+        | 版本控制                 |
| golangci-lint    | 1.54+        | 代码检查                 |
| oapi-codegen     | 2.0+         | OpenAPI 代码生成         |

### 安装

```bash
# 克隆仓库
git clone https://github.com/vanling1111/amazon-ads-api-go-sdk.git
cd amazon-ads-api-go-sdk

# 安装依赖
go mod download

# 安装开发工具
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

---

## 📁 项目结构

### 目录职责

```
amazon-ads-api-go-sdk/
├── pkg/adsapi/                # 公共 API（外部可导入）
│   ├── client.go              # 主客户端
│   ├── config.go              # 配置管理
│   └── {api-name}-{version}/  # 各 API 实现
│
├── internal/                  # 内部实现（外部不可导入）
│   ├── auth/                  # 认证模块（继承 SP-API）
│   ├── transport/             # HTTP 传输（继承 SP-API）
│   ├── ratelimit/             # 速率限制（ADS-API 特定）
│   ├── codegen/               # 自动生成代码
│   └── utils/                 # 工具函数
│
├── cmd/                       # 命令行工具
│   ├── api-monitor/           # API 监控工具
│   └── generator/             # 代码生成器
│
├── examples/                  # 示例代码
├── tests/                     # 集成测试
├── docs/                      # 文档
└── scripts/                   # 脚本工具
```

### 包命名规范

| 包类型           | 命名规则                     | 示例                              |
|------------------|------------------------------|-----------------------------------|
| API 客户端       | `{api-name}-{version}`       | `sponsored-products-v3`           |
| 内部模块         | 小写单数                     | `auth`, `transport`, `ratelimit`  |
| 生成代码         | `codegen/{api-name}`         | `codegen/sponsored-products-v3`   |

---

## 📝 代码规范

### Go 官方规范

**强制遵循**:
- ✅ [Effective Go](https://golang.org/doc/effective_go)
- ✅ [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- ✅ [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### 注释规范

**文件头注释**（强制）:

```go
// Copyright 2025 Amazon Ads API Go SDK Authors.
//
// Licensed under AGPL-3.0 or Commercial License.
// See LICENSE file for details.

// Package sponsored_products_v3 provides Amazon Sponsored Products API v3 client.
//
// 本包实现了 Amazon Sponsored Products API v3 的完整功能，包括：
//   - 广告活动管理 (Campaigns)
//   - 广告组管理 (Ad Groups)
//   - 关键词管理 (Keywords)
//   - 商品定向管理 (Product Targets)
//
// 官方文档: https://advertising.amazon.com/API/docs/en-us/sponsored-products/3-0/openapi
package sponsored_products_v3
```

**函数注释**（Google 风格，强制中文）:

```go
// CreateCampaigns 批量创建 Sponsored Products 广告活动。
//
// 参数:
//   - ctx: 请求上下文，用于超时控制和取消
//   - requests: 广告活动创建请求列表（最多 100 个）
//
// 返回:
//   - *CreateCampaignsResponse: 创建结果，包含成功和失败的活动
//   - error: 错误信息（网络错误、API 错误、参数校验错误）
//
// 错误类型:
//   - ErrInvalidRequest: 请求参数无效
//   - ErrRateLimitExceeded: 超过速率限制
//   - ErrUnauthorized: 认证失败
//
// 示例:
//
//	campaigns := []*CreateCampaignRequest{
//	    {Name: "Q4促销", Budget: 1000.0, TargetingType: "AUTO"},
//	}
//	resp, err := client.CreateCampaigns(ctx, campaigns)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) CreateCampaigns(ctx context.Context, requests []*CreateCampaignRequest) (*CreateCampaignsResponse, error) {
    // ...
}
```

### 错误处理

**自定义错误类型**:

```go
// pkg/adsapi/errors.go

// APIError 表示 Amazon Ads API 返回的错误
type APIError struct {
    Code       string `json:"code"`       // 错误代码
    Message    string `json:"message"`    // 错误消息
    StatusCode int    `json:"-"`          // HTTP 状态码
}

func (e *APIError) Error() string {
    return fmt.Sprintf("ads api error: %s (%s)", e.Message, e.Code)
}

// 预定义错误
var (
    ErrUnauthorized       = &APIError{Code: "UNAUTHORIZED", Message: "认证失败"}
    ErrRateLimitExceeded  = &APIError{Code: "RATE_LIMIT", Message: "超过速率限制"}
    ErrInvalidRequest     = &APIError{Code: "INVALID_REQUEST", Message: "请求参数无效"}
)
```

**错误处理示例**:

```go
func (c *Client) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("http request failed: %w", err)
    }

    // 检查 HTTP 状态码
    if resp.StatusCode >= 400 {
        return nil, parseAPIError(resp)
    }

    return resp, nil
}
```

### 类型定义

**优先使用值类型**:

```go
// ✅ 好的做法
type Campaign struct {
    CampaignID   int64   `json:"campaignId"`
    Name         string  `json:"name"`
    State        string  `json:"state"`
    Budget       float64 `json:"budget"`
}

// ❌ 避免
type Campaign struct {
    CampaignID   *int64   `json:"campaignId"`   // 不必要的指针
    Name         *string  `json:"name"`
    // ...
}
```

**可选字段使用指针**:

```go
// ✅ 正确：可选字段使用指针
type UpdateCampaignRequest struct {
    CampaignID int64    `json:"campaignId"`           // 必需
    Name       *string  `json:"name,omitempty"`       // 可选
    Budget     *float64 `json:"budget,omitempty"`     // 可选
}
```

---

## ✅ 测试规范

### 测试覆盖率要求

| 类型           | 覆盖率要求 | 说明                     |
|----------------|------------|--------------------------|
| 单元测试       | 90%+       | 所有公共函数必须测试     |
| 集成测试       | 核心流程   | 关键业务流程必须测试     |
| 基准测试       | 性能关键   | 性能敏感代码必须测试     |

### 单元测试

**文件命名**: `*_test.go`

**测试结构**:

```go
// pkg/adsapi/sponsored-products-v3/campaigns_test.go
package sponsored_products_v3

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestClient_CreateCampaigns(t *testing.T) {
    tests := []struct {
        name    string
        input   []*CreateCampaignRequest
        want    *CreateCampaignsResponse
        wantErr bool
    }{
        {
            name: "成功创建单个广告活动",
            input: []*CreateCampaignRequest{
                {Name: "Test Campaign", Budget: 100.0, TargetingType: "AUTO"},
            },
            want: &CreateCampaignsResponse{
                Campaigns: []*Campaign{{CampaignID: 123, Name: "Test Campaign"}},
            },
            wantErr: false,
        },
        {
            name: "请求参数为空",
            input: nil,
            want: nil,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Setup
            client := setupTestClient(t)
            ctx := context.Background()

            // Execute
            got, err := client.CreateCampaigns(ctx, tt.input)

            // Assert
            if tt.wantErr {
                require.Error(t, err)
                return
            }

            require.NoError(t, err)
            assert.Equal(t, tt.want, got)
        })
    }
}
```

### Mock 测试

**使用 httptest**:

```go
func setupTestClient(t *testing.T) *Client {
    t.Helper()

    // 创建 mock HTTP 服务器
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Mock 响应
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "campaigns": []map[string]interface{}{
                {"campaignId": 123, "name": "Test Campaign"},
            },
        })
    }))
    t.Cleanup(server.Close)

    // 创建测试客户端
    cfg := &Config{
        Endpoint: server.URL,
        // ...
    }
    client, err := NewClient(cfg)
    require.NoError(t, err)

    return client
}
```

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行指定包测试
go test ./pkg/adsapi/sponsored-products-v3

# 查看覆盖率
go test -cover ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

---

## 📚 文档规范

### 必需文档

| 文档                  | 位置                          | 说明                     |
|-----------------------|-------------------------------|--------------------------|
| README                | 根目录                        | 项目介绍和快速开始       |
| 架构设计              | `docs/ARCHITECTURE.md`        | 技术架构和设计决策       |
| API 覆盖              | `docs/API_COVERAGE.md`        | API 实现进度追踪         |
| 开发指南              | `docs/DEVELOPMENT.md`         | 本文档                   |
| 模块文档              | `docs/modules/{api}.md`       | 每个 API 的详细文档      |

### 模块文档模板

```markdown
# Sponsored Products API v3

## 概述

Sponsored Products API v3 提供...

## 快速开始

\`\`\`go
// 示例代码
\`\`\`

## API 方法

### CreateCampaigns

创建广告活动...

**参数**:
- ...

**返回**:
- ...

**示例**:
\`\`\`go
// 示例代码
\`\`\`

## 错误处理

...

## 最佳实践

...
```

---

## 📝 提交规范

### Commit Message 格式

遵循 [Conventional Commits](https://www.conventionalcommits.org/)：

```
<type>(<scope>): <subject>

<body>

<footer>
```

**类型**:
- `feat`: 新功能
- `fix`: Bug 修复
- `docs`: 文档变更
- `style`: 代码格式（不影响功能）
- `refactor`: 重构
- `test`: 测试相关
- `chore`: 构建/工具变更

**示例**:

```
feat(sp-v3): 实现 Campaigns 创建和列表功能

- 添加 CreateCampaigns 方法
- 添加 ListCampaigns 方法
- 完成单元测试（覆盖率 95%）

Closes #123
```

---

## 🚀 发布流程

### 版本号规范

遵循 [Semantic Versioning](https://semver.org/)：

- **MAJOR**: 不兼容的 API 变更
- **MINOR**: 向后兼容的功能新增
- **PATCH**: 向后兼容的 Bug 修复

### 发布步骤

1. **更新 CHANGELOG.md**

```markdown
## [0.4.0] - 2025-11-15

### Added
- Profiles API v2 完整实现
- Sponsored Products API v3 Campaigns 和 Ad Groups

### Changed
- 优化速率限制算法

### Fixed
- 修复认证 Token 刷新问题
```

2. **创建 Git Tag**

```bash
git tag -a v0.4.0 -m "Release v0.4.0"
git push origin v0.4.0
```

3. **GitHub Release**

通过 GitHub Actions 自动发布。

---

## 🔍 代码审查清单

### 提交 PR 前检查

- [ ] 代码符合 Go 官方规范
- [ ] 所有测试通过（`go test ./...`）
- [ ] 代码覆盖率 ≥ 90%
- [ ] golangci-lint 无警告（`golangci-lint run`）
- [ ] 添加了完整的中文注释
- [ ] 更新了相关文档
- [ ] Commit message 符合规范

---

**文档版本**: v0.1.0  
**最后更新**: 2025-10-07

