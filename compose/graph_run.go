package compose

import (
	"context"
	"reflect"

	"github.com/cloudwego/eino/internal/core"
)

type chanCall struct {
	action          *composableRunnable
	writeTo         []string
	writeToBranches []*GraphBranch

	controls []string

	preProcessor, postProcessor *composableRunnable
}

type chanBuilder func(dependencies []string, indirectDependencies []string, zeroValue func() any, emptyStream func() streamReader) channel

type runner struct {
	chanSubscribeTo map[string]*chanCall

	successors          map[string][]string
	dataPredecessors    map[string][]string
	controlPredecessors map[string][]string

	inputChannels *chanCall

	chanBuilder chanBuilder
	eager       bool
	dag         bool

	runCtx func(ctx context.Context) context.Context

	options graphCompileOptions

	inputType  reflect.Type
	outputType reflect.Type

	inputStreamFilter          streamMapFilter
	inputConverter             handlerPair
	inputFieldMappingConverter handlerPair

	*genericHelper

	runtimeCheckEdges    map[string]map[string]bool
	runtimeCheckBranches map[string][]bool

	edgeHandlerManager      *edgeHandlerManager
	preNodeHandlerManager   *preNodeHandlerManager
	preBranchHandlerManager *preBranchHandlerManager

	checkPointer         *checkPointer
	interruptBeforeNodes []string
	interruptAfterNodes  []string

	mergeConfigs map[string]FanInMergeConfig
}

func (r *runner) invoke(ctx context.Context, input any, opts ...Option) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (r *runner) transform(ctx context.Context, input streamReader, opts ...Option) (streamReader, error) {
	_ = "STUB: not implemented"
	return *new(streamReader), nil
}

type runnableCallWrapper func(context.Context, *composableRunnable, any, ...any) (any, error)

func runnableInvoke(ctx context.Context, r *composableRunnable, input any, opts ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func runnableTransform(ctx context.Context, r *composableRunnable, input any, opts ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (r *runner) run(ctx context.Context, isStream bool, input any, opts ...Option) (result any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (r *runner) resolveMaxSteps(maxSteps int, opts []Option) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *runner) restoreCheckPointState(
	ctx context.Context,
	path NodePath,
	sm StateModifier,
	cp *checkpoint,
	isStream bool,
	cm *channelManager,
) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func newInterruptTempInfo() *interruptTempInfo { _ = "STUB: not implemented"; return nil }

type interruptTempInfo struct {
	subGraphInterrupts   map[string]*subGraphInterruptError
	interruptTaskNodes   map[string]struct{}
	interruptRerunNodes  []string
	interruptBeforeNodes []string
	interruptAfterNodes  []string
	interruptRerunExtra  map[string]any

	signals []*core.InterruptSignal
}

func (ti *interruptTempInfo) appendSignal(signal *core.InterruptSignal) {
	_ = "STUB: not implemented"
	return
}

func (ti *interruptTempInfo) collectCanceledInfo(canceled bool, canceledTasks, completedTasks []*task) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) resolveInterruptCompletedTasks(tempInfo *interruptTempInfo, completedTasks []*task) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) interruptOriginNodeKey(completedTask *task, address Address) string {
	_ = "STUB: not implemented"
	return ""
}

func getHitKey(tasks []*task, keys []string) []string { _ = "STUB: not implemented"; return nil }

func (r *runner) handleInterrupt(
	ctx context.Context,
	tempInfo *interruptTempInfo,
	nextTasks []*task,
	channels map[string]channel,
	isStream bool,
	isSubGraph bool,
	checkPointID *string,
	publishSubGraphCheckpoint func(*subGraphInterruptError),
) error {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyState(state any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (r *runner) handleInterruptWithSubGraphAndRerunNodes(
	ctx context.Context,
	tempInfo *interruptTempInfo,
	completeTasks []*task,
	checkPointID *string,
	isSubGraph bool,
	cm *channelManager,
	isStream bool,
	publishSubGraphCheckpoint func(*subGraphInterruptError),
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) calculateNextTasks(ctx context.Context, completedTasks []*task, isStream bool, cm *channelManager, optMap map[string][]any) ([]*task, any, bool, error) {
	_ = "STUB: not implemented"
	return nil, *new(any), false, nil
}

func (r *runner) createTasks(ctx context.Context, nodeMap map[string]any, optMap map[string][]any) ([]*task, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getCheckPointInfo(opts ...Option) (checkPointID *string, writeToCheckPointID *string, stateModifier StateModifier, forceNewRun bool) {
	_ = "STUB: not implemented"
	return nil, nil, *new(StateModifier), false
}

func (r *runner) restoreTasks(
	ctx context.Context,
	inputs map[string]any,
	skipPreHandler map[string]bool,
	rerunNodes []string,
	isStream bool,
	optMap map[string][]any) ([]*task, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *runner) validateCheckpointIntegrity(cp *checkpoint) error {
	_ = "STUB: not implemented"
	return nil
}

func isSubGraphCall(call *chanCall) bool { _ = "STUB: not implemented"; return false }

func (r *runner) resolveCompletedTasks(ctx context.Context, completedTasks []*task, isStream bool, cm *channelManager) (map[string]map[string]any, map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (r *runner) calculateBranch(ctx context.Context, curNodeKey string, startChan *chanCall, input []any, isStream bool, cm *channelManager) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *runner) initTaskManager(runWrapper runnableCallWrapper, cancelVal *graphCancelSignal, opts ...Option) *taskManager {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) initChannelManager(isStream bool) *channelManager {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) toComposableRunnable() *composableRunnable { _ = "STUB: not implemented"; return nil }

func copyItem(item any, n int) []any { _ = "STUB: not implemented"; return nil }

func printTask(ts []*task) string { _ = "STUB: not implemented"; return "" }
