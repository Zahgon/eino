package compose

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

type Invoke[I, O, TOption any] func(ctx context.Context, input I, opts ...TOption) (output O, err error)

type Stream[I, O, TOption any] func(ctx context.Context,
	input I, opts ...TOption) (output *schema.StreamReader[O], err error)

type Collect[I, O, TOption any] func(ctx context.Context,
	input *schema.StreamReader[I], opts ...TOption) (output O, err error)

type Transform[I, O, TOption any] func(ctx context.Context,
	input *schema.StreamReader[I], opts ...TOption) (output *schema.StreamReader[O], err error)

type InvokeWOOpt[I, O any] func(ctx context.Context, input I) (output O, err error)

type StreamWOOpt[I, O any] func(ctx context.Context,
	input I) (output *schema.StreamReader[O], err error)

type CollectWOOpt[I, O any] func(ctx context.Context,
	input *schema.StreamReader[I]) (output O, err error)

type TransformWOOpts[I, O any] func(ctx context.Context,
	input *schema.StreamReader[I]) (output *schema.StreamReader[O], err error)

type Lambda struct {
	executor *composableRunnable
}

type lambdaOpts struct {
	enableComponentCallback bool

	componentImplType string
}

type LambdaOpt func(o *lambdaOpts)

func WithLambdaCallbackEnable(y bool) LambdaOpt { _ = "STUB: not implemented"; return *new(LambdaOpt) }

func WithLambdaType(t string) LambdaOpt { _ = "STUB: not implemented"; return *new(LambdaOpt) }

type unreachableOption struct{}

func InvokableLambdaWithOption[I, O, TOption any](i Invoke[I, O, TOption], opts ...LambdaOpt) *Lambda {
	_ = "STUB: not implemented"
	return nil
}

func InvokableLambda[I, O any](i InvokeWOOpt[I, O], opts ...LambdaOpt) *Lambda {
	_ = "STUB: not implemented"
	return nil
}

func StreamableLambdaWithOption[I, O, TOption any](s Stream[I, O, TOption], opts ...LambdaOpt) *Lambda {
	_ = "STUB: not implemented"
	return nil
}

func StreamableLambda[I, O any](s StreamWOOpt[I, O], opts ...LambdaOpt) *Lambda {
	_ = "STUB: not implemented"
	return nil
}

func CollectableLambdaWithOption[I, O, TOption any](c Collect[I, O, TOption], opts ...LambdaOpt) *Lambda {
	_ = "STUB: not implemented"
	return nil
}

func CollectableLambda[I, O any](c CollectWOOpt[I, O], opts ...LambdaOpt) *Lambda {
	_ = "STUB: not implemented"
	return nil
}

func TransformableLambdaWithOption[I, O, TOption any](t Transform[I, O, TOption], opts ...LambdaOpt) *Lambda {
	_ = "STUB: not implemented"
	return nil
}

func TransformableLambda[I, O any](t TransformWOOpts[I, O], opts ...LambdaOpt) *Lambda {
	_ = "STUB: not implemented"
	return nil
}

func AnyLambda[I, O, TOption any](i Invoke[I, O, TOption], s Stream[I, O, TOption],
	c Collect[I, O, TOption], t Transform[I, O, TOption], opts ...LambdaOpt) (*Lambda, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func anyLambda[I, O, TOption any](i Invoke[I, O, TOption], s Stream[I, O, TOption],
	c Collect[I, O, TOption], t Transform[I, O, TOption], opts ...LambdaOpt) *Lambda {
	_ = "STUB: not implemented"
	return nil
}

func getLambdaOpt(opts ...LambdaOpt) *lambdaOpts { _ = "STUB: not implemented"; return nil }

func ToList[I any](opts ...LambdaOpt) *Lambda { _ = "STUB: not implemented"; return nil }

func MessageParser[T any](p schema.MessageParser[T], opts ...LambdaOpt) *Lambda {
	_ = "STUB: not implemented"
	return nil
}
