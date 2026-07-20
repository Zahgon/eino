package model

import "github.com/cloudwego/eino/schema"

type Options struct {
	Temperature *float32

	Model *string

	TopP *float32

	Tools []*schema.ToolInfo

	DeferredTools []*schema.ToolInfo

	ToolSearchTool *schema.ToolInfo

	MaxTokens *int

	Stop []string

	ToolChoice *schema.ToolChoice

	AllowedToolNames []string

	AgenticToolChoice *schema.AgenticToolChoice
}

type Option struct {
	apply func(opts *Options)

	implSpecificOptFn any
}

func WithTemperature(temperature float32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxTokens(maxTokens int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithModel(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTopP(topP float32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithStop(stop []string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTools(tools []*schema.ToolInfo) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithToolSearchTool(tool *schema.ToolInfo) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithDeferredTools(tools []*schema.ToolInfo) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithToolChoice(toolChoice schema.ToolChoice, allowedToolNames ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithAgenticToolChoice(toolChoice *schema.AgenticToolChoice) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WrapImplSpecificOptFn[T any](optFn func(*T)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func GetCommonOptions(base *Options, opts ...Option) *Options {
	_ = "STUB: not implemented"
	return nil
}

func GetImplSpecificOptions[T any](base *T, opts ...Option) *T {
	_ = "STUB: not implemented"
	return nil
}
