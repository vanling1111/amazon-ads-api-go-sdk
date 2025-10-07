# Amazon Ads API Go SDK - Examples

示例代码集合，展示如何使用 SDK 的各种功能。

## 📁 示例列表

### 1. [basic_usage](basic_usage/) - 基础使用

最简单的入门示例，展示如何：
- 配置客户端
- 进行认证
- 发起基本请求

**运行**:
```bash
export ADS_CLIENT_ID="amzn1.application-oa2-client.xxx"
export ADS_CLIENT_SECRET="your-secret"
export ADS_REFRESH_TOKEN="Atzr|xxx"
export ADS_PROFILE_ID="123456789"

go run examples/basic_usage/main.go
```

### 2. [campaign_management](campaign_management/) - 广告活动管理

展示如何管理 Sponsored Products 广告活动：
- 创建广告活动
- 列出广告活动
- 更新广告活动
- 删除广告活动

**运行**:
```bash
go run examples/campaign_management/main.go
```

### 3. [reporting](reporting/) - 报告生成

展示如何使用 Reporting API：
- 创建报告
- 查询报告状态
- 下载报告数据
- 解析报告内容

**运行**:
```bash
go run examples/reporting/main.go
```

### 4. [advanced_usage](advanced_usage/) - 高级功能

展示高级功能：
- 批量操作
- 错误处理
- 重试策略
- 速率限制管理

**运行**:
```bash
go run examples/advanced_usage/main.go
```

## 🔑 环境变量

所有示例都需要以下环境变量：

| 变量 | 说明 | 示例 |
|------|------|------|
| `ADS_CLIENT_ID` | LWA 客户端 ID | `amzn1.application-oa2-client.xxx` |
| `ADS_CLIENT_SECRET` | LWA 客户端密钥 | `your-secret-key` |
| `ADS_REFRESH_TOKEN` | LWA 刷新令牌 | `Atzr|xxx` |
| `ADS_PROFILE_ID` | Advertising Profile ID | `123456789` |
| `ADS_REGION` | API 区域 (可选) | `NA` (默认), `EU`, `FE` |

**获取这些凭证**:

1. 访问 [Amazon Advertising Console](https://advertising.amazon.com/)
2. 创建 LWA 应用程序
3. 获取 Client ID 和 Client Secret
4. 生成 Refresh Token

详见：[认证指南](../docs/AUTHENTICATION.md)

## 📚 更多资源

- [API 文档](https://pkg.go.dev/github.com/vanling1111/amazon-ads-api-go-sdk)
- [架构设计](../docs/ARCHITECTURE.md)
- [开发指南](../docs/DEVELOPMENT.md)
- [官方文档](https://advertising.amazon.com/API/docs/)

## 🤝 贡献

欢迎贡献新的示例！请参考 [贡献指南](../CONTRIBUTING.md)。

## ⚖️ 许可证

所有示例代码同样遵循项目的双许可证模式（AGPL-3.0 或 Commercial License）。

