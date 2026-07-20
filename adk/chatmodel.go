package adk

import (
	"context"
	"sync"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

var _ ResumableAgent = &TypedChatModelAgent[*schema.Message]{}
var _ TypedResumableAgent[*schema.AgenticMessage] = &TypedChatModelAgent[*schema.AgenticMessage]{}

type typedChatModelAgentExecCtx[M MessageType] struct {
	runtimeReturnDirectly map[string]bool
	generator             *AsyncGenerator[*TypedAgentEvent[M]]
	cancelCtx             *cancelContext

	failoverLastSuccessModel model.BaseModel[M]

	suppressEventSend  bool
	retryVerdictSignal *retryVerdictSignal

	afterToolCallsHook func(ctx context.Context) error
}

func (e *typedChatModelAgentExecCtx[M]) send(event *TypedAgentEvent[M]) {
	_ = "STUB: not implemented"
	return
}

type chatModelAgentExecCtx = typedChatModelAgentExecCtx[*schema.Message]

type typedChatModelAgentExecCtxKey[M MessageType] struct{}

func withTypedChatModelAgentExecCtx[M MessageType](ctx context.Context, execCtx *typedChatModelAgentExecCtx[M]) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getTypedChatModelAgentExecCtx[M MessageType](ctx context.Context) *typedChatModelAgentExecCtx[M] {
	_ = "STUB: not implemented"
	return nil
}

type chatModelAgentRunOptions struct {
	chatModelOptions []model.Option
	toolOptions      []tool.Option
	agentToolOptions map[string][]AgentRunOption

	historyModifier func(context.Context, []Message) []Message

	afterToolCallsHook func(ctx context.Context) error
}

func WithChatModelOptions(opts []model.Option) AgentRunOption {
	_ = "STUB: not implemented"
	return *new(AgentRunOption)
}

func WithToolOptions(opts []tool.Option) AgentRunOption {
	_ = "STUB: not implemented"
	return *new(AgentRunOption)
}

func WithAgentToolRunOptions(opts map[string][]AgentRunOption) AgentRunOption {
	_ = "STUB: not implemented"
	return *new(AgentRunOption)
}

func WithHistoryModifier(f func(context.Context, []Message) []Message) AgentRunOption {
	_ = "STUB: not implemented"
	return *new(AgentRunOption)
}

func WithAfterToolCallsHook(fn func(ctx context.Context) error) AgentRunOption {
	_ = "STUB: not implemented"
	return *new(AgentRunOption)
}

type ToolsConfig struct {
	compose.ToolsNodeConfig

	ReturnDirectly map[string]bool

	EmitInternalEvents bool
}

type TypedGenModelInput[M MessageType] func(ctx context.Context, instruction string, input *TypedAgentInput[M]) ([]M, error)

type GenModelInput = TypedGenModelInput[*schema.Message]

func defaultGenModelInput(ctx context.Context, instruction string, input *AgentInput) ([]Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newDefaultGenModelInput[M MessageType]() TypedGenModelInput[M] {
	_ = "STUB: not implemented"
	return nil
}

type TypedChatModelAgentState[M MessageType] struct {
	Messages []M

	ToolInfos []*schema.ToolInfo

	DeferredToolInfos []*schema.ToolInfo
}

type ChatModelAgentState = TypedChatModelAgentState[*schema.Message]

type AgentMiddleware struct {
	AdditionalInstruction string

	AdditionalTools []tool.BaseTool

	BeforeChatModel func(context.Context, *ChatModelAgentState) error

	AfterChatModel func(context.Context, *ChatModelAgentState) error

	WrapToolCall compose.ToolMiddleware
}

type TypedChatModelAgentConfig[M MessageType] struct {
	Name string

	Description string

	Instruction string

	Model model.BaseModel[M]

	ToolsConfig ToolsConfig

	GenModelInput TypedGenModelInput[M]

	Exit tool.BaseTool

	OutputKey string

	MaxIterations int

	Middlewares []AgentMiddleware

	Handlers []TypedChatModelAgentMiddleware[M]

	ModelRetryConfig *TypedModelRetryConfig[M]

	ModelFailoverConfig *ModelFailoverConfig[M]
}

type ChatModelAgentConfig = TypedChatModelAgentConfig[*schema.Message]

type TypedChatModelAgent[M MessageType] struct {
	name        string
	description string
	instruction string

	model       model.BaseModel[M]
	toolsConfig ToolsConfig

	genModelInput TypedGenModelInput[M]

	outputKey     string
	maxIterations int

	subAgents   []TypedAgent[M]
	parentAgent TypedAgent[M]

	disallowTransferToParent bool

	exit tool.BaseTool

	handlers    []TypedChatModelAgentMiddleware[M]
	middlewares []AgentMiddleware

	modelRetryConfig    *TypedModelRetryConfig[M]
	modelFailoverConfig *ModelFailoverConfig[M]

	once   sync.Once
	run    typedRunFunc[M]
	frozen uint32
	exeCtx *execContext
}

type ChatModelAgent = TypedChatModelAgent[*schema.Message]

type typedRunParams[M MessageType] struct {
	input          *TypedAgentInput[M]
	generator      *AsyncGenerator[*TypedAgentEvent[M]]
	store          *bridgeStore
	instruction    string
	returnDirectly map[string]bool
	cancelCtx      *cancelContext
	cancelCtxOwned bool
	composeOpts    []compose.Option

	afterToolCallsHook func(ctx context.Context) error
}

type typedRunFunc[M MessageType] func(ctx context.Context, p *typedRunParams[M])

func isCheckpointAwareCancelScope(cc *cancelContext) bool { _ = "STUB: not implemented"; return false }

func resolveRunCancelContext(ctx context.Context, o *options) (*cancelContext, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func NewChatModelAgent(ctx context.Context, config *ChatModelAgentConfig) (*ChatModelAgent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTypedChatModelAgent[M MessageType](ctx context.Context, config *TypedChatModelAgentConfig[M]) (*TypedChatModelAgent[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectToolMiddlewaresFromMiddlewares(mws []AgentMiddleware) []compose.ToolMiddleware {
	_ = "STUB: not implemented"
	return nil
}

const (
	TransferToAgentToolName        = "transfer_to_agent"
	TransferToAgentToolDesc        = "Transfer the question to another agent."
	TransferToAgentToolDescChinese = "将问题移交给其他 Agent。"
)

var (
	toolInfoTransferToAgent = &schema.ToolInfo{
		Name: TransferToAgentToolName,

		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"agent_name": {
				Desc:     "the name of the agent to transfer to",
				Required: true,
				Type:     schema.String,
			},
		}),
	}

	ToolInfoExit = &schema.ToolInfo{
		Name: "exit",
		Desc: "Exit the agent process and return the final result.",

		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"final_result": {
				Desc:     "the final result to return",
				Required: true,
				Type:     schema.String,
			},
		}),
	}
)

type ExitTool struct{}

func (et ExitTool) Info(_ context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (et ExitTool) InvokableRun(ctx context.Context, argumentsInJSON string, _ ...tool.Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type transferToAgent struct{}

func (tta transferToAgent) Info(_ context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func transferToAgentToolOutput(destName string) string { _ = "STUB: not implemented"; return "" }

func (tta transferToAgent) InvokableRun(ctx context.Context, argumentsInJSON string, _ ...tool.Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (a *TypedChatModelAgent[M]) Name(_ context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

func (a *TypedChatModelAgent[M]) Description(_ context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

func (a *TypedChatModelAgent[M]) GetType() string { _ = "STUB: not implemented"; return "" }

func (a *TypedChatModelAgent[M]) OnSetSubAgents(_ context.Context, subAgents []TypedAgent[M]) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *TypedChatModelAgent[M]) OnSetAsSubAgent(_ context.Context, parent TypedAgent[M]) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *TypedChatModelAgent[M]) OnDisallowTransferToParent(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type ChatModelAgentInterruptInfo struct {
	Info *compose.InterruptInfo
	Data []byte
}

func init() {
	schema.RegisterName[*ChatModelAgentInterruptInfo]("_eino_adk_chat_model_agent_interrupt_info")
}

func extractTextContent[M MessageType](msg M) string { _ = "STUB: not implemented"; return "" }

func setOutputToSession[M MessageType](ctx context.Context, msg M, msgStream *schema.StreamReader[M], outputKey string) error {
	_ = "STUB: not implemented"
	return nil
}

func typedErrFunc[M MessageType](err error) typedRunFunc[M] { _ = "STUB: not implemented"; return nil }

type ChatModelAgentResumeData struct {
	HistoryModifier func(ctx context.Context, history []Message) []Message
}

type execContext struct {
	instruction    string
	toolsNodeConf  compose.ToolsNodeConfig
	returnDirectly map[string]bool

	toolInfos      []*schema.ToolInfo
	unwrappedTools []tool.BaseTool

	toolSearchTool *schema.ToolInfo

	rebuildGraph bool
	toolUpdated  bool
}

func (a *TypedChatModelAgent[M]) applyBeforeAgent(ctx context.Context, ec *execContext) (context.Context, *execContext, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func (a *TypedChatModelAgent[M]) applyAfterAgent(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (a *TypedChatModelAgent[M]) prepareExecContext(ctx context.Context) (*execContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *TypedChatModelAgent[M]) handleRunFuncError(
	ctx context.Context,
	err error,
	cancelCtx *cancelContext,
	cancelCtxOwned bool,
	store *bridgeStore,
	generator *AsyncGenerator[*TypedAgentEvent[M]],
) {
	_ = "STUB: not implemented"
	return
}

type typedNoToolsInput[M MessageType] struct {
	input       *TypedAgentInput[M]
	instruction string
}

func appendModelToChain[I, O any, M MessageType](chain *compose.Chain[I, O], m model.BaseModel[M]) {
	_ = "STUB: not implemented"
	return
}

func (a *TypedChatModelAgent[M]) buildNoToolsRunFunc(_ context.Context) (typedRunFunc[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *TypedChatModelAgent[M]) buildReActRunFunc(ctx context.Context, bc *execContext) (typedRunFunc[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type reactRunInput struct {
	input       *AgentInput
	instruction string
}

func (a *TypedChatModelAgent[M]) buildMessageReActRunFunc(_ context.Context, bc *execContext) (typedRunFunc[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type agenticReactRunInput struct {
	input       *TypedAgentInput[*schema.AgenticMessage]
	instruction string
}

func (a *TypedChatModelAgent[M]) buildAgenticReActRunFunc(_ context.Context, bc *execContext) (typedRunFunc[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *TypedChatModelAgent[M]) buildRunFunc(ctx context.Context) typedRunFunc[M] {
	_ = "STUB: not implemented"
	return nil
}

func (a *TypedChatModelAgent[M]) getRunFunc(ctx context.Context) (context.Context, typedRunFunc[M], *execContext, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil, nil
}

func (a *TypedChatModelAgent[M]) Run(ctx context.Context, input *TypedAgentInput[M], opts ...AgentRunOption) *AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}

func (a *TypedChatModelAgent[M]) Resume(ctx context.Context, info *ResumeInfo, opts ...AgentRunOption) *AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}

func getComposeOptions(opts []AgentRunOption) []compose.Option {
	_ = "STUB: not implemented"
	return nil
}

type gobSerializer struct{}

func (g *gobSerializer) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (g *gobSerializer) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

func preprocessComposeCheckpoint(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
