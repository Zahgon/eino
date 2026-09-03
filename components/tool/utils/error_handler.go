package utils

import (
	"context"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type ErrorHandler func(context.Context, error) string

func WrapToolWithErrorHandler(t tool.BaseTool, h ErrorHandler) tool.BaseTool {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool)
}

func WrapInvokableToolWithErrorHandler(t tool.InvokableTool, h ErrorHandler) tool.InvokableTool {
	_ = "STUB: not implemented"
	return *new(tool.InvokableTool)
}

func WrapStreamableToolWithErrorHandler(t tool.StreamableTool, h ErrorHandler) tool.StreamableTool {
	_ = "STUB: not implemented"
	return *new(tool.StreamableTool)
}

type errorWrapper struct {
	*infoHelper
	*errorHelper
}

type streamErrorWrapper struct {
	*infoHelper
	*streamErrorHelper
}

type combinedErrorWrapper struct {
	*infoHelper
	*errorHelper
	*streamErrorHelper
}

type infoHelper struct {
	info func(ctx context.Context) (*schema.ToolInfo, error)
}

func (i *infoHelper) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type errorHelper struct {
	i func(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error)
	h ErrorHandler
}

func (s *errorHelper) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type streamErrorHelper struct {
	s func(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (*schema.StreamReader[string], error)
	h ErrorHandler
}

func (s *streamErrorHelper) StreamableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (*schema.StreamReader[string], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
