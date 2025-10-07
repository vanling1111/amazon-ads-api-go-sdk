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

// Package models 提供通用的数据模型。
package models

// Region 表示 Amazon Advertising API 的区域。
//
// Amazon Ads API 支持多个区域，每个区域有独立的端点和 LWA 认证服务器。
//
// 参考文档:
//   - https://advertising.amazon.com/API/docs/en-us/info/api-overview#api-endpoints
type Region struct {
	// Code 是区域代码（如 "NA", "EU", "FE"）
	Code string

	// Name 是区域名称
	Name string

	// Endpoint 是 API 端点 URL
	Endpoint string

	// LWAEndpoint 是 LWA 认证端点 URL
	LWAEndpoint string
}

// 预定义的区域常量。
var (
	// RegionNorthAmerica 北美区域（美国、加拿大、墨西哥）
	RegionNorthAmerica = Region{
		Code:        "NA",
		Name:        "North America",
		Endpoint:    "https://advertising-api.amazon.com",
		LWAEndpoint: "https://api.amazon.com/auth/o2/token",
	}

	// RegionEurope 欧洲区域（英国、法国、德国、意大利、西班牙等）
	RegionEurope = Region{
		Code:        "EU",
		Name:        "Europe",
		Endpoint:    "https://advertising-api-eu.amazon.com",
		LWAEndpoint: "https://api.amazon.co.uk/auth/o2/token",
	}

	// RegionFarEast 远东区域（日本、澳大利亚）
	RegionFarEast = Region{
		Code:        "FE",
		Name:        "Far East",
		Endpoint:    "https://advertising-api-fe.amazon.com",
		LWAEndpoint: "https://api.amazon.co.jp/auth/o2/token",
	}
)

// AllRegions 返回所有支持的区域。
func AllRegions() []Region {
	return []Region{
		RegionNorthAmerica,
		RegionEurope,
		RegionFarEast,
	}
}

// GetRegionByCode 通过区域代码获取区域。
//
// 参数:
//   - code: 区域代码（如 "NA", "EU", "FE"）
//
// 返回值:
//   - *Region: 区域实例，如果未找到返回 nil
func GetRegionByCode(code string) *Region {
	for _, r := range AllRegions() {
		if r.Code == code {
			return &r
		}
	}
	return nil
}
