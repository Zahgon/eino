package compose

import (
	"context"
	"time"

	"github.com/cloudwego/eino/internal"
)

type channel interface {
	reportValues(map[string]any) error
	reportDependencies([]string)
	reportSkip([]string) bool
	get(bool, string, *edgeHandlerManager) (any, bool, error)
	convertValues(fn func(map[string]any) error) error
	load(channel) error

	setMergeConfig(FanInMergeConfig)
}

type edgeHandlerManager struct {
	h map[string]map[string][]handlerPair
}

func (e *edgeHandlerManager) handle(from, to string, value any, isStream bool) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type preNodeHandlerManager struct {
	h map[string][]handlerPair
}

func (p *preNodeHandlerManager) handle(nodeKey string, value any, isStream bool) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type preBranchHandlerManager struct {
	h map[string][][]handlerPair
}

func (p *preBranchHandlerManager) handle(nodeKey string, idx int, value any, isStream bool) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type channelManager struct {
	isStream bool
	channels map[string]channel

	successors          map[string][]string
	dataPredecessors    map[string]map[string]struct{}
	controlPredecessors map[string]map[string]struct{}

	edgeHandlerManager    *edgeHandlerManager
	preNodeHandlerManager *preNodeHandlerManager
}

func (c *channelManager) loadChannels(channels map[string]channel) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *channelManager) updateValues(_ context.Context, values map[string]map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *channelManager) updateDependencies(_ context.Context, dependenciesMap map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *channelManager) getFromReadyChannels(_ context.Context) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *channelManager) updateAndGet(ctx context.Context, values map[string]map[string]any, dependencies map[string][]string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *channelManager) reportBranch(from string, skippedNodes []string) error {
	_ = "STUB: not implemented"
	return nil
}

func appendIfNotExist(s []string, elem string) []string { _ = "STUB: not implemented"; return nil }

type task struct {
	ctx            context.Context
	nodeKey        string
	call           *chanCall
	input          any
	originalInput  any
	output         any
	option         []any
	err            error
	skipPreHandler bool
}

type taskManager struct {
	runWrapper runnableCallWrapper
	opts       []Option
	needAll    bool

	num          uint32
	done         *internal.UnboundedChan[*task]
	runningTasks map[string]*task

	cancelCh chan *time.Duration
	canceled bool
	deadline *time.Time

	persistRerunInput bool
}

func (t *taskManager) execute(currentTask *task) { _ = "STUB: not implemented"; return }

func (t *taskManager) submit(tasks []*task) error { _ = "STUB: not implemented"; return nil }

func (t *taskManager) wait() (tasks []*task, canceled bool, canceledTasks []*task) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (t *taskManager) waitOne() (ta *task, success bool, canceled bool) {
	_ = "STUB: not implemented"
	return nil, false, false
}

func (t *taskManager) waitAll() (successTasks []*task, canceledTasks []*task) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *taskManager) receive(recv func() (*task, bool)) (ta *task, closed bool, canceled bool) {
	_ = "STUB: not implemented"
	return nil, false, false
}

func receiveWithDeadline(recv func() (*task, bool), deadline time.Time) (ta *task, closed bool, canceled bool) {
	_ = "STUB: not implemented"
	return nil, false, false
}

func receiveWithListening(recv func() (*task, bool), cancel chan *time.Duration) (*task, bool, bool, bool, *time.Time) {
	_ = "STUB: not implemented"
	return nil, false, false, false, nil
}

func runPreHandler(ta *task, runWrapper runnableCallWrapper) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func runPostHandler(ta *task, runWrapper runnableCallWrapper) { _ = "STUB: not implemented"; return }
