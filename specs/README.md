# Amazon Ads API OpenAPI 规范文件

本目录包含从 Amazon Ads API 官方下载的 OpenAPI 规范文件。

## 📋 文件清单（共 40 个 API 规范）

### 🎯 统一 API（新一代）
| 文件名 | API | 版本 | 优先级 | 说明 |
|--------|-----|------|--------|------|
| `amazon-ads-v1-campaign-management.json` | Amazon Ads API v1 | v1 | ⭐⭐⭐⭐⭐ | **统一 Campaign Management API**（跨 SP/SB/DSP） |

### 核心账户管理
| 文件名 | API | 版本 | 优先级 | 说明 |
|--------|-----|------|--------|------|
| `profiles-v3.yaml` | Profiles | v3 | ⭐⭐⭐⭐⭐ | **用户配置管理**（必需） |
| `portfolios-v2.json` | Portfolios | v2 | ⭐⭐⭐⭐ | 投资组合管理 |
| `manager-accounts.json` | Manager Accounts | - | ⭐⭐⭐ | 管理账户 |
| `advertising-accounts.json` | Advertising Accounts | - | ⭐⭐⭐ | 广告账户 |
| `billing.json` | Billing | - | ⭐⭐ | 账单管理 |
| `account-budgets.json` | Account Budgets | - | ⭐⭐ | 平均每日预算 |
| `test-accounts.json` | Test Accounts | - | ⭐ | 测试账户 |

### 核心广告产品
| 文件名 | API | 版本 | 优先级 | 说明 |
|--------|-----|------|--------|------|
| `sponsored-products-v3.json` | Sponsored Products | v3 | ⭐⭐⭐⭐⭐ | **商品推广**（最核心） |
| `sponsored-brands-v4.json` | Sponsored Brands | v4 | ⭐⭐⭐⭐ | **品牌推广**（新版） |
| `sponsored-brands-v3.yaml` | Sponsored Brands | v3 | ⭐⭐⭐ | 品牌推广（旧版兼容） |
| `sponsored-display-v3.yaml` | Sponsored Display | v3 | ⭐⭐⭐⭐ | **展示型推广** |

### 核心报表与分析
| 文件名 | API | 版本 | 优先级 | 说明 |
|--------|-----|------|--------|------|
| `reporting-v3.json` | Asynchronous Reporting | v3 | ⭐⭐⭐⭐⭐ | **异步报表**（必需） |
| `brand-metrics.json` | Brand Metrics | - | ⭐⭐⭐ | 品牌指标 |
| `stores-analytics.json` | Stores Analytics | - | ⭐⭐ | 店铺分析 |

### 推荐与洞察
| 文件名 | API | 版本 | 优先级 | 说明 |
|--------|-----|------|--------|------|
| `audiences-discovery.json` | Audience Discovery | - | ⭐⭐⭐ | 受众发现 |

### 实时数据流
| 文件名 | API | 版本 | 优先级 | 说明 |
|--------|-----|------|--------|------|
| `marketing-stream.json` | Amazon Marketing Stream | - | ⭐⭐⭐⭐ | **实时数据流** |

### 数据导出与创意
| 文件名 | API | 版本 | 优先级 | 说明 |
|--------|-----|------|--------|------|
| `exports.json` | Exports | - | ⭐⭐⭐ | 数据导出 |
| `creative-assets.yaml` | Creative Assets | - | ⭐⭐ | 创意资产库 |

---

## 🎯 开发优先级

### 第一阶段（MVP - 最小可用产品）
1. **Profiles v3** - 用户配置管理（必需）
2. **Sponsored Products v3** - 商品推广（核心业务）
3. **Reporting v3** - 异步报表（核心数据）

### 第二阶段（核心功能）
4. **Sponsored Brands v4** - 品牌推广
5. **Sponsored Display v3** - 展示型推广
6. **Portfolios v2** - 投资组合管理

### 第三阶段（高级功能）
7. **Marketing Stream** - 实时数据流
8. **Audiences Discovery** - 受众发现
9. **Exports** - 数据导出

### 第四阶段（扩展功能）
10. **Brand Metrics** - 品牌指标
11. **Manager Accounts** - 管理账户
12. **Creative Assets** - 创意资产

---

## 📊 文件统计

- **总文件数**: 15
- **JSON 文件**: 9
- **YAML 文件**: 6
- **总大小**: ~2.6 MB

---

## 🔄 更新日志

- **2025-10-07**: 初始下载所有核心 API 的 OpenAPI 规范

---

## 📚 参考资料

- [OpenAPI 官方目录](https://advertising.amazon.com/API/docs/en-us/reference/openapi-download)
- [Amazon Ads API 文档](https://advertising.amazon.com/API/docs/en-us)

