package patchtoolcalls

import (
	"context"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

type Config struct {
	PatchedContentGenerator func(ctx context.Context, toolName, toolCallID string) (string, error)
}

func NewTyped[M adk.MessageType](_ context.Context, cfg *Config) (adk.TypedChatModelAgentMiddleware[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(ctx context.Context, cfg *Config) (adk.ChatModelAgentMiddleware, error) {
	_ = "STUB: not implemented"
	return *new(adk.ChatModelAgentMiddleware), nil
}

type typedMiddleware[M adk.MessageType] struct {
	*adk.TypedBaseChatModelAgentMiddleware[M]
	gen func(ctx context.Context, toolName, toolCallID string) (string, error)
}

func (m *typedMiddleware[M]) BeforeModelRewriteState(ctx context.Context, state *adk.TypedChatModelAgentState[M],
	mc *adk.TypedModelContext[M],
) (context.Context, *adk.TypedChatModelAgentState[M], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func patchToolCallsForMessage[M adk.MessageType](ctx context.Context,
	gen func(ctx context.Context, toolName, toolCallID string) (string, error),
	state *adk.TypedChatModelAgentState[*schema.Message],
	_ *adk.TypedModelContext[M],
) (context.Context, *adk.TypedChatModelAgentState[M], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func patchToolCallsForAgenticMessage[M adk.MessageType](ctx context.Context,
	gen func(ctx context.Context, toolName, toolCallID string) (string, error),
	state *adk.TypedChatModelAgentState[*schema.AgenticMessage],
	_ *adk.TypedModelContext[M],
) (context.Context, *adk.TypedChatModelAgentState[M], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func hasCorrespondingToolMessage(messages []*schema.Message, toolCallID string) bool {
	_ = "STUB: not implemented"
	return false
}

func hasCorrespondingAgenticToolResult(messages []*schema.AgenticMessage, toolCallID string) bool {
	_ = "STUB: not implemented"
	return false
}

func createPatchedToolMessage(ctx context.Context, gen func(ctx context.Context, toolName, toolCallID string) (string, error), tc schema.ToolCall) (*schema.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createPatchedAgenticToolMessage(ctx context.Context, gen func(ctx context.Context, toolName, toolCallID string) (string, error), toolName, callID string) (*schema.AgenticMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	defaultPatchedToolMessageTemplate        = "Tool call %s with id %s was canceled - another message came in before it could be completed."
	defaultPatchedToolMessageTemplateChinese = "工具调用 %s（ID 为 %s）已被取消——在其完成之前收到了另一条消息。"
)
