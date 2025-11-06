package async

import (
	"context"
	"errors"
	"sync"
	"time"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
)

var _ AwaitableAsyncCall[any] = (*asyncCallImpl[any])(nil)

// asyncCallImpl implements [AsyncCall]
type asyncCallImpl[T any] struct {
	done    chan struct{} // 完成信号
	retCode int
	reqId   uint64
	once    sync.Once // 保证只有一次结果

	result T
	err    error
}

func newAsyncCallImpl[T any](retCode int, reqId uint64) *asyncCallImpl[T] {
	var zeroT T
	asyncCtx := &asyncCallImpl[T]{
		done:    make(chan struct{}),
		retCode: retCode,
		reqId:   reqId,

		result: zeroT,
		err:    nil,
	}

	if retCode != 0 {
		// retCode非0，不需要等待 callback
		asyncCtx.errorWithCode(retCode)
		return asyncCtx
	}

	go func() {
		<-time.After(callbackTimeout) // callback timeout
		asyncCtx.error(ErrCallbackTimeout)
	}()

	return asyncCtx
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
	case <-ctx.Done(): // context cancel
		var zeroT T
		return zeroT, context.Cause(ctx)
	}
}

// complete 完成异步调用
func (a *asyncCallImpl[T]) complete(result T) {
	a.once.Do(func() {
		a.result = result
		close(a.done) // 通知等待方完成
	})
}

// error 完成异步调用
func (a *asyncCallImpl[T]) error(err error) {
	a.once.Do(func() {
		a.err = err
		close(a.done) // 通知等待方完成
	})
}

// errorWithCode 完成异步调用
func (a *asyncCallImpl[T]) errorWithCode(errorCode int) {
	err := errors.New(agrtm.GetErrorReason(int(errorCode)))
	a.error(err)
}
