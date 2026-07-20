package adk

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

type workflowAgentMode int

const (
	workflowAgentModeUnknown workflowAgentMode = iota
	workflowAgentModeSequential
	workflowAgentModeLoop
	workflowAgentModeParallel
)

type workflowAgent struct {
	name        string
	description string
	subAgents   []*flowAgent

	mode workflowAgentMode

	maxIterations int
}

func (a *workflowAgent) Name(_ context.Context) string { _ = "STUB: not implemented"; return "" }

func (a *workflowAgent) Description(_ context.Context) string { _ = "STUB: not implemented"; return "" }

func (a *workflowAgent) GetType() string { _ = "STUB: not implemented"; return "" }

func (a *workflowAgent) Run(ctx context.Context, _ *AgentInput, opts ...AgentRunOption) *AsyncIterator[*AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

type sequentialWorkflowState struct {
	InterruptIndex int
}

type parallelWorkflowState struct {
	SubAgentEvents map[int][]*agentEventWrapper
}

type loopWorkflowState struct {
	LoopIterations int
	SubAgentIndex  int
}

func init() {
	schema.RegisterName[*sequentialWorkflowState]("eino_adk_sequential_workflow_state")
	schema.RegisterName[*parallelWorkflowState]("eino_adk_parallel_workflow_state")
	schema.RegisterName[*loopWorkflowState]("eino_adk_loop_workflow_state")
}

func (a *workflowAgent) Resume(ctx context.Context, info *ResumeInfo, opts ...AgentRunOption) *AsyncIterator[*AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

type WorkflowInterruptInfo struct {
	OrigInput *AgentInput

	SequentialInterruptIndex int
	SequentialInterruptInfo  *InterruptInfo

	LoopIterations int

	ParallelInterruptInfo map[int]*InterruptInfo
}

func (a *workflowAgent) runSequential(ctx context.Context,
	generator *AsyncGenerator[*AgentEvent], seqState *sequentialWorkflowState, info *ResumeInfo,
	opts ...AgentRunOption) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type BreakLoopAction struct {
	From string

	Done bool

	CurrentIterations int
}

func NewBreakLoopAction(agentName string) *AgentAction { _ = "STUB: not implemented"; return nil }

func (a *workflowAgent) runLoop(ctx context.Context, generator *AsyncGenerator[*AgentEvent],
	loopState *loopWorkflowState, resumeInfo *ResumeInfo, opts ...AgentRunOption) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (a *workflowAgent) runParallel(ctx context.Context, generator *AsyncGenerator[*AgentEvent],
	parState *parallelWorkflowState, resumeInfo *ResumeInfo, opts ...AgentRunOption) error {
	_ = "STUB: not implemented"
	return nil
}

func cancelAtTransition(ctx context.Context, info string, state any) *AgentEvent {
	_ = "STUB: not implemented"
	return nil
}

type SequentialAgentConfig struct {
	Name        string
	Description string
	SubAgents   []Agent
}

type ParallelAgentConfig struct {
	Name        string
	Description string
	SubAgents   []Agent
}

type LoopAgentConfig struct {
	Name        string
	Description string
	SubAgents   []Agent

	MaxIterations int
}

func newWorkflowAgent(ctx context.Context, name, desc string,
	subAgents []Agent, mode workflowAgentMode, maxIterations int) (*flowAgent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSequentialAgent(ctx context.Context, config *SequentialAgentConfig) (ResumableAgent, error) {
	_ = "STUB: not implemented"
	return *new(ResumableAgent), nil
}

func NewParallelAgent(ctx context.Context, config *ParallelAgentConfig) (ResumableAgent, error) {
	_ = "STUB: not implemented"
	return *new(ResumableAgent), nil
}

func NewLoopAgent(ctx context.Context, config *LoopAgentConfig) (ResumableAgent, error) {
	_ = "STUB: not implemented"
	return *new(ResumableAgent), nil
}
