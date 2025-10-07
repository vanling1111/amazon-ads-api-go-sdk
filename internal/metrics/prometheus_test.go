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

package metrics

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewPrometheusMetrics 测试创建 Prometheus 指标收集器
func TestNewPrometheusMetrics(t *testing.T) {
	metrics := NewPrometheusMetrics("test_adsapi")

	require.NotNil(t, metrics, "指标收集器不应该为 nil")
	assert.NotNil(t, metrics.RequestsTotal, "RequestsTotal 不应该为 nil")
	assert.NotNil(t, metrics.RequestDuration, "RequestDuration 不应该为 nil")
	assert.NotNil(t, metrics.ActiveRequests, "ActiveRequests 不应该为 nil")
	assert.NotNil(t, metrics.ErrorsTotal, "ErrorsTotal 不应该为 nil")
	assert.NotNil(t, metrics.RateLimitHits, "RateLimitHits 不应该为 nil")
	assert.NotNil(t, metrics.TokenCacheHits, "TokenCacheHits 不应该为 nil")
}

// TestPrometheusMetrics_RecordRequest 测试记录请求
func TestPrometheusMetrics_RecordRequest(t *testing.T) {
	metrics := NewPrometheusMetrics("test_record_request")

	// 测试记录成功请求
	t.Run("记录成功请求", func(t *testing.T) {
		metrics.RecordRequest("sponsored-products", "GET", 200, 100*time.Millisecond)
		// 指标已被记录，不会 panic
	})

	// 测试记录失败请求
	t.Run("记录失败请求", func(t *testing.T) {
		metrics.RecordRequest("sponsored-products", "POST", 400, 50*time.Millisecond)
		// 指标已被记录，不会 panic
	})

	// 测试记录不同的 HTTP 方法
	t.Run("记录不同的 HTTP 方法", func(t *testing.T) {
		methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
		for _, method := range methods {
			metrics.RecordRequest("sponsored-brands", method, 200, 100*time.Millisecond)
		}
	})

	// 测试记录不同的状态码
	t.Run("记录不同的状态码", func(t *testing.T) {
		statusCodes := []int{200, 201, 204, 400, 401, 403, 404, 429, 500, 502, 503}
		for _, code := range statusCodes {
			metrics.RecordRequest("profiles", "GET", code, 100*time.Millisecond)
		}
	})
}

// TestPrometheusMetrics_RecordError 测试记录错误
func TestPrometheusMetrics_RecordError(t *testing.T) {
	metrics := NewPrometheusMetrics("test_record_error")

	tests := []struct {
		name      string
		api       string
		errorType string
	}{
		{
			name:      "超时错误",
			api:       "sponsored-products",
			errorType: "timeout",
		},
		{
			name:      "认证失败",
			api:       "profiles",
			errorType: "auth_failed",
		},
		{
			name:      "网络错误",
			api:       "sponsored-brands",
			errorType: "network_error",
		},
		{
			name:      "解析错误",
			api:       "reports",
			errorType: "parse_error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metrics.RecordError(tt.api, tt.errorType)
			// 指标已被记录，不会 panic
		})
	}
}

// TestPrometheusMetrics_RecordRateLimitHit 测试记录速率限制命中
func TestPrometheusMetrics_RecordRateLimitHit(t *testing.T) {
	metrics := NewPrometheusMetrics("test_rate_limit")

	apis := []string{"sponsored-products", "sponsored-brands", "profiles", "reports"}
	for _, api := range apis {
		t.Run("记录 "+api+" 速率限制", func(t *testing.T) {
			metrics.RecordRateLimitHit(api)
			// 指标已被记录，不会 panic
		})
	}
}

// TestPrometheusMetrics_RecordTokenCacheHit 测试记录令牌缓存命中
func TestPrometheusMetrics_RecordTokenCacheHit(t *testing.T) {
	metrics := NewPrometheusMetrics("test_token_cache")

	t.Run("记录缓存命中", func(t *testing.T) {
		metrics.RecordTokenCacheHit(true)
		// 指标已被记录，不会 panic
	})

	t.Run("记录缓存未命中", func(t *testing.T) {
		metrics.RecordTokenCacheHit(false)
		// 指标已被记录，不会 panic
	})

	// 多次记录
	t.Run("多次记录", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			metrics.RecordTokenCacheHit(i%2 == 0)
		}
	})
}

// TestPrometheusMetrics_ActiveRequests 测试活跃请求计数
func TestPrometheusMetrics_ActiveRequests(t *testing.T) {
	metrics := NewPrometheusMetrics("test_active_requests")

	t.Run("增加活跃请求数", func(t *testing.T) {
		metrics.IncActiveRequests()
		metrics.IncActiveRequests()
		metrics.IncActiveRequests()
		// 指标已被更新，不会 panic
	})

	t.Run("减少活跃请求数", func(t *testing.T) {
		metrics.DecActiveRequests()
		metrics.DecActiveRequests()
		// 指标已被更新，不会 panic
	})

	t.Run("模拟请求生命周期", func(t *testing.T) {
		// 开始请求
		metrics.IncActiveRequests()

		// 模拟请求处理
		time.Sleep(10 * time.Millisecond)

		// 结束请求
		metrics.DecActiveRequests()
	})
}

// TestStatusCodeToString 测试状态码转换
func TestStatusCodeToString(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		expected   string
	}{
		{
			name:       "2xx - 200 OK",
			statusCode: 200,
			expected:   "2xx",
		},
		{
			name:       "2xx - 201 Created",
			statusCode: 201,
			expected:   "2xx",
		},
		{
			name:       "2xx - 204 No Content",
			statusCode: 204,
			expected:   "2xx",
		},
		{
			name:       "3xx - 301 Moved Permanently",
			statusCode: 301,
			expected:   "3xx",
		},
		{
			name:       "3xx - 302 Found",
			statusCode: 302,
			expected:   "3xx",
		},
		{
			name:       "4xx - 400 Bad Request",
			statusCode: 400,
			expected:   "4xx",
		},
		{
			name:       "4xx - 401 Unauthorized",
			statusCode: 401,
			expected:   "4xx",
		},
		{
			name:       "4xx - 404 Not Found",
			statusCode: 404,
			expected:   "4xx",
		},
		{
			name:       "4xx - 429 Too Many Requests",
			statusCode: 429,
			expected:   "4xx",
		},
		{
			name:       "5xx - 500 Internal Server Error",
			statusCode: 500,
			expected:   "5xx",
		},
		{
			name:       "5xx - 502 Bad Gateway",
			statusCode: 502,
			expected:   "5xx",
		},
		{
			name:       "5xx - 503 Service Unavailable",
			statusCode: 503,
			expected:   "5xx",
		},
		{
			name:       "未知状态码 - 100",
			statusCode: 100,
			expected:   "unknown",
		},
		{
			name:       "未知状态码 - 0",
			statusCode: 0,
			expected:   "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := statusCodeToString(tt.statusCode)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestPrometheusMetrics_ConcurrentUsage 测试并发使用
func TestPrometheusMetrics_ConcurrentUsage(t *testing.T) {
	metrics := NewPrometheusMetrics("test_concurrent")

	// 启动多个 goroutine 并发记录指标
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(id int) {
			// 记录请求
			metrics.RecordRequest("test-api", "GET", 200, 100*time.Millisecond)

			// 记录错误
			metrics.RecordError("test-api", "test_error")

			// 记录速率限制
			metrics.RecordRateLimitHit("test-api")

			// 记录令牌缓存
			metrics.RecordTokenCacheHit(id%2 == 0)

			// 活跃请求计数
			metrics.IncActiveRequests()
			time.Sleep(10 * time.Millisecond)
			metrics.DecActiveRequests()

			done <- true
		}(i)
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 10; i++ {
		<-done
	}
}

// Benchmark 测试

// BenchmarkNewPrometheusMetrics 基准测试创建指标收集器
func BenchmarkNewPrometheusMetrics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewPrometheusMetrics("bench_test")
	}
}

// BenchmarkPrometheusMetrics_RecordRequest 基准测试记录请求
func BenchmarkPrometheusMetrics_RecordRequest(b *testing.B) {
	metrics := NewPrometheusMetrics("bench_record_request")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		metrics.RecordRequest("test-api", "GET", 200, 100*time.Millisecond)
	}
}

// BenchmarkPrometheusMetrics_RecordError 基准测试记录错误
func BenchmarkPrometheusMetrics_RecordError(b *testing.B) {
	metrics := NewPrometheusMetrics("bench_record_error")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		metrics.RecordError("test-api", "test_error")
	}
}

// BenchmarkPrometheusMetrics_ActiveRequests 基准测试活跃请求计数
func BenchmarkPrometheusMetrics_ActiveRequests(b *testing.B) {
	metrics := NewPrometheusMetrics("bench_active_requests")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		metrics.IncActiveRequests()
		metrics.DecActiveRequests()
	}
}

// BenchmarkStatusCodeToString 基准测试状态码转换
func BenchmarkStatusCodeToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = statusCodeToString(200)
	}
}
