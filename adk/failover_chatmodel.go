package adk

import (
	"context"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

type failoverCurrentModelKey struct{}

func typedSetFailoverCurrentModel[M MessageType](ctx context.Context, currentModel model.BaseModel[M]) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func typedGetFailoverCurrentModel[M MessageType](ctx context.Context) (model.BaseModel[M], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type failoverHasMoreAttemptsKey struct{}

func withFailoverHasMoreAttempts(ctx context.Context, hasMore bool) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getFailoverHasMoreAttempts(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

type typedFailoverProxyModel[M MessageType] struct {
}

func (m *typedFailoverProxyModel[M]) prepareTarget(ctx context.Context) (model.BaseModel[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *typedFailoverProxyModel[M]) Generate(ctx context.Context, input []M, opts ...model.Option) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func (m *typedFailoverProxyModel[M]) Stream(ctx context.Context, input []M, opts ...model.Option) (*schema.StreamReader[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *typedFailoverProxyModel[M]) IsCallbacksEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (m *typedFailoverProxyModel[M]) GetType() string { _ = "STUB: not implemented"; return "" }

type failoverProxyModel = typedFailoverProxyModel[*schema.Message]

type FailoverContext[M MessageType] struct {
	FailoverAttempt uint

	InputMessages []M

	LastOutputMessage M

	LastErr error
}

type ModelFailoverConfig[M MessageType] struct {
	MaxRetries uint

	ShouldFailover func(ctx context.Context, outputMessage M, outputErr error) bool

	GetFailoverModel func(ctx context.Context, failoverCtx *FailoverContext[M]) (
		failoverModel model.BaseModel[M], failoverModelInputMessages []M, failoverErr error)
}

func typedGetFailoverLastSuccessModel[M MessageType](ctx context.Context) model.BaseModel[M] {
	_ = "STUB: not implemented"
	return nil
}

func typedSetFailoverLastSuccessModel[M MessageType](ctx context.Context, m model.BaseModel[M]) {
	_ = "STUB: not implemented"
	return
}

type failoverModelWrapper[M MessageType] struct {
	config *ModelFailoverConfig[M]
	inner  model.BaseModel[M]
}

func newFailoverModelWrapper[M MessageType](inner model.BaseModel[M], config *ModelFailoverConfig[M]) *failoverModelWrapper[M] {
	_ = "STUB: not implemented"
	return nil
}

func (f *failoverModelWrapper[M]) needFailover(ctx context.Context, outputMessage M, outputErr error) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *failoverModelWrapper[M]) getFailoverModel(ctx context.Context, failoverCtx *FailoverContext[M]) (model.BaseModel[M], []M, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (f *failoverModelWrapper[M]) Generate(ctx context.Context, input []M, opts ...model.Option) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func (f *failoverModelWrapper[M]) Stream(ctx context.Context, input []M, opts ...model.Option) (
	*schema.StreamReader[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func typedConsumeStream[M MessageType](stream *schema.StreamReader[M]) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}
