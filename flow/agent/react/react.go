package react

import (
	"context"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent"
	"github.com/cloudwego/eino/schema"
)

type toolResultSender func(toolName, callID, result string)

type enhancedToolResultSender func(toolName, callID string, result *schema.ToolResult)
type streamToolResultSender func(toolName, callID string, resultStream *schema.StreamReader[string])
type enhancedStreamToolResultSender func(toolName, callID string, resultStream *schema.StreamReader[*schema.ToolResult])
type toolResultSenders struct {
	sender       toolResultSender
	streamSender streamToolResultSender

	enhancedResultSender           enhancedToolResultSender
	enhancedStreamToolResultSender enhancedStreamToolResultSender
}

type toolResultSenderCtxKey struct{}

func setToolResultSendersToCtx(ctx context.Context, senders *toolResultSenders) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getToolResultSendersFromCtx(ctx context.Context) *toolResultSenders {
	_ = "STUB: not implemented"
	return nil
}

type state struct {
	Messages                 []*schema.Message
	ReturnDirectlyToolCallID string
}

func init() {
	schema.RegisterName[*state]("_eino_react_state")
}

func newToolResultCollectorMiddleware() compose.ToolMiddleware {
	_ = "STUB: not implemented"
	return *new(compose.ToolMiddleware)
}

const (
	nodeKeyTools = "tools"
	nodeKeyModel = "chat"
)

type MessageModifier func(ctx context.Context, input []*schema.Message) []*schema.Message

type AgentConfig struct {
	ToolCallingModel model.ToolCallingChatModel

	Model model.ChatModel

	ToolsConfig compose.ToolsNodeConfig

	MessageModifier MessageModifier

	MessageRewriter MessageModifier

	MaxStep int `json:"max_step"`

	ToolReturnDirectly map[string]struct{}

	StreamToolCallChecker func(ctx context.Context, modelOutput *schema.StreamReader[*schema.Message]) (bool, error)

	GraphName string

	ModelNodeName string

	ToolsNodeName string
}

func NewPersonaModifier(persona string) MessageModifier {
	_ = "STUB: not implemented"
	return *new(MessageModifier)
}

func firstChunkStreamToolCallChecker(_ context.Context, sr *schema.StreamReader[*schema.Message]) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

const (
	GraphName     = "ReActAgent"
	ModelNodeName = "ChatModel"
	ToolsNodeName = "Tools"
)

func SetReturnDirectly(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type Agent struct {
	runnable         compose.Runnable[[]*schema.Message, *schema.Message]
	graph            *compose.Graph[[]*schema.Message, *schema.Message]
	graphAddNodeOpts []compose.GraphAddNodeOpt
}

func NewAgent(ctx context.Context, config *AgentConfig) (_ *Agent, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildReturnDirectly(graph *compose.Graph[[]*schema.Message, *schema.Message]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func genToolInfos(ctx context.Context, config compose.ToolsNodeConfig) ([]*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getReturnDirectlyToolCallID(input *schema.Message, toolReturnDirectly map[string]struct{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *Agent) Generate(ctx context.Context, input []*schema.Message, opts ...agent.AgentOption) (*schema.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Agent) Stream(ctx context.Context, input []*schema.Message, opts ...agent.AgentOption) (output *schema.StreamReader[*schema.Message], err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Agent) ExportGraph() (compose.AnyGraph, []compose.GraphAddNodeOpt) {
	_ = "STUB: not implemented"
	return *new(compose.AnyGraph), nil
}
