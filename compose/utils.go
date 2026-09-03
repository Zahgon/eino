package compose

import (
	"context"
	"reflect"

	"github.com/cloudwego/eino/callbacks"
	icb "github.com/cloudwego/eino/internal/callbacks"
	"github.com/cloudwego/eino/schema"
)

type on[T any] func(context.Context, T) (context.Context, T)

func onStart[T any](ctx context.Context, input T) (context.Context, T) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(T)
}

func onEnd[T any](ctx context.Context, output T) (context.Context, T) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(T)
}

func onStartWithStreamInput[T any](ctx context.Context, input *schema.StreamReader[T]) (
	context.Context, *schema.StreamReader[T]) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func genericOnStartWithStreamInputHandle(ctx context.Context, input streamReader,
	runInfo *icb.RunInfo, handlers []icb.Handler) (context.Context, streamReader) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(streamReader)
}

func genericOnStartWithStreamInput(ctx context.Context, input streamReader) (context.Context, streamReader) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(streamReader)
}

func onEndWithStreamOutput[T any](ctx context.Context, output *schema.StreamReader[T]) (
	context.Context, *schema.StreamReader[T]) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func genericOnEndWithStreamOutputHandle(ctx context.Context, output streamReader,
	runInfo *icb.RunInfo, handlers []icb.Handler) (context.Context, streamReader) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(streamReader)
}

func genericOnEndWithStreamOutput(ctx context.Context, output streamReader) (context.Context, streamReader) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(streamReader)
}

func onError(ctx context.Context, err error) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func runWithCallbacks[I, O, TOption any](r func(context.Context, I, ...TOption) (O, error),
	onStart on[I], onEnd on[O], onError on[error]) func(context.Context, I, ...TOption) (O, error) {
	_ = "STUB: not implemented"
	return nil
}

func invokeWithCallbacks[I, O, TOption any](i Invoke[I, O, TOption]) Invoke[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func onGraphStart(ctx context.Context, input any, isStream bool) (context.Context, any) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(any)
}

func onGraphEnd(ctx context.Context, output any, isStream bool) (context.Context, any) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(any)
}

func onGraphError(ctx context.Context, err error) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func streamWithCallbacks[I, O, TOption any](s Stream[I, O, TOption]) Stream[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func collectWithCallbacks[I, O, TOption any](c Collect[I, O, TOption]) Collect[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func transformWithCallbacks[I, O, TOption any](t Transform[I, O, TOption]) Transform[I, O, TOption] {
	_ = "STUB: not implemented"
	return nil
}

func initGraphCallbacks(ctx context.Context, info *nodeInfo, meta *executorMeta, opts ...Option) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func initNodeCallbacks(ctx context.Context, key string, info *nodeInfo, meta *executorMeta, opts ...Option) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func streamChunkConvertForCBOutput[O any](o O) (callbacks.CallbackOutput, error) {
	_ = "STUB: not implemented"
	return *new(callbacks.CallbackOutput), nil
}

func streamChunkConvertForCBInput[I any](i I) (callbacks.CallbackInput, error) {
	_ = "STUB: not implemented"
	return *new(callbacks.CallbackInput), nil
}

func toAnyList[T any](in []T) []any { _ = "STUB: not implemented"; return nil }

type assignableType uint8

const (
	assignableTypeMustNot assignableType = iota
	assignableTypeMust
	assignableTypeMay
)

func checkAssignable(input, arg reflect.Type) assignableType {
	_ = "STUB: not implemented"
	return *new(assignableType)
}

func extractOption(nodes map[string]*chanCall, opts ...Option) (map[string][]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mapToList(m map[string]any) []any { _ = "STUB: not implemented"; return nil }
