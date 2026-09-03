package patchtoolcalls

import (
	"context"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

type PatchedToolResult struct {
	Content    string
	ToolResult *schema.ToolResult
}

type Config struct {
	PatchedContentGenerator func(ctx context.Context, toolName, toolCallID string) (string, error)

	PatchedToolResultGenerator func(
		ctx context.Context,
		toolName string,
		toolCallID string,
		toolArgument *schema.ToolArgument,
	) (*PatchedToolResult, error)
}

type patchedGenerators struct {
	content func(ctx context.Context, toolName, toolCallID string) (string, error)
	result  func(ctx context.Context, toolName, toolCallID string, toolArgument *schema.ToolArgument) (*PatchedToolResult, error)
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
	gens patchedGenerators
}

func (m *typedMiddleware[M]) BeforeModelRewriteState(ctx context.Context, state *adk.TypedChatModelAgentState[M],
	mc *adk.TypedModelContext[M],
) (context.Context, *adk.TypedChatModelAgentState[M], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func patchToolCallsForMessage[M adk.MessageType](ctx context.Context,
	gens patchedGenerators,
	state *adk.TypedChatModelAgentState[*schema.Message],
	_ *adk.TypedModelContext[M],
) (context.Context, *adk.TypedChatModelAgentState[M], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func patchToolCallsForAgenticMessage[M adk.MessageType](ctx context.Context,
	gens patchedGenerators,
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

func createPatchedToolMessage(ctx context.Context, gens patchedGenerators, tc schema.ToolCall) (*schema.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createPatchedAgenticToolMessage(ctx context.Context, gens patchedGenerators, toolName, callID, arguments string) (*schema.AgenticMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func patchedToolResultToMessage(toolName, callID string, result *PatchedToolResult) (*schema.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func patchedToolResultToAgenticMessage(toolName, callID string, result *PatchedToolResult) (*schema.AgenticMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validatePatchedToolResult(result *PatchedToolResult) error {
	_ = "STUB: not implemented"
	return nil
}

func agenticTextToolResultMessage(toolName, callID, content string) *schema.AgenticMessage {
	_ = "STUB: not implemented"
	return nil
}

func toolResultToMessage(toolName, callID string, result *schema.ToolResult) (*schema.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toolResultToAgenticMessage(toolName, callID string, result *schema.ToolResult) (*schema.AgenticMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func singleTextToolResult(result *schema.ToolResult) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func toolResultToFunctionBlocks(result *schema.ToolResult) ([]*schema.FunctionToolResultContentBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func derefString(s *string) string { _ = "STUB: not implemented"; return "" }

const (
	defaultPatchedToolMessageTemplate        = "Tool call %s with id %s was canceled - another message came in before it could be completed."
	defaultPatchedToolMessageTemplateChinese = "工具调用 %s（ID 为 %s）已被取消——在其完成之前收到了另一条消息。"
)
