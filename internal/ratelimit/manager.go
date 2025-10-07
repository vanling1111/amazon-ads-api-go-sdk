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

// Package ratelimit 提供速率限制功能。
package ratelimit

import (
	"context"
	"sync"

	"golang.org/x/time/rate"
)

// Manager 速率限制管理器。
type Manager struct {
	limiter *rate.Limiter
	mu      sync.RWMutex
}

// NewManager 创建速率限制管理器。
//
// 参数:
//   - requestsPerSecond: 每秒请求数
//   - burst: 突发请求数
//
// 返回:
//   - *Manager: 管理器实例
func NewManager(requestsPerSecond float64, burst int) *Manager {
	return &Manager{
		limiter: rate.NewLimiter(rate.Limit(requestsPerSecond), burst),
	}
}

// Wait 等待直到可以发送请求。
//
// 参数:
//   - ctx: 请求上下文
//
// 返回:
//   - error: 错误信息
func (m *Manager) Wait(ctx context.Context) error {
	return m.limiter.Wait(ctx)
}

// Allow 检查是否允许请求。
//
// 返回:
//   - bool: 是否允许
func (m *Manager) Allow() bool {
	return m.limiter.Allow()
}

// SetRate 动态更新速率限制。
//
// 参数:
//   - requestsPerSecond: 每秒请求数
//   - burst: 突发请求数
func (m *Manager) SetRate(requestsPerSecond float64, burst int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.limiter.SetLimit(rate.Limit(requestsPerSecond))
	m.limiter.SetBurst(burst)
}

