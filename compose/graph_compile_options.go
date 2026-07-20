package compose

type graphCompileOptions struct {
	maxRunSteps     int
	graphName       string
	nodeTriggerMode NodeTriggerMode

	callbacks []GraphCompileCallback

	origOpts []GraphCompileOption

	checkPointStore      CheckPointStore
	serializer           Serializer
	interruptBeforeNodes []string
	interruptAfterNodes  []string

	eagerDisabled bool

	mergeConfigs map[string]FanInMergeConfig
}

func newGraphCompileOptions(opts ...GraphCompileOption) *graphCompileOptions {
	_ = "STUB: not implemented"
	return nil
}

type GraphCompileOption func(*graphCompileOptions)

func WithMaxRunSteps(maxSteps int) GraphCompileOption {
	_ = "STUB: not implemented"
	return *new(GraphCompileOption)
}

func WithGraphName(graphName string) GraphCompileOption {
	_ = "STUB: not implemented"
	return *new(GraphCompileOption)
}

func WithEagerExecution() GraphCompileOption {
	_ = "STUB: not implemented"
	return *new(GraphCompileOption)
}

func WithEagerExecutionDisabled() GraphCompileOption {
	_ = "STUB: not implemented"
	return *new(GraphCompileOption)
}

func WithNodeTriggerMode(triggerMode NodeTriggerMode) GraphCompileOption {
	_ = "STUB: not implemented"
	return *new(GraphCompileOption)
}

func WithGraphCompileCallbacks(cbs ...GraphCompileCallback) GraphCompileOption {
	_ = "STUB: not implemented"
	return *new(GraphCompileOption)
}

type FanInMergeConfig struct {
	StreamMergeWithSourceEOF bool
}

func WithFanInMergeConfig(confs map[string]FanInMergeConfig) GraphCompileOption {
	_ = "STUB: not implemented"
	return *new(GraphCompileOption)
}

func InitGraphCompileCallbacks(cbs []GraphCompileCallback) { _ = "STUB: not implemented"; return }

var globalGraphCompileCallbacks []GraphCompileCallback
