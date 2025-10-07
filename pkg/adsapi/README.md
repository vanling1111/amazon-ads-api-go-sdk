# Amazon Ads API Go SDK - 主客户端

本包提供 Amazon Advertising API 的主客户端实现。

## 功能特性

- 🔐 **自动认证** - LWA OAuth 2.0 自动令牌管理
- 🚀 **速率限制** - 自动速率限制管理，防止 API 限流
- 🔄 **自动重试** - 智能重试机制，处理临时性错误
- 📊 **指标监控** - 可选的 Prometheus 指标集成
- 🛡️ **类型安全** - 完整的类型定义和参数验证

## 快速开始

```go
package main

import (
    "context"
    "log"

    "github.com/vanling1111/amazon-ads-api-go-sdk/internal/models"
    "github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi"
)

func main() {
    // 创建客户端
    client, err := adsapi.NewClient(
        adsapi.WithRegion(models.RegionNorthAmerica),
        adsapi.WithCredentials(
            "your-client-id",
            "your-client-secret",
            "your-refresh-token",
        ),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // 使用客户端...
}
```

## 配置选项

详见 `config.go` 文件。

