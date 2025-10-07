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

// Package metrics 提供 Prometheus 指标收集功能。
package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// PrometheusMetrics Prometheus 指标收集器。
//
// 提供预定义的 Ads API 相关指标：
//   - 请求总数
//   - 请求延迟直方图
//   - 活跃请求数
//   - 错误总数
//   - 速率限制命中次数
type PrometheusMetrics struct {
	// RequestsTotal 请求总数计数器
	RequestsTotal *prometheus.CounterVec

	// RequestDuration 请求延迟直方图
	RequestDuration *prometheus.HistogramVec

	// ActiveRequests 当前活跃请求数
	ActiveRequests prometheus.Gauge

	// ErrorsTotal 错误总数计数器
	ErrorsTotal *prometheus.CounterVec

	// RateLimitHits 速率限制命中次数
	RateLimitHits *prometheus.CounterVec

	// TokenCacheHits 令牌缓存命中次数
	TokenCacheHits *prometheus.CounterVec
}

// NewPrometheusMetrics 创建 Prometheus 指标收集器。
//
// 参数:
//   - namespace: 指标命名空间（如 "adsapi"）
//
// 返回:
//   - *PrometheusMetrics: 指标收集器实例
//
// 示例:
//
//	metrics := NewPrometheusMetrics("adsapi")
//	// 在 HTTP handler 中暴露指标
//	http.Handle("/metrics", promhttp.Handler())
func NewPrometheusMetrics(namespace string) *PrometheusMetrics {
	return &PrometheusMetrics{
		RequestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "requests_total",
				Help:      "Total number of API requests",
			},
			[]string{"api", "method", "status"},
		),
		RequestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:      "request_duration_seconds",
				Help:      "API request duration in seconds",
				Buckets:   prometheus.DefBuckets, // 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10
			},
			[]string{"api", "method"},
		),
		ActiveRequests: promauto.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "active_requests",
				Help:      "Number of currently active requests",
			},
		),
		ErrorsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "errors_total",
				Help:      "Total number of errors",
			},
			[]string{"api", "error_type"},
		),
		RateLimitHits: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "rate_limit_hits_total",
				Help:      "Total number of rate limit hits",
			},
			[]string{"api"},
		),
		TokenCacheHits: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "token_cache_hits_total",
				Help:      "Total number of token cache hits",
			},
			[]string{"hit"},
		),
	}
}

// RecordRequest 记录 API 请求。
//
// 参数:
//   - api: API 名称（如 "sponsored-products"）
//   - method: HTTP 方法
//   - status: HTTP 状态码
//   - duration: 请求持续时间
func (m *PrometheusMetrics) RecordRequest(api, method string, status int, duration time.Duration) {
	m.RequestsTotal.WithLabelValues(api, method, statusCodeToString(status)).Inc()
	m.RequestDuration.WithLabelValues(api, method).Observe(duration.Seconds())
}

// RecordError 记录错误。
//
// 参数:
//   - api: API 名称
//   - errorType: 错误类型（如 "timeout", "auth_failed"）
func (m *PrometheusMetrics) RecordError(api, errorType string) {
	m.ErrorsTotal.WithLabelValues(api, errorType).Inc()
}

// RecordRateLimitHit 记录速率限制命中。
//
// 参数:
//   - api: API 名称
func (m *PrometheusMetrics) RecordRateLimitHit(api string) {
	m.RateLimitHits.WithLabelValues(api).Inc()
}

// RecordTokenCacheHit 记录令牌缓存命中。
//
// 参数:
//   - hit: 是否命中（"hit" 或 "miss"）
func (m *PrometheusMetrics) RecordTokenCacheHit(hit bool) {
	if hit {
		m.TokenCacheHits.WithLabelValues("hit").Inc()
	} else {
		m.TokenCacheHits.WithLabelValues("miss").Inc()
	}
}

// IncActiveRequests 增加活跃请求数。
func (m *PrometheusMetrics) IncActiveRequests() {
	m.ActiveRequests.Inc()
}

// DecActiveRequests 减少活跃请求数。
func (m *PrometheusMetrics) DecActiveRequests() {
	m.ActiveRequests.Dec()
}

// statusCodeToString 将 HTTP 状态码转换为字符串。
func statusCodeToString(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "2xx"
	case code >= 300 && code < 400:
		return "3xx"
	case code >= 400 && code < 500:
		return "4xx"
	case code >= 500:
		return "5xx"
	default:
		return "unknown"
	}
}

