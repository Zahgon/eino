package compose

import (
	"context"
	"reflect"
)

type newGraphOptions struct {
	withState func(ctx context.Context) any
	stateType reflect.Type
}

type NewGraphOption func(ngo *newGraphOptions)

func WithGenLocalState[S any](gls GenLocalState[S]) NewGraphOption {
	_ = "STUB: not implemented"
	return *new(NewGraphOption)
}

func NewGraph[I, O any](opts ...NewGraphOption) *Graph[I, O] { _ = "STUB: not implemented"; return nil }

type Graph[I, O any] struct {
	*graph
}

func (g *Graph[I, O]) AddEdge(startNode, endNode string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (g *Graph[I, O]) Compile(ctx context.Context, opts ...GraphCompileOption) (Runnable[I, O], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func compileAnyGraph[I, O any](ctx context.Context, g AnyGraph, opts ...GraphCompileOption) (Runnable[I, O], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
