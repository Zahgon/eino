package prompt

import (
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/schema"
)

type AgenticCallbackInput struct {
	Variables map[string]any

	Templates []schema.AgenticMessagesTemplate

	Extra map[string]any
}

type AgenticCallbackOutput struct {
	Result []*schema.AgenticMessage

	Templates []schema.AgenticMessagesTemplate

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
