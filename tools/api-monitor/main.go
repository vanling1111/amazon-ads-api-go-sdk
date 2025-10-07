// Copyright 2025 Amazon Ads API Go SDK Authors.
//
// API 监控工具 - 检测 Amazon Advertising API 规范的变更

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	
	"gopkg.in/yaml.v3"
)

// APISpec 表示一个 API 规范
type APISpec struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Checksum string `json:"checksum"`
	Version  string `json:"version"`
	LastCheck time.Time `json:"last_check"`
}

// MonitorConfig 监控配置
type MonitorConfig struct {
	Specs []APISpec `json:"specs"`
}

// 所有需要监控的 API 规范
// 使用 CloudFront CDN 的直接 JSON/YAML 链接，确保内容稳定、无动态元素
// 来源：https://advertising.amazon.com/API/docs/en-us/reference/openapi-download
var apiSpecs = []struct {
	Name string
	URL  string
}{
	// ========== 账户管理 (7个) ==========
	{"Profiles v3", "https://d3a0d0y2hgofx6.cloudfront.net/openapi/en-us/profiles/3-0/openapi.yaml"},
	{"Manager Accounts", "https://dtrnk0o2zy01c.cloudfront.net/openapi/en-us/dest/ManagerAccount_prod_3p.json"},
	{"Advertising Accounts", "https://dtrnk0o2zy01c.cloudfront.net/openapi/en-us/dest/AdvertisingAccounts_prod_3p.json"},
	{"Portfolios", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/Portfolios_prod_3p.json"},
	{"Billing", "https://dtrnk0o2zy01c.cloudfront.net/openapi/en-us/dest/AdvertisingBilling_prod_3p.json"},
	{"Account Budgets", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/Advertisers_prod_3p.json"},
	{"Test Accounts", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/AdvertisingTestAccount_prod_3p.json"},
	
	// ========== 广告产品 (5个) ==========
	{"Sponsored Products v3", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/SponsoredProducts_prod_3p.json"},
	{"Sponsored Products v2", "https://d3a0d0y2hgofx6.cloudfront.net/openapi/en-us/sponsored-products/2-0/openapi.yaml"},
	{"Sponsored Brands v4", "https://d3a0d0y2hgofx6.cloudfront.net/openapi/en-us/sponsored-brands/4-0/openapi.json"},
	{"Sponsored Brands v3", "https://d3a0d0y2hgofx6.cloudfront.net/openapi/en-us/sponsored-brands/3-0/openapi.yaml"},
	{"Sponsored Display v3", "https://d3a0d0y2hgofx6.cloudfront.net/openapi/en-us/sponsored-display/3-0/openapi.yaml"},
	
	// ========== 报告和分析 (5个) ==========
	{"Reporting v3", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/OfflineReport_prod_3p.json"},
	{"Brand Metrics", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/BrandMetrics_prod_3p.json"},
	{"Stores Analytics", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/Stores_prod_3p.json"},
	{"Marketing Mix Modeling", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/MarketingMixModeling_prod_3p.json"},
	{"Reach Forecasting", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/ReachPlanningService_prod_3p.json"},
	
	// ========== Amazon DSP (5个) ==========
	{"DSP Measurement", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/Measurement_prod_3p.json"},
	{"DSP Advertisers", "https://d3a0d0y2hgofx6.cloudfront.net/openapi/en-us/dsp/3-0/advertiser.yaml"},
	{"DSP Audiences", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/ADSPAudiences_prod_3p.json"},
	{"DSP Conversions", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/ConversionsAPI_prod_3p.json"},
	{"DSP Target KPI", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/GoalSeekingBidderTargetKPIRecommendation_prod_3p.json"},
	
	// ========== 推荐和洞察 (5个) ==========
	{"Amazon Attribution", "https://dtrnk0o2zy01c.cloudfront.net/openapi/en-us/dest/AmazonAttribution_prod_3p.json"},
	{"Insights", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/Insights_prod_3p.json"},
	{"Partner Opportunities", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/PartnerOpportunities_prod_3p.json"},
	{"Tactical Recommendations", "https://dtrnk0o2zy01c.cloudfront.net/openapi/en-us/dest/Recommendations_prod_3p.json"},
	{"Persona Builder", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/PersonaBuilderAPI_prod_3p.json"},
	
	// ========== 创意和内容 (3个) ==========
	{"Creative Assets", "https://d3a0d0y2hgofx6.cloudfront.net/openapi/en-us/creative-asset-library/creative-asset-library-openapi.yaml"},
	{"Pre-Moderation", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/PreModeration_prod_3p.json"},
	{"Moderation", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/Moderation_prod_3p.json"},
	
	// ========== 数据管理 (6个) ==========
	{"Audiences Discovery", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/Audiences_prod_3p.json"},
	{"Amazon Marketing Stream", "https://dtrnk0o2zy01c.cloudfront.net/openapi/en-us/dest/AmazonMarketingStream_prod_3p.json"},
	{"Locations", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/Locations_prod_3p.json"},
	{"Exports", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/AmazonAdsAPIExports_prod_3p.json"},
	{"Change History", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/Changehistory_prod_3p.json"},
	{"Data Provider", "https://d3a0d0y2hgofx6.cloudfront.net/openapi/en-us/data-provider/openapi.yaml"},
	{"Hashed Records", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/HashedRecords_prod_3p.json"},
	
	// ========== 产品 (2个) ==========
	{"Product Metadata", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/ProductSelector_prod_3p.json"},
	{"Product Eligibility", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/Eligibility_prod_3p.json"},
	
	// ========== 统一 API (1个) ==========
	{"Amazon Ads API v1", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/AmazonAdsAPIALL_prod_3p.json"},
	
	// ========== Retail Ad Service (2个) ==========
	{"Retail Ad Service", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/AmazonAdvertiserAPIforRetailAdService_prod_3p.json"},
	{"Retail Ad Service Identity", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/RetailerIdentityAPIforRetailAdService_prod_3p.json"},
	
	// ========== Sponsored TV (1个) ==========
	{"Sponsored TV", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/SponsoredTV_prod_3p.json"},
	
	// ========== 已弃用 (2个) ==========
	{"Posts (Deprecated)", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/Posts_prod_3p.json"},
	{"Ad Library", "https://d1y2lf8k3vrkfu.cloudfront.net/openapi/en-us/dest/AdLibraryAPI_prod_3p.json"},
}

func main() {
	log.Println("🔍 Starting Amazon Ads API Monitor...")
	log.Printf("📊 Monitoring %d APIs in total", len(apiSpecs))
	
	configPath := "monitor-state.json"
	
	// 加载之前的状态
	config, err := loadConfig(configPath)
	if err != nil {
		log.Printf("⚠️  No previous state found, creating new: %v", err)
		config = &MonitorConfig{Specs: []APISpec{}}
	}
	
	hasChanges := false
	changedAPIs := []string{}
	newAPIs := []string{}
	deletedAPIs := []string{}
	
	// 先检测是否有 API 被删除（之前有，现在没有）
	currentAPINames := make(map[string]bool)
	for _, spec := range apiSpecs {
		currentAPINames[spec.Name] = true
	}
	
	for _, prevSpec := range config.Specs {
		if !currentAPINames[prevSpec.Name] {
			log.Printf("🗑️  API removed or deprecated: %s", prevSpec.Name)
			deletedAPIs = append(deletedAPIs, prevSpec.Name)
			hasChanges = true
		}
	}
	
	// 检查每个 API 规范
	for _, spec := range apiSpecs {
		log.Printf("📡 Checking %s...", spec.Name)
		
		// 下载规范
		content, err := downloadSpec(spec.URL)
		if err != nil {
			log.Printf("❌ Failed to download %s: %v", spec.Name, err)
			continue
		}
		
		// 计算 checksum
		checksum := calculateChecksum(content)
		
		// 提取版本信息
		version := extractVersion(content)
		
		// 查找之前的记录
		prevSpec := findSpec(config.Specs, spec.Name)
		
		if prevSpec == nil {
			// 新的 API
			log.Printf("🆕 New API detected: %s (version: %s)", spec.Name, version)
			config.Specs = append(config.Specs, APISpec{
				Name:      spec.Name,
				URL:       spec.URL,
				Checksum:  checksum,
				Version:   version,
				LastCheck: time.Now(),
			})
			hasChanges = true
			newAPIs = append(newAPIs, fmt.Sprintf("%s (v%s)", spec.Name, version))
			changedAPIs = append(changedAPIs, fmt.Sprintf("NEW: %s (v%s)", spec.Name, version))
		} else if prevSpec.Checksum != checksum {
			// API 规范的 checksum 已变更
			log.Printf("🔄 Checksum change detected in %s", spec.Name)
			log.Printf("   Old version: %s", prevSpec.Version)
			log.Printf("   New version: %s", version)
			log.Printf("   Old checksum: %s", prevSpec.Checksum[:16])
			log.Printf("   New checksum: %s", checksum[:16])
			
			// 更新 checksum 和时间戳
			prevSpec.Checksum = checksum
			prevSpec.LastCheck = time.Now()
			
			// 只有当版本号真正改变时才算作重要变更
			if prevSpec.Version != version && version != "unknown" {
				log.Printf("   ⚠️  VERSION CHANGED: %s → %s", prevSpec.Version, version)
				prevSpec.Version = version
				hasChanges = true
				changedAPIs = append(changedAPIs, fmt.Sprintf(
					"%s: %s → %s", 
					spec.Name, 
					prevSpec.Version, 
					version,
				))
			} else {
				log.Printf("   ℹ️  Version unchanged (%s), likely formatting/timestamp difference", version)
				// 只更新 version 字段（可能是 unknown → 具体版本）
				if version != "unknown" && prevSpec.Version == "unknown" {
					prevSpec.Version = version
				}
			}
		} else {
			log.Printf("✅ No changes in %s (v%s)", spec.Name, version)
			prevSpec.LastCheck = time.Now()
		}
		
		// 避免请求过快
		time.Sleep(2 * time.Second)
	}
	
	// 保存新状态
	if err := saveConfig(configPath, config); err != nil {
		log.Printf("❌ Failed to save config: %v", err)
	}
	
	// 生成变更报告
	if hasChanges {
		log.Println("\n" + strings.Repeat("=", 60))
		log.Println("📢 API CHANGES DETECTED!")
		log.Println(strings.Repeat("=", 60))
		
		if len(newAPIs) > 0 {
			log.Println("\n🆕 New APIs:")
			for _, api := range newAPIs {
				log.Printf("  • %s", api)
			}
		}
		
		if len(deletedAPIs) > 0 {
			log.Println("\n🗑️  Deleted/Deprecated APIs:")
			for _, api := range deletedAPIs {
				log.Printf("  • %s", api)
			}
		}
		
		updatedAPIs := []string{}
		for _, change := range changedAPIs {
			if !strings.HasPrefix(change, "NEW:") {
				updatedAPIs = append(updatedAPIs, change)
			}
		}
		
		if len(updatedAPIs) > 0 {
			log.Println("\n🔄 Updated APIs:")
			for _, change := range updatedAPIs {
				log.Printf("  • %s", change)
			}
		}
		
		log.Println(strings.Repeat("=", 60))
		
		// 生成详细报告
		generateReport(changedAPIs, newAPIs, deletedAPIs)
		
		// 退出码 1 表示检测到变更（触发 GitHub Actions 创建 Issue）
		os.Exit(1)
	}
	
	log.Println("\n✅ No API changes detected. All specifications are up to date.")
}

// downloadSpec 下载 API 规范
func downloadSpec(url string) ([]byte, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	
	return io.ReadAll(resp.Body)
}

// calculateChecksum 计算内容的 SHA256 校验和
func calculateChecksum(content []byte) string {
	hash := sha256.Sum256(content)
	return hex.EncodeToString(hash[:])
}

// extractVersion 从规范中提取版本号
func extractVersion(content []byte) string {
	var data map[string]interface{}
	
	// 先尝试解析 JSON
	if err := json.Unmarshal(content, &data); err != nil {
		// JSON 解析失败，尝试 YAML
		if err := yaml.Unmarshal(content, &data); err != nil {
			return "unknown"
		}
	}
	
	// 提取 info.version
	if info, ok := data["info"].(map[string]interface{}); ok {
		if version, ok := info["version"].(string); ok {
			return version
		}
	}
	
	return "unknown"
}

// findSpec 查找指定名称的规范
func findSpec(specs []APISpec, name string) *APISpec {
	for i := range specs {
		if specs[i].Name == name {
			return &specs[i]
		}
	}
	return nil
}

// loadConfig 加载配置
func loadConfig(path string) (*MonitorConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	
	var config MonitorConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	
	return &config, nil
}

// saveConfig 保存配置
func saveConfig(path string, config *MonitorConfig) error {
	// 确保目录存在
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(path, data, 0644)
}

// generateReport 生成详细的变更报告
func generateReport(changes []string, newAPIs []string, deletedAPIs []string) {
	reportPath := "latest-changes.md"
	
	report := fmt.Sprintf(`# API Change Report

**Date**: %s

`, time.Now().Format("2006-01-02 15:04:05"))
	
	// 新增的 API
	if len(newAPIs) > 0 {
		report += "## 🆕 New APIs\n\n"
		for _, api := range newAPIs {
			report += fmt.Sprintf("- %s\n", api)
		}
		report += "\n"
	}
	
	// 删除/弃用的 API
	if len(deletedAPIs) > 0 {
		report += "## 🗑️ Deleted/Deprecated APIs\n\n"
		for _, api := range deletedAPIs {
			report += fmt.Sprintf("- %s\n", api)
		}
		report += "\n"
	}
	
	// 更新的 API（排除新增的）
	updatedAPIs := []string{}
	for _, change := range changes {
		if !strings.HasPrefix(change, "NEW:") {
			updatedAPIs = append(updatedAPIs, change)
		}
	}
	
	if len(updatedAPIs) > 0 {
		report += "## 🔄 Updated APIs (Version/Specification Changes)\n\n"
		for _, change := range updatedAPIs {
			report += fmt.Sprintf("- %s\n", change)
		}
		report += "\n"
	}
	
	report += `## ⚠️ Action Items

1. ✅ Review the changed API specifications
2. ✅ Check for breaking changes (parameters, response formats, etc.)
3. ✅ Update affected code in SDK
4. ✅ Regenerate client code using code generation tool
5. ✅ Update unit tests
6. ✅ Update integration tests
7. ✅ Update documentation and examples
8. ✅ Update CHANGELOG.md
9. ✅ Bump version number (patch/minor/major based on changes)
10. ✅ Create release notes

## 🔧 Next Steps

### 1. Regenerate All APIs

Run the following command to regenerate all API clients:

` + "```powershell\n.\\scripts\\generate-all-apis.ps1\n```\n\n"

	report += `### 2. Test Changes

` + "```powershell\ngo test ./...\n```\n\n"

	report += `### 3. Update Documentation

- Update affected ` + "`doc.go`" + ` files
- Update ` + "`README.md`" + ` if new APIs added
- Update ` + "`API_COVERAGE.md`" + `

### 4. Version Bump

- **Patch** (x.x.X): Bug fixes, parameter updates
- **Minor** (x.X.0): New APIs, new features, backward compatible
- **Major** (X.0.0): Breaking changes

---

**Generated by**: Amazon Ads API Monitor  
**Repository**: https://github.com/vanling1111/amazon-ads-api-go-sdk
`
	
	if err := os.WriteFile(reportPath, []byte(report), 0644); err != nil {
		log.Printf("❌ Failed to write report: %v", err)
	} else {
		log.Printf("📄 Report saved to: %s", reportPath)
	}
}

