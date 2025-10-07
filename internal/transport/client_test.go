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

package transport

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := NewClient(30*time.Second, 3, "test-user-agent")

	assert.NotNil(t, client)
	assert.NotNil(t, client.httpClient)
	assert.Equal(t, "test-user-agent", client.userAgent)
	assert.Equal(t, 3, client.maxRetries)
}

func TestDo_Success(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "test-user-agent", r.Header.Get("User-Agent"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	client := NewClient(30*time.Second, 3, "test-user-agent")

	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDo_UserAgentOverride(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "custom-user-agent", r.Header.Get("User-Agent"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(30*time.Second, 3, "test-user-agent")

	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	// 手动设置 User-Agent
	req.Header.Set("User-Agent", "custom-user-agent")

	resp, err := client.Do(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetHTTPClient(t *testing.T) {
	client := NewClient(30*time.Second, 3, "test-user-agent")

	httpClient := client.GetHTTPClient()
	assert.NotNil(t, httpClient)
	assert.Equal(t, 30*time.Second, httpClient.Timeout)
}
