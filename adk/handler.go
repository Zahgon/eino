package adk

import (
	"context"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type InvokableToolCallEndpoint func(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error)

type StreamableToolCallEndpoint func(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (*schema.StreamReader[string], error)

type EnhancedInvokableToolCallEndpoint func(ctx context.Context, toolArgument *schema.ToolArgument, opts ...tool.Option) (*schema.ToolResult, error)

type EnhancedStreamableToolCallEndpoint func(ctx context.Context, toolArgument *schema.ToolArgument, opts ...tool.Option) (*schema.StreamReader[*schema.ToolResult], error)

type ToolContext struct {
	Name   string
	CallID string
}

type ToolCallsContext struct {
	ToolCalls []ToolContext
}

type TypedModelContext[M MessageType] struct {
	Tools []*schema.ToolInfo

	ModelRetryConfig *TypedModelRetryConfig[M]

	ModelFailoverConfig *ModelFailoverConfig[M]

	cancelContext *cancelContext
}

type ModelContext = TypedModelContext[*schema.Message]

type ChatModelAgentContext struct {
	Instruction string

	Tools []tool.BaseTool

	ReturnDirectly map[string]bool

	ToolSearchTool *schema.ToolInfo
}

type TypedChatModelAgentMiddleware[M MessageType] interface {
	BeforeAgent(ctx context.Context, runCtx *ChatModelAgentContext) (context.Context, *ChatModelAgentContext, error)

	AfterAgent(ctx context.Context, state *TypedChatModelAgentState[M]) (context.Context, error)

	BeforeModelRewriteState(ctx context.Context, state *TypedChatModelAgentState[M], mc *TypedModelContext[M]) (context.Context, *TypedChatModelAgentState[M], error)

	AfterModelRewriteState(ctx context.Context, state *TypedChatModelAgentState[M], mc *TypedModelContext[M]) (context.Context, *TypedChatModelAgentState[M], error)

	WrapInvokableToolCall(ctx context.Context, endpoint InvokableToolCallEndpoint, tCtx *ToolContext) (InvokableToolCallEndpoint, error)

	WrapStreamableToolCall(ctx context.Context, endpoint StreamableToolCallEndpoint, tCtx *ToolContext) (StreamableToolCallEndpoint, error)

	WrapEnhancedInvokableToolCall(ctx context.Context, endpoint EnhancedInvokableToolCallEndpoint, tCtx *ToolContext) (EnhancedInvokableToolCallEndpoint, error)

	WrapEnhancedStreamableToolCall(ctx context.Context, endpoint EnhancedStreamableToolCallEndpoint, tCtx *ToolContext) (EnhancedStreamableToolCallEndpoint, error)

	WrapModel(ctx context.Context, m model.BaseModel[M], mc *TypedModelContext[M]) (model.BaseModel[M], error)
}

type ChatModelAgentMiddleware = TypedChatModelAgentMiddleware[*schema.Message]

type TypedBaseChatModelAgentMiddleware[M MessageType] struct{}

type BaseChatModelAgentMiddleware = TypedBaseChatModelAgentMiddleware[*schema.Message]

func (b *TypedBaseChatModelAgentMiddleware[M]) WrapInvokableToolCall(_ context.Context, endpoint InvokableToolCallEndpoint, _ *ToolContext) (InvokableToolCallEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(InvokableToolCallEndpoint), nil
}

func (b *TypedBaseChatModelAgentMiddleware[M]) WrapStreamableToolCall(_ context.Context, endpoint StreamableToolCallEndpoint, _ *ToolContext) (StreamableToolCallEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(StreamableToolCallEndpoint), nil
}

func (b *TypedBaseChatModelAgentMiddleware[M]) WrapEnhancedInvokableToolCall(_ context.Context, endpoint EnhancedInvokableToolCallEndpoint, _ *ToolContext) (EnhancedInvokableToolCallEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(EnhancedInvokableToolCallEndpoint), nil
}

func (b *TypedBaseChatModelAgentMiddleware[M]) WrapEnhancedStreamableToolCall(_ context.Context, endpoint EnhancedStreamableToolCallEndpoint, _ *ToolContext) (EnhancedStreamableToolCallEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(EnhancedStreamableToolCallEndpoint), nil
}

func (b *TypedBaseChatModelAgentMiddleware[M]) WrapModel(_ context.Context, m model.BaseModel[M], _ *TypedModelContext[M]) (model.BaseModel[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *TypedBaseChatModelAgentMiddleware[M]) BeforeAgent(ctx context.Context, runCtx *ChatModelAgentContext) (context.Context, *ChatModelAgentContext, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func (b *TypedBaseChatModelAgentMiddleware[M]) AfterAgent(ctx context.Context, state *TypedChatModelAgentState[M]) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (b *TypedBaseChatModelAgentMiddleware[M]) BeforeModelRewriteState(ctx context.Context, state *TypedChatModelAgentState[M], mc *TypedModelContext[M]) (context.Context, *TypedChatModelAgentState[M], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func (b *TypedBaseChatModelAgentMiddleware[M]) AfterModelRewriteState(ctx context.Context, state *TypedChatModelAgentState[M], mc *TypedModelContext[M]) (context.Context, *TypedChatModelAgentState[M], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func processTypedState(ctx context.Context, fn func(extra map[string]any) map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func SetRunLocalValue(ctx context.Context, key string, value any) error {
	_ = "STUB: not implemented"
	return nil
}

func GetRunLocalValue(ctx context.Context, key string) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

func DeleteRunLocalValue(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func TypedSendEvent[M MessageType](ctx context.Context, event *TypedAgentEvent[M]) error {
	_ = "STUB: not implemented"
	return nil
}

func SendEvent(ctx context.Context, event *AgentEvent) error { _ = "STUB: not implemented"; return nil }

func checkGobEncodability(key string, value any) error { _ = "STUB: not implemented"; return nil }
