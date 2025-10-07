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

// Package auth 提供 LWA OAuth 2.0 认证功能。
package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// Client LWA 认证客户端。
type Client struct {
	clientID     string
	clientSecret string
	refreshToken string
	endpoint     string
	httpClient   *http.Client
	
	// 令牌缓存
	mu           sync.RWMutex
	cachedToken  string
	tokenExpiry  time.Time
}

// NewClient 创建 LWA 认证客户端。
//
// 参数:
//   - clientID: LWA 客户端 ID
//   - clientSecret: LWA 客户端密钥
//   - refreshToken: LWA 刷新令牌
//   - endpoint: LWA 端点 URL
//
// 返回:
//   - *Client: 认证客户端实例
func NewClient(clientID, clientSecret, refreshToken, endpoint string) *Client {
	return &Client{
		clientID:     clientID,
		clientSecret: clientSecret,
		refreshToken: refreshToken,
		endpoint:     endpoint,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetAccessToken 获取访问令牌。
//
// 如果缓存中有有效令牌，直接返回；否则从 LWA 获取新令牌。
//
// 参数:
//   - ctx: 请求上下文
//
// 返回:
//   - string: 访问令牌
//   - error: 错误信息
func (c *Client) GetAccessToken(ctx context.Context) (string, error) {
	// 检查缓存
	c.mu.RLock()
	if c.cachedToken != "" && time.Now().Before(c.tokenExpiry) {
		token := c.cachedToken
		c.mu.RUnlock()
		return token, nil
	}
	c.mu.RUnlock()

	// 获取新令牌
	c.mu.Lock()
	defer c.mu.Unlock()

	// 双重检查
	if c.cachedToken != "" && time.Now().Before(c.tokenExpiry) {
		return c.cachedToken, nil
	}

	// 构建请求
	data := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {c.refreshToken},
		"client_id":     {c.clientID},
		"client_secret": {c.clientSecret},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("LWA error (HTTP %d): %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var tokenResp struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}

	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("parse response: %w", err)
	}

	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("empty access token")
	}

	// 缓存令牌（提前 60 秒过期）
	c.cachedToken = tokenResp.AccessToken
	c.tokenExpiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn-60) * time.Second)

	return tokenResp.AccessToken, nil
}

// ClearCache 清除令牌缓存。
func (c *Client) ClearCache() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cachedToken = ""
	c.tokenExpiry = time.Time{}
}

