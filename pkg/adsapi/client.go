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

// Package adsapi provides Amazon Advertising API client.
//
// 本包实现了 Amazon Advertising API 的主客户端，提供统一的认证、
// 速率限制和错误处理。各个具体 API（如 Sponsored Products v3）
// 通过子包提供。
//
// 示例:
//
//	cfg := &adsapi.Config{
//	    ClientID:     "amzn1.application-oa2-client.xxx",
//	    ClientSecret: "your-client-secret",
//	    RefreshToken: "Atzr|xxx",
//	    ProfileID:    123456789,
//	    Region:       adsapi.RegionNA,
//	}
//
//	client, err := adsapi.NewClient(cfg)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// 官方文档: https://advertising.amazon.com/API/docs/
package adsapi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/vanling1111/amazon-ads-api-go-sdk/internal/auth"
	"github.com/vanling1111/amazon-ads-api-go-sdk/internal/ratelimit"
	"github.com/vanling1111/amazon-ads-api-go-sdk/internal/transport"
	"go.uber.org/zap"
)

// Client 是 Amazon Advertising API 的主客户端。
//
// 提供统一的认证、HTTP 传输、速率限制等基础功能。
// 各个具体 API（如 Sponsored Products）通过子包提供。
type Client struct {
	// config 是客户端配置
	config *Config

	// lwaClient 是 LWA 认证客户端（复用 SP-API）
	lwaClient *auth.Client

	// httpClient 是 HTTP 传输客户端（复用 SP-API）
	httpClient *transport.Client

	// rateLimitManager 是速率限制管理器（ADS-API 特定）
	rateLimitManager *ratelimit.Manager
}

// NewClient 创建新的 Amazon Advertising API 客户端。
//
// 参数:
//   - opts: 客户端配置选项（使用 Functional Options 模式）
//
// 返回:
//   - *Client: 客户端实例
//   - error: 错误信息（配置无效、认证失败等）
//
// 示例:
//
//	client, err := adsapi.NewClient(
//	    adsapi.WithRegion(models.RegionNorthAmerica),
//	    adsapi.WithCredentials("client-id", "secret", "token"),
//	    adsapi.WithProfileID(123456789),
//	    adsapi.WithDebug(),
//	)
func NewClient(opts ...ClientOption) (*Client, error) {
	// 1. 创建默认配置
	cfg := DefaultConfig()

	// 2. 应用用户选项
	for _, opt := range opts {
		opt(cfg)
	}

	// 3. 参数校验
	if err := cfg.Validate(); err != nil {
		return nil, Wrap(err, "config validation failed")
	}

	// 4. 初始化日志器
	logger := cfg.Logger
	if logger == nil {
		logger = zap.NewNop()
	}
	logger.Info("initializing Ads API client",
		zap.String("region", cfg.Region.Code),
		zap.Int64("profile_id", cfg.ProfileID),
	)

	// 3. 创建 LWA 认证客户端
	logger.Debug("creating LWA authentication client")
	lwaClient := auth.NewClient(
		cfg.ClientID,
		cfg.ClientSecret,
		cfg.RefreshToken,
		cfg.Region.LWAEndpoint,
	)
	logger.Info("LWA client created successfully")

	// 4. 创建 HTTP 传输客户端
	logger.Debug("creating HTTP transport client",
		zap.Duration("timeout", cfg.HTTPTimeout),
		zap.Int("max_retries", cfg.MaxRetries),
	)
	httpClient := transport.NewClient(
		cfg.HTTPTimeout,
		cfg.MaxRetries,
		"amazon-ads-api-go-sdk/1.0.0",
	)
	logger.Info("HTTP client created successfully")

	// 5. 创建速率限制管理器（默认 2 req/s, burst 10）
	rateLimitManager := ratelimit.NewManager(2.0, 10)

	return &Client{
		config:           cfg,
		lwaClient:        lwaClient,
		httpClient:       httpClient,
		rateLimitManager: rateLimitManager,
	}, nil
}

// DoRequest 执行 HTTP 请求（内部方法，供子包使用）。
//
// 功能:
//   - 自动添加认证头
//   - 速率限制控制
//   - 自动重试
//   - 错误处理
//
// 参数:
//   - ctx: 请求上下文
//   - req: HTTP 请求
//
// 返回:
//   - *http.Response: HTTP 响应
//   - error: 错误信息
func (c *Client) DoRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	startTime := time.Now()
	logger := c.config.Logger
	
	// 记录请求开始
	logger.Debug("starting API request",
		zap.String("method", req.Method),
		zap.String("url", req.URL.String()),
	)
	
	// Prometheus: 增加活跃请求数
	if c.config.Metrics != nil {
		c.config.Metrics.IncActiveRequests()
		defer c.config.Metrics.DecActiveRequests()
	}

	// 1. 获取访问令牌
	token, err := c.lwaClient.GetAccessToken(ctx)
	if err != nil {
		logger.Error("failed to get access token", zap.Error(err))
		if c.config.Metrics != nil {
			c.config.Metrics.RecordError("auth", "token_fetch_failed")
		}
		return nil, Wrap(err, "failed to get access token")
	}
	logger.Debug("access token obtained successfully")

	// 2. 添加认证头
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Amazon-Advertising-API-ClientId", c.config.ClientID)
	if c.config.ProfileID > 0 {
		req.Header.Set("Amazon-Advertising-API-Scope", fmt.Sprintf("%d", c.config.ProfileID))
	}

	// 3. 速率限制
	logger.Debug("checking rate limit")
	if err := c.rateLimitManager.Wait(ctx); err != nil {
		logger.Warn("rate limit exceeded", zap.Error(err))
		if c.config.Metrics != nil {
			c.config.Metrics.RecordRateLimitHit("ads_api")
		}
		return nil, Wrap(err, "rate limit exceeded")
	}

	// 4. 执行请求
	resp, err := c.httpClient.Do(req)
	duration := time.Since(startTime)
	
	if err != nil {
		logger.Error("HTTP request failed",
			zap.Error(err),
			zap.Duration("duration", duration),
		)
		if c.config.Metrics != nil {
			c.config.Metrics.RecordError("http", "request_failed")
		}
		return nil, Wrap(err, "HTTP request failed")
	}

	// 5. 记录成功的请求
	logger.Info("API request completed",
		zap.Int("status_code", resp.StatusCode),
		zap.Duration("duration", duration),
	)
	
	// Prometheus: 记录请求指标
	if c.config.Metrics != nil {
		c.config.Metrics.RecordRequest("ads_api", req.Method, resp.StatusCode, duration)
	}

	// 6. 检查错误
	if resp.StatusCode >= 400 {
		logger.Warn("API request returned error status",
			zap.Int("status_code", resp.StatusCode),
		)
		if c.config.Metrics != nil {
			c.config.Metrics.RecordError("api", fmt.Sprintf("status_%d", resp.StatusCode))
		}
		return resp, parseAPIError(resp)
	}

	return resp, nil
}

// GetConfig 返回客户端配置（只读）。
func (c *Client) GetConfig() *Config {
	return c.config
}

// parseAPIError 解析 API 错误响应（内部方法）。
func parseAPIError(resp *http.Response) error {
	// TODO: 实现详细的错误解析
	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    resp.Status,
	}
}

