package adk

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func init() {
	schema.RegisterName[*CancelError]("_eino_adk_cancel_error")
	schema.RegisterName[*AgentCancelInfo]("_eino_adk_agent_cancel_info")
	schema.RegisterName[*StreamCanceledError]("_eino_adk_stream_cancelled_error")
}

type CancelMode int

const (
	CancelImmediate CancelMode = 0

	CancelAfterChatModel CancelMode = 1 << iota

	CancelAfterToolCalls
)

type CancelHandle struct {
	wait func() error
}

func (h *CancelHandle) Wait() error { _ = "STUB: not implemented"; return nil }

type AgentCancelFunc func(...AgentCancelOption) (*CancelHandle, bool)

type agentCancelConfig struct {
	Mode      CancelMode
	Recursive bool
	Timeout   *time.Duration
}

type AgentCancelOption func(*agentCancelConfig)

func WithAgentCancelMode(mode CancelMode) AgentCancelOption {
	_ = "STUB: not implemented"
	return *new(AgentCancelOption)
}

func WithAgentCancelTimeout(timeout time.Duration) AgentCancelOption {
	_ = "STUB: not implemented"
	return *new(AgentCancelOption)
}

func WithRecursive() AgentCancelOption { _ = "STUB: not implemented"; return *new(AgentCancelOption) }

type AgentCancelInfo struct {
	Mode      CancelMode
	Escalated bool
	Timeout   bool
}

type CancelError struct {
	Info *AgentCancelInfo

	InterruptContexts []*InterruptCtx

	interruptSignal *InterruptSignal
}

func (e *CancelError) Error() string { _ = "STUB: not implemented"; return "" }

var (
	ErrCancelTimeout = errors.New("cancel timed out")

	ErrExecutionEnded = errors.New("execution already ended")

	ErrStreamCanceled error = &StreamCanceledError{}
)

type StreamCanceledError struct{}

func (e *StreamCanceledError) Error() string { _ = "STUB: not implemented"; return "" }

func WithCancel() (AgentRunOption, AgentCancelFunc) {
	_ = "STUB: not implemented"
	return *new(AgentRunOption), *new(AgentCancelFunc)
}

const (
	stateRunning int32 = 0

	stateCancelling int32 = 1

	stateDone int32 = 2

	stateCancelHandled int32 = 5
)

const (
	interruptNotSent int32 = 0

	interruptImmediate int32 = 1
)

const defaultCancelImmediateGracePeriod = 1 * time.Second

type cancelContextKey struct{}

func withCancelContext(ctx context.Context, cc *cancelContext) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getCancelContext(ctx context.Context) *cancelContext { _ = "STUB: not implemented"; return nil }

type cancelContext struct {
	mode int32

	cancelChan    chan struct{}
	immediateChan chan struct{}
	doneChan      chan struct{}
	doneOnce      sync.Once

	state            int32
	interruptSent    int32
	escalated        int32
	timeoutEscalated int32
	startedMode      int32
	deadlineUnixNano int64

	recursive     int32
	recursiveChan chan struct{}

	root      bool
	abortOnly bool
	parent    *cancelContext

	checkpointAwareDescendant int32

	cancelMu      sync.Mutex
	timeoutOnce   sync.Once
	timeoutNotify chan struct{}

	mu                  sync.Mutex
	graphInterruptFuncs []func(...compose.GraphInterruptOption)
}

func newCancelContext() *cancelContext { _ = "STUB: not implemented"; return nil }

func (cc *cancelContext) isRoot() bool { _ = "STUB: not implemented"; return false }

func (cc *cancelContext) isRecursive() bool { _ = "STUB: not implemented"; return false }

func (cc *cancelContext) setRecursive(v bool) { _ = "STUB: not implemented"; return }

func (cc *cancelContext) deriveCheckpointAwareCancelContext(ctx context.Context) *cancelContext {
	_ = "STUB: not implemented"
	return nil
}

func deriveAbortOnlyCancelContext(ctx context.Context, parent *cancelContext) *cancelContext {
	_ = "STUB: not implemented"
	return nil
}

func withAbortOnlyCancelContext(ctx context.Context, cc *cancelContext) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func (cc *cancelContext) triggerCancel(mode CancelMode) { _ = "STUB: not implemented"; return }

func (cc *cancelContext) triggerImmediateCancel() { _ = "STUB: not implemented"; return }

func (cc *cancelContext) getMode() CancelMode { _ = "STUB: not implemented"; return *new(CancelMode) }

func (cc *cancelContext) setMode(mode CancelMode) { _ = "STUB: not implemented"; return }

func (cc *cancelContext) getDeadlineUnixNano() int64 { _ = "STUB: not implemented"; return 0 }

func (cc *cancelContext) setDeadlineUnixNano(v int64) { _ = "STUB: not implemented"; return }

func (cc *cancelContext) wakeTimeoutController() { _ = "STUB: not implemented"; return }

func (cc *cancelContext) shouldCancel() bool { _ = "STUB: not implemented"; return false }

func (cc *cancelContext) isImmediateCancelled() bool { _ = "STUB: not implemented"; return false }

func (cc *cancelContext) sendImmediateInterrupt() bool { _ = "STUB: not implemented"; return false }

func (cc *cancelContext) setGraphInterruptFunc(interrupt func(...compose.GraphInterruptOption)) {
	_ = "STUB: not implemented"
	return
}

func (cc *cancelContext) markDone() { _ = "STUB: not implemented"; return }

func (cc *cancelContext) hasCheckpointAwareDescendant() bool {
	_ = "STUB: not implemented"
	return false
}

func (cc *cancelContext) markCheckpointAwareDescendant() { _ = "STUB: not implemented"; return }

func deriveCheckpointAwareSubAgentCancelContext(ctx context.Context, opts []AgentRunOption) *cancelContext {
	_ = "STUB: not implemented"
	return nil
}

func appendCancelContextOption(opts []AgentRunOption, cancelCtx *cancelContext) []AgentRunOption {
	_ = "STUB: not implemented"
	return nil
}

func (cc *cancelContext) markCancelHandled() bool { _ = "STUB: not implemented"; return false }

func (cc *cancelContext) createCancelError() *CancelError { _ = "STUB: not implemented"; return nil }

func (cc *cancelContext) createAndMarkCancelHandled() (*CancelError, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func isAbortOnlyCancelled(cc *cancelContext) bool { _ = "STUB: not implemented"; return false }

func checkPreExecCancel[M MessageType](cc *cancelContext, gen *AsyncGenerator[*TypedAgentEvent[M]]) (terminated bool) {
	_ = "STUB: not implemented"
	return false
}

func (cc *cancelContext) buildCancelFunc() AgentCancelFunc {
	_ = "STUB: not implemented"
	return *new(AgentCancelFunc)
}

func wrapIterWithCancelCtx[M MessageType](iter *AsyncIterator[*TypedAgentEvent[M]], cancelCtx *cancelContext) *AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}

type typedCancelMonitoredModel[M MessageType] struct {
	inner         model.BaseModel[M]
	cancelContext *cancelContext
}

type recvResult[T any] struct {
	data T
	err  error
}

func (m *typedCancelMonitoredModel[M]) Generate(ctx context.Context, input []M, opts ...model.Option) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func (m *typedCancelMonitoredModel[M]) Stream(ctx context.Context, input []M, opts ...model.Option) (*schema.StreamReader[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func wrapStreamWithCancelMonitoring[T any](stream *schema.StreamReader[T], cc *cancelContext) *schema.StreamReader[T] {
	_ = "STUB: not implemented"
	return nil
}

type cancelMonitoredToolHandler struct{}

func (h *cancelMonitoredToolHandler) WrapStreamableToolCall(next compose.StreamableToolEndpoint) compose.StreamableToolEndpoint {
	_ = "STUB: not implemented"
	return *new(compose.StreamableToolEndpoint)
}

func (h *cancelMonitoredToolHandler) WrapEnhancedStreamableToolCall(next compose.EnhancedStreamableToolEndpoint) compose.EnhancedStreamableToolEndpoint {
	_ = "STUB: not implemented"
	return *new(compose.EnhancedStreamableToolEndpoint)
}
