package adk

import (
	"context"
	"sync"
	"time"
)

type stopPhase uint8

const (
	stopOpen stopPhase = iota
	stopIdleWaiting
	stopCommitted
)

type preemptTurnPhase uint8

const (
	preemptTurnIdle preemptTurnPhase = iota
	preemptTurnPlanning
	preemptTurnActive
)

func (p preemptTurnPhase) String() string { _ = "STUB: not implemented"; return "" }

type preemptTurnSnapshot struct {
	hasTargetTurn bool
	turnID        uint64
	ctx           context.Context
	tc            any
}

type cancelRequestState struct {
	cfg             agentCancelConfig
	timeoutDeadline *time.Time
}

type preemptRequest struct {
	cancel   cancelRequestState
	ackChans []chan struct{}
}

func parseAgentCancelOptions(opts ...AgentCancelOption) agentCancelConfig {
	_ = "STUB: not implemented"
	return *new(agentCancelConfig)
}

func newCancelRequestState(opts []AgentCancelOption, now time.Time) cancelRequestState {
	_ = "STUB: not implemented"
	return *new(cancelRequestState)
}

func (s *cancelRequestState) merge(opts []AgentCancelOption, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (s cancelRequestState) cancelOptions(now time.Time) []AgentCancelOption {
	_ = "STUB: not implemented"
	return nil
}

func newPreemptRequest(ack chan struct{}, opts []AgentCancelOption, now time.Time) *preemptRequest {
	_ = "STUB: not implemented"
	return nil
}

func (r *preemptRequest) ack() { _ = "STUB: not implemented"; return }

func (r *preemptRequest) merge(ack chan struct{}, opts []AgentCancelOption, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (r *preemptRequest) cancelOptions(now time.Time) []AgentCancelOption {
	_ = "STUB: not implemented"
	return nil
}

type preemptController struct {
	mu   sync.Mutex
	cond *sync.Cond

	turnPhase     preemptTurnPhase
	turnID        uint64
	currentTC     any
	currentRunCtx context.Context

	pushInFlight int
	pending      *preemptRequest
	notify       chan struct{}
	closed       bool
}

func newPreemptController() *preemptController { _ = "STUB: not implemented"; return nil }

func (c *preemptController) beginPlanningTurn() { _ = "STUB: not implemented"; return }

func (c *preemptController) abortPlanningTurn() *preemptRequest {
	_ = "STUB: not implemented"
	return nil
}

func (c *preemptController) beginActiveTurn(ctx context.Context, tc any) {
	_ = "STUB: not implemented"
	return
}

func (c *preemptController) endActiveTurn() *preemptRequest { _ = "STUB: not implemented"; return nil }

func (c *preemptController) requirePhaseLocked(expected preemptTurnPhase, op string) {
	_ = "STUB: not implemented"
	return
}

func (c *preemptController) requireNoPendingLocked(op string) { _ = "STUB: not implemented"; return }

func (c *preemptController) beginPush() preemptTurnSnapshot {
	_ = "STUB: not implemented"
	return *new(preemptTurnSnapshot)
}

func (c *preemptController) endPush() { _ = "STUB: not implemented"; return }

func (c *preemptController) waitForPushes() { _ = "STUB: not implemented"; return }

func (c *preemptController) requestPreempt(target preemptTurnSnapshot, ack chan struct{}, opts ...AgentCancelOption) {
	_ = "STUB: not implemented"
	return
}

func (c *preemptController) receivePreempt() (*preemptRequest, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *preemptController) closeForLoopExit() { _ = "STUB: not implemented"; return }

func (c *preemptController) notifyWatcherLocked() { _ = "STUB: not implemented"; return }

type stopDecision struct {
	commit   bool
	wakeIdle bool
}

type stopCancelRequest struct {
	cancel cancelRequestState
}

func newStopCancelRequest(opts []AgentCancelOption, now time.Time) *stopCancelRequest {
	_ = "STUB: not implemented"
	return nil
}

func (r *stopCancelRequest) merge(opts []AgentCancelOption, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (r *stopCancelRequest) cancelOptions(now time.Time) []AgentCancelOption {
	_ = "STUB: not implemented"
	return nil
}

type stopController struct {
	mu sync.Mutex

	phase stopPhase

	hasActiveCancelTarget bool
	pending               *stopCancelRequest
	notify                chan struct{}

	idleFor        time.Duration
	skipCheckpoint bool
	stopCause      string

	closed bool
}

func newStopController() *stopController { _ = "STUB: not implemented"; return nil }

func (c *stopController) requestStop(cfg *stopConfig) stopDecision {
	_ = "STUB: not implemented"
	return *new(stopDecision)
}

func (c *stopController) commit() bool { _ = "STUB: not implemented"; return false }

func (c *stopController) commitLocked() bool { _ = "STUB: not implemented"; return false }

func (c *stopController) isCommitted() bool { _ = "STUB: not implemented"; return false }

func (c *stopController) idleDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *stopController) skipCheckpointEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *stopController) cause() string { _ = "STUB: not implemented"; return "" }

func (c *stopController) beginActiveTurn() { _ = "STUB: not implemented"; return }

func (c *stopController) endActiveTurn() *stopCancelRequest { _ = "STUB: not implemented"; return nil }

func (c *stopController) receiveCancel() (*stopCancelRequest, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *stopController) closeForLoopExit() { _ = "STUB: not implemented"; return }

func (c *stopController) notifyWatcherLocked() { _ = "STUB: not implemented"; return }

type TurnLoopConfig[T any, M MessageType] struct {
	GenInput func(ctx context.Context, loop *TurnLoop[T, M], items []T) (*GenInputResult[T, M], error)

	GenResume func(ctx context.Context, loop *TurnLoop[T, M], interruptedItems, unhandledItems, newItems []T) (*GenResumeResult[T, M], error)

	PrepareAgent func(ctx context.Context, loop *TurnLoop[T, M], consumed []T) (TypedAgent[M], error)

	OnAgentEvents func(ctx context.Context, tc *TurnContext[T, M], events *AsyncIterator[*TypedAgentEvent[M]]) error

	Store CheckPointStore

	CheckpointID string
}

type GenInputResult[T any, M MessageType] struct {
	RunCtx context.Context

	Input *TypedAgentInput[M]

	RunOpts []AgentRunOption

	Consumed []T

	Remaining []T
}

type GenResumeResult[T any, M MessageType] struct {
	RunCtx context.Context

	RunOpts []AgentRunOption

	ResumeParams *ResumeParams

	Consumed []T

	Remaining []T
}

type turnRunSpec[T any, M MessageType] struct {
	runCtx       context.Context
	input        *TypedAgentInput[M]
	runOpts      []AgentRunOption
	resumeParams *ResumeParams
	isResume     bool
	consumed     []T
	resumeBytes  []byte
}

type turnPlan[T any, M MessageType] struct {
	turnCtx   context.Context
	remaining []T
	spec      *turnRunSpec[T, M]
}

func (l *TurnLoop[T, M]) planTurn(
	ctx context.Context,
	isResume bool,
	items []T,
	pr *turnLoopPendingResume[T],
) (*turnPlan[T, M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type InterruptError struct {
	InterruptContexts []*InterruptCtx
}

func (e *InterruptError) Error() string { _ = "STUB: not implemented"; return "" }

type TurnLoopExitState[T any, M MessageType] struct {
	ExitReason error

	UnhandledItems []T

	InterruptedItems []T

	StopCause string

	CheckpointAttempted bool

	CheckpointErr error

	TakeLateItems func() []T
}

type TurnContext[T any, M MessageType] struct {
	Loop *TurnLoop[T, M]

	Consumed []T

	Preempted <-chan struct{}

	Stopped <-chan struct{}

	StopCause func() string
}

type TurnLoop[T any, M MessageType] struct {
	config TurnLoopConfig[T, M]

	buffer *turnBuffer[T]

	stopped int32
	started int32

	done chan struct{}

	result *TurnLoopExitState[T, M]

	runOnce sync.Once

	stopCtrl *stopController

	preemptCtrl *preemptController

	runErr error

	interruptedItems []T

	checkPointRunnerBytes []byte
	interruptContexts     []*InterruptCtx
	capturedCancelErr     *CancelError

	pendingResume *turnLoopPendingResume[T]

	loadCheckpointID string

	onAgentEvents func(ctx context.Context, tc *TurnContext[T, M], events *AsyncIterator[*TypedAgentEvent[M]]) error

	lateMu     sync.Mutex
	lateItems  []T
	lateSealed bool
}

func (l *TurnLoop[T, M]) appendLate(item T) { _ = "STUB: not implemented"; return }

type turnLoopCheckpoint[T any] struct {
	RunnerCheckpoint []byte

	HasRunnerState bool
	UnhandledItems []T
	CanceledItems  []T
}

func marshalTurnLoopCheckpoint[T any](c *turnLoopCheckpoint[T]) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalTurnLoopCheckpoint[T any](data []byte) (*turnLoopCheckpoint[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *TurnLoop[T, M]) saveTurnLoopCheckpoint(ctx context.Context, checkPointID string, c *turnLoopCheckpoint[T]) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *TurnLoop[T, M]) deleteTurnLoopCheckpoint(ctx context.Context, checkPointID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *TurnLoop[T, M]) tryLoadCheckpoint(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type turnLoopPendingResume[T any] struct {
	interrupted []T
	unhandled   []T
	newItems    []T
	resumeBytes []byte
}

type SafePoint int

const (
	AfterChatModel SafePoint = 1 << iota

	AfterToolCalls

	AnySafePoint = AfterChatModel | AfterToolCalls
)

func (sp SafePoint) toCancelMode() CancelMode { _ = "STUB: not implemented"; return *new(CancelMode) }

type stopConfig struct {
	agentCancelOpts []AgentCancelOption
	skipCheckpoint  bool
	stopCause       string
	idleFor         time.Duration
}

type StopOption func(*stopConfig)

func WithGraceful() StopOption { _ = "STUB: not implemented"; return *new(StopOption) }

func WithImmediate() StopOption { _ = "STUB: not implemented"; return *new(StopOption) }

func WithGracefulTimeout(gracePeriod time.Duration) StopOption {
	_ = "STUB: not implemented"
	return *new(StopOption)
}

func WithSkipCheckpoint() StopOption { _ = "STUB: not implemented"; return *new(StopOption) }

func WithStopCause(cause string) StopOption { _ = "STUB: not implemented"; return *new(StopOption) }

func UntilIdleFor(duration time.Duration) StopOption {
	_ = "STUB: not implemented"
	return *new(StopOption)
}

type pushConfig[T any, M MessageType] struct {
	preempt         bool
	preemptDelay    time.Duration
	agentCancelOpts []AgentCancelOption
	pushStrategy    func(context.Context, *TurnContext[T, M]) []PushOption[T, M]
}

type PushOption[T any, M MessageType] func(*pushConfig[T, M])

func WithPreempt[T any, M MessageType](safePoint SafePoint) PushOption[T, M] {
	_ = "STUB: not implemented"
	return nil
}

func WithPreemptTimeout[T any, M MessageType](safePoint SafePoint, timeout time.Duration) PushOption[T, M] {
	_ = "STUB: not implemented"
	return nil
}

func WithPreemptDelay[T any, M MessageType](delay time.Duration) PushOption[T, M] {
	_ = "STUB: not implemented"
	return nil
}

func WithPushStrategy[T any, M MessageType](fn func(ctx context.Context, tc *TurnContext[T, M]) []PushOption[T, M]) PushOption[T, M] {
	_ = "STUB: not implemented"
	return nil
}

func defaultTurnLoopOnAgentEvents[T any, M MessageType](_ context.Context, _ *TurnContext[T, M], events *AsyncIterator[*TypedAgentEvent[M]]) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTurnLoop[T any, M MessageType](cfg TurnLoopConfig[T, M]) *TurnLoop[T, M] {
	_ = "STUB: not implemented"
	return nil
}

func (l *TurnLoop[T, M]) start(ctx context.Context) { _ = "STUB: not implemented"; return }

func (l *TurnLoop[T, M]) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (l *TurnLoop[T, M]) Push(item T, opts ...PushOption[T, M]) (bool, <-chan struct{}) {
	_ = "STUB: not implemented"
	return false, nil
}

func (l *TurnLoop[T, M]) pushWithStrategy(item T, cfg *pushConfig[T, M]) (bool, <-chan struct{}) {
	_ = "STUB: not implemented"
	return false, nil
}

func (l *TurnLoop[T, M]) pushWithConfig(item T, cfg *pushConfig[T, M]) (bool, <-chan struct{}) {
	_ = "STUB: not implemented"
	return false, nil
}

func (l *TurnLoop[T, M]) Stop(opts ...StopOption) { _ = "STUB: not implemented"; return }

func (l *TurnLoop[T, M]) commitStop() { _ = "STUB: not implemented"; return }

func (l *TurnLoop[T, M]) finishStopCommit() { _ = "STUB: not implemented"; return }

func (l *TurnLoop[T, M]) Wait() *TurnLoopExitState[T, M] { _ = "STUB: not implemented"; return nil }

func (l *TurnLoop[T, M]) run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (l *TurnLoop[T, M]) setupBridgeStore(spec *turnRunSpec[T, M], runOpts []AgentRunOption) ([]AgentRunOption, *bridgeStore, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (l *TurnLoop[T, M]) watchPreempt(done <-chan struct{}, agentCancelFunc AgentCancelFunc, preemptDone chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (l *TurnLoop[T, M]) watchStop(done <-chan struct{}, agentCancelFunc AgentCancelFunc, stoppedDone chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (l *TurnLoop[T, M]) runAgentAndHandleEvents(
	ctx context.Context,
	agent TypedAgent[M],
	spec *turnRunSpec[T, M],
) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *TurnLoop[T, M]) applyFrameworkCapturedError(handleErr error) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *TurnLoop[T, M]) cleanup(ctx context.Context) { _ = "STUB: not implemented"; return }
