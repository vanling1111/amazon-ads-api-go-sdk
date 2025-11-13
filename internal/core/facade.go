// Package core 提供SDK的核心协调层。
//
// Facade模式封装所有内部组件，提供统一的接口。
package core

import (
	"github.com/vanling1111/amazon-ads-api-go-sdk/internal/auth"
	"github.com/vanling1111/amazon-ads-api-go-sdk/internal/ratelimit"
	"github.com/vanling1111/amazon-ads-api-go-sdk/internal/transport"
)

// Facade 封装所有内部组件，提供统一的访问接口。
//
// 这是Facade设计模式的实现，隐藏内部复杂性，
// 为pkg/adsapi提供简洁的接口。
type Facade struct {
	// lwaClient LWA认证客户端
	lwaClient *auth.Client

	// httpClient HTTP传输客户端
	httpClient *transport.Client

	// rateLimitManager 速率限制管理器
	rateLimitManager *ratelimit.Manager
}

// NewFacade 创建核心门面实例。
//
// 参数:
//   - lwaClient: LWA认证客户端
//   - httpClient: HTTP传输客户端
//   - rateLimitManager: 速率限制管理器
//
// 返回值:
//   - *Facade: 门面实例
func NewFacade(
	lwaClient *auth.Client,
	httpClient *transport.Client,
	rateLimitManager *ratelimit.Manager,
) *Facade {
	return &Facade{
		lwaClient:        lwaClient,
		httpClient:       httpClient,
		rateLimitManager: rateLimitManager,
	}
}

// GetLWAClient 返回LWA认证客户端。
func (f *Facade) GetLWAClient() *auth.Client {
	return f.lwaClient
}

// GetHTTPClient 返回HTTP客户端。
//
// 用于需要直接访问HTTP客户端的场景。
func (f *Facade) GetHTTPClient() *transport.Client {
	return f.httpClient
}

// GetRateLimitManager 返回速率限制管理器。
//
// 用于需要直接访问速率限制管理器的场景。
func (f *Facade) GetRateLimitManager() *ratelimit.Manager {
	return f.rateLimitManager
}

