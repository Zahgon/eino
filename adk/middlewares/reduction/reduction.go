package reduction

import (
	"context"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

type TypedConfig[M adk.MessageType] struct {
	Backend Backend

	SkipTruncation bool

	SkipClear bool

	ReadFileToolName string

	RootDir string

	GenTruncOffloadFilePath func(ctx context.Context, toolDetail *ToolDetail) (filePath string, err error)

	GenClearOffloadFilePath func(ctx context.Context, toolDetail *ToolDetail) (filePath string, err error)

	MaxLengthForTrunc int

	TruncExcludeTools []string

	TokenCounter func(ctx context.Context, msg []M, tools []*schema.ToolInfo) (int64, error)

	MaxTokensForClear int64

	ClearRetentionSuffixLimit int

	ClearAtLeastTokens int64

	ClearExcludeTools []string

	ClearMessageRewriter func(ctx context.Context, toolCallMsg M, toolResponseMsgs []M) (messagesAfterRewrite []M, err error)

	ClearPostProcess func(ctx context.Context, state *adk.TypedChatModelAgentState[M]) context.Context

	ToolConfig map[string]*ToolReductionConfig
}

type Config = TypedConfig[*schema.Message]

type ToolReductionConfig struct {
	Backend Backend

	SkipTruncation bool

	TruncHandler func(ctx context.Context, detail *ToolDetail) (*TruncResult, error)

	SkipClear bool

	ClearHandler func(ctx context.Context, detail *ToolDetail) (*ClearResult, error)
}

type ToolDetail struct {
	ToolContext *adk.ToolContext

	ToolArgument *schema.ToolArgument

	ToolResult *schema.ToolResult

	StreamToolResult *schema.StreamReader[*schema.ToolResult]
}

type TruncResult struct {
	NeedTrunc bool

	ToolResult *schema.ToolResult

	StreamToolResult *schema.StreamReader[*schema.ToolResult]

	NeedOffload bool

	OffloadFilePath string

	OffloadContent string
}

type ClearResult struct {
	NeedClear bool

	ToolArgument *schema.ToolArgument

	ToolResult *schema.ToolResult

	NeedOffload bool

	OffloadFilePath string

	OffloadContent string
}

func (t *TypedConfig[M]) copyAndFillDefaults() (*TypedConfig[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTyped[M adk.MessageType](_ context.Context, config *TypedConfig[M]) (adk.TypedChatModelAgentMiddleware[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(ctx context.Context, config *Config) (adk.ChatModelAgentMiddleware, error) {
	_ = "STUB: not implemented"
	return *new(adk.ChatModelAgentMiddleware), nil
}

type typedToolReductionMiddleware[M adk.MessageType] struct {
	*adk.TypedBaseChatModelAgentMiddleware[M]

	config        *TypedConfig[M]
	defaultConfig *ToolReductionConfig

	excludeTruncTools map[string]struct{}
	excludeClearTools map[string]struct{}
}

func getDefaultTokenCounter[M adk.MessageType]() func(ctx context.Context, msgs []M, tools []*schema.ToolInfo) (int64, error) {
	_ = "STUB: not implemented"
	return nil
}

func defaultAgenticTokenCounter(_ context.Context, msgs []*schema.AgenticMessage, tools []*schema.ToolInfo) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (t *typedToolReductionMiddleware[M]) getToolConfig(toolName string, sc scene) *ToolReductionConfig {
	_ = "STUB: not implemented"
	return nil
}

func (t *typedToolReductionMiddleware[M]) WrapInvokableToolCall(_ context.Context, endpoint adk.InvokableToolCallEndpoint, tCtx *adk.ToolContext) (adk.InvokableToolCallEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(adk.InvokableToolCallEndpoint), nil
}

func (t *typedToolReductionMiddleware[M]) WrapStreamableToolCall(_ context.Context, endpoint adk.StreamableToolCallEndpoint, tCtx *adk.ToolContext) (adk.StreamableToolCallEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(adk.StreamableToolCallEndpoint), nil
}

func (t *typedToolReductionMiddleware[M]) wrapDefaultStreamableTruncation(
	ctx context.Context,
	cfg *ToolReductionConfig,
	tCtx *adk.ToolContext,
	argumentsInJSON string,
	output *schema.StreamReader[string],
) *schema.StreamReader[string] {
	_ = "STUB: not implemented"
	return nil
}

func (t *typedToolReductionMiddleware[M]) WrapEnhancedInvokableToolCall(_ context.Context, endpoint adk.EnhancedInvokableToolCallEndpoint, tCtx *adk.ToolContext) (adk.EnhancedInvokableToolCallEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(adk.EnhancedInvokableToolCallEndpoint), nil
}

func (t *typedToolReductionMiddleware[M]) WrapEnhancedStreamableToolCall(_ context.Context, endpoint adk.EnhancedStreamableToolCallEndpoint, tCtx *adk.ToolContext) (adk.EnhancedStreamableToolCallEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(adk.EnhancedStreamableToolCallEndpoint), nil
}

func (t *typedToolReductionMiddleware[M]) wrapDefaultEnhancedStreamableTruncation(
	ctx context.Context,
	cfg *ToolReductionConfig,
	tCtx *adk.ToolContext,
	toolArgument *schema.ToolArgument,
	output *schema.StreamReader[*schema.ToolResult],
) *schema.StreamReader[*schema.ToolResult] {
	_ = "STUB: not implemented"
	return nil
}

func formatStreamTruncNotice(previewSize int, offloadNotify, errorMsgNotify string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func toolResultTextPrefix(result *schema.ToolResult, maxLength int) *schema.ToolResult {
	_ = "STUB: not implemented"
	return nil
}

func toolResultTextLength(result *schema.ToolResult) int { _ = "STUB: not implemented"; return 0 }

func flattenToolResultParts(chunks []*schema.ToolResult) []schema.ToolOutputPart {
	_ = "STUB: not implemented"
	return nil
}

func writeTruncOffload(ctx context.Context, cfg *ToolReductionConfig, truncResult *TruncResult) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *typedToolReductionMiddleware[M]) BeforeModelRewriteState(ctx context.Context, state *adk.TypedChatModelAgentState[M], mc *adk.TypedModelContext[M]) (
	context.Context, *adk.TypedChatModelAgentState[M], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func (t *typedToolReductionMiddleware[M]) beforeModelRewriteStateGeneric(ctx context.Context, state *adk.TypedChatModelAgentState[M], _ *adk.TypedModelContext[M]) (
	context.Context, *adk.TypedChatModelAgentState[M], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func (t *typedToolReductionMiddleware[M]) applyClearRewriteGeneric(ctx context.Context, state *adk.TypedChatModelAgentState[M], start, end int, clearAtLeastTokens int64) (
	[]M, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func agenticResultCallID(block *schema.ContentBlock) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func findToolResultByCallID[M adk.MessageType](messages []M, startIndex, endIndex int, callID string) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func getToolResultCallID[M adk.MessageType](msg M) string { _ = "STUB: not implemented"; return "" }

type offloadStashItem struct {
	config      *ToolReductionConfig
	offloadInfo *ClearResult
}

type toolCallInfo struct {
	BlockIndex int
	CallID     string
	Name       string
	Arguments  string
}

func isAssistantMsg[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func isSystemMsg[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func isUserMsg[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func hasToolCalls[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func isToolResultMsg[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func isToolResultOnlyMsg[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func getMsgClearedFlagGeneric[M adk.MessageType](msg M) bool {
	_ = "STUB: not implemented"
	return false
}

func setMsgClearedFlagGeneric[M adk.MessageType](msg M) { _ = "STUB: not implemented"; return }

func getToolCallsGeneric[M adk.MessageType](msg M) []toolCallInfo {
	_ = "STUB: not implemented"
	return nil
}

func setToolCallArguments[M adk.MessageType](msg M, blockIndex int, args string) {
	_ = "STUB: not implemented"
	return
}

func toolResultFromMsgGeneric[M adk.MessageType](msg M) (result *schema.ToolResult, fromContent bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func setToolResultContent[M adk.MessageType](msg M, toolResult *schema.ToolResult, fromContent bool) {
	_ = "STUB: not implemented"
	return
}

func copyMessagesGeneric[M adk.MessageType](msgs []M) []M { _ = "STUB: not implemented"; return nil }

func copyAgenticMessages(msgs []*schema.AgenticMessage) []*schema.AgenticMessage {
	_ = "STUB: not implemented"
	return nil
}

func copyMessages(msgs []*schema.Message) []*schema.Message { _ = "STUB: not implemented"; return nil }

func defaultTokenCounter(_ context.Context, msgs []*schema.Message, tools []*schema.ToolInfo) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func defaultTruncHandler(
	genOffloadFilePathFn func(ctx context.Context, toolDetail *ToolDetail) (filePath string, err error),
	truncMaxLength int,
	readFileToolName string,
) func(ctx context.Context, detail *ToolDetail) (truncResult *TruncResult, err error) {
	_ = "STUB: not implemented"
	return nil
}

func clampPrefixToUTF8Boundary(s string, n int) string { _ = "STUB: not implemented"; return "" }

func clampSuffixToUTF8Boundary(s string, n int) string { _ = "STUB: not implemented"; return "" }

func defaultClearHandler(
	genOffloadFilePathFn func(ctx context.Context, toolDetail *ToolDetail) (filePath string, err error),
	needOffload bool,
	readFileToolName string,
) func(ctx context.Context, detail *ToolDetail) (*ClearResult, error) {
	_ = "STUB: not implemented"
	return nil
}

func getJointToolResult(toolDetail *ToolDetail) (toolOutputParts []schema.ToolOutputPart, needProcess bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func stringifyToolOutputParts(toolOutputParts []schema.ToolOutputPart) string {
	_ = "STUB: not implemented"
	return ""
}

func getMsgClearedFlag(msg *schema.Message) (offloaded bool) {
	_ = "STUB: not implemented"
	return false
}

func setMsgClearedFlag(msg *schema.Message) { _ = "STUB: not implemented"; return }

func toolResultFromMessage(msg *schema.Message) (result *schema.ToolResult, fromContent bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func convMessageInputPartToToolOutputPart(msgPart schema.MessageInputPart) (schema.ToolOutputPart, error) {
	_ = "STUB: not implemented"
	return *new(schema.ToolOutputPart), nil
}

func toolResultToOutputParts(f *schema.FunctionToolResult) []schema.ToolOutputPart {
	_ = "STUB: not implemented"
	return nil
}

func setToolResultFromOutputParts(f *schema.FunctionToolResult, parts []schema.ToolOutputPart) {
	_ = "STUB: not implemented"
	return
}

func strPtr(s string) *string { _ = "STUB: not implemented"; return nil }

func ptrStr(p *string) string { _ = "STUB: not implemented"; return "" }
