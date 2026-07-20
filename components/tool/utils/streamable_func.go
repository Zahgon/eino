package utils

import (
	"context"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type StreamFunc[T, D any] func(ctx context.Context, input T) (output *schema.StreamReader[D], err error)

type OptionableStreamFunc[T, D any] func(ctx context.Context, input T, opts ...tool.Option) (output *schema.StreamReader[D], err error)

func InferStreamTool[T, D any](toolName, toolDesc string, s StreamFunc[T, D], opts ...Option) (tool.StreamableTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.StreamableTool), nil
}

func InferOptionableStreamTool[T, D any](toolName, toolDesc string, s OptionableStreamFunc[T, D], opts ...Option) (tool.StreamableTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.StreamableTool), nil
}

func NewStreamTool[T, D any](desc *schema.ToolInfo, s StreamFunc[T, D], opts ...Option) tool.StreamableTool {
	_ = "STUB: not implemented"
	return *new(tool.StreamableTool)
}

func newOptionableStreamTool[T, D any](desc *schema.ToolInfo, s OptionableStreamFunc[T, D], opts ...Option) tool.StreamableTool {
	_ = "STUB: not implemented"
	return *new(tool.StreamableTool)
}

type streamableTool[T, D any] struct {
	info *schema.ToolInfo

	um UnmarshalArguments
	m  MarshalOutput

	Fn OptionableStreamFunc[T, D]
}

func (s *streamableTool[T, D]) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *streamableTool[T, D]) StreamableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (
	outStream *schema.StreamReader[string], err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *streamableTool[T, D]) GetType() string { _ = "STUB: not implemented"; return "" }

func (s *streamableTool[T, D]) getToolName() string { _ = "STUB: not implemented"; return "" }

type EnhancedStreamFunc[T any] func(ctx context.Context, input T) (output *schema.StreamReader[*schema.ToolResult], err error)

type OptionableEnhancedStreamFunc[T any] func(ctx context.Context, input T, opts ...tool.Option) (output *schema.StreamReader[*schema.ToolResult], err error)

func InferEnhancedStreamTool[T any](toolName, toolDesc string, s EnhancedStreamFunc[T], opts ...Option) (tool.EnhancedStreamableTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.EnhancedStreamableTool), nil
}

func InferOptionableEnhancedStreamTool[T any](toolName, toolDesc string, s OptionableEnhancedStreamFunc[T], opts ...Option) (tool.EnhancedStreamableTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.EnhancedStreamableTool), nil
}

func NewEnhancedStreamTool[T any](desc *schema.ToolInfo, s EnhancedStreamFunc[T], opts ...Option) tool.EnhancedStreamableTool {
	_ = "STUB: not implemented"
	return *new(tool.EnhancedStreamableTool)
}

func newOptionableEnhancedStreamTool[T any](desc *schema.ToolInfo, s OptionableEnhancedStreamFunc[T], opts ...Option) tool.EnhancedStreamableTool {
	_ = "STUB: not implemented"
	return *new(tool.EnhancedStreamableTool)
}

type enhancedStreamableTool[T any] struct {
	info *schema.ToolInfo

	um UnmarshalArguments

	Fn OptionableEnhancedStreamFunc[T]
}

func (s *enhancedStreamableTool[T]) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *enhancedStreamableTool[T]) StreamableRun(ctx context.Context, toolArgument *schema.ToolArgument, opts ...tool.Option) (
	outStream *schema.StreamReader[*schema.ToolResult], err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *enhancedStreamableTool[T]) GetType() string { _ = "STUB: not implemented"; return "" }

func (s *enhancedStreamableTool[T]) getToolName() string { _ = "STUB: not implemented"; return "" }
