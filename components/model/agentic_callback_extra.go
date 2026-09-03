package model

import (
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/schema"
)

type AgenticConfig struct {
	Model string

	MaxTokens int

	Temperature float32

	TopP float32
}

type AgenticCallbackInput struct {
	Messages []*schema.AgenticMessage

	Tools []*schema.ToolInfo

	Config *AgenticConfig

	Extra map[string]any
}

type AgenticCallbackOutput struct {
	Message *schema.AgenticMessage

	Config *AgenticConfig

	TokenUsage *TokenUsage

	Extra map[string]any
}

func ConvAgenticCallbackInput(src callbacks.CallbackInput) *AgenticCallbackInput {
	_ = "STUB: not implemented"
	return nil
}

func ConvAgenticCallbackOutput(src callbacks.CallbackOutput) *AgenticCallbackOutput {
	_ = "STUB: not implemented"
	return nil
}
