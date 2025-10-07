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
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vanling1111/amazon-ads-api-go-sdk/internal/models"
)

// TestNewClient 测试创建客户端
func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		opts    []ClientOption
		wantErr bool
		errMsg  string
	}{
		{
			name: "成功创建客户端 - Regular 操作",
			opts: []ClientOption{
				WithRegion(models.RegionNorthAmerica),
				WithCredentials("test-client-id", "test-secret", "test-token"),
				WithProfileID(123456789),
			},
			wantErr: false,
		},
		{
			name: "成功创建客户端 - Grantless 操作",
			opts: []ClientOption{
				WithRegion(models.RegionNorthAmerica),
				WithGrantlessCredentials("test-client-id", "test-secret", []string{"advertising::campaign_management"}),
			},
			wantErr: false,
		},
		{
			name: "成功创建客户端 - 使用所有选项",
			opts: []ClientOption{
				WithRegion(models.RegionNorthAmerica),
				WithCredentials("test-client-id", "test-secret", "test-token"),
				WithProfileID(123456789),
				WithHTTPTimeout(60 * time.Second),
				WithMaxRetries(5),
				WithRateLimitBuffer(0.2),
				WithDebug(),
			},
			wantErr: false,
		},
		{
			name: "配置验证失败 - 缺少凭证",
			opts: []ClientOption{
				WithRegion(models.RegionNorthAmerica),
			},
			wantErr: true,
			errMsg:  "config validation failed",
		},
		{
			name: "配置验证失败 - 无效的 Region",
			opts: []ClientOption{
				WithCredentials("test-client-id", "test-secret", "test-token"),
			},
			wantErr: true,
			errMsg:  "invalid region",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.opts...)

			if tt.wantErr {
				require.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
				assert.Nil(t, client)
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)
				assert.NotNil(t, client.config)
				assert.NotNil(t, client.lwaClient)
				assert.NotNil(t, client.httpClient)
				assert.NotNil(t, client.rateLimitManager)
			}
		})
	}
}

// TestClient_GetConfig 测试获取配置
func TestClient_GetConfig(t *testing.T) {
	client, err := NewClient(
		WithRegion(models.RegionNorthAmerica),
		WithCredentials("test-client-id", "test-secret", "test-token"),
		WithProfileID(123456789),
	)
	require.NoError(t, err)

	cfg := client.GetConfig()
	require.NotNil(t, cfg)
	assert.Equal(t, models.RegionNorthAmerica, cfg.Region)
	assert.Equal(t, "test-client-id", cfg.ClientID)
	assert.Equal(t, int64(123456789), cfg.ProfileID)
}

// TestClient_DoRequest 测试执行 HTTP 请求
func TestClient_DoRequest(t *testing.T) {
	// 创建 mock LWA 服务器
	lwaServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token":"mock-token","token_type":"bearer","expires_in":3600}`))
	}))
	defer lwaServer.Close()

	tests := []struct {
		name           string
		setupServer    func() *httptest.Server
		setupRequest   func(serverURL string) *http.Request
		expectedStatus int
		wantErr        bool
	}{
		{
			name: "成功的 GET 请求",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// 验证认证头
					assert.NotEmpty(t, r.Header.Get("Authorization"))
					assert.NotEmpty(t, r.Header.Get("Amazon-Advertising-API-ClientId"))
					assert.NotEmpty(t, r.Header.Get("Amazon-Advertising-API-Scope"))

					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"success":true}`))
				}))
			},
			setupRequest: func(serverURL string) *http.Request {
				req, _ := http.NewRequest("GET", serverURL+"/campaigns", nil)
				return req
			},
			expectedStatus: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "API 返回 400 错误",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(`{"code":"INVALID_ARGUMENT","message":"Invalid request"}`))
				}))
			},
			setupRequest: func(serverURL string) *http.Request {
				req, _ := http.NewRequest("POST", serverURL+"/campaigns", nil)
				return req
			},
			expectedStatus: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "API 返回 401 未授权",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte(`{"code":"UNAUTHORIZED","message":"Invalid token"}`))
				}))
			},
			setupRequest: func(serverURL string) *http.Request {
				req, _ := http.NewRequest("GET", serverURL+"/campaigns", nil)
				return req
			},
			expectedStatus: http.StatusUnauthorized,
			wantErr:        true,
		},
		{
			name: "API 返回 429 速率限制",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusTooManyRequests)
					w.Write([]byte(`{"code":"RATE_LIMIT_EXCEEDED","message":"Too many requests"}`))
				}))
			},
			setupRequest: func(serverURL string) *http.Request {
				req, _ := http.NewRequest("GET", serverURL+"/campaigns", nil)
				return req
			},
			expectedStatus: http.StatusTooManyRequests,
			wantErr:        true,
		},
		{
			name: "API 返回 500 服务器错误",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(`{"code":"INTERNAL_ERROR","message":"Server error"}`))
				}))
			},
			setupRequest: func(serverURL string) *http.Request {
				req, _ := http.NewRequest("GET", serverURL+"/campaigns", nil)
				return req
			},
			expectedStatus: http.StatusInternalServerError,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建 API mock 服务器
			apiServer := tt.setupServer()
			defer apiServer.Close()

			// 创建客户端
			client, err := NewClient(
				WithRegion(models.Region{
					Code:        "NA",
					Endpoint:    apiServer.URL,
					LWAEndpoint: lwaServer.URL,
				}),
				WithCredentials("test-client-id", "test-secret", "test-token"),
				WithProfileID(123456789),
			)
			require.NoError(t, err)

			// 创建请求
			req := tt.setupRequest(apiServer.URL)
			req = req.WithContext(context.Background())

			// 执行请求
			resp, err := client.DoRequest(context.Background(), req)

			if tt.wantErr {
				// 应该返回错误或者响应状态码 >= 400
				if err != nil {
					assert.Error(t, err)
				} else {
					require.NotNil(t, resp)
					assert.Equal(t, tt.expectedStatus, resp.StatusCode)
					resp.Body.Close()
				}
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				assert.Equal(t, tt.expectedStatus, resp.StatusCode)
				resp.Body.Close()
			}
		})
	}
}

// TestClient_DoRequest_ContextCancellation 测试上下文取消
func TestClient_DoRequest_ContextCancellation(t *testing.T) {
	// 创建一个会延迟响应的服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// 创建 LWA mock 服务器
	lwaServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token":"mock-token","token_type":"bearer","expires_in":3600}`))
	}))
	defer lwaServer.Close()

	// 创建客户端
	client, err := NewClient(
		WithRegion(models.Region{
			Code:        "NA",
			Endpoint:    server.URL,
			LWAEndpoint: lwaServer.URL,
		}),
		WithCredentials("test-client-id", "test-secret", "test-token"),
		WithHTTPTimeout(5 * time.Second),
	)
	require.NoError(t, err)

	// 创建会立即取消的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequest("GET", server.URL+"/campaigns", nil)
	req = req.WithContext(ctx)

	// 执行请求（应该因为上下文取消而失败）
	_, err = client.DoRequest(ctx, req)

	// 应该返回上下文相关的错误
	assert.Error(t, err)
}

// TestParseAPIError 测试 API 错误解析
func TestParseAPIError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		wantCode   int
	}{
		{
			name:       "400 错误",
			statusCode: http.StatusBadRequest,
			body:       `{"code":"INVALID_ARGUMENT","message":"Invalid request"}`,
			wantCode:   http.StatusBadRequest,
		},
		{
			name:       "401 错误",
			statusCode: http.StatusUnauthorized,
			body:       `{"code":"UNAUTHORIZED","message":"Invalid token"}`,
			wantCode:   http.StatusUnauthorized,
		},
		{
			name:       "429 错误",
			statusCode: http.StatusTooManyRequests,
			body:       `{"code":"RATE_LIMIT_EXCEEDED","message":"Too many requests"}`,
			wantCode:   http.StatusTooManyRequests,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建模拟响应
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Status:     http.StatusText(tt.statusCode),
			}

			// 解析错误
			err := parseAPIError(resp)

			require.Error(t, err)
			apiErr, ok := err.(*APIError)
			require.True(t, ok, "错误应该是 APIError 类型")
			assert.Equal(t, tt.wantCode, apiErr.StatusCode)
		})
	}
}

// Benchmark 测试

// BenchmarkNewClient 基准测试创建客户端
func BenchmarkNewClient(b *testing.B) {
	opts := []ClientOption{
		WithRegion(models.RegionNorthAmerica),
		WithCredentials("test-client-id", "test-secret", "test-token"),
		WithProfileID(123456789),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewClient(opts...)
	}
}

// BenchmarkClient_GetConfig 基准测试获取配置
func BenchmarkClient_GetConfig(b *testing.B) {
	client, _ := NewClient(
		WithRegion(models.RegionNorthAmerica),
		WithCredentials("test-client-id", "test-secret", "test-token"),
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.GetConfig()
	}
}

