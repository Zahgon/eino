package adk

import (
	"context"

	"github.com/cloudwego/eino/callbacks"
)

type AgentCallbackInput struct {
	Input *AgentInput

	ResumeInfo *ResumeInfo
}

type AgentCallbackOutput struct {
	Events *AsyncIterator[*AgentEvent]
}

func copyTypedEventIterator[M MessageType](iter *AsyncIterator[*TypedAgentEvent[M]], n int) []*AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}

func copyAgentCallbackOutput(out *AgentCallbackOutput, n int) []*AgentCallbackOutput {
	_ = "STUB: not implemented"
	return nil
}

func ConvAgentCallbackInput(input callbacks.CallbackInput) *AgentCallbackInput {
	_ = "STUB: not implemented"
	return nil
}

func ConvAgentCallbackOutput(output callbacks.CallbackOutput) *AgentCallbackOutput {
	_ = "STUB: not implemented"
	return nil
}

func initAgentCallbacks(ctx context.Context, agentName, agentType string, opts ...AgentRunOption) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getAgentType(agent Agent) string { _ = "STUB: not implemented"; return "" }

type TypedAgentCallbackInput[M MessageType] struct {
	Input *TypedAgentInput[M]

	ResumeInfo *ResumeInfo
}

type TypedAgentCallbackOutput[M MessageType] struct {
	Events *AsyncIterator[*TypedAgentEvent[M]]
}

func ConvTypedCallbackInput[M MessageType](input callbacks.CallbackInput) *TypedAgentCallbackInput[M] {
	_ = "STUB: not implemented"
	return nil
}

func ConvTypedCallbackOutput[M MessageType](output callbacks.CallbackOutput) *TypedAgentCallbackOutput[M] {
	_ = "STUB: not implemented"
	return nil
}

func copyTypedCallbackOutput[M MessageType](out *TypedAgentCallbackOutput[M], n int) []*TypedAgentCallbackOutput[M] {
	_ = "STUB: not implemented"
	return nil
}

func initAgenticCallbacks(ctx context.Context, agentName, agentType string, opts ...AgentRunOption) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
