package utils

import (
	"context"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type InvokeFunc[T, D any] func(ctx context.Context, input T) (output D, err error)

type OptionableInvokeFunc[T, D any] func(ctx context.Context, input T, opts ...tool.Option) (output D, err error)

func InferTool[T, D any](toolName, toolDesc string, i InvokeFunc[T, D], opts ...Option) (tool.InvokableTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.InvokableTool), nil
}

func InferOptionableTool[T, D any](toolName, toolDesc string, i OptionableInvokeFunc[T, D], opts ...Option) (tool.InvokableTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.InvokableTool), nil
}

type EnhancedInvokeFunc[T any] func(ctx context.Context, input T) (output *schema.ToolResult, err error)

type OptionableEnhancedInvokeFunc[T any] func(ctx context.Context, input T, opts ...tool.Option) (output *schema.ToolResult, err error)

func InferEnhancedTool[T any](toolName, toolDesc string, i EnhancedInvokeFunc[T], opts ...Option) (tool.EnhancedInvokableTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.EnhancedInvokableTool), nil
}

func InferOptionableEnhancedTool[T any](toolName, toolDesc string, i OptionableEnhancedInvokeFunc[T], opts ...Option) (tool.EnhancedInvokableTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.EnhancedInvokableTool), nil
}

func GoStruct2ParamsOneOf[T any](opts ...Option) (*schema.ParamsOneOf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GoStruct2ToolInfo[T any](toolName, toolDesc string, opts ...Option) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func goStruct2ToolInfo[T any](toolName, toolDesc string, opts ...Option) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func goStruct2ParamsOneOf[T any](opts ...Option) (*schema.ParamsOneOf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTool[T, D any](desc *schema.ToolInfo, i InvokeFunc[T, D], opts ...Option) tool.InvokableTool {
	_ = "STUB: not implemented"
	return *new(tool.InvokableTool)
}

func newOptionableTool[T, D any](desc *schema.ToolInfo, i OptionableInvokeFunc[T, D], opts ...Option) tool.InvokableTool {
	_ = "STUB: not implemented"
	return *new(tool.InvokableTool)
}

type invokableTool[T, D any] struct {
	info *schema.ToolInfo

	um UnmarshalArguments
	m  MarshalOutput

	Fn OptionableInvokeFunc[T, D]
}

func (i *invokableTool[T, D]) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *invokableTool[T, D]) InvokableRun(ctx context.Context, arguments string, opts ...tool.Option) (output string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (i *invokableTool[T, D]) GetType() string { _ = "STUB: not implemented"; return "" }

func (i *invokableTool[T, D]) getToolName() string { _ = "STUB: not implemented"; return "" }

func snakeToCamel(s string) string { _ = "STUB: not implemented"; return "" }

func NewEnhancedTool[T any](desc *schema.ToolInfo, i EnhancedInvokeFunc[T], opts ...Option) tool.EnhancedInvokableTool {
	_ = "STUB: not implemented"
	return *new(tool.EnhancedInvokableTool)
}

func newOptionableEnhancedTool[T any](desc *schema.ToolInfo, i OptionableEnhancedInvokeFunc[T], opts ...Option) tool.EnhancedInvokableTool {
	_ = "STUB: not implemented"
	return *new(tool.EnhancedInvokableTool)
}

type enhancedInvokableTool[T any] struct {
	info *schema.ToolInfo

	um UnmarshalArguments

	Fn OptionableEnhancedInvokeFunc[T]
}

func (e *enhancedInvokableTool[T]) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *enhancedInvokableTool[T]) InvokableRun(ctx context.Context, toolArgument *schema.ToolArgument, opts ...tool.Option) (*schema.ToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *enhancedInvokableTool[T]) GetType() string { _ = "STUB: not implemented"; return "" }

func (e *enhancedInvokableTool[T]) getToolName() string { _ = "STUB: not implemented"; return "" }
