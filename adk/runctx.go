package adk

import (
	"context"
	"sync"
)

type runSession struct {
	Values    map[string]any
	valuesMtx *sync.Mutex

	Events     []*agentEventWrapper
	LaneEvents *laneEvents
	mtx        sync.Mutex

	TypedEvents any
}

type laneEvents struct {
	Events []*agentEventWrapper
	Parent *laneEvents
}

type agentEventWrapper struct {
	*AgentEvent
	mu                  sync.Mutex
	concatenatedMessage Message

	TS int64

	StreamErr error
}

type typedAgentEventWrapper[M MessageType] struct {
	event               *TypedAgentEvent[M]
	mu                  sync.Mutex
	concatenatedMessage M
	TS                  int64
	StreamErr           error
}

type typedAgentEventWrapperForGob[M MessageType] struct {
	Event *TypedAgentEvent[M]
	TS    int64
}

func (e *typedAgentEventWrapper[M]) GobEncode() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *typedAgentEventWrapper[M]) GobDecode(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *typedAgentEventWrapper[M]) consumeStream() { _ = "STUB: not implemented"; return }

type otherAgentEventWrapperForEncode agentEventWrapper

func (a *agentEventWrapper) GobEncode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *agentEventWrapper) GobDecode(b []byte) error { _ = "STUB: not implemented"; return nil }

func newRunSession() *runSession { _ = "STUB: not implemented"; return nil }

func GetSessionValues(ctx context.Context) map[string]any { _ = "STUB: not implemented"; return nil }

func AddSessionValue(ctx context.Context, key string, value any) { _ = "STUB: not implemented"; return }

func AddSessionValues(ctx context.Context, kvs map[string]any) { _ = "STUB: not implemented"; return }

func GetSessionValue(ctx context.Context, key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (rs *runSession) addEvent(event *AgentEvent) { _ = "STUB: not implemented"; return }

func (rs *runSession) getEvents() []*agentEventWrapper { _ = "STUB: not implemented"; return nil }

func addTypedEvent[M MessageType](session *runSession, event *TypedAgentEvent[M]) {
	_ = "STUB: not implemented"
	return
}

func (rs *runSession) getValues() map[string]any { _ = "STUB: not implemented"; return nil }

func (rs *runSession) addValue(key string, value any) { _ = "STUB: not implemented"; return }

func (rs *runSession) addValues(kvs map[string]any) { _ = "STUB: not implemented"; return }

func (rs *runSession) getValue(key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

type runContext struct {
	RootInput *AgentInput
	RunPath   []RunStep

	AgenticRootInput any

	Session *runSession
}

func (rc *runContext) isRoot() bool { _ = "STUB: not implemented"; return false }

func (rc *runContext) deepCopy() *runContext { _ = "STUB: not implemented"; return nil }

type runCtxKey struct{}

func getRunCtx(ctx context.Context) *runContext { _ = "STUB: not implemented"; return nil }

func setRunCtx(ctx context.Context, runCtx *runContext) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func initRunCtx(ctx context.Context, agentName string, input *AgentInput) (context.Context, *runContext) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func initTypedRunCtx[M MessageType](ctx context.Context, agentName string, input *TypedAgentInput[M]) (context.Context, *runContext) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func joinRunCtxs(parentCtx context.Context, childCtxs ...context.Context) {
	_ = "STUB: not implemented"
	return
}

func commitEvents(ctx context.Context, newEvents []*agentEventWrapper) {
	_ = "STUB: not implemented"
	return
}

func unwindLaneEvents(ctxs ...context.Context) []*agentEventWrapper {
	_ = "STUB: not implemented"
	return nil
}

func forkRunCtx(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func updateRunPathOnly(ctx context.Context, agentNames ...string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ClearRunCtx(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ctxWithNewTypedRunCtx[M MessageType](ctx context.Context, input *TypedAgentInput[M], sharedParentSession bool) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getSession(ctx context.Context) *runSession { _ = "STUB: not implemented"; return nil }
