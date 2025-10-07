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

// 所有需要监控的 API 规范 - 完整的 45 个 API
var apiSpecs = []struct {
	Name string
	URL  string
}{
	// ========== 核心广告产品 (4) ==========
	{"Sponsored Products v3", "https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-products/openapi.json"},
	{"Sponsored Brands v4", "https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-brands/openapi.json"},
	{"Sponsored Display v3", "https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-display/openapi.yaml"},
	{"Sponsored TV", "https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-tv/openapi.json"},
	
	// ========== 核心账户管理 (4) ==========
	{"Profiles v3", "https://advertising.amazon.com/API/docs/en-us/openapi/profiles/openapi.yaml"},
	{"Portfolios v2", "https://advertising.amazon.com/API/docs/en-us/openapi/portfolios/openapi.json"},
	{"Advertising Accounts", "https://advertising.amazon.com/API/docs/en-us/openapi/advertising-accounts/openapi.json"},
	{"Manager Accounts", "https://advertising.amazon.com/API/docs/en-us/openapi/manager-accounts/openapi.json"},
	
	// ========== 报告和分析 (5) ==========
	{"Reporting v3", "https://advertising.amazon.com/API/docs/en-us/openapi/reporting/openapi.json"},
	{"Brand Metrics", "https://advertising.amazon.com/API/docs/en-us/openapi/brand-metrics/openapi.json"},
	{"Insights", "https://advertising.amazon.com/API/docs/en-us/openapi/insights/openapi.json"},
	{"Stores Analytics", "https://advertising.amazon.com/API/docs/en-us/openapi/stores-analytics/openapi.json"},
	{"Marketing Mix Modeling", "https://advertising.amazon.com/API/docs/en-us/openapi/marketing-mix-modeling/openapi.json"},
	
	// ========== DSP (5) ==========
	{"DSP Audiences", "https://advertising.amazon.com/API/docs/en-us/openapi/dsp-audiences/openapi.json"},
	{"DSP Conversions", "https://advertising.amazon.com/API/docs/en-us/openapi/dsp-conversions/openapi.json"},
	{"DSP Measurement", "https://advertising.amazon.com/API/docs/en-us/openapi/dsp-measurement/openapi.json"},
	{"DSP Target KPI", "https://advertising.amazon.com/API/docs/en-us/openapi/dsp-target-kpi/openapi.json"},
	{"DSP Advertisers", "https://advertising.amazon.com/API/docs/en-us/openapi/dsp-advertisers/openapi.yaml"},
	
	// ========== 受众和定向 (3) ==========
	{"Audiences Discovery", "https://advertising.amazon.com/API/docs/en-us/openapi/audiences-discovery/openapi.json"},
	{"Persona Builder", "https://advertising.amazon.com/API/docs/en-us/openapi/persona-builder/openapi.json"},
	{"Locations", "https://advertising.amazon.com/API/docs/en-us/openapi/locations/openapi.json"},
	
	// ========== 创意和素材 (4) ==========
	{"Creative Assets", "https://advertising.amazon.com/API/docs/en-us/openapi/creative-assets/openapi.yaml"},
	{"Moderation", "https://advertising.amazon.com/API/docs/en-us/openapi/moderation/openapi.json"},
	{"Pre-Moderation", "https://advertising.amazon.com/API/docs/en-us/openapi/pre-moderation/openapi.json"},
	{"Ad Library", "https://advertising.amazon.com/API/docs/en-us/openapi/ad-library/openapi.json"},
	
	// ========== 产品和目录 (2) ==========
	{"Product Metadata", "https://advertising.amazon.com/API/docs/en-us/openapi/product-metadata/openapi.json"},
	{"Product Eligibility", "https://advertising.amazon.com/API/docs/en-us/openapi/product-eligibility/openapi.json"},
	
	// ========== 归因和测量 (2) ==========
	{"Amazon Attribution", "https://advertising.amazon.com/API/docs/en-us/openapi/amazon-attribution/openapi.json"},
	{"Reach Forecasting", "https://advertising.amazon.com/API/docs/en-us/openapi/reach-forecasting/openapi.json"},
	
	// ========== 财务和预算 (2) ==========
	{"Billing", "https://advertising.amazon.com/API/docs/en-us/openapi/billing/openapi.json"},
	{"Account Budgets", "https://advertising.amazon.com/API/docs/en-us/openapi/account-budgets/openapi.json"},
	
	// ========== 数据管理 (4) ==========
	{"Exports", "https://advertising.amazon.com/API/docs/en-us/openapi/exports/openapi.json"},
	{"Marketing Stream", "https://advertising.amazon.com/API/docs/en-us/openapi/marketing-stream/openapi.json"},
	{"Hashed Records", "https://advertising.amazon.com/API/docs/en-us/openapi/hashed-records/openapi.json"},
	{"Data Provider", "https://advertising.amazon.com/API/docs/en-us/openapi/data-provider/openapi.yaml"},
	
	// ========== 优化 (1) ==========
	{"Tactical Recommendations", "https://advertising.amazon.com/API/docs/en-us/openapi/tactical-recommendations/openapi.json"},
	
	// ========== 审计和历史 (1) ==========
	{"Change History", "https://advertising.amazon.com/API/docs/en-us/openapi/change-history/openapi.json"},
	
	// ========== 合作伙伴 (1) ==========
	{"Partner Opportunities", "https://advertising.amazon.com/API/docs/en-us/openapi/partner-opportunities/openapi.json"},
	
	// ========== 测试 (1) ==========
	{"Test Accounts", "https://advertising.amazon.com/API/docs/en-us/openapi/test-accounts/openapi.json"},
	
	// ========== 统一 API (1) ==========
	{"Amazon Ads v1", "https://advertising.amazon.com/API/docs/en-us/openapi/amazon-ads-v1/openapi.json"},
	
	// ========== Retail Ad Service (2) ==========
	{"Retail Ad Service", "https://advertising.amazon.com/API/docs/en-us/openapi/retail-ad-service/openapi.json"},
	{"Retail Ad Service Identity", "https://advertising.amazon.com/API/docs/en-us/openapi/retail-ad-service-retailer-identity/openapi.json"},
	
	// ========== 已弃用/旧版本 (3) ==========
	{"Posts (Deprecated)", "https://advertising.amazon.com/API/docs/en-us/openapi/posts/openapi.json"},
	{"Sponsored Products v2", "https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-products-v2/openapi.yaml"},
	{"Sponsored Brands v3", "https://advertising.amazon.com/API/docs/en-us/openapi/sponsored-brands-v3/openapi.yaml"},
}

func main() {
	log.Println("🔍 Starting Amazon Ads API Monitor...")
	log.Printf("📊 Monitoring %d APIs in total", len(apiSpecs))
	
	configPath := "tools/api-monitor/monitor-state.json"
	
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
			// API 规范已变更
			log.Printf("🔄 Change detected in %s!", spec.Name)
			log.Printf("   Old version: %s", prevSpec.Version)
			log.Printf("   New version: %s", version)
			log.Printf("   Old checksum: %s", prevSpec.Checksum[:16])
			log.Printf("   New checksum: %s", checksum[:16])
			
			prevSpec.Checksum = checksum
			prevSpec.Version = version
			prevSpec.LastCheck = time.Now()
			
			hasChanges = true
			changedAPIs = append(changedAPIs, fmt.Sprintf(
				"%s: %s → %s", 
				spec.Name, 
				prevSpec.Version, 
				version,
			))
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
	// 尝试解析 JSON
	var data map[string]interface{}
	if err := json.Unmarshal(content, &data); err == nil {
		if info, ok := data["info"].(map[string]interface{}); ok {
			if version, ok := info["version"].(string); ok {
				return version
			}
		}
	}
	
	// 如果是 YAML 或解析失败，返回 unknown
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
	reportPath := "tools/api-monitor/latest-changes.md"
	
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

