# Amazon Ads API Coverage

> **Version**: v0.2.0  
> **Last Updated**: 2025-10-07

This document tracks the Amazon Ads API Go SDK coverage of official APIs.

---

## 📊 Overall Coverage

### Downloaded API Specifications (specs/ directory)

**Total**: 45 OpenAPI specification files (38 JSON + 7 YAML)

### Implemented APIs

**Total Implemented**: 45/45 (100%)  
**Core API Coverage**: 6/6 (100%)  
**Overall Coverage**: 45/45 (100%)

---

## ✅ All Implemented APIs (45)

### Core Advertising Products (4)
- ✅ **Sponsored Products v3** - `sponsored-products-v3.json` - Product advertising (Core)
- ✅ **Sponsored Brands v4** - `sponsored-brands-v4.json` - Brand advertising
- ✅ **Sponsored Display v3** - `sponsored-display-v3.yaml` - Display advertising
- ✅ **Sponsored TV** - `sponsored-tv.json` - Streaming TV advertising (Open Beta)

### Core Account Management (4)
- ✅ **Profiles v3** - `profiles-v3.yaml` - Profile management (Required)
- ✅ **Portfolios v2** - `portfolios-v2.json` - Portfolio management
- ✅ **Advertising Accounts** - `advertising-accounts.json` - Ad account management
- ✅ **Manager Accounts** - `manager-accounts.json` - Manager account features

### Reporting & Analytics (5)
- ✅ **Reporting v3** - `reporting-v3.json` - Async reporting (Core)
- ✅ **Brand Metrics** - `brand-metrics.json` - Brand awareness metrics
- ✅ **Insights** - `insights.json` - Search insights
- ✅ **Stores Analytics** - `stores-analytics.json` - Store analytics
- ✅ **Marketing Mix Modeling** - `marketing-mix-modeling.json` - Attribution modeling

### DSP APIs (5)
- ✅ **DSP Audiences** - `dsp-audiences.json` - Audience management
- ✅ **DSP Conversions** - `dsp-conversions.json` - Conversion tracking
- ✅ **DSP Measurement** - `dsp-measurement.json` - Brand lift studies
- ✅ **DSP Target KPI** - `dsp-target-kpi.json` - KPI optimization
- ✅ **DSP Advertisers** - `dsp-advertisers.yaml` - DSP advertiser management

### Audiences & Targeting (3)
- ✅ **Audiences Discovery** - `audiences-discovery.json` - Audience insights
- ✅ **Persona Builder** - `persona-builder.json` - Persona construction
- ✅ **Locations** - `locations.json` - Geographic targeting

### Creative & Assets (4)
- ✅ **Creative Assets** - `creative-assets.yaml` - Asset management
- ✅ **Moderation** - `moderation.json` - Ad moderation
- ✅ **Pre-Moderation** - `pre-moderation.json` - Pre-approval checks
- ✅ **Ad Library** - `ad-library.json` - Ad library management

### Product & Catalog (2)
- ✅ **Product Metadata** - `product-metadata.json` - Product information
- ✅ **Product Eligibility** - `product-eligibility.json` - Eligibility checks

### Attribution & Measurement (2)
- ✅ **Amazon Attribution** - `amazon-attribution.json` - Cross-channel attribution
- ✅ **Reach Forecasting** - `reach-forecasting.json` - Reach estimation

### Finance & Budget (2)
- ✅ **Billing** - `billing.json` - Billing management
- ✅ **Account Budgets** - `account-budgets.json` - Budget management

### Data Management (4)
- ✅ **Exports** - `exports.json` - Data export
- ✅ **Marketing Stream** - `marketing-stream.json` - Real-time data streaming
- ✅ **Hashed Records** - `hashed-records.json` - Privacy-safe data
- ✅ **Data Provider** - `data-provider.yaml` - Third-party data integration

### Optimization (1)
- ✅ **Tactical Recommendations** - `tactical-recommendations.json` - Bid/budget recommendations

### Audit & History (1)
- ✅ **Change History** - `change-history.json` - Change tracking

### Partners (1)
- ✅ **Partner Opportunities** - `partner-opportunities.json` - Partner program

### Testing (1)
- ✅ **Test Accounts** - `test-accounts.json` - Sandbox accounts

### Unified API (1)
- ✅ **Amazon Ads API v1** - `amazon-ads-v1-campaign-management.json` - Unified campaign management (Open Beta)

### Retail Ad Service (2)
- ✅ **Retail Ad Service** - `retail-ad-service.json` - Retail advertising
- ✅ **Retail Ad Service Identity** - `retail-ad-service-retailer-identity.json` - Retailer identity

### Deprecated/Older Versions (3)
- ✅ **Posts** - `posts.json` - Post management (Deprecated, shuts down 2025-07-31)
- ✅ **Sponsored Products v2** - `sponsored-products-v2.yaml` - Legacy SP API
- ✅ **Sponsored Brands v3** - `sponsored-brands-v3.yaml` - Legacy SB API

---

## 🚀 Implementation Status

### v0.2.0 (Current)
- ✅ All 45 APIs code generated
- ✅ Auto-fix for known issues (Creative Assets, Billing, Amazon Ads v1)
- ✅ Package documentation (`doc.go`) for all APIs
- ✅ Compilation verified
- ⏳ Unit tests in progress
- ⏳ Integration tests in progress

### v0.1.0 (Previous)
- ✅ 6 core APIs implemented
- ✅ 90%+ test coverage for core modules
- ✅ Basic usage examples

---

## 📝 Notes

### Code Generation

All APIs are auto-generated using `swagger-codegen`:

```bash
./scripts/generate-all-apis.ps1
```

Features:
- Automatic generation from OpenAPI specs
- Built-in fixes for known issues
- Complete type safety
- Full API coverage

### Known Issues & Fixes

The generation script automatically handles:

1. **Creative Assets** - Union types → `interface{}`
2. **Billing** - Missing `CountryCode` type definition
3. **Amazon Ads v1** - Invalid identifier names (numbers at start)

### API Versions

- **Latest versions** are recommended for new projects
- **Older versions** (v2, v3) are included for backward compatibility
- **Deprecated APIs** (Posts) will be removed after their sunset date

---

## 🔗 Related Documentation

- [Architecture Overview](ARCHITECTURE.md)
- [Development Guide](../DEVELOPMENT.md)
- [Project README](../README.md)

---

**Last Updated**: 2025-10-07  
**SDK Version**: v0.2.0
