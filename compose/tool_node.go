package compose

import (
	"context"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type toolsNodeOptions struct {
	ToolOptions []tool.Option
	ToolList    []tool.BaseTool

	ToolAliases map[string]ToolAliasConfig
}

type ToolsNodeOption func(o *toolsNodeOptions)

func WithToolOption(opts ...tool.Option) ToolsNodeOption {
	_ = "STUB: not implemented"
	return *new(ToolsNodeOption)
}

func WithToolList(tool ...tool.BaseTool) ToolsNodeOption {
	_ = "STUB: not implemented"
	return *new(ToolsNodeOption)
}

func WithToolAliases(toolAliases map[string]ToolAliasConfig) ToolsNodeOption {
	_ = "STUB: not implemented"
	return *new(ToolsNodeOption)
}

type ToolsNode struct {
	tuple                             *toolsTuple
	tools                             []tool.BaseTool
	unknownToolHandler                func(ctx context.Context, name, input string) (string, error)
	executeSequentially               bool
	toolArgumentsHandler              func(ctx context.Context, name, input string) (string, error)
	toolCallMiddlewares               []InvokableToolMiddleware
	streamToolCallMiddlewares         []StreamableToolMiddleware
	enhancedToolCallMiddlewares       []EnhancedInvokableToolMiddleware
	enhancedStreamToolCallMiddlewares []EnhancedStreamableToolMiddleware
	toolAliasConfigs                  map[string]ToolAliasConfig
}

type ToolInput struct {
	Name string

	Arguments string

	CallID string

	CallOptions []tool.Option
}

type ToolOutput struct {
	Result string
}

type StreamToolOutput struct {
	Result *schema.StreamReader[string]
}

type EnhancedInvokableToolOutput struct {
	Result *schema.ToolResult
}

type EnhancedStreamableToolOutput struct {
	Result *schema.StreamReader[*schema.ToolResult]
}

type InvokableToolEndpoint func(ctx context.Context, input *ToolInput) (*ToolOutput, error)

type StreamableToolEndpoint func(ctx context.Context, input *ToolInput) (*StreamToolOutput, error)

type EnhancedInvokableToolEndpoint func(ctx context.Context, input *ToolInput) (*EnhancedInvokableToolOutput, error)

type EnhancedStreamableToolEndpoint func(ctx context.Context, input *ToolInput) (*EnhancedStreamableToolOutput, error)

type InvokableToolMiddleware func(InvokableToolEndpoint) InvokableToolEndpoint

type StreamableToolMiddleware func(StreamableToolEndpoint) StreamableToolEndpoint

type EnhancedInvokableToolMiddleware func(EnhancedInvokableToolEndpoint) EnhancedInvokableToolEndpoint

type EnhancedStreamableToolMiddleware func(EnhancedStreamableToolEndpoint) EnhancedStreamableToolEndpoint

type ToolMiddleware struct {
	Invokable InvokableToolMiddleware

	Streamable StreamableToolMiddleware

	EnhancedInvokable EnhancedInvokableToolMiddleware

	EnhancedStreamable EnhancedStreamableToolMiddleware
}

type ToolAliasConfig struct {
	NameAliases []string

	ArgumentsAliases map[string][]string
}

type ToolsNodeConfig struct {
	Tools []tool.BaseTool

	ToolAliases map[string]ToolAliasConfig

	UnknownToolsHandler func(ctx context.Context, name, input string) (string, error)

	ExecuteSequentially bool

	ToolArgumentsHandler func(ctx context.Context, name, arguments string) (string, error)

	ToolCallMiddlewares []ToolMiddleware
}

func NewToolNode(ctx context.Context, conf *ToolsNodeConfig) (*ToolsNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ToolsInterruptAndRerunExtra struct {
	ToolCalls []schema.ToolCall

	ExecutedTools map[string]string

	ExecutedEnhancedTools map[string]*schema.ToolResult

	RerunTools []string

	RerunExtraMap map[string]any
}

func init() {
	schema.RegisterName[*ToolsInterruptAndRerunExtra]("_eino_compose_tools_interrupt_and_rerun_extra")
	schema.RegisterName[*toolsInterruptAndRerunState]("_eino_compose_tools_interrupt_and_rerun_state")
}

type toolsInterruptAndRerunState struct {
	Input                 *schema.Message
	ExecutedTools         map[string]string
	ExecutedEnhancedTools map[string]*schema.ToolResult
	RerunTools            []string
}

type toolsTuple struct {
	indexes                     map[string]int
	meta                        []*executorMeta
	endpoints                   []InvokableToolEndpoint
	streamEndpoints             []StreamableToolEndpoint
	enhancedInvokableEndpoints  []EnhancedInvokableToolEndpoint
	enhancedStreamableEndpoints []EnhancedStreamableToolEndpoint

	argsAliasMap map[string]map[string]string

	canonicalNames []string

	toolInfos []*schema.ToolInfo
}

func remapArgs(args string, aliasMap map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type convToolsParams struct {
	tools       []tool.BaseTool
	middlewares struct {
		invokable          []InvokableToolMiddleware
		streamable         []StreamableToolMiddleware
		enhancedInvokable  []EnhancedInvokableToolMiddleware
		enhancedStreamable []EnhancedStreamableToolMiddleware
	}
	aliasConfigs map[string]ToolAliasConfig
}

func (t *toolsTuple) applyAliasConfigs(aliasConfigs map[string]ToolAliasConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *toolsTuple) applyNameAliases(toolName string, toolIdx int, nameAliases []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *toolsTuple) applyArgsAliases(toolName string, toolIdx int, argumentsAliases map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func convTools(ctx context.Context, params convToolsParams) (*toolsTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func wrapToolCall(it tool.InvokableTool, middlewares []InvokableToolMiddleware, needCallback bool) InvokableToolEndpoint {
	_ = "STUB: not implemented"
	return *new(InvokableToolEndpoint)
}

func wrapStreamToolCall(st tool.StreamableTool, middlewares []StreamableToolMiddleware, needCallback bool) StreamableToolEndpoint {
	_ = "STUB: not implemented"
	return *new(StreamableToolEndpoint)
}

func wrapEnhancedInvokableToolCall(eiTool tool.EnhancedInvokableTool, middlewares []EnhancedInvokableToolMiddleware, needCallback bool) EnhancedInvokableToolEndpoint {
	_ = "STUB: not implemented"
	return *new(EnhancedInvokableToolEndpoint)
}

func wrapEnhancedStreamableToolCall(est tool.EnhancedStreamableTool, middlewares []EnhancedStreamableToolMiddleware, needCallback bool) EnhancedStreamableToolEndpoint {
	_ = "STUB: not implemented"
	return *new(EnhancedStreamableToolEndpoint)
}

type invokableToolWithCallback struct {
	it tool.InvokableTool
}

func (i *invokableToolWithCallback) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *invokableToolWithCallback) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type streamableToolWithCallback struct {
	st tool.StreamableTool
}

func (s *streamableToolWithCallback) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *streamableToolWithCallback) StreamableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (*schema.StreamReader[string], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type enhancedInvokableToolWithCallback struct {
	eiTool tool.EnhancedInvokableTool
}

func (e *enhancedInvokableToolWithCallback) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *enhancedInvokableToolWithCallback) InvokableRun(ctx context.Context, toolArgument *schema.ToolArgument, opts ...tool.Option) (*schema.ToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type enhancedStreamableToolWithCallback struct {
	est tool.EnhancedStreamableTool
}

func (e *enhancedStreamableToolWithCallback) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *enhancedStreamableToolWithCallback) StreamableRun(ctx context.Context, toolArgument *schema.ToolArgument, opts ...tool.Option) (*schema.StreamReader[*schema.ToolResult], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func streamableToInvokable(e StreamableToolEndpoint) InvokableToolEndpoint {
	_ = "STUB: not implemented"
	return *new(InvokableToolEndpoint)
}

func invokableToStreamable(e InvokableToolEndpoint) StreamableToolEndpoint {
	_ = "STUB: not implemented"
	return *new(StreamableToolEndpoint)
}

func enhancedStreamableToEnhancedInvokable(e EnhancedStreamableToolEndpoint) EnhancedInvokableToolEndpoint {
	_ = "STUB: not implemented"
	return *new(EnhancedInvokableToolEndpoint)
}

func enhancedInvokableToEnhancedStreamable(e EnhancedInvokableToolEndpoint) EnhancedStreamableToolEndpoint {
	_ = "STUB: not implemented"
	return *new(EnhancedStreamableToolEndpoint)
}

func invokeEnhancedWithCallbacks(i func(ctx context.Context, toolArgument *schema.ToolArgument, opts ...tool.Option) (*schema.ToolResult, error)) func(ctx context.Context, toolArgument *schema.ToolArgument, opts ...tool.Option) (*schema.ToolResult, error) {
	_ = "STUB: not implemented"
	return nil
}

func streamEnhancedWithCallbacks(s func(ctx context.Context, toolArgument *schema.ToolArgument, opts ...tool.Option) (*schema.StreamReader[*schema.ToolResult], error)) func(ctx context.Context, toolArgument *schema.ToolArgument, opts ...tool.Option) (*schema.StreamReader[*schema.ToolResult], error) {
	_ = "STUB: not implemented"
	return nil
}

type toolCallTask struct {
	endpoint                   InvokableToolEndpoint
	streamEndpoint             StreamableToolEndpoint
	enhancedInvokableEndpoint  EnhancedInvokableToolEndpoint
	enhancedStreamableEndpoint EnhancedStreamableToolEndpoint
	meta                       *executorMeta
	name                       string
	arg                        string
	callID                     string
	useEnhanced                bool

	executed        bool
	output          string
	sOutput         *schema.StreamReader[string]
	enhancedOutput  *schema.ToolResult
	enhancedSOutput *schema.StreamReader[*schema.ToolResult]
	err             error
}

func (tn *ToolsNode) genToolCallTasks(ctx context.Context, tuple *toolsTuple,
	input *schema.Message, executedTools map[string]string, executedEnhancedTools map[string]*schema.ToolResult, isStream bool) ([]toolCallTask, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newUnknownToolTask(name, arg, callID string, unknownToolHandler func(ctx context.Context, name, input string) (string, error)) toolCallTask {
	_ = "STUB: not implemented"
	return *new(toolCallTask)
}

func runToolCallTaskByInvoke(ctx context.Context, task *toolCallTask, opts ...tool.Option) {
	_ = "STUB: not implemented"
	return
}

func runToolCallTaskByStream(ctx context.Context, task *toolCallTask, opts ...tool.Option) {
	_ = "STUB: not implemented"
	return
}

func sequentialRunToolCall(ctx context.Context,
	run func(ctx2 context.Context, callTask *toolCallTask, opts ...tool.Option),
	tasks []toolCallTask, opts ...tool.Option) {
	_ = "STUB: not implemented"
	return
}

func parallelRunToolCall(ctx context.Context,
	run func(ctx2 context.Context, callTask *toolCallTask, opts ...tool.Option),
	tasks []toolCallTask, opts ...tool.Option) {
	_ = "STUB: not implemented"
	return
}

func (tn *ToolsNode) buildTupleFromOpts(ctx context.Context, opt *toolsNodeOptions) (*toolsTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tn *ToolsNode) Invoke(ctx context.Context, input *schema.Message,
	opts ...ToolsNodeOption) ([]*schema.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tn *ToolsNode) Stream(ctx context.Context, input *schema.Message,
	opts ...ToolsNodeOption) (*schema.StreamReader[[]*schema.Message], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tn *ToolsNode) GetType() string { _ = "STUB: not implemented"; return "" }

func getToolsNodeOptions(opts ...ToolsNodeOption) *toolsNodeOptions {
	_ = "STUB: not implemented"
	return nil
}

type toolCallInfoKey struct{}
type toolCallInfo struct {
	toolCallID string
}

func setToolCallInfo(ctx context.Context, toolCallInfo *toolCallInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetToolCallID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }
