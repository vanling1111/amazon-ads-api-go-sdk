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

package sponsoredproductsv3

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewAPIClient 测试创建 API 客户端
func TestNewAPIClient(t *testing.T) {
	tests := []struct {
		name string
		cfg  *Configuration
		want bool // 是否应该成功创建
	}{
		{
			name: "使用默认配置创建客户端",
			cfg:  NewConfiguration(),
			want: true,
		},
		{
			name: "使用自定义 HTTP 客户端",
			cfg: &Configuration{
				HTTPClient: &http.Client{},
				BasePath:   "https://advertising-api.amazon.com",
				UserAgent:  "test-agent/1.0",
			},
			want: true,
		},
		{
			name: "配置为 nil 时使用默认值",
			cfg:  nil,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.cfg
			if cfg == nil {
				cfg = NewConfiguration()
			}

			client := NewAPIClient(cfg)

			if tt.want {
				require.NotNil(t, client, "客户端不应该为 nil")
				assert.NotNil(t, client.cfg, "配置不应该为 nil")
				assert.NotNil(t, client.CampaignsApi, "CampaignsApi 不应该为 nil")
				assert.NotNil(t, client.AdGroupsApi, "AdGroupsApi 不应该为 nil")
				assert.NotNil(t, client.KeywordsApi, "KeywordsApi 不应该为 nil")
				assert.NotNil(t, client.ProductAdsApi, "ProductAdsApi 不应该为 nil")
			}
		})
	}
}

// TestNewConfiguration 测试创建配置
func TestNewConfiguration(t *testing.T) {
	cfg := NewConfiguration()

	require.NotNil(t, cfg, "配置不应该为 nil")
	assert.Equal(t, "/", cfg.BasePath, "默认 BasePath 应该是 /")
	assert.NotNil(t, cfg.DefaultHeader, "DefaultHeader 应该被初始化")
	assert.Contains(t, cfg.UserAgent, "Swagger-Codegen", "UserAgent 应该包含 Swagger-Codegen")
}

// TestConfiguration_AddDefaultHeader 测试添加默认 Header
func TestConfiguration_AddDefaultHeader(t *testing.T) {
	cfg := NewConfiguration()

	// 添加单个 header
	cfg.AddDefaultHeader("X-Custom-Header", "test-value")
	assert.Equal(t, "test-value", cfg.DefaultHeader["X-Custom-Header"])

	// 覆盖已有 header
	cfg.AddDefaultHeader("X-Custom-Header", "new-value")
	assert.Equal(t, "new-value", cfg.DefaultHeader["X-Custom-Header"])

	// 添加多个 headers
	cfg.AddDefaultHeader("X-Another-Header", "another-value")
	assert.Equal(t, "another-value", cfg.DefaultHeader["X-Another-Header"])
	assert.Len(t, cfg.DefaultHeader, 2)
}

// TestAPIClient_SelectHeaderContentType 测试选择 Content-Type
func TestAPIClient_SelectHeaderContentType(t *testing.T) {
	client := NewAPIClient(NewConfiguration())

	tests := []struct {
		name         string
		contentTypes []string
		want         string
	}{
		{
			name:         "空列表返回空字符串",
			contentTypes: []string{},
			want:         "",
		},
		{
			name:         "包含 application/json 时优先选择",
			contentTypes: []string{"text/plain", "application/json", "application/xml"},
			want:         "application/json",
		},
		{
			name:         "只有一个类型时返回该类型",
			contentTypes: []string{"application/xml"},
			want:         "application/xml",
		},
		{
			name:         "不包含 application/json 时返回第一个",
			contentTypes: []string{"text/plain", "application/xml"},
			want:         "text/plain",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := client.selectHeaderContentType(tt.contentTypes)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestAPIClient_SelectHeaderAccept 测试选择 Accept header
func TestAPIClient_SelectHeaderAccept(t *testing.T) {
	client := NewAPIClient(NewConfiguration())

	tests := []struct {
		name    string
		accepts []string
		want    string
	}{
		{
			name:    "空列表返回空字符串",
			accepts: []string{},
			want:    "",
		},
		{
			name:    "包含 application/json 时优先选择",
			accepts: []string{"text/plain", "application/json", "application/xml"},
			want:    "application/json",
		},
		{
			name:    "不包含 application/json 时返回逗号分隔的所有类型",
			accepts: []string{"text/plain", "application/xml"},
			want:    "text/plain,application/xml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := client.selectHeaderAccept(tt.accepts)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestParameterToString 测试参数转字符串
func TestParameterToString(t *testing.T) {
	tests := []struct {
		name             string
		obj              interface{}
		collectionFormat string
		want             string
	}{
		{
			name:             "字符串类型",
			obj:              "test",
			collectionFormat: "",
			want:             "test",
		},
		{
			name:             "整数类型",
			obj:              123,
			collectionFormat: "",
			want:             "123",
		},
		{
			name:             "浮点数类型",
			obj:              123.45,
			collectionFormat: "",
			want:             "123.45",
		},
		{
			name:             "布尔类型",
			obj:              true,
			collectionFormat: "",
			want:             "true",
		},
		{
			name:             "字符串切片 - csv 格式",
			obj:              []string{"a", "b", "c"},
			collectionFormat: "csv",
			want:             "a,b,c",
		},
		{
			name:             "字符串切片 - pipes 格式",
			obj:              []string{"a", "b", "c"},
			collectionFormat: "pipes",
			want:             "a|b|c",
		},
		{
			name:             "字符串切片 - ssv 格式",
			obj:              []string{"a", "b", "c"},
			collectionFormat: "ssv",
			want:             "a b c",
		},
		{
			name:             "字符串切片 - tsv 格式",
			obj:              []string{"a", "b", "c"},
			collectionFormat: "tsv",
			want:             "a\tb\tc",
		},
		{
			name:             "整数切片",
			obj:              []int{1, 2, 3},
			collectionFormat: "csv",
			want:             "1,2,3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parameterToString(tt.obj, tt.collectionFormat)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestContains 测试字符串包含检查
func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		haystack []string
		needle   string
		want     bool
	}{
		{
			name:     "找到匹配项",
			haystack: []string{"apple", "banana", "orange"},
			needle:   "banana",
			want:     true,
		},
		{
			name:     "未找到匹配项",
			haystack: []string{"apple", "banana", "orange"},
			needle:   "grape",
			want:     false,
		},
		{
			name:     "空列表",
			haystack: []string{},
			needle:   "test",
			want:     false,
		},
		{
			name:     "大小写不敏感匹配",
			haystack: []string{"Apple", "Banana", "Orange"},
			needle:   "apple",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := contains(tt.haystack, tt.needle)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestTypeCheckParameter 测试参数类型检查
func TestTypeCheckParameter(t *testing.T) {
	tests := []struct {
		name     string
		obj      interface{}
		expected string
		wantErr  bool
	}{
		{
			name:     "类型匹配 - string",
			obj:      "test",
			expected: "string",
			wantErr:  false,
		},
		{
			name:     "类型匹配 - int",
			obj:      123,
			expected: "int",
			wantErr:  false,
		},
		{
			name:     "类型不匹配",
			obj:      "test",
			expected: "int",
			wantErr:  true,
		},
		{
			name:     "nil 对象总是通过",
			obj:      nil,
			expected: "string",
			wantErr:  false,
		},
		{
			name:     "指针类型",
			obj:      new(string),
			expected: "*string",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := typeCheckParameter(tt.obj, tt.expected, "test_param")
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "test_param")
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestAPIClient_CallAPI 测试 HTTP 调用
func TestAPIClient_CallAPI(t *testing.T) {
	// 创建 mock HTTP 服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	// 创建客户端
	cfg := NewConfiguration()
	cfg.HTTPClient = server.Client()
	client := NewAPIClient(cfg)

	// 创建请求
	req, err := http.NewRequest("GET", server.URL, nil)
	require.NoError(t, err)

	// 执行请求
	resp, err := client.callAPI(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestAPIClient_ChangeBasePath 测试修改 BasePath
func TestAPIClient_ChangeBasePath(t *testing.T) {
	client := NewAPIClient(NewConfiguration())

	originalPath := client.cfg.BasePath
	newPath := "https://new-api.example.com"

	client.ChangeBasePath(newPath)

	assert.NotEqual(t, originalPath, client.cfg.BasePath)
	assert.Equal(t, newPath, client.cfg.BasePath)
}

// TestGenericSwaggerError 测试通用错误类型
func TestGenericSwaggerError(t *testing.T) {
	err := GenericSwaggerError{
		body:  []byte(`{"error": "test"}`),
		error: "Test error message",
		model: map[string]interface{}{"code": "TEST_ERROR"},
	}

	// 测试 Error() 方法
	assert.Equal(t, "Test error message", err.Error())

	// 测试 Body() 方法
	assert.Equal(t, []byte(`{"error": "test"}`), err.Body())

	// 测试 Model() 方法
	model := err.Model().(map[string]interface{})
	assert.Equal(t, "TEST_ERROR", model["code"])
}

// TestAPIClient_AllServicesInitialized 测试所有 API 服务都已初始化
func TestAPIClient_AllServicesInitialized(t *testing.T) {
	client := NewAPIClient(NewConfiguration())

	// 验证所有 24 个 API 服务都已初始化
	services := []interface{}{
		client.AdGroupsApi,
		client.BudgetRecommendationNewCampaignsApi,
		client.BudgetRecommendationsAndMissedOpportunitiesApi,
		client.BudgetRulesApi,
		client.BudgetRulesRecommendationApi,
		client.BudgetUsageApi,
		client.CampaignNegativeKeywordsApi,
		client.CampaignNegativeTargetingClausesApi,
		client.CampaignOptimizationRulesApi,
		client.CampaignsApi,
		client.ConsolidatedRecommendationsApi,
		client.KeywordGroupTargetingRecommendationsApi,
		client.KeywordTargetsApi,
		client.KeywordsApi,
		client.MultiCountryThemeBasedBidRecommendationsApi,
		client.NegativeKeywordsApi,
		client.NegativeTargetingClausesApi,
		client.OptimizationRulesApi,
		client.ProductAdsApi,
		client.ProductRecommendationServiceApi,
		client.ProductTargetingApi,
		client.TargetPromotionGroupsApi,
		client.TargetingClausesApi,
		client.ThemeBasedBidRecommendationsApi,
	}

	for i, service := range services {
		assert.NotNil(t, service, "API 服务 #%d 不应该为 nil", i)
	}
}

// Benchmark 测试

// BenchmarkParameterToString 基准测试参数转字符串
func BenchmarkParameterToString(b *testing.B) {
	obj := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parameterToString(obj, "csv")
	}
}

// BenchmarkContains 基准测试字符串包含检查
func BenchmarkContains(b *testing.B) {
	haystack := []string{"apple", "banana", "orange", "grape", "melon"}
	needle := "grape"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = contains(haystack, needle)
	}
}

// BenchmarkNewAPIClient 基准测试创建 API 客户端
func BenchmarkNewAPIClient(b *testing.B) {
	cfg := NewConfiguration()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = NewAPIClient(cfg)
	}
}

