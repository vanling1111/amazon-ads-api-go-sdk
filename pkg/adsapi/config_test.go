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

package adsapi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vanling1111/amazon-ads-api-go-sdk/internal/models"
	"go.uber.org/zap"
)

// TestDefaultConfig 测试默认配置
func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	require.NotNil(t, cfg, "默认配置不应该为 nil")
	assert.Equal(t, 30*time.Second, cfg.HTTPTimeout, "默认超时应该是 30 秒")
	assert.Equal(t, 3, cfg.MaxRetries, "默认重试次数应该是 3")
	assert.Equal(t, 0.1, cfg.RateLimitBuffer, "默认速率限制缓冲应该是 0.1")
	assert.NotNil(t, cfg.Logger, "默认日志器不应该为 nil")
	assert.False(t, cfg.Debug, "默认不应该启用调试模式")
}

// TestConfig_Validate 测试配置验证
func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() *Config
		wantErr bool
		errMsg  string
	}{
		{
			name: "有效配置 - Regular 操作",
			setup: func() *Config {
				cfg := DefaultConfig()
				cfg.Region = models.RegionNorthAmerica
				cfg.ClientID = "test-client-id"
				cfg.ClientSecret = "test-client-secret"
				cfg.RefreshToken = "test-refresh-token"
				return cfg
			},
			wantErr: false,
		},
		{
			name: "有效配置 - Grantless 操作",
			setup: func() *Config {
				cfg := DefaultConfig()
				cfg.Region = models.RegionNorthAmerica
				cfg.ClientID = "test-client-id"
				cfg.ClientSecret = "test-client-secret"
				cfg.Scopes = []string{"advertising::campaign_management"}
				return cfg
			},
			wantErr: false,
		},
		{
			name: "缺少 ClientID",
			setup: func() *Config {
				cfg := DefaultConfig()
				cfg.Region = models.RegionNorthAmerica
				cfg.ClientSecret = "test-client-secret"
				cfg.RefreshToken = "test-refresh-token"
				return cfg
			},
			wantErr: true,
			errMsg:  "ClientID is required",
		},
		{
			name: "缺少 ClientSecret",
			setup: func() *Config {
				cfg := DefaultConfig()
				cfg.Region = models.RegionNorthAmerica
				cfg.ClientID = "test-client-id"
				cfg.RefreshToken = "test-refresh-token"
				return cfg
			},
			wantErr: true,
			errMsg:  "ClientSecret is required",
		},
		{
			name: "既没有 RefreshToken 也没有 Scopes",
			setup: func() *Config {
				cfg := DefaultConfig()
				cfg.Region = models.RegionNorthAmerica
				cfg.ClientID = "test-client-id"
				cfg.ClientSecret = "test-client-secret"
				return cfg
			},
			wantErr: true,
			errMsg:  "RefreshToken is required",
		},
		{
			name: "无效的 Region",
			setup: func() *Config {
				cfg := DefaultConfig()
				cfg.ClientID = "test-client-id"
				cfg.ClientSecret = "test-client-secret"
				cfg.RefreshToken = "test-refresh-token"
				cfg.Region = models.Region{} // 空 region
				return cfg
			},
			wantErr: true,
			errMsg:  "invalid region",
		},
		{
			name: "HTTPTimeout 超出范围 - 太小",
			setup: func() *Config {
				cfg := DefaultConfig()
				cfg.Region = models.RegionNorthAmerica
				cfg.ClientID = "test-client-id"
				cfg.ClientSecret = "test-client-secret"
				cfg.RefreshToken = "test-refresh-token"
				cfg.HTTPTimeout = 500 * time.Millisecond // 小于 1s
				return cfg
			},
			wantErr: true,
			errMsg:  "HTTPTimeout must be at least 1s",
		},
		{
			name: "HTTPTimeout 超出范围 - 太大",
			setup: func() *Config {
				cfg := DefaultConfig()
				cfg.Region = models.RegionNorthAmerica
				cfg.ClientID = "test-client-id"
				cfg.ClientSecret = "test-client-secret"
				cfg.RefreshToken = "test-refresh-token"
				cfg.HTTPTimeout = 10 * time.Minute // 大于 5m
				return cfg
			},
			wantErr: true,
			errMsg:  "HTTPTimeout must not exceed 5m",
		},
		{
			name: "MaxRetries 为负数",
			setup: func() *Config {
				cfg := DefaultConfig()
				cfg.Region = models.RegionNorthAmerica
				cfg.ClientID = "test-client-id"
				cfg.ClientSecret = "test-client-secret"
				cfg.RefreshToken = "test-refresh-token"
				cfg.MaxRetries = -1
				return cfg
			},
			wantErr: true,
			errMsg:  "MaxRetries must be at least 0",
		},
		{
			name: "MaxRetries 超过上限",
			setup: func() *Config {
				cfg := DefaultConfig()
				cfg.Region = models.RegionNorthAmerica
				cfg.ClientID = "test-client-id"
				cfg.ClientSecret = "test-client-secret"
				cfg.RefreshToken = "test-refresh-token"
				cfg.MaxRetries = 15
				return cfg
			},
			wantErr: true,
			errMsg:  "MaxRetries must not exceed 10",
		},
		{
			name: "RateLimitBuffer 为负数",
			setup: func() *Config {
				cfg := DefaultConfig()
				cfg.Region = models.RegionNorthAmerica
				cfg.ClientID = "test-client-id"
				cfg.ClientSecret = "test-client-secret"
				cfg.RefreshToken = "test-refresh-token"
				cfg.RateLimitBuffer = -0.1
				return cfg
			},
			wantErr: true,
			errMsg:  "RateLimitBuffer must be at least 0",
		},
		{
			name: "RateLimitBuffer 超过 1",
			setup: func() *Config {
				cfg := DefaultConfig()
				cfg.Region = models.RegionNorthAmerica
				cfg.ClientID = "test-client-id"
				cfg.ClientSecret = "test-client-secret"
				cfg.RefreshToken = "test-refresh-token"
				cfg.RateLimitBuffer = 1.5
				return cfg
			},
			wantErr: true,
			errMsg:  "RateLimitBuffer must not exceed 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.setup()
			err := cfg.Validate()

			if tt.wantErr {
				require.Error(t, err, "期望验证失败")
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg, "错误消息应该包含预期内容")
				}
			} else {
				require.NoError(t, err, "期望验证成功")
			}
		})
	}
}

// TestClientOptions 测试客户端选项
func TestClientOptions(t *testing.T) {
	t.Run("WithRegion", func(t *testing.T) {
		cfg := DefaultConfig()
		opt := WithRegion(models.RegionEurope)
		opt(cfg)
		assert.Equal(t, models.RegionEurope, cfg.Region)
	})

	t.Run("WithCredentials", func(t *testing.T) {
		cfg := DefaultConfig()
		opt := WithCredentials("client-id", "client-secret", "refresh-token")
		opt(cfg)
		assert.Equal(t, "client-id", cfg.ClientID)
		assert.Equal(t, "client-secret", cfg.ClientSecret)
		assert.Equal(t, "refresh-token", cfg.RefreshToken)
	})

	t.Run("WithGrantlessCredentials", func(t *testing.T) {
		cfg := DefaultConfig()
		scopes := []string{"advertising::campaign_management"}
		opt := WithGrantlessCredentials("client-id", "client-secret", scopes)
		opt(cfg)
		assert.Equal(t, "client-id", cfg.ClientID)
		assert.Equal(t, "client-secret", cfg.ClientSecret)
		assert.Equal(t, scopes, cfg.Scopes)
	})

	t.Run("WithProfileID", func(t *testing.T) {
		cfg := DefaultConfig()
		opt := WithProfileID(123456789)
		opt(cfg)
		assert.Equal(t, int64(123456789), cfg.ProfileID)
	})

	t.Run("WithHTTPTimeout", func(t *testing.T) {
		cfg := DefaultConfig()
		opt := WithHTTPTimeout(60 * time.Second)
		opt(cfg)
		assert.Equal(t, 60*time.Second, cfg.HTTPTimeout)
	})

	t.Run("WithMaxRetries", func(t *testing.T) {
		cfg := DefaultConfig()
		opt := WithMaxRetries(5)
		opt(cfg)
		assert.Equal(t, 5, cfg.MaxRetries)
	})

	t.Run("WithRateLimitBuffer", func(t *testing.T) {
		cfg := DefaultConfig()
		opt := WithRateLimitBuffer(0.2)
		opt(cfg)
		assert.Equal(t, 0.2, cfg.RateLimitBuffer)
	})

	t.Run("WithLogger", func(t *testing.T) {
		cfg := DefaultConfig()
		logger, _ := zap.NewDevelopment()
		opt := WithLogger(logger)
		opt(cfg)
		assert.Equal(t, logger, cfg.Logger)
	})

	t.Run("WithLogger - nil 不应该覆盖默认值", func(t *testing.T) {
		cfg := DefaultConfig()
		originalLogger := cfg.Logger
		opt := WithLogger(nil)
		opt(cfg)
		assert.Equal(t, originalLogger, cfg.Logger)
	})

	t.Run("WithDebug", func(t *testing.T) {
		cfg := DefaultConfig()
		opt := WithDebug()
		opt(cfg)
		assert.True(t, cfg.Debug)
		assert.NotNil(t, cfg.Logger)
		// Debug 模式会创建开发日志器
	})
}

// TestClientOptions_ChainedUsage 测试链式使用选项
func TestClientOptions_ChainedUsage(t *testing.T) {
	cfg := DefaultConfig()

	// 应用多个选项
	options := []ClientOption{
		WithRegion(models.RegionNorthAmerica),
		WithCredentials("client-id", "secret", "token"),
		WithProfileID(123456789),
		WithHTTPTimeout(45 * time.Second),
		WithMaxRetries(5),
		WithRateLimitBuffer(0.15),
		WithDebug(),
	}

	for _, opt := range options {
		opt(cfg)
	}

	// 验证所有选项都生效
	assert.Equal(t, models.RegionNorthAmerica, cfg.Region)
	assert.Equal(t, "client-id", cfg.ClientID)
	assert.Equal(t, "secret", cfg.ClientSecret)
	assert.Equal(t, "token", cfg.RefreshToken)
	assert.Equal(t, int64(123456789), cfg.ProfileID)
	assert.Equal(t, 45*time.Second, cfg.HTTPTimeout)
	assert.Equal(t, 5, cfg.MaxRetries)
	assert.Equal(t, 0.15, cfg.RateLimitBuffer)
	assert.True(t, cfg.Debug)
	assert.NotNil(t, cfg.Logger)
}

// TestFormatValidationErrors 测试验证错误格式化
func TestFormatValidationErrors(t *testing.T) {
	// 测试单个错误
	t.Run("单个验证错误", func(t *testing.T) {
		cfg := DefaultConfig()
		cfg.Region = models.RegionNorthAmerica
		// 缺少所有必需的凭证字段
		cfg.ClientID = ""

		err := cfg.Validate()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "config validation failed")
		assert.Contains(t, err.Error(), "ClientID")
	})

	// 测试多个错误
	t.Run("多个验证错误", func(t *testing.T) {
		cfg := DefaultConfig()
		cfg.Region = models.RegionNorthAmerica
		cfg.ClientID = ""
		cfg.ClientSecret = ""

		err := cfg.Validate()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "config validation failed")
	})
}

// Benchmark 测试

// BenchmarkDefaultConfig 基准测试创建默认配置
func BenchmarkDefaultConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DefaultConfig()
	}
}

// BenchmarkConfig_Validate 基准测试配置验证
func BenchmarkConfig_Validate(b *testing.B) {
	cfg := DefaultConfig()
	cfg.Region = models.RegionNorthAmerica
	cfg.ClientID = "test-client-id"
	cfg.ClientSecret = "test-client-secret"
	cfg.RefreshToken = "test-refresh-token"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.Validate()
	}
}

// BenchmarkClientOptions 基准测试客户端选项应用
func BenchmarkClientOptions(b *testing.B) {
	options := []ClientOption{
		WithRegion(models.RegionNorthAmerica),
		WithCredentials("client-id", "secret", "token"),
		WithProfileID(123456789),
		WithHTTPTimeout(45 * time.Second),
		WithMaxRetries(5),
		WithRateLimitBuffer(0.15),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cfg := DefaultConfig()
		for _, opt := range options {
			opt(cfg)
		}
	}
}
