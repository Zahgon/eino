package adk

import (
	"context"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

type typedGenerateEndpoint[M MessageType] func(ctx context.Context, input []M, opts ...model.Option) (M, error)
type typedStreamEndpoint[M MessageType] func(ctx context.Context, input []M, opts ...model.Option) (*schema.StreamReader[M], error)

type typedModelWrapperConfig[M MessageType] struct {
	handlers       []TypedChatModelAgentMiddleware[M]
	middlewares    []AgentMiddleware
	retryConfig    *TypedModelRetryConfig[M]
	failoverConfig *ModelFailoverConfig[M]
	toolInfos      []*schema.ToolInfo
	cancelContext  *cancelContext
}

type modelWrapperConfig = typedModelWrapperConfig[*schema.Message]

func buildModelWrappers[M MessageType](m model.BaseModel[M], config *typedModelWrapperConfig[M]) model.BaseModel[M] {
	_ = "STUB: not implemented"
	return nil
}

func buildModelWrappersImpl[M MessageType](m model.BaseModel[M], config *typedModelWrapperConfig[M]) model.BaseModel[M] {
	_ = "STUB: not implemented"
	return nil
}

type typedCallbackInjectionModelWrapper[M MessageType] struct{}

func (w typedCallbackInjectionModelWrapper[M]) wrapModel(m model.BaseModel[M]) model.BaseModel[M] {
	_ = "STUB: not implemented"
	return nil
}

type typedCallbackInjectedModel[M MessageType] struct {
	inner model.BaseModel[M]
}

func (m *typedCallbackInjectedModel[M]) Generate(ctx context.Context, input []M, opts ...model.Option) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func (m *typedCallbackInjectedModel[M]) Stream(ctx context.Context, input []M, opts ...model.Option) (*schema.StreamReader[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handlersToToolMiddlewares[M MessageType](handlers []TypedChatModelAgentMiddleware[M]) []compose.ToolMiddleware {
	_ = "STUB: not implemented"
	return nil
}

type typedEventSenderModelWrapper[M MessageType] struct {
	*TypedBaseChatModelAgentMiddleware[M]
}

func NewEventSenderModelWrapper() ChatModelAgentMiddleware {
	_ = "STUB: not implemented"
	return *new(ChatModelAgentMiddleware)
}

func (w *typedEventSenderModelWrapper[M]) WrapModel(_ context.Context, m model.BaseModel[M], mc *TypedModelContext[M]) (model.BaseModel[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type typedEventSenderModel[M MessageType] struct {
	inner               model.BaseModel[M]
	modelRetryConfig    *TypedModelRetryConfig[M]
	modelFailoverConfig *ModelFailoverConfig[M]
}

func (m *typedEventSenderModel[M]) Generate(ctx context.Context, input []M, opts ...model.Option) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func (m *typedEventSenderModel[M]) Stream(ctx context.Context, input []M, opts ...model.Option) (*schema.StreamReader[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *typedEventSenderModel[M]) buildStreamConvertOptions(ctx context.Context) []schema.ConvertOption {
	_ = "STUB: not implemented"
	return nil
}

func copyMessage[M MessageType](msg M) M { _ = "STUB: not implemented"; return *new(M) }

func typedSetMessageID[M MessageType](msg M, id string) { _ = "STUB: not implemented"; return }

func GetMessageID[M MessageType](msg M) string { _ = "STUB: not implemented"; return "" }

func EnsureMessageID[M MessageType](msg M) { _ = "STUB: not implemented"; return }

func typedPopToolGenAction[M MessageType](ctx context.Context, toolName string) *AgentAction {
	_ = "STUB: not implemented"
	return nil
}

type typedEventSenderToolWrapper[M MessageType] struct {
	*TypedBaseChatModelAgentMiddleware[M]
}

func (*typedEventSenderToolWrapper[M]) isEventSenderToolWrapper() {
	_ = "STUB: not implemented"
	return
}

type eventSenderToolWrapperMarker interface{ isEventSenderToolWrapper() }

func NewEventSenderToolWrapper() ChatModelAgentMiddleware {
	_ = "STUB: not implemented"
	return *new(ChatModelAgentMiddleware)
}

func newTypedEventSenderToolWrapper[M MessageType]() *typedEventSenderToolWrapper[M] {
	_ = "STUB: not implemented"
	return nil
}

func textToFunctionToolResultBlocks(text string) []*schema.FunctionToolResultContentBlock {
	_ = "STUB: not implemented"
	return nil
}

func functionToolResultAgenticMessage(callID, name string, content []*schema.FunctionToolResultContentBlock) *schema.AgenticMessage {
	_ = "STUB: not implemented"
	return nil
}

func markAgenticMessageStreamingMeta(msg *schema.AgenticMessage, index int) {
	_ = "STUB: not implemented"
	return
}

func toolSearchResultAgenticMessage(callID, name string, tr *schema.ToolResult) (*schema.AgenticMessage, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func toolResultAgenticMessage(callID, name string, tr *schema.ToolResult) *schema.AgenticMessage {
	_ = "STUB: not implemented"
	return nil
}

func toolResultToBlocks(tr *schema.ToolResult) []*schema.FunctionToolResultContentBlock {
	_ = "STUB: not implemented"
	return nil
}

func derefString(s *string) string { _ = "STUB: not implemented"; return "" }

func typedToolInvokeEvent[M MessageType](callID, toolName, result, toolMsgID string) *TypedAgentEvent[M] {
	_ = "STUB: not implemented"
	return nil
}

func typedToolStreamEvent[M MessageType](callID, toolName, toolMsgID string, stream *schema.StreamReader[string]) *TypedAgentEvent[M] {
	_ = "STUB: not implemented"
	return nil
}

func typedToolEnhancedInvokeEvent[M MessageType](callID, toolName, toolMsgID string, result *schema.ToolResult) (*TypedAgentEvent[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func typedToolEnhancedStreamEvent[M MessageType](callID, toolName, toolMsgID string, stream *schema.StreamReader[*schema.ToolResult]) *TypedAgentEvent[M] {
	_ = "STUB: not implemented"
	return nil
}

func (w *typedEventSenderToolWrapper[M]) WrapInvokableToolCall(_ context.Context, endpoint InvokableToolCallEndpoint, tCtx *ToolContext) (InvokableToolCallEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(InvokableToolCallEndpoint), nil
}

func (w *typedEventSenderToolWrapper[M]) WrapStreamableToolCall(_ context.Context, endpoint StreamableToolCallEndpoint, tCtx *ToolContext) (StreamableToolCallEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(StreamableToolCallEndpoint), nil
}

func (w *typedEventSenderToolWrapper[M]) WrapEnhancedInvokableToolCall(_ context.Context, endpoint EnhancedInvokableToolCallEndpoint, tCtx *ToolContext) (EnhancedInvokableToolCallEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(EnhancedInvokableToolCallEndpoint), nil
}

func (w *typedEventSenderToolWrapper[M]) WrapEnhancedStreamableToolCall(_ context.Context, endpoint EnhancedStreamableToolCallEndpoint, tCtx *ToolContext) (EnhancedStreamableToolCallEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(EnhancedStreamableToolCallEndpoint), nil
}

func hasUserEventSenderToolWrapper[M MessageType](handlers []TypedChatModelAgentMiddleware[M]) bool {
	_ = "STUB: not implemented"
	return false
}

type typedStateModelWrapper[M MessageType] struct {
	inner               model.BaseModel[M]
	original            model.BaseModel[M]
	handlers            []TypedChatModelAgentMiddleware[M]
	middlewares         []AgentMiddleware
	toolInfos           []*schema.ToolInfo
	modelRetryConfig    *TypedModelRetryConfig[M]
	modelFailoverConfig *ModelFailoverConfig[M]
	cancelContext       *cancelContext
}

type stateModelWrapper = typedStateModelWrapper[*schema.Message]

func (w *typedStateModelWrapper[M]) IsCallbacksEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (w *typedStateModelWrapper[M]) GetType() string { _ = "STUB: not implemented"; return "" }

func (w *typedStateModelWrapper[M]) hasUserEventSender() bool {
	_ = "STUB: not implemented"
	return false
}

func (w *typedStateModelWrapper[M]) wrapGenerateEndpoint(endpoint typedGenerateEndpoint[M]) typedGenerateEndpoint[M] {
	_ = "STUB: not implemented"
	return nil
}

func (w *typedStateModelWrapper[M]) wrapStreamEndpoint(endpoint typedStreamEndpoint[M]) typedStreamEndpoint[M] {
	_ = "STUB: not implemented"
	return nil
}

func (w *typedStateModelWrapper[M]) Generate(ctx context.Context, _ []M, opts ...model.Option) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func (w *typedStateModelWrapper[M]) Stream(ctx context.Context, _ []M, opts ...model.Option) (*schema.StreamReader[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type typedEndpointModel[M MessageType] struct {
	generate typedGenerateEndpoint[M]
	stream   typedStreamEndpoint[M]
}

func (m *typedEndpointModel[M]) Generate(ctx context.Context, input []M, opts ...model.Option) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func (m *typedEndpointModel[M]) Stream(ctx context.Context, input []M, opts ...model.Option) (*schema.StreamReader[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
