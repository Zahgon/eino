package document

import (
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/schema"
)

type LoaderCallbackInput struct {
	Source Source

	Extra map[string]any
}

type LoaderCallbackOutput struct {
	Source Source

	Docs []*schema.Document

	Extra map[string]any
}

func ConvLoaderCallbackInput(src callbacks.CallbackInput) *LoaderCallbackInput {
	_ = "STUB: not implemented"
	return nil
}

func ConvLoaderCallbackOutput(src callbacks.CallbackOutput) *LoaderCallbackOutput {
	_ = "STUB: not implemented"
	return nil
}
