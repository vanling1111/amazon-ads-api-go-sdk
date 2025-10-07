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
// terms of the applicable license. All rights reserved.

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRegionConstants 测试预定义的区域常量
func TestRegionConstants(t *testing.T) {
	tests := []struct {
		name         string
		region       Region
		expectedCode string
		expectedName string
	}{
		{
			name:         "North America Region",
			region:       RegionNorthAmerica,
			expectedCode: "NA",
			expectedName: "North America",
		},
		{
			name:         "Europe Region",
			region:       RegionEurope,
			expectedCode: "EU",
			expectedName: "Europe",
		},
		{
			name:         "Far East Region",
			region:       RegionFarEast,
			expectedCode: "FE",
			expectedName: "Far East",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedCode, tt.region.Code)
			assert.Equal(t, tt.expectedName, tt.region.Name)
			assert.NotEmpty(t, tt.region.Endpoint, "Endpoint 不应该为空")
			assert.NotEmpty(t, tt.region.LWAEndpoint, "LWAEndpoint 不应该为空")
		})
	}
}

// TestRegionNorthAmerica 测试北美区域
func TestRegionNorthAmerica(t *testing.T) {
	region := RegionNorthAmerica

	assert.Equal(t, "NA", region.Code)
	assert.Equal(t, "North America", region.Name)
	assert.Equal(t, "https://advertising-api.amazon.com", region.Endpoint)
	assert.Equal(t, "https://api.amazon.com/auth/o2/token", region.LWAEndpoint)
}

// TestRegionEurope 测试欧洲区域
func TestRegionEurope(t *testing.T) {
	region := RegionEurope

	assert.Equal(t, "EU", region.Code)
	assert.Equal(t, "Europe", region.Name)
	assert.Equal(t, "https://advertising-api-eu.amazon.com", region.Endpoint)
	assert.Equal(t, "https://api.amazon.co.uk/auth/o2/token", region.LWAEndpoint)
}

// TestRegionFarEast 测试远东区域
func TestRegionFarEast(t *testing.T) {
	region := RegionFarEast

	assert.Equal(t, "FE", region.Code)
	assert.Equal(t, "Far East", region.Name)
	assert.Equal(t, "https://advertising-api-fe.amazon.com", region.Endpoint)
	assert.Equal(t, "https://api.amazon.co.jp/auth/o2/token", region.LWAEndpoint)
}

// TestAllRegions 测试获取所有区域
func TestAllRegions(t *testing.T) {
	regions := AllRegions()

	require.Len(t, regions, 3, "应该有 3 个区域")

	// 验证每个区域都存在
	codes := make(map[string]bool)
	for _, r := range regions {
		codes[r.Code] = true
		assert.NotEmpty(t, r.Code, "区域代码不应该为空")
		assert.NotEmpty(t, r.Name, "区域名称不应该为空")
		assert.NotEmpty(t, r.Endpoint, "API 端点不应该为空")
		assert.NotEmpty(t, r.LWAEndpoint, "LWA 端点不应该为空")
	}

	// 验证所有预期的区域都存在
	assert.True(t, codes["NA"], "应该包含北美区域")
	assert.True(t, codes["EU"], "应该包含欧洲区域")
	assert.True(t, codes["FE"], "应该包含远东区域")
}

// TestGetRegionByCode 测试通过代码获取区域
func TestGetRegionByCode(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected *Region
		wantNil  bool
	}{
		{
			name:     "获取北美区域",
			code:     "NA",
			expected: &RegionNorthAmerica,
			wantNil:  false,
		},
		{
			name:     "获取欧洲区域",
			code:     "EU",
			expected: &RegionEurope,
			wantNil:  false,
		},
		{
			name:     "获取远东区域",
			code:     "FE",
			expected: &RegionFarEast,
			wantNil:  false,
		},
		{
			name:     "无效的区域代码",
			code:     "INVALID",
			expected: nil,
			wantNil:  true,
		},
		{
			name:     "空字符串",
			code:     "",
			expected: nil,
			wantNil:  true,
		},
		{
			name:     "小写代码（大小写敏感）",
			code:     "na",
			expected: nil,
			wantNil:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			region := GetRegionByCode(tt.code)

			if tt.wantNil {
				assert.Nil(t, region, "应该返回 nil")
			} else {
				require.NotNil(t, region, "不应该返回 nil")
				assert.Equal(t, tt.expected.Code, region.Code)
				assert.Equal(t, tt.expected.Name, region.Name)
				assert.Equal(t, tt.expected.Endpoint, region.Endpoint)
				assert.Equal(t, tt.expected.LWAEndpoint, region.LWAEndpoint)
			}
		})
	}
}

// TestRegion_Endpoints 测试所有区域的端点格式
func TestRegion_Endpoints(t *testing.T) {
	regions := AllRegions()

	for _, region := range regions {
		t.Run(region.Name, func(t *testing.T) {
			// 验证 API 端点格式
			assert.Contains(t, region.Endpoint, "https://", "API 端点应该使用 HTTPS")
			assert.Contains(t, region.Endpoint, "advertising-api", "API 端点应该包含 'advertising-api'")

			// 验证 LWA 端点格式
			assert.Contains(t, region.LWAEndpoint, "https://", "LWA 端点应该使用 HTTPS")
			assert.Contains(t, region.LWAEndpoint, "api.amazon", "LWA 端点应该包含 'api.amazon'")
			assert.Contains(t, region.LWAEndpoint, "/auth/o2/token", "LWA 端点应该包含认证路径")
		})
	}
}

// TestRegion_Uniqueness 测试区域的唯一性
func TestRegion_Uniqueness(t *testing.T) {
	regions := AllRegions()

	// 检查代码唯一性
	codes := make(map[string]bool)
	for _, r := range regions {
		assert.False(t, codes[r.Code], "区域代码 %s 重复", r.Code)
		codes[r.Code] = true
	}

	// 检查端点唯一性
	endpoints := make(map[string]bool)
	for _, r := range regions {
		assert.False(t, endpoints[r.Endpoint], "API 端点 %s 重复", r.Endpoint)
		endpoints[r.Endpoint] = true
	}

	// 检查 LWA 端点唯一性
	lwaEndpoints := make(map[string]bool)
	for _, r := range regions {
		assert.False(t, lwaEndpoints[r.LWAEndpoint], "LWA 端点 %s 重复", r.LWAEndpoint)
		lwaEndpoints[r.LWAEndpoint] = true
	}
}

// Benchmark 测试

// BenchmarkGetRegionByCode 基准测试通过代码获取区域
func BenchmarkGetRegionByCode(b *testing.B) {
	b.Run("Found", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = GetRegionByCode("NA")
		}
	})

	b.Run("NotFound", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = GetRegionByCode("INVALID")
		}
	})
}

// BenchmarkAllRegions 基准测试获取所有区域
func BenchmarkAllRegions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AllRegions()
	}
}

