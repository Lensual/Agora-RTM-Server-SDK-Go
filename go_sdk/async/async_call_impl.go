package async

import (
	"context"
	"errors"
	"sync"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
)

var _ AwaitableAsyncCall[any] = (*asyncCallImpl[any])(nil)

// asyncCallImpl implements [AsyncCall]
type asyncCallImpl[T any] struct {
	done         chan struct{} // 通知通道
	resultLocker sync.Mutex    // 保护 result & err
	retCode      int
	reqId        uint64

	result T
	err    error
}

// ReturnCode implements [AsyncCall]
func (a *asyncCallImpl[T]) ReturnCode() int {
	return a.retCode
}

// RequestId implements [AsyncCall]
func (a *asyncCallImpl[T]) RequestId() uint64 {
	return a.reqId
}

// Await implements [AwaitableAsyncCall].
func (a *asyncCallImpl[T]) Await(ctx context.Context) (T, error) {
	select {
	case <-a.done:
		return a.result, a.err
	case <-ctx.Done(): // await timeout
		var zero T
		return zero, context.Cause(ctx)
	}
}

// complete 完成异步调用
func (a *asyncCallImpl[T]) complete(result T) {
	a.result = result
	close(a.done) // 唤醒等待方
}

// error 完成异步调用
func (a *asyncCallImpl[T]) error(errorCode int) {
	a.err = errors.New(agrtm.GetErrorReason(int(errorCode)))
	close(a.done) // 唤醒等待方
}
