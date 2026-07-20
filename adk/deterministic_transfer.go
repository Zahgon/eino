package adk

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

func init() {
	schema.RegisterName[*deterministicTransferState]("_eino_adk_deterministic_transfer_state")
}

type deterministicTransferState struct {
	EventList []*agentEventWrapper
}

func AgentWithDeterministicTransferTo(_ context.Context, config *DeterministicTransferConfig) Agent {
	_ = "STUB: not implemented"
	return *new(Agent)
}

type agentWithDeterministicTransferTo struct {
	agent        Agent
	toAgentNames []string
}

func (a *agentWithDeterministicTransferTo) Description(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

func (a *agentWithDeterministicTransferTo) Name(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

func (a *agentWithDeterministicTransferTo) GetType() string { _ = "STUB: not implemented"; return "" }

func (a *agentWithDeterministicTransferTo) Run(ctx context.Context,
	input *AgentInput, options ...AgentRunOption) *AsyncIterator[*AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

type resumableAgentWithDeterministicTransferTo struct {
	agent        ResumableAgent
	toAgentNames []string
}

func (a *resumableAgentWithDeterministicTransferTo) Description(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

func (a *resumableAgentWithDeterministicTransferTo) Name(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

func (a *resumableAgentWithDeterministicTransferTo) GetType() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *resumableAgentWithDeterministicTransferTo) Run(ctx context.Context,
	input *AgentInput, options ...AgentRunOption) *AsyncIterator[*AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

func (a *resumableAgentWithDeterministicTransferTo) Resume(ctx context.Context, info *ResumeInfo, opts ...AgentRunOption) *AsyncIterator[*AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

func forwardEventsAndAppendTransfer(iter *AsyncIterator[*AgentEvent],
	generator *AsyncGenerator[*AgentEvent], toAgentNames []string) {
	_ = "STUB: not implemented"
	return
}

func runFlowAgentWithIsolatedSession(ctx context.Context, fa *flowAgent, input *AgentInput,
	toAgentNames []string, options ...AgentRunOption) *AsyncIterator[*AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

func resumeFlowAgentWithIsolatedSession(ctx context.Context, fa *flowAgent, info *ResumeInfo,
	toAgentNames []string, opts ...AgentRunOption) *AsyncIterator[*AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

func handleFlowAgentEvents(ctx context.Context, iter *AsyncIterator[*AgentEvent],
	generator *AsyncGenerator[*AgentEvent], isolatedSession, parentSession *runSession, toAgentNames []string) {
	_ = "STUB: not implemented"
	return
}

func sendTransferEvents(generator *AsyncGenerator[*AgentEvent], toAgentNames []string) {
	_ = "STUB: not implemented"
	return
}
