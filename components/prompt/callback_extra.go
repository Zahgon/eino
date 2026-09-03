package prompt

import (
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/schema"
)

type CallbackInput struct {
	Variables map[string]any

	Templates []schema.MessagesTemplate

	Extra map[string]any
}

type CallbackOutput struct {
	Result []*schema.Message

	Templates []schema.MessagesTemplate

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
