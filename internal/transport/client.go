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

// Package transport 提供 HTTP 传输层功能。
package transport

import (
	"net/http"
	"time"
)

// Client HTTP 传输客户端。
type Client struct {
	httpClient *http.Client
	userAgent  string
	maxRetries int
}

// NewClient 创建 HTTP 传输客户端。
//
// 参数:
//   - timeout: 请求超时时间
//   - maxRetries: 最大重试次数
//   - userAgent: User-Agent 头
//
// 返回:
//   - *Client: HTTP 客户端实例
func NewClient(timeout time.Duration, maxRetries int, userAgent string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
		userAgent:  userAgent,
		maxRetries: maxRetries,
	}
}

// Do 执行 HTTP 请求。
//
// 参数:
//   - req: HTTP 请求
//
// 返回:
//   - *http.Response: HTTP 响应
//   - error: 错误信息
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	// 设置 User-Agent
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	// 执行请求（暂不实现重试，后续可添加）
	return c.httpClient.Do(req)
}

// GetHTTPClient 获取底层 HTTP 客户端。
func (c *Client) GetHTTPClient() *http.Client {
	return c.httpClient
}

