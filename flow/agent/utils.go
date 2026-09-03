package agent

import (
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

func ChatModelWithTools(cm model.ChatModel, toolCallingModel model.ToolCallingChatModel, toolInfos []*schema.ToolInfo) (
	model.BaseChatModel, error) {
	_ = "STUB: not implemented"
	return *new(model.BaseChatModel), nil
}
