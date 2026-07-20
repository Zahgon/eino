package retriever

import (
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/schema"
)

type CallbackInput struct {
	Query string

	TopK int

	Filter string

	ScoreThreshold *float64

	Extra map[string]any
}

type CallbackOutput struct {
	Docs []*schema.Document

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
