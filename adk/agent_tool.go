package adk

import (
	"context"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

var (
	defaultAgentToolParam = schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
		"request": {
			Desc:     "request to be processed",
			Required: true,
			Type:     schema.String,
		},
	})
)

type AgentToolOptions struct {
	fullChatHistoryAsInput bool
	agentInputSchema       *schema.ParamsOneOf
}

type AgentToolOption func(*AgentToolOptions)

func WithFullChatHistoryAsInput() AgentToolOption {
	_ = "STUB: not implemented"
	return *new(AgentToolOption)
}

func WithAgentInputSchema(schema *schema.ParamsOneOf) AgentToolOption {
	_ = "STUB: not implemented"
	return *new(AgentToolOption)
}

func withAgentToolEnableStreaming(enabled bool) tool.Option {
	_ = "STUB: not implemented"
	return *new(tool.Option)
}

func NewAgentTool(_ context.Context, agent Agent, options ...AgentToolOption) tool.BaseTool {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool)
}

func NewTypedAgentTool[M MessageType](_ context.Context, agent TypedAgent[M], options ...AgentToolOption) tool.BaseTool {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool)
}

type typedAgentTool[M MessageType] struct {
	agent TypedAgent[M]

	fullChatHistoryAsInput bool
	inputSchema            *schema.ParamsOneOf
}

type agentTool = typedAgentTool[*schema.Message]

type agentToolRequest struct {
	Request string `json:"request"`
}

func (at *typedAgentTool[M]) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (at *typedAgentTool[M]) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type agentToolOptions struct {
	agentName       string
	opts            []AgentRunOption
	enableStreaming bool
}

type typedAgentToolEventOptions[M MessageType] struct {
	generator *AsyncGenerator[*TypedAgentEvent[M]]
}

func withAgentToolOptions(agentName string, opts []AgentRunOption) tool.Option {
	_ = "STUB: not implemented"
	return *new(tool.Option)
}

func withAgentToolEventGenerator(gen *AsyncGenerator[*AgentEvent]) tool.Option {
	_ = "STUB: not implemented"
	return *new(tool.Option)
}

func withTypedAgentToolEventGenerator[M MessageType](gen *AsyncGenerator[*TypedAgentEvent[M]]) tool.Option {
	_ = "STUB: not implemented"
	return *new(tool.Option)
}

func getOptionsByAgentName(agentName string, opts []tool.Option) []AgentRunOption {
	_ = "STUB: not implemented"
	return nil
}

func extractAndDeriveAgentToolCancelCtx(ctx context.Context, agentName string, opts []tool.Option) []AgentRunOption {
	_ = "STUB: not implemented"
	return nil
}

func getEmitGeneratorAndEnableStreaming[M MessageType](opts []tool.Option) (*AsyncGenerator[*TypedAgentEvent[M]], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func getReactChatHistory(ctx context.Context, destAgentName string) ([]Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTypedUserMessages[M MessageType](text string) []M { _ = "STUB: not implemented"; return nil }

func newTypedInvokableAgentToolRunner[M MessageType](agent TypedAgent[M], store compose.CheckPointStore, enableStreaming bool) *TypedRunner[M] {
	_ = "STUB: not implemented"
	return nil
}
