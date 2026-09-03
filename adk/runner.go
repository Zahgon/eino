package adk

import (
	"context"

	"github.com/cloudwego/eino/internal/core"
	"github.com/cloudwego/eino/schema"
)

func errorIterator[M MessageType](err error) *AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}

func newUserMessage[M MessageType](query string) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

type TypedRunner[M MessageType] struct {
	a               TypedAgent[M]
	enableStreaming bool
	store           CheckPointStore
}

type Runner = TypedRunner[*schema.Message]

type CheckPointStore = core.CheckPointStore

type CheckPointDeleter = core.CheckPointDeleter

type TypedRunnerConfig[M MessageType] struct {
	Agent           TypedAgent[M]
	EnableStreaming bool

	CheckPointStore CheckPointStore
}

type RunnerConfig = TypedRunnerConfig[*schema.Message]

type ResumeParams struct {
	Targets map[string]any
}

func NewRunner(_ context.Context, conf RunnerConfig) *Runner { _ = "STUB: not implemented"; return nil }

func NewTypedRunner[M MessageType](conf TypedRunnerConfig[M]) *TypedRunner[M] {
	_ = "STUB: not implemented"
	return nil
}

func (r *TypedRunner[M]) Run(ctx context.Context, messages []M,
	opts ...AgentRunOption) *AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}

func (r *TypedRunner[M]) Query(ctx context.Context,
	query string, opts ...AgentRunOption) *AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}

func (r *TypedRunner[M]) Resume(ctx context.Context, checkPointID string, opts ...AgentRunOption) (
	*AsyncIterator[*TypedAgentEvent[M]], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TypedRunner[M]) ResumeWithParams(ctx context.Context, checkPointID string, params *ResumeParams, opts ...AgentRunOption) (*AsyncIterator[*TypedAgentEvent[M]], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TypedRunner[M]) resumeInternal(ctx context.Context, checkPointID string, resumeData map[string]any,
	opts ...AgentRunOption) (*AsyncIterator[*TypedAgentEvent[M]], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func typedRunnerRunImpl[M MessageType](a TypedAgent[M], enableStreaming bool, store CheckPointStore, ctx context.Context, messages []M, opts ...AgentRunOption) *AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}

func typedRunnerResumeInternalImpl[M MessageType](a TypedAgent[M], store CheckPointStore, ctx context.Context, checkPointID string, resumeData map[string]any, //nolint:revive // argument-limit
	opts ...AgentRunOption) (*AsyncIterator[*TypedAgentEvent[M]], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func typedRunnerHandleIterImpl[M MessageType](enableStreaming bool, store CheckPointStore, ctx context.Context, aIter *AsyncIterator[*TypedAgentEvent[M]], //nolint:revive // argument-limit
	gen *AsyncGenerator[*TypedAgentEvent[M]], checkPointID *string, cancelCtx *cancelContext) {
	_ = "STUB: not implemented"
	return
}
