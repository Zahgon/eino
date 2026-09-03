package compose

import (
	"context"
	"reflect"
)

type executorMeta struct {
	component component

	isComponentCallbackEnabled bool

	componentImplType string
}

type nodeInfo struct {
	name string

	inputKey  string
	outputKey string

	preProcessor, postProcessor *composableRunnable

	compileOption *graphCompileOptions
}

type graphNode struct {
	cr *composableRunnable

	g AnyGraph

	nodeInfo     *nodeInfo
	executorMeta *executorMeta

	instance any
	opts     []GraphAddNodeOpt
}

func (gn *graphNode) getGenericHelper() *genericHelper { _ = "STUB: not implemented"; return nil }

func (gn *graphNode) inputType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (gn *graphNode) outputType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (gn *graphNode) compileIfNeeded(ctx context.Context) (*composableRunnable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseExecutorInfoFromComponent(c component, executor any) *executorMeta {
	_ = "STUB: not implemented"
	return nil
}

func getNodeInfo(opts ...GraphAddNodeOpt) (*nodeInfo, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}
