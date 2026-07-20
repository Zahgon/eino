package host

import (
	"context"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent"
	"github.com/cloudwego/eino/schema"
)

type MultiAgent struct {
	runnable         compose.Runnable[[]*schema.Message, *schema.Message]
	graph            *compose.Graph[[]*schema.Message, *schema.Message]
	graphAddNodeOpts []compose.GraphAddNodeOpt
}

func (ma *MultiAgent) Generate(ctx context.Context, input []*schema.Message, opts ...agent.AgentOption) (*schema.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ma *MultiAgent) Stream(ctx context.Context, input []*schema.Message, opts ...agent.AgentOption) (*schema.StreamReader[*schema.Message], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ma *MultiAgent) ExportGraph() (compose.AnyGraph, []compose.GraphAddNodeOpt) {
	_ = "STUB: not implemented"
	return *new(compose.AnyGraph), nil
}

func (ma *MultiAgent) HostNodeKey() string { _ = "STUB: not implemented"; return "" }

type MultiAgentConfig struct {
	Host        Host
	Specialists []*Specialist

	Name         string
	HostNodeName string

	StreamToolCallChecker func(ctx context.Context, modelOutput *schema.StreamReader[*schema.Message]) (bool, error)

	Summarizer *Summarizer
}

func (conf *MultiAgentConfig) validate() error { _ = "STUB: not implemented"; return nil }

type AgentMeta struct {
	Name        string
	IntendedUse string
}

func (am AgentMeta) validate() error { _ = "STUB: not implemented"; return nil }

type Host struct {
	ToolCallingModel model.ToolCallingChatModel

	ChatModel    model.ChatModel
	SystemPrompt string
}

type Specialist struct {
	AgentMeta

	ChatModel    model.BaseChatModel
	SystemPrompt string

	Invokable  compose.Invoke[[]*schema.Message, *schema.Message, agent.AgentOption]
	Streamable compose.Stream[[]*schema.Message, *schema.Message, agent.AgentOption]
}

type Summarizer struct {
	ChatModel    model.BaseChatModel
	SystemPrompt string
}

func firstChunkStreamToolCallChecker(_ context.Context, sr *schema.StreamReader[*schema.Message]) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
