package compose

import (
	"context"
	"sync"

	"github.com/cloudwego/eino/schema"
)

type GenLocalState[S any] func(ctx context.Context) (state S)

type stateKey struct{}

type internalState struct {
	state  any
	mu     sync.Mutex
	parent *internalState
}

type StatePreHandler[I, S any] func(ctx context.Context, in I, state S) (I, error)

type StatePostHandler[O, S any] func(ctx context.Context, out O, state S) (O, error)

type StreamStatePreHandler[I, S any] func(ctx context.Context, in *schema.StreamReader[I], state S) (*schema.StreamReader[I], error)

type StreamStatePostHandler[O, S any] func(ctx context.Context, out *schema.StreamReader[O], state S) (*schema.StreamReader[O], error)

func convertPreHandler[I, S any](handler StatePreHandler[I, S]) *composableRunnable {
	_ = "STUB: not implemented"
	return nil
}

func convertPostHandler[O, S any](handler StatePostHandler[O, S]) *composableRunnable {
	_ = "STUB: not implemented"
	return nil
}

func streamConvertPreHandler[I, S any](handler StreamStatePreHandler[I, S]) *composableRunnable {
	_ = "STUB: not implemented"
	return nil
}

func streamConvertPostHandler[O, S any](handler StreamStatePostHandler[O, S]) *composableRunnable {
	_ = "STUB: not implemented"
	return nil
}

func ProcessState[S any](ctx context.Context, handler func(context.Context, S) error) error {
	_ = "STUB: not implemented"
	return nil
}

func getState[S any](ctx context.Context) (S, *sync.Mutex, error) {
	_ = "STUB: not implemented"
	return *new(S), nil, nil
}
