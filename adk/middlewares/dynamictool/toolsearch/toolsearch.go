package toolsearch

import (
	"context"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type Config struct {
	DynamicTools []tool.BaseTool

	UseModelToolSearch bool
}

func NewTyped[M adk.MessageType](ctx context.Context, config *Config) (adk.TypedChatModelAgentMiddleware[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(ctx context.Context, config *Config) (adk.ChatModelAgentMiddleware, error) {
	_ = "STUB: not implemented"
	return *new(adk.ChatModelAgentMiddleware), nil
}

type systemReminder struct {
	Tools []string
}

type typedMiddleware[M adk.MessageType] struct {
	*adk.TypedBaseChatModelAgentMiddleware[M]
	dynamicTools       []tool.BaseTool
	mapOfDynamicTools  map[string]*schema.ToolInfo
	dynamicToolInfos   []*schema.ToolInfo
	useModelToolSearch bool
	sr                 string
}

func (m *typedMiddleware[M]) BeforeAgent(ctx context.Context, runCtx *adk.ChatModelAgentContext) (context.Context, *adk.ChatModelAgentContext, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

const toolSearchInitializedKey = "__toolsearch_initialized__"
const toolSearchReminderExtraKey = "__toolsearch_reminder__"

func (m *typedMiddleware[M]) isInitialized(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *typedMiddleware[M]) markInitialized(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (m *typedMiddleware[M]) ensureReminder(msgs []M) []M { _ = "STUB: not implemented"; return nil }

func isSystemRoleTS[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func makeReminderMsg[M adk.MessageType](content string) M {
	_ = "STUB: not implemented"
	return *new(M)
}

func hasToolSearchReminderExtra[M adk.MessageType](msg M) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *typedMiddleware[M]) extractDynamicTools(tools []*schema.ToolInfo) []*schema.ToolInfo {
	_ = "STUB: not implemented"
	return nil
}

func (m *typedMiddleware[M]) stripDynamicTools(tools []*schema.ToolInfo) []*schema.ToolInfo {
	_ = "STUB: not implemented"
	return nil
}

func removeTool(tools []*schema.ToolInfo, name string) []*schema.ToolInfo {
	_ = "STUB: not implemented"
	return nil
}

func toolNameSet(tools []*schema.ToolInfo) map[string]bool { _ = "STUB: not implemented"; return nil }

func (m *typedMiddleware[M]) BeforeModelRewriteState(ctx context.Context, state *adk.TypedChatModelAgentState[M], _ *adk.TypedModelContext[M]) (context.Context, *adk.TypedChatModelAgentState[M], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func extractToolSearchResult[M adk.MessageType](msg M, toolName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func newToolSearchTool(tools map[string]*schema.ToolInfo, useModelToolSearch bool) tool.BaseTool {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool)
}

type toolSearchArgs struct {
	Query      string `json:"query"`
	MaxResults *int   `json:"max_results,omitempty"`
}

type toolSearchResult struct {
	Matches []string `json:"matches"`
}

type toolSearchTool struct {
	tools map[string]*schema.ToolInfo
}

func (t *toolSearchTool) Info(_ context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *toolSearchTool) InvokableRun(_ context.Context, argumentsInJSON string, _ ...tool.Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type modelToolSearchTool struct {
	tools map[string]*schema.ToolInfo
}

func (t *modelToolSearchTool) Info(_ context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *modelToolSearchTool) InvokableRun(_ context.Context, argumentsInJSON *schema.ToolArgument, _ ...tool.Option) (*schema.ToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	toolSearchToolName = "tool_search"
	defaultMaxResults  = 5
)

func getToolSearchToolInfo() *schema.ToolInfo { _ = "STUB: not implemented"; return nil }

func search(argumentsInJSON string, tools map[string]*schema.ToolInfo) ([]*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func intMax(a, b int) int { _ = "STUB: not implemented"; return 0 }

func intMin(a, b int) int { _ = "STUB: not implemented"; return 0 }

type scoredTool struct {
	name  string
	score int
}

func keywordSearch(query string, maxResults int, tools map[string]*schema.ToolInfo) []string {
	_ = "STUB: not implemented"
	return nil
}

type keyword struct {
	word     string
	required bool
}

func parseKeywords(query string) (keywords []keyword) { _ = "STUB: not implemented"; return nil }

func splitToolName(name string) []string { _ = "STUB: not implemented"; return nil }

func splitCamelCase(s string) []string { _ = "STUB: not implemented"; return nil }
