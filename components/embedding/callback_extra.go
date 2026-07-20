package embedding

import (
	"github.com/cloudwego/eino/callbacks"
)

type TokenUsage struct {
	PromptTokens int

	CompletionTokens int

	TotalTokens int
}

type Config struct {
	Model string

	EncodingFormat string
}

type ComponentExtra struct {
	Config *Config

	TokenUsage *TokenUsage
}

type CallbackInput struct {
	Texts []string

	Config *Config

	Extra map[string]any
}

type CallbackOutput struct {
	Embeddings [][]float64

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
