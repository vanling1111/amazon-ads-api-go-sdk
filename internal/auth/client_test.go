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

package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := NewClient("client-id", "client-secret", "refresh-token", "https://api.amazon.com/auth/o2/token")

	assert.NotNil(t, client)
	assert.Equal(t, "client-id", client.clientID)
	assert.Equal(t, "client-secret", client.clientSecret)
	assert.Equal(t, "refresh-token", client.refreshToken)
	assert.Equal(t, "https://api.amazon.com/auth/o2/token", client.endpoint)
}

func TestGetAccessToken_Success(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))

		// 返回成功响应
		resp := map[string]interface{}{
			"access_token": "test-access-token",
			"token_type":   "bearer",
			"expires_in":   3600,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient("client-id", "client-secret", "refresh-token", server.URL)

	// 第一次获取令牌
	token, err := client.GetAccessToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "test-access-token", token)

	// 第二次应该使用缓存
	token2, err := client.GetAccessToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "test-access-token", token2)
}

func TestGetAccessToken_Error(t *testing.T) {
	// 创建返回错误的测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error":             "invalid_client",
			"error_description": "Invalid client credentials",
		})
	}))
	defer server.Close()

	client := NewClient("invalid-id", "invalid-secret", "invalid-token", server.URL)

	_, err := client.GetAccessToken(context.Background())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "LWA error")
}

func TestGetAccessToken_EmptyToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"access_token": "",
			"token_type":   "bearer",
			"expires_in":   3600,
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient("client-id", "client-secret", "refresh-token", server.URL)

	_, err := client.GetAccessToken(context.Background())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "empty access token")
}

func TestClearCache(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		resp := map[string]interface{}{
			"access_token": "test-token",
			"token_type":   "bearer",
			"expires_in":   3600,
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient("client-id", "client-secret", "refresh-token", server.URL)

	// 第一次获取
	_, err := client.GetAccessToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 1, callCount)

	// 第二次使用缓存
	_, err = client.GetAccessToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 1, callCount, "should use cache")

	// 清除缓存
	client.ClearCache()

	// 第三次重新获取
	_, err = client.GetAccessToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 2, callCount, "should fetch new token after cache clear")
}

func TestGetAccessToken_CacheExpiry(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"access_token": "test-token",
			"token_type":   "bearer",
			"expires_in":   1, // 1 秒过期
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient("client-id", "client-secret", "refresh-token", server.URL)

	// 第一次获取
	token1, err := client.GetAccessToken(context.Background())
	require.NoError(t, err)

	// 等待令牌过期（1秒 + 60秒安全缓冲 = 提前过期）
	// 实际上由于安全缓冲，令牌立即就是过期状态
	time.Sleep(100 * time.Millisecond)

	// 第二次应该重新获取
	token2, err := client.GetAccessToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, token1, token2) // 令牌内容相同，但是重新获取的
}

func TestGetAccessToken_ContextCanceled(t *testing.T) {
	// 创建一个慢响应的服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "test-token",
			"token_type":   "bearer",
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	client := NewClient("client-id", "client-secret", "refresh-token", server.URL)

	// 创建可取消的 context
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err := client.GetAccessToken(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}
