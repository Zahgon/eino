package callbacks

import (
	"context"

	"github.com/cloudwego/eino/components"
	"github.com/cloudwego/eino/schema"
)

func OnStart[T any](ctx context.Context, input T) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func OnEnd[T any](ctx context.Context, output T) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func OnStartWithStreamInput[T any](ctx context.Context, input *schema.StreamReader[T]) (
	nextCtx context.Context, newStreamReader *schema.StreamReader[T]) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func OnEndWithStreamOutput[T any](ctx context.Context, output *schema.StreamReader[T]) (
	nextCtx context.Context, newStreamReader *schema.StreamReader[T]) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func OnError(ctx context.Context, err error) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func EnsureRunInfo(ctx context.Context, typ string, comp components.Component) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ReuseHandlers(ctx context.Context, info *RunInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func InitCallbacks(ctx context.Context, info *RunInfo, handlers ...Handler) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
