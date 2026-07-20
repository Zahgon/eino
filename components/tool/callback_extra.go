package tool

import (
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/schema"
)

type CallbackInput struct {
	ArgumentsInJSON string

	Extra map[string]any
}

type CallbackOutput struct {
	Response string

	ToolOutput *schema.ToolResult

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
