package callbacks

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

type HandlerBuilder struct {
	onStartFn                func(ctx context.Context, info *RunInfo, input CallbackInput) context.Context
	onEndFn                  func(ctx context.Context, info *RunInfo, output CallbackOutput) context.Context
	onErrorFn                func(ctx context.Context, info *RunInfo, err error) context.Context
	onStartWithStreamInputFn func(ctx context.Context, info *RunInfo, input *schema.StreamReader[CallbackInput]) context.Context
	onEndWithStreamOutputFn  func(ctx context.Context, info *RunInfo, output *schema.StreamReader[CallbackOutput]) context.Context
}

type handlerImpl struct {
	HandlerBuilder
}

func (hb *handlerImpl) OnStart(ctx context.Context, info *RunInfo, input CallbackInput) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (hb *handlerImpl) OnEnd(ctx context.Context, info *RunInfo, output CallbackOutput) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (hb *handlerImpl) OnError(ctx context.Context, info *RunInfo, err error) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (hb *handlerImpl) OnStartWithStreamInput(ctx context.Context, info *RunInfo,
	input *schema.StreamReader[CallbackInput]) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (hb *handlerImpl) OnEndWithStreamOutput(ctx context.Context, info *RunInfo,
	output *schema.StreamReader[CallbackOutput]) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (hb *handlerImpl) Needed(_ context.Context, _ *RunInfo, timing CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

func NewHandlerBuilder() *HandlerBuilder { _ = "STUB: not implemented"; return nil }

func (hb *HandlerBuilder) OnStartFn(
	fn func(ctx context.Context, info *RunInfo, input CallbackInput) context.Context) *HandlerBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (hb *HandlerBuilder) OnEndFn(
	fn func(ctx context.Context, info *RunInfo, output CallbackOutput) context.Context) *HandlerBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (hb *HandlerBuilder) OnErrorFn(
	fn func(ctx context.Context, info *RunInfo, err error) context.Context) *HandlerBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (hb *HandlerBuilder) OnStartWithStreamInputFn(
	fn func(ctx context.Context, info *RunInfo, input *schema.StreamReader[CallbackInput]) context.Context) *HandlerBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (hb *HandlerBuilder) OnEndWithStreamOutputFn(
	fn func(ctx context.Context, info *RunInfo, output *schema.StreamReader[CallbackOutput]) context.Context) *HandlerBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (hb *HandlerBuilder) Build() Handler { _ = "STUB: not implemented"; return *new(Handler) }
