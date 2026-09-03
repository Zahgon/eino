package summarization

import (
	"context"
	"regexp"
	"time"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

func init() {
	schema.RegisterName[*TypedCustomizedAction[*schema.Message]]("_eino_adk_summarization_mw_customized_action")
	schema.RegisterName[*TypedCustomizedAction[*schema.AgenticMessage]]("_eino_adk_summarization_mw_customized_action_agentic")
}

type TypedTokenCounterFunc[M adk.MessageType] func(ctx context.Context, input *TypedTokenCounterInput[M]) (int, error)
type TypedGenModelInputFunc[M adk.MessageType] func(ctx context.Context, sysInstruction, userInstruction M, originalMsgs []M) ([]M, error)
type TypedGetFailoverModelFunc[M adk.MessageType] func(ctx context.Context, failoverCtx *TypedFailoverContext[M]) (failoverModel model.BaseModel[M], failoverModelInputMsgs []M, failoverErr error)
type TypedFinalizeFunc[M adk.MessageType] func(ctx context.Context, originalMessages []M, summary M) ([]M, error)
type TypedCallbackFunc[M adk.MessageType] func(ctx context.Context, before, after adk.TypedChatModelAgentState[M]) error

type TokenCounterFunc = TypedTokenCounterFunc[*schema.Message]
type GenModelInputFunc = TypedGenModelInputFunc[*schema.Message]
type GetFailoverModelFunc = TypedGetFailoverModelFunc[*schema.Message]
type FinalizeFunc = TypedFinalizeFunc[*schema.Message]
type CallbackFunc = TypedCallbackFunc[*schema.Message]

type TypedConfig[M adk.MessageType] struct {
	Model model.BaseModel[M]

	ModelOptions []model.Option

	TokenCounter TypedTokenCounterFunc[M]

	Trigger *TriggerCondition

	EmitInternalEvents bool

	UserInstruction string

	TranscriptFilePath string

	GenModelInput TypedGenModelInputFunc[M]

	Finalize TypedFinalizeFunc[M]

	Callback TypedCallbackFunc[M]

	Retry *TypedRetryConfig[M]

	Failover *TypedFailoverConfig[M]
}

type Config = TypedConfig[*schema.Message]

type TypedTokenCounterInput[M adk.MessageType] struct {
	Messages []M

	Tools []*schema.ToolInfo
}

type TokenCounterInput = TypedTokenCounterInput[*schema.Message]

type TriggerCondition struct {
	ContextTokens int

	ContextMessages int
}

type TypedRetryConfig[M adk.MessageType] struct {
	MaxRetries *int

	ShouldRetry func(ctx context.Context, resp M, err error) bool

	BackoffFunc func(ctx context.Context, attempt int, resp M, err error) time.Duration
}

type RetryConfig = TypedRetryConfig[*schema.Message]

type TypedFailoverConfig[M adk.MessageType] struct {
	MaxRetries *int

	ShouldFailover func(ctx context.Context, resp M, err error) bool

	BackoffFunc func(ctx context.Context, attempt int, resp M, err error) time.Duration

	GetFailoverModel TypedGetFailoverModelFunc[M]
}

type FailoverConfig = TypedFailoverConfig[*schema.Message]

type TypedFailoverContext[M adk.MessageType] struct {
	Attempt int

	SystemInstruction M

	UserInstruction M

	OriginalMessages []M

	LastModelResponse M

	LastErr error
}

type FailoverContext = TypedFailoverContext[*schema.Message]

func NewTyped[M adk.MessageType](_ context.Context, cfg *TypedConfig[M]) (adk.TypedChatModelAgentMiddleware[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(ctx context.Context, cfg *Config) (adk.ChatModelAgentMiddleware, error) {
	_ = "STUB: not implemented"
	return *new(adk.ChatModelAgentMiddleware), nil
}

type TypedMiddleware[M adk.MessageType] struct {
	*adk.TypedBaseChatModelAgentMiddleware[M]
	cfg *TypedConfig[M]
}

func (m *TypedMiddleware[M]) Summarize(ctx context.Context, state *adk.TypedChatModelAgentState[M]) ([]M, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *TypedMiddleware[M]) BeforeModelRewriteState(ctx context.Context, state *adk.TypedChatModelAgentState[M],
	_ *adk.TypedModelContext[M]) (context.Context, *adk.TypedChatModelAgentState[M], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func (m *TypedMiddleware[M]) shouldSummarize(ctx context.Context, input *TypedTokenCounterInput[M]) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *TypedMiddleware[M]) getTriggerContextTokens() int { _ = "STUB: not implemented"; return 0 }

func (m *TypedMiddleware[M]) emitEvent(ctx context.Context, action *TypedCustomizedAction[M]) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *TypedMiddleware[M]) emitGenerateSummaryEvent(ctx context.Context, attempt int, phase GenerateSummaryPhase,
	resp M, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *TypedMiddleware[M]) countTokens(ctx context.Context, input *TypedTokenCounterInput[M]) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func defaultTypedTokenCounter[M adk.MessageType](_ context.Context, input *TypedTokenCounterInput[M]) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getAssistantTotalTokens[M adk.MessageType](msg M) int { _ = "STUB: not implemented"; return 0 }

func estimateTokenCount(charLen int) int { _ = "STUB: not implemented"; return 0 }

func estimateTokenBytes(tokens int) int { _ = "STUB: not implemented"; return 0 }

func (m *TypedMiddleware[M]) summarize(ctx context.Context, originalMsgs []M) (M, []M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil, nil
}

func splitSystemAndContextMsgs[M adk.MessageType](msgs []M) ([]M, []M) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *TypedMiddleware[M]) runFailover(ctx context.Context, originalMsgs, defaultInput []M, lastResp M,
	lastErr error) (M, []M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil, nil
}

func (m *TypedMiddleware[M]) getFailoverModel(ctx context.Context, failoverCtx *TypedFailoverContext[M], defaultInput []M) (model.BaseModel[M], []M, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (m *TypedMiddleware[M]) buildSummarizationModelInput(ctx context.Context, originMsgs, contextMsgs []M) ([]M, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *TypedMiddleware[M]) getModelInstructions() (M, M) {
	_ = "STUB: not implemented"
	return *new(M), *new(M)
}

func buildInternalFinalizer[M adk.MessageType](cfg *TypedConfig[M]) TypedFinalizeFunc[M] {
	_ = "STUB: not implemented"
	return nil
}

func getAssistantTextContent[M adk.MessageType](msg M) string { _ = "STUB: not implemented"; return "" }

type postProcessSummaryParams[M adk.MessageType] struct {
	contextMsgs    []M
	summaryContent string
	transcriptPath string
}

func postProcessSummary[M adk.MessageType](ctx context.Context, p *postProcessSummaryParams[M]) (processed M, err error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

type replaceUserMessagesInSummaryParams[M adk.MessageType] struct {
	contextMsgs []M
	summaryText string
}

func replaceUserMessagesInSummary[M adk.MessageType](ctx context.Context, p *replaceUserMessagesInSummaryParams[M]) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func findLastMatch(re *regexp.Regexp, s string) []int { _ = "STUB: not implemented"; return nil }

func appendSection(base, section string) string { _ = "STUB: not implemented"; return "" }

func (m *TypedMiddleware[M]) generateAndEmit(ctx context.Context, chatModel model.BaseModel[M], input []M,
	opts []model.Option, attempt int, phase GenerateSummaryPhase) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func (m *TypedMiddleware[M]) generateWithRetry(ctx context.Context, chatModel model.BaseModel[M], input []M,
	opts []model.Option, retryCfg *TypedRetryConfig[M]) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func truncateTextByChars(text string) string { _ = "STUB: not implemented"; return "" }

func (c *TypedConfig[M]) check() error { _ = "STUB: not implemented"; return nil }

func (c *TypedRetryConfig[M]) check() error { _ = "STUB: not implemented"; return nil }

func (c *TypedFailoverConfig[M]) check() error { _ = "STUB: not implemented"; return nil }

func (c *TriggerCondition) check() error { _ = "STUB: not implemented"; return nil }

func isSystemRole[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func isUserRole[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func messageUserTextContent(m *schema.Message) string { _ = "STUB: not implemented"; return "" }

func getUserMsgTextContent[M adk.MessageType](msg M) string { _ = "STUB: not implemented"; return "" }

const multimodalTokenEstimate = 2000

func estimateMessageTokens(msg *schema.Message) int { _ = "STUB: not implemented"; return 0 }

func estimateAgenticMessageTokens(msg *schema.AgenticMessage) int {
	_ = "STUB: not implemented"
	return 0
}

func getMsgExtra[M adk.MessageType](msg M) map[string]any { _ = "STUB: not implemented"; return nil }

func setMsgExtra[M adk.MessageType](msg M, key string, value any) {
	_ = "STUB: not implemented"
	return
}

func makeSystemMsg[M adk.MessageType](text string) M { _ = "STUB: not implemented"; return *new(M) }

func makeUserMsg[M adk.MessageType](text string) M { _ = "STUB: not implemented"; return *new(M) }

func newTypedSummaryMessage[M adk.MessageType](content string) M {
	_ = "STUB: not implemented"
	return *new(M)
}

func typedGetContentType[M adk.MessageType](msg M) summarizationContentType {
	_ = "STUB: not implemented"
	return *new(summarizationContentType)
}

func isInternalUserMessage[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func isPreservedMessage[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func typedShouldFailover[M adk.MessageType](ctx context.Context, cfg *TypedFailoverConfig[M], resp M, err error) bool {
	_ = "STUB: not implemented"
	return false
}

func defaultTypedShouldRetry[M adk.MessageType](_ context.Context, _ M, err error) bool {
	_ = "STUB: not implemented"
	return false
}

func defaultTypedBackoffFunc[M adk.MessageType](_ context.Context, attempt int, _ M, _ error) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func defaultBackoffDuration(attempt int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func defaultTypedTrimUserMessage[M adk.MessageType](msg M, remainingTokens int) M {
	_ = "STUB: not implemented"
	return *new(M)
}
