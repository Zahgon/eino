package compose

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

func NewAgenticToolsNode(ctx context.Context, conf *ToolsNodeConfig) (*AgenticToolsNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type AgenticToolsNode struct {
	inner *ToolsNode
}

func (a *AgenticToolsNode) Invoke(ctx context.Context, input *schema.AgenticMessage, opts ...ToolsNodeOption) ([]*schema.AgenticMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AgenticToolsNode) Stream(ctx context.Context, input *schema.AgenticMessage,
	opts ...ToolsNodeOption) (*schema.StreamReader[[]*schema.AgenticMessage], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func agenticMessageToToolCallMessage(input *schema.AgenticMessage) *schema.Message {
	_ = "STUB: not implemented"
	return nil
}

func toolMessageToAgenticMessage(input []*schema.Message) []*schema.AgenticMessage {
	_ = "STUB: not implemented"
	return nil
}

func streamToolMessageToAgenticMessage(input *schema.StreamReader[[]*schema.Message]) *schema.StreamReader[[]*schema.AgenticMessage] {
	_ = "STUB: not implemented"
	return nil
}

func toolSearchResultMessageToAgenticMessage(m *schema.Message, meta *schema.StreamingMeta) (*schema.AgenticMessage, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func messageInputPartsToFunctionToolBlocks(parts []schema.MessageInputPart) []*schema.FunctionToolResultContentBlock {
	_ = "STUB: not implemented"
	return nil
}

type userInputVariant interface {
	schema.UserInputText | schema.UserInputImage | schema.UserInputAudio | schema.UserInputVideo | schema.UserInputFile
}

func newFuncToolResultContentBlock[T userInputVariant](content *T) *schema.FunctionToolResultContentBlock {
	_ = "STUB: not implemented"
	return nil
}

func derefString(s *string) string { _ = "STUB: not implemented"; return "" }

func (a *AgenticToolsNode) GetType() string { _ = "STUB: not implemented"; return "" }
