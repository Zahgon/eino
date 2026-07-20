package model

import (
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/schema"
)

type TokenUsage struct {
	PromptTokens int

	PromptTokenDetails PromptTokenDetails

	CompletionTokens int

	TotalTokens int

	CompletionTokensDetails CompletionTokensDetails `json:"completion_token_details"`
}

type CompletionTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens,omitempty"`
}

type PromptTokenDetails struct {
	CachedTokens int
}

type Config struct {
	Model string

	MaxTokens int

	Temperature float32

	TopP float32

	Stop []string
}

type CallbackInput struct {
	Messages []*schema.Message

	Tools []*schema.ToolInfo

	ToolChoice *schema.ToolChoice

	Config *Config

	Extra map[string]any
}

type CallbackOutput struct {
	Message *schema.Message

	Config *Config

	TokenUsage *TokenUsage

	Extra map[string]any
}

func ConvCallbackInput(src callbacks.CallbackInput) *CallbackInput {
	_ = "STUB: not implemented"
	return nil
}

func ConvCallbackOutput(src callbacks.CallbackOutput) *CallbackOutput {
	_ = "STUB: not implemented"
	return nil
}
