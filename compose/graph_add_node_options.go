package compose

import (
	"reflect"
)

type graphAddNodeOpts struct {
	nodeOptions *nodeOptions
	processor   *processorOpts

	needState bool
}

type GraphAddNodeOpt func(o *graphAddNodeOpts)

type nodeOptions struct {
	nodeName string

	nodeKey string

	inputKey  string
	outputKey string

	graphCompileOption []GraphCompileOption
}

func WithNodeName(n string) GraphAddNodeOpt {
	_ = "STUB: not implemented"
	return *new(GraphAddNodeOpt)
}

func WithNodeKey(key string) GraphAddNodeOpt {
	_ = "STUB: not implemented"
	return *new(GraphAddNodeOpt)
}

func WithInputKey(k string) GraphAddNodeOpt {
	_ = "STUB: not implemented"
	return *new(GraphAddNodeOpt)
}

func WithOutputKey(k string) GraphAddNodeOpt {
	_ = "STUB: not implemented"
	return *new(GraphAddNodeOpt)
}

func WithGraphCompileOptions(opts ...GraphCompileOption) GraphAddNodeOpt {
	_ = "STUB: not implemented"
	return *new(GraphAddNodeOpt)
}

func WithStatePreHandler[I, S any](pre StatePreHandler[I, S]) GraphAddNodeOpt {
	_ = "STUB: not implemented"
	return *new(GraphAddNodeOpt)
}

func WithStatePostHandler[O, S any](post StatePostHandler[O, S]) GraphAddNodeOpt {
	_ = "STUB: not implemented"
	return *new(GraphAddNodeOpt)
}

func WithStreamStatePreHandler[I, S any](pre StreamStatePreHandler[I, S]) GraphAddNodeOpt {
	_ = "STUB: not implemented"
	return *new(GraphAddNodeOpt)
}

func WithStreamStatePostHandler[O, S any](post StreamStatePostHandler[O, S]) GraphAddNodeOpt {
	_ = "STUB: not implemented"
	return *new(GraphAddNodeOpt)
}

type processorOpts struct {
	statePreHandler  *composableRunnable
	preStateType     reflect.Type
	statePostHandler *composableRunnable
	postStateType    reflect.Type
}

func getGraphAddNodeOpts(opts ...GraphAddNodeOpt) *graphAddNodeOpts {
	_ = "STUB: not implemented"
	return nil
}
