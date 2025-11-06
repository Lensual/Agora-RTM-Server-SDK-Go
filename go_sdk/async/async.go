// Package async 提供了 RTM 可等待的异步调用方式。
//
// 该包主要用于处理 RTM 异步回调，避免业务逻辑跨越至基础设施 EventHandler 中。
//
// 回调上下文存储至公共的 reqIdMap 中，并由 reqIdMapLocker 进行线程安全保护。
//
// 使用方法：
//  1. [NewAsyncCallRtmEventHandler] 创建一个新的 RTM 事件处理器。你可能需要调用 utils 包中的 NewCompositeRtmEventHandler 来组合多个 EventHandler。
//  2. 使用 [Call] 包装 RTM 的异步方法，他会保存 requestId 并返回一个 [AwaitableAsyncCall]。
//  3. [AwaitableAsyncCall].Await() 等待异步调用完成。
package async

import (
	"context"
	"sync"
)

var reqIdMapLocker sync.Mutex
var reqIdMap map[uint64]AsyncCall = make(map[uint64]AsyncCall) //TODO 如果 asyncCall.Complete 没有被调用，这里可能会产生内存泄露。例如retCode是非0的情况

type RtmAsyncFunc func() (ret int, reqId uint64)

// Call 异步调用 RTM ，返回一个可等待的 [AwaitableAsyncCall]
// 需要搭配 `NewCompositeRtmEventHandler` 使用
func Call[T any](callFunc RtmAsyncFunc) AwaitableAsyncCall[T] {
	reqIdMapLocker.Lock()
	defer reqIdMapLocker.Unlock()

	var zeroT T
	asyncCtx := &asyncCallImpl[T]{
		done:         make(chan struct{}),
		resultLocker: sync.Mutex{},
		retCode:      0,
		reqId:        0,

		result: zeroT,
		err:    nil,
	}
	asyncCtx.retCode, asyncCtx.reqId = callFunc()
	reqIdMap[asyncCtx.reqId] = asyncCtx

	return asyncCtx
}

// AsyncCall RTM 异步调用
type AsyncCall interface {
	ReturnCode() int
	RequestId() uint64
}

// AwaitableAsyncCall 可等待的 RTM 异步调用
type AwaitableAsyncCall[T any] interface {
	AsyncCall
	Await(ctx context.Context) (T, error)
}
