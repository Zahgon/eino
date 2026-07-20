package document

import (
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/schema"
)

type TransformerCallbackInput struct {
	Input []*schema.Document

	Extra map[string]any
}

type TransformerCallbackOutput struct {
	Output []*schema.Document

	Extra map[string]any
}

func ConvTransformerCallbackInput(src callbacks.CallbackInput) *TransformerCallbackInput {
	_ = "STUB: not implemented"
	return nil
}

func ConvTransformerCallbackOutput(src callbacks.CallbackOutput) *TransformerCallbackOutput {
	_ = "STUB: not implemented"
	return nil
}
