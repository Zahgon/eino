package adk

import (
	"context"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

type HistoryEntry struct {
	IsUserInput bool
	AgentName   string
	Message     Message
}

type HistoryRewriter func(ctx context.Context, entries []*HistoryEntry) ([]Message, error)

type flowAgent struct {
	Agent

	subAgents   []*flowAgent
	parentAgent *flowAgent

	disallowTransferToParent bool
	historyRewriter          HistoryRewriter

	checkPointStore compose.CheckPointStore
}

func (a *flowAgent) deepCopy() *flowAgent { _ = "STUB: not implemented"; return nil }

func SetSubAgents(ctx context.Context, agent Agent, subAgents []Agent) (ResumableAgent, error) {
	_ = "STUB: not implemented"
	return *new(ResumableAgent), nil
}

type AgentOption func(options *flowAgent)

func WithDisallowTransferToParent() AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

func WithHistoryRewriter(h HistoryRewriter) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

func toFlowAgent(ctx context.Context, agent Agent, opts ...AgentOption) *flowAgent {
	_ = "STUB: not implemented"
	return nil
}

func AgentWithOptions(ctx context.Context, agent Agent, opts ...AgentOption) Agent {
	_ = "STUB: not implemented"
	return *new(Agent)
}

func setSubAgents(ctx context.Context, agent Agent, subAgents []Agent) (*flowAgent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *flowAgent) getAgent(ctx context.Context, name string) *flowAgent {
	_ = "STUB: not implemented"
	return nil
}

func rewriteMessage(msg Message, agentName string) Message {
	_ = "STUB: not implemented"
	return *new(Message)
}

//nolint:staticcheck // backward compat with deprecated MultiContent field

func genMsg(entry *HistoryEntry, agentName string) (Message, error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}

func deepCopyAgentInput(ai *AgentInput) *AgentInput { _ = "STUB: not implemented"; return nil }

func (a *flowAgent) genAgentInput(ctx context.Context, runCtx *runContext, skipTransferMessages bool) (*AgentInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildDefaultHistoryRewriter(agentName string) HistoryRewriter {
	_ = "STUB: not implemented"
	return *new(HistoryRewriter)
}

func (a *flowAgent) Run(ctx context.Context, input *AgentInput, opts ...AgentRunOption) *AsyncIterator[*AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

func (a *flowAgent) Resume(ctx context.Context, info *ResumeInfo, opts ...AgentRunOption) *AsyncIterator[*AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

type DeterministicTransferConfig struct {
	Agent        Agent
	ToAgentNames []string
}

func (a *flowAgent) run(
	ctx context.Context,
	ctxForSubAgents context.Context,
	runCtx *runContext,
	aIter *AsyncIterator[*AgentEvent],
	generator *AsyncGenerator[*AgentEvent],
	opts ...AgentRunOption) {
	_ = "STUB: not implemented"
	return
}

func exactRunPathMatch(aPath, bPath []RunStep) bool { _ = "STUB: not implemented"; return false }

func wrapIterWithOnEnd(ctx context.Context, iter *AsyncIterator[*AgentEvent]) *AsyncIterator[*AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

type typedFlowAgent[M MessageType] struct {
	TypedAgent[M]

	checkPointStore compose.CheckPointStore
}

func toTypedFlowAgent[M MessageType](agent TypedAgent[M]) *typedFlowAgent[M] {
	_ = "STUB: not implemented"
	return nil
}

func getTypedAgentType[M MessageType](agent TypedAgent[M]) string {
	_ = "STUB: not implemented"
	return ""
}

func (a *typedFlowAgent[M]) Run(ctx context.Context, input *TypedAgentInput[M], opts ...AgentRunOption) *AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}

func (a *typedFlowAgent[M]) Resume(ctx context.Context, info *ResumeInfo, opts ...AgentRunOption) *AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}

func (a *typedFlowAgent[M]) run(
	ctx context.Context,
	_ context.Context,
	runCtx *runContext,
	aIter *AsyncIterator[*TypedAgentEvent[M]],
	generator *AsyncGenerator[*TypedAgentEvent[M]],
	_ ...AgentRunOption) {
	_ = "STUB: not implemented"
	return
}

func wrapAgenticIterWithOnEnd(ctx context.Context, iter *AsyncIterator[*TypedAgentEvent[*schema.AgenticMessage]]) *AsyncIterator[*TypedAgentEvent[*schema.AgenticMessage]] {
	_ = "STUB: not implemented"
	return nil
}

func genAgenticErrorIter(err error) *AsyncIterator[*TypedAgentEvent[*schema.AgenticMessage]] {
	_ = "STUB: not implemented"
	return nil
}

func typedWrapIterWithOnEnd[M MessageType](ctx context.Context, iter *AsyncIterator[*TypedAgentEvent[M]]) *AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}

func typedErrorIterWithOnEnd[M MessageType](ctx context.Context, err error) *AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}
