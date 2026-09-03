package compose

import (
	"context"
	"reflect"

	"github.com/cloudwego/eino/schema"
)

type Runnable[I, O any] interface {
	Invoke(ctx context.Context, input I, opts ...Option) (output O, err error)
	Stream(ctx context.Context, input I, opts ...Option) (output *schema.StreamReader[O], err error)
	Collect(ctx context.Context, input *schema.StreamReader[I], opts ...Option) (output O, err error)
	Transform(ctx context.Context, input *schema.StreamReader[I], opts ...Option) (output *schema.StreamReader[O], err error)
}

type invoke func(ctx context.Context, input any, opts ...any) (output any, err error)
type transform func(ctx context.Context, input streamReader, opts ...any) (output streamReader, err error)

type composableRunnable struct {
	i invoke
	t transform

	inputType  reflect.Type
	outputType reflect.Type
	optionType reflect.Type

	*genericHelper

	isPassthrough bool

	meta *executorMeta

	nodeInfo *nodeInfo
}

func runnableLambda[I, O, TOption any](i Invoke[I, O, TOption], s Stream[I, O, TOption], c Collect[I, O, TOption],
	t Transform[I, O, TOption], enableCallback bool) *composableRunnable {
	_ = "STUB: not implemented"
	return nil
}

func isNilAssignableType(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

type runnablePacker[I, O, TOption any] struct {
	i Invoke[I, O, TOption]
	s Stream[I, O, TOption]
	c Collect[I, O, TOption]
	t Transform[I, O, TOption]
}

func (rp *runnablePacker[I, O, TOption]) wrapRunnableCtx(ctxWrapper func(ctx context.Context, opts ...TOption) context.Context) {
	_ = "STUB: not implemented"
	return
}

func (rp *runnablePacker[I, O, TOption]) toComposableRunnable() *composableRunnable {
	_ = "STUB: not implemented"
	return nil
}

func (rp *runnablePacker[I, O, TOption]) Invoke(ctx context.Context,
	input I, opts ...TOption) (output O, err error) {
	_ = "STUB: not implemented"
	return *new(O), nil
}

func (rp *runnablePacker[I, O, TOption]) Stream(ctx context.Context,
	input I, opts ...TOption) (output *schema.StreamReader[O], err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rp *runnablePacker[I, O, TOption]) Collect(ctx context.Context,
	input *schema.StreamReader[I], opts ...TOption) (output O, err error) {
	_ = "STUB: not implemented"
	return *new(O), nil
}

func (rp *runnablePacker[I, O, TOption]) Transform(ctx context.Context,
	input *schema.StreamReader[I], opts ...TOption) (output *schema.StreamReader[O], err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func defaultImplConcatStreamReader[T any](
	sr *schema.StreamReader[T]) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func invokeByStream[I, O, TOption any](s Stream[I, O, TOption]) Invoke[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func invokeByCollect[I, O, TOption any](c Collect[I, O, TOption]) Invoke[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func invokeByTransform[I, O, TOption any](t Transform[I, O, TOption]) Invoke[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func streamByTransform[I, O, TOption any](t Transform[I, O, TOption]) Stream[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func streamByInvoke[I, O, TOption any](i Invoke[I, O, TOption]) Stream[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func streamByCollect[I, O, TOption any](c Collect[I, O, TOption]) Stream[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func collectByTransform[I, O, TOption any](t Transform[I, O, TOption]) Collect[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func collectByInvoke[I, O, TOption any](i Invoke[I, O, TOption]) Collect[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func collectByStream[I, O, TOption any](s Stream[I, O, TOption]) Collect[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func transformByStream[I, O, TOption any](s Stream[I, O, TOption]) Transform[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func transformByCollect[I, O, TOption any](c Collect[I, O, TOption]) Transform[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func transformByInvoke[I, O, TOption any](i Invoke[I, O, TOption]) Transform[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func newRunnablePacker[I, O, TOption any](i Invoke[I, O, TOption], s Stream[I, O, TOption],
	c Collect[I, O, TOption], t Transform[I, O, TOption], enableCallback bool) *runnablePacker[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func toGenericRunnable[I, O any](cr *composableRunnable, ctxWrapper func(ctx context.Context, opts ...Option) context.Context) (
	*runnablePacker[I, O, Option], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inputKeyedComposableRunnable(key string, r *composableRunnable) *composableRunnable {
	_ = "STUB: not implemented"
	return nil
}

func outputKeyedComposableRunnable(key string, r *composableRunnable) *composableRunnable {
	_ = "STUB: not implemented"
	return nil
}

func composablePassthrough() *composableRunnable { _ = "STUB: not implemented"; return nil }
