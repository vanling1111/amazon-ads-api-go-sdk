# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.3.0] - 2025-11-13

### Added

#### 接口抽象层
- 新增 `interfaces.go` - 定义 Logger, MetricsCollector, Tracer 接口
- 新增 `noop.go` - 提供 No-Op 默认实现
- 新增 `fields.go` - 日志字段辅助函数

#### Facade设计模式
- 新增 `internal/core/facade.go` - 核心门面层
- 封装内部组件（auth, transport, ratelimit）
- 提供统一的访问接口

### Changed

#### 架构优化
- `Config.Logger` 从 `*zap.Logger` 改为 `Logger` 接口
- `Config.Metrics` 从 `*metrics.PrometheusMetrics` 改为 `MetricsCollector` 接口
- 移除对 zap 和 prometheus 的硬依赖
- 更新所有日志调用使用接口

### Benefits

- ✅ **零依赖友好** - No-Op 默认实现
- ✅ **更好的可扩展性** - 用户可以自定义实现
- ✅ **架构一致性** - 与 SP-API SDK 保持一致
- ✅ **符合Go最佳实践** - 面向接口编程

### 示例

```go
// 使用默认的 No-Op 实现（零依赖）
client := adsapi.NewClient(
    adsapi.WithRegion(models.RegionNA),
    adsapi.WithCredentials(...),
)

// 或使用自定义实现
client := adsapi.NewClient(
    adsapi.WithLogger(myLogger),
    adsapi.WithMetrics(myMetrics),
)
```

## [0.2.0] - 2025-10-07

### 🎉 Major Achievement: 100% API Coverage

This release completes the generation of **all 45 Amazon Ads APIs**, achieving 100% API coverage.

### Added
- ✅ **All 45 APIs generated** - Complete Amazon Ads API coverage
  - Core Advertising Products (4 APIs)
  - Account Management (4 APIs)
  - Reporting & Analytics (5 APIs)
  - DSP (5 APIs)
  - Audiences & Targeting (3 APIs)
  - Creative & Assets (4 APIs)
  - Product & Catalog (2 APIs)
  - Attribution & Measurement (2 APIs)
  - Finance & Budget (2 APIs)
  - Data Management (4 APIs)
  - Optimization, Audit, Partners, Testing (4 APIs)
  - Unified API (1 API)
  - Retail Ad Service (2 APIs)
  - Deprecated/Older Versions (3 APIs)

- ✅ **One-click generation script** - `scripts/generate-all-apis.ps1`
  - Automated code generation for all APIs
  - Built-in auto-fix for known issues
  - Progress tracking and error handling

- ✅ **Auto-fix for generated code**
  - Creative Assets: Union types → `interface{}`
  - Billing: Missing `CountryCode` type definition
  - Amazon Ads v1: Invalid identifier names (numbers at start)

- ✅ **Complete package documentation**
  - `doc.go` files for all 45 APIs
  - License headers
  - Usage examples
  - API structure documentation

- ✅ **Additional API specifications downloaded**
  - Sponsored TV (Open Beta)
  - Ad Library
  - Posts (Deprecated)
  - Retail Ad Service (2 APIs)
  - Total: 45 OpenAPI files (38 JSON + 7 YAML)

### Changed
- 📝 Updated `README.md` to reflect 100% API coverage
- 📝 Updated `docs/API_COVERAGE.md` with all 45 APIs
- 📦 Project status upgraded from Alpha to Beta

### Technical Details
- **Total APIs**: 45/45 (100%)
- **Generated Files**: 8,449 files
- **OpenAPI Specifications**: 45 files (38 JSON + 7 YAML)
- **Code Generation Tool**: swagger-codegen-cli-3.0.62
- **Compilation Status**: ✅ All pass

---

## [0.1.0] - 2025-10-07

### Initial Release

### Added
- ✅ **6 Core APIs implemented**
  - Profiles v3
  - Sponsored Products v3
  - Reporting v3
  - Sponsored Brands v4
  - Sponsored Display v3
  - Portfolios v2

- ✅ **Core SDK modules**
  - `internal/auth` - LWA OAuth 2.0 authentication
  - `internal/transport` - HTTP transport with retry logic
  - `internal/ratelimit` - Rate limit management
  - `internal/metrics` - Prometheus metrics
  - `internal/signer` - Request signing
  - `internal/models` - Region and common models
  - `internal/utils` - Utility functions

- ✅ **High test coverage**
  - Core modules: 90%+ coverage
  - `internal/auth`: 89.2%
  - `internal/models`: 100%
  - `internal/metrics`: 100%
  - `internal/ratelimit`: 100%
  - `internal/transport`: 100%
  - `pkg/adsapi`: 76.7%

- ✅ **Comprehensive documentation**
  - Package-level `doc.go` for all core modules
  - Architecture documentation
  - Development guidelines
  - API coverage tracking

- ✅ **Usage examples**
  - Basic usage example
  - Advanced usage example
  - Integration examples

- ✅ **Build and development tools**
  - Code generation scripts
  - Test coverage scripts
  - Benchmark tools
  - Performance monitoring

### Technical Details
- **Core APIs**: 6/6 (100%)
- **Overall APIs**: 6/41 (15%)
- **OpenAPI Specifications**: 41 files downloaded
- **Test Coverage**: 90%+ for core modules
- **Go Version**: 1.21+

---

## Version History

- [0.2.0] - 2025-10-07 - 100% API Coverage Achievement
- [0.1.0] - 2025-10-07 - Initial Release with 6 Core APIs

---

**Note**: This project uses dual licensing:
- AGPL-3.0 for open source projects
- Commercial License for proprietary use

For licensing inquiries, contact: vanling1111@gmail.com
