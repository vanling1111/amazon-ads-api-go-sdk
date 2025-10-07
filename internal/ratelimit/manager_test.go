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

package ratelimit

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewManager(t *testing.T) {
	manager := NewManager(10.0, 20)

	assert.NotNil(t, manager)
	assert.NotNil(t, manager.limiter)
}

func TestWait_Success(t *testing.T) {
	manager := NewManager(100.0, 10) // 高速率，方便测试

	ctx := context.Background()
	err := manager.Wait(ctx)

	require.NoError(t, err)
}

func TestWait_ContextCanceled(t *testing.T) {
	manager := NewManager(0.1, 1) // 低速率

	// 创建已取消的 context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := manager.Wait(ctx)

	assert.Error(t, err)
	assert.Equal(t, context.Canceled, err)
}

func TestAllow(t *testing.T) {
	manager := NewManager(10.0, 5)

	// 前5个请求应该立即允许（burst）
	for i := 0; i < 5; i++ {
		allowed := manager.Allow()
		assert.True(t, allowed, "request %d should be allowed", i)
	}

	// 第6个请求应该被限制
	allowed := manager.Allow()
	assert.False(t, allowed, "request 6 should be rate limited")
}

func TestSetRate(t *testing.T) {
	manager := NewManager(1.0, 1)

	// 消耗初始令牌
	assert.True(t, manager.Allow())
	assert.False(t, manager.Allow())

	// 更新速率为更高值
	manager.SetRate(100.0, 10)

	// 等待一小段时间让新速率生效
	time.Sleep(20 * time.Millisecond)

	// 现在应该有更多令牌可用
	assert.True(t, manager.Allow())
}

func TestRateLimit_Timing(t *testing.T) {
	// 创建 2 req/s 的限制器
	manager := NewManager(2.0, 1)

	ctx := context.Background()

	// 第一个请求应该立即通过
	start := time.Now()
	err := manager.Wait(ctx)
	require.NoError(t, err)
	elapsed := time.Since(start)
	assert.Less(t, elapsed, 10*time.Millisecond, "first request should be immediate")

	// 第二个请求需要等待（因为 burst=1）
	start = time.Now()
	err = manager.Wait(ctx)
	require.NoError(t, err)
	elapsed = time.Since(start)
	// 应该等待约 0.5 秒（1/2 req/s）
	assert.GreaterOrEqual(t, elapsed, 400*time.Millisecond, "should wait for rate limit")
	assert.Less(t, elapsed, 600*time.Millisecond, "should not wait too long")
}

func TestConcurrentAccess(t *testing.T) {
	manager := NewManager(10.0, 10)

	ctx := context.Background()
	done := make(chan bool, 20)

	// 并发访问
	for i := 0; i < 20; i++ {
		go func() {
			err := manager.Wait(ctx)
			assert.NoError(t, err)
			done <- true
		}()
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 20; i++ {
		select {
		case <-done:
			// Success
		case <-time.After(5 * time.Second):
			t.Fatal("timeout waiting for goroutines")
		}
	}
}
