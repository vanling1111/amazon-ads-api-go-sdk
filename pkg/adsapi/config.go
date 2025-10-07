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

// Package adsapi 提供 Amazon Advertising API 的 Go SDK。
//
// 此包是 SDK 的公开接口，提供简洁、易用的 API 访问方式。
//
// 基于官方 Ads API 文档:
//   - https://advertising.amazon.com/API/docs/
package adsapi

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/vanling1111/amazon-ads-api-go-sdk/internal/metrics"
	"github.com/vanling1111/amazon-ads-api-go-sdk/internal/models"
	"go.uber.org/zap"
)

// Config 定义 Ads API 客户端的配置。
type Config struct {
	// Region 是 Ads API 的区域（如 NA, EU, FE）。
	Region models.Region

	// ClientID 是 LWA 客户端 ID。
	ClientID string `validate:"required"`

	// ClientSecret 是 LWA 客户端密钥。
	ClientSecret string `validate:"required"`

	// RefreshToken 是 LWA 刷新令牌（用于 Regular 操作）。
	// 如果使用 Grantless 操作，可以为空。
	RefreshToken string `validate:"required_without=Scopes"`

	// Scopes 是 Grantless 操作所需的权限范围。
	// 如果为空，则使用 RefreshToken 进行 Regular 操作。
	Scopes []string `validate:"required_without=RefreshToken,dive,required"`

	// ProfileID 是广告主 Profile ID（可选，但大多数 API 需要）。
	ProfileID int64

	// HTTPTimeout 是 HTTP 请求超时时间。
	HTTPTimeout time.Duration `validate:"min=1s,max=5m"`

	// MaxRetries 是请求失败时的最大重试次数。
	MaxRetries int `validate:"min=0,max=10"`

	// RateLimitBuffer 是速率限制的缓冲比例（0.0-1.0）。
	// 例如 0.1 表示保留 10% 的速率限制作为缓冲。
	RateLimitBuffer float64 `validate:"min=0,max=1"`

	// Logger 是结构化日志记录器（使用 Uber Zap）。
	// 如果为 nil，使用 zap.NewNop() 作为默认值（不输出日志）。
	Logger *zap.Logger `validate:"-"`

	// Metrics 是 Prometheus 指标收集器（可选）。
	// 如果设置，SDK 会自动收集请求、错误、延迟等指标。
	Metrics *metrics.PrometheusMetrics `validate:"-"`

	// Debug 启用调试模式（详细日志）。
	Debug bool
}

// validate 全局验证器实例
var validate = validator.New()

// DefaultConfig 返回默认配置。
//
// 默认配置：
//   - HTTPTimeout: 30s
//   - MaxRetries: 3
//   - RateLimitBuffer: 0.1 (10%)
//   - Logger: zap.NewNop() (不输出日志)
//   - Debug: false
func DefaultConfig() *Config {
	return &Config{
		HTTPTimeout:     30 * time.Second,
		MaxRetries:      3,
		RateLimitBuffer: 0.1,
		Logger:          zap.NewNop(), // 默认不输出日志
		Debug:           false,
	}
}

// Validate 验证配置的有效性。
//
// 使用 validator 库进行声明式验证，支持丰富的验证规则。
//
// 返回值:
//   - error: 如果配置无效，返回错误
func (c *Config) Validate() error {
	// 使用 validator 进行结构体验证
	if err := validate.Struct(c); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			return formatValidationErrors(validationErrs)
		}
		return err
	}

	// 额外的业务逻辑验证
	if c.Region.Code == "" || c.Region.Endpoint == "" || c.Region.LWAEndpoint == "" {
		return ErrInvalidRegion
	}

	return nil
}

// formatValidationErrors 格式化验证错误
func formatValidationErrors(errs validator.ValidationErrors) error {
	messages := make([]string, 0, len(errs))

	for _, err := range errs {
		var msg string
		field := err.Field()

		switch err.Tag() {
		case "required":
			msg = fmt.Sprintf("%s is required", field)
		case "required_without":
			msg = fmt.Sprintf("%s is required when %s is not provided", field, err.Param())
		case "min":
			msg = fmt.Sprintf("%s must be at least %s", field, err.Param())
		case "max":
			msg = fmt.Sprintf("%s must not exceed %s", field, err.Param())
		case "dive":
			msg = fmt.Sprintf("%s contains invalid elements", field)
		default:
			msg = fmt.Sprintf("%s validation failed: %s", field, err.Tag())
		}

		messages = append(messages, msg)
	}

	if len(messages) == 1 {
		return fmt.Errorf("config validation failed: %s", messages[0])
	}

	return fmt.Errorf("config validation failed: %v", messages)
}

// ClientOption 定义客户端配置选项函数。
type ClientOption func(*Config)

// WithRegion 设置 API 区域。
//
// 参数:
//   - region: Ads API 区域（NA, EU, FE）
//
// 示例:
//
//	client := adsapi.NewClient(adsapi.WithRegion(models.RegionNorthAmerica))
func WithRegion(region models.Region) ClientOption {
	return func(c *Config) {
		c.Region = region
	}
}

// WithCredentials 设置 LWA 凭证（用于 Regular 操作）。
//
// 参数:
//   - clientID: LWA 客户端 ID
//   - clientSecret: LWA 客户端密钥
//   - refreshToken: LWA 刷新令牌
//
// 示例:
//
//	client := adsapi.NewClient(
//	    adsapi.WithCredentials("client-id", "client-secret", "refresh-token"),
//	)
func WithCredentials(clientID, clientSecret, refreshToken string) ClientOption {
	return func(c *Config) {
		c.ClientID = clientID
		c.ClientSecret = clientSecret
		c.RefreshToken = refreshToken
	}
}

// WithGrantlessCredentials 设置 Grantless 凭证。
//
// 参数:
//   - clientID: LWA 客户端 ID
//   - clientSecret: LWA 客户端密钥
//   - scopes: 权限范围
//
// 示例:
//
//	client := adsapi.NewClient(
//	    adsapi.WithGrantlessCredentials("client-id", "client-secret", []string{
//	        "advertising::campaign_management",
//	    }),
//	)
func WithGrantlessCredentials(clientID, clientSecret string, scopes []string) ClientOption {
	return func(c *Config) {
		c.ClientID = clientID
		c.ClientSecret = clientSecret
		c.Scopes = scopes
	}
}

// WithProfileID 设置广告主 Profile ID。
//
// Profile ID 是大多数 Ads API 调用的必需参数。
//
// 参数:
//   - profileID: 广告主 Profile ID
//
// 示例:
//
//	client := adsapi.NewClient(
//	    adsapi.WithRegion(models.RegionNA),
//	    adsapi.WithCredentials("client-id", "client-secret", "refresh-token"),
//	    adsapi.WithProfileID(123456789),
//	)
func WithProfileID(profileID int64) ClientOption {
	return func(c *Config) {
		c.ProfileID = profileID
	}
}

// WithHTTPTimeout 设置 HTTP 请求超时时间。
//
// 参数:
//   - timeout: 超时时间
//
// 示例:
//
//	client := adsapi.NewClient(adsapi.WithHTTPTimeout(60 * time.Second))
func WithHTTPTimeout(timeout time.Duration) ClientOption {
	return func(c *Config) {
		c.HTTPTimeout = timeout
	}
}

// WithMaxRetries 设置最大重试次数。
//
// 参数:
//   - maxRetries: 最大重试次数
//
// 示例:
//
//	client := adsapi.NewClient(adsapi.WithMaxRetries(5))
func WithMaxRetries(maxRetries int) ClientOption {
	return func(c *Config) {
		c.MaxRetries = maxRetries
	}
}

// WithRateLimitBuffer 设置速率限制缓冲比例。
//
// 参数:
//   - buffer: 缓冲比例（0.0-1.0）
//
// 示例:
//
//	client := adsapi.NewClient(adsapi.WithRateLimitBuffer(0.2)) // 20% 缓冲
func WithRateLimitBuffer(buffer float64) ClientOption {
	return func(c *Config) {
		c.RateLimitBuffer = buffer
	}
}

// WithLogger 设置结构化日志记录器。
//
// 参数:
//   - logger: Zap 日志记录器
//
// 示例:
//
//	logger, _ := zap.NewProduction()
//	client := adsapi.NewClient(adsapi.WithLogger(logger))
func WithLogger(logger *zap.Logger) ClientOption {
	return func(c *Config) {
		if logger != nil {
			c.Logger = logger
		}
	}
}

// WithMetrics 设置 Prometheus 指标收集器。
//
// 参数:
//   - metrics: Prometheus 指标收集器
//
// 示例:
//
//	metrics := metrics.NewPrometheusMetrics("adsapi")
//	client := adsapi.NewClient(adsapi.WithMetrics(metrics))
//	// 暴露指标端点
//	http.Handle("/metrics", promhttp.Handler())
//	go http.ListenAndServe(":9090", nil)
func WithMetrics(m *metrics.PrometheusMetrics) ClientOption {
	return func(c *Config) {
		if m != nil {
			c.Metrics = m
		}
	}
}

// WithDebug 启用调试模式（自动配置 Development 级别的日志）。
//
// 示例:
//
//	client := adsapi.NewClient(adsapi.WithDebug())
func WithDebug() ClientOption {
	return func(c *Config) {
		c.Debug = true
		// 自动创建开发模式的日志器
		if c.Logger == nil || c.Logger == zap.NewNop() {
			if logger, err := zap.NewDevelopment(); err == nil {
				c.Logger = logger
			}
		}
	}
}
