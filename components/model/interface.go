package model

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

type messageType interface {
	*schema.Message | *schema.AgenticMessage
}

type BaseModel[M messageType] interface {
	Generate(ctx context.Context, input []M, opts ...Option) (M, error)
	Stream(ctx context.Context, input []M, opts ...Option) (*schema.StreamReader[M], error)
}

//go:generate  mockgen -destination ../../internal/mock/components/model/ChatModel_mock.go --package model github.com/cloudwego/eino/components/model BaseChatModel,ChatModel,ToolCallingChatModel
type BaseChatModel = BaseModel[*schema.Message]

type ChatModel interface {
	BaseChatModel

	BindTools(tools []*schema.ToolInfo) error
}

type ToolCallingChatModel interface {
	BaseChatModel

	WithTools(tools []*schema.ToolInfo) (ToolCallingChatModel, error)
}

type AgenticModel = BaseModel[*schema.AgenticMessage]
