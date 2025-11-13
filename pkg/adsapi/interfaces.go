// Copyright 2025 Amazon SP-API Go SDK Authors.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//   - Free for personal, educational, and open source projects
//   - Your project must also be open sourced under AGPL-3.0
//   - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//   - Required for any commercial, enterprise, or proprietary use
//   - Allows closed source distribution
//   - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license. All rights reserved.

// Package adsapi provides interfaces for extensibility.
package adsapi

import (
	"context"
	"net/http"
	"time"
)

// Logger 定义日志接口。
//
// 用户可以提供自己的日志实现，或使用SDK提供的默认实现。
// 默认情况下，SDK使用no-op logger（不输出任何日志）。
//
// 示例:
//
//	// 使用Zap日志
//	zapLogger := NewZapLogger(zap.NewProduction())
//	client := adsapi.NewClient(
//	    adsapi.WithLogger(zapLogger),
//	)
type Logger interface {
	// Debug 记录调试级别的日志
	Debug(msg string, fields ...Field)

	// Info 记录信息级别的日志
	Info(msg string, fields ...Field)

	// Warn 记录警告级别的日志
	Warn(msg string, fields ...Field)

	// Error 记录错误级别的日志
	Error(msg string, fields ...Field)

	// With 创建带有额外字段的子logger
	With(fields ...Field) Logger
}

// Field 表示日志字段。
type Field struct {
	Key   string
	Value interface{}
}

// MetricsCollector 定义指标收集接口。
//
// 用户可以提供自己的指标收集实现，或使用SDK提供的Prometheus实现。
// 默认情况下，SDK使用no-op collector（不收集任何指标）。
//
// 示例:
//
//	// 使用Prometheus
//	prometheusCollector := NewPrometheusCollector()
//	client := spapi.NewClient(
//	    spapi.WithMetrics(prometheusCollector),
//	)
type MetricsCollector interface {
	// RecordRequest 记录API请求
	RecordRequest(api, method string, statusCode int, duration time.Duration)

	// RecordError 记录错误
	RecordError(api, errorType string)

	// RecordRateLimitHit 记录速率限制命中
	RecordRateLimitHit(api string)

	// IncActiveRequests 增加活跃请求数
	IncActiveRequests()

	// DecActiveRequests 减少活跃请求数
	DecActiveRequests()
}

// HTTPClient 定义HTTP客户端接口。
//
// 这允许用户提供自定义的HTTP客户端实现，用于测试或特殊需求。
// 默认情况下，SDK使用标准的http.Client。
type HTTPClient interface {
	// Do 执行HTTP请求
	Do(req *http.Request) (*http.Response, error)
}

// Signer 定义请求签名接口。
//
// 这是内部接口，用于抽象LWA签名逻辑。
type Signer interface {
	// Sign 对HTTP请求进行签名
	Sign(ctx context.Context, req *http.Request) error
}

// RateLimiter 定义速率限制接口。
//
// 这是内部接口，用于抽象速率限制逻辑。
type RateLimiter interface {
	// Wait 等待直到可以发送请求
	Wait(ctx context.Context, api string) error

	// Update 更新API的速率限制
	Update(api string, rate float64, burst int)
}

// Tracer 定义分布式追踪接口。
//
// 用户可以提供自己的追踪实现，或使用SDK提供的OpenTelemetry实现。
// 默认情况下，SDK使用no-op tracer（不进行追踪）。
type Tracer interface {
	// StartSpan 开始一个新的span
	StartSpan(ctx context.Context, name string) (context.Context, Span)
}

// Span 表示一个追踪span。
type Span interface {
	// End 结束span
	End()

	// SetAttribute 设置属性
	SetAttribute(key string, value interface{})

	// RecordError 记录错误
	RecordError(err error)
}
