package compose

func dagChannelBuilder(controlDependencies []string, dataDependencies []string, zeroValue func() any, emptyStream func() streamReader) channel {
	_ = "STUB: not implemented"
	return *new(channel)
}

type dependencyState uint8

const (
	dependencyStateWaiting dependencyState = iota
	dependencyStateReady
	dependencyStateSkipped
)

type dagChannel struct {
	zeroValue   func() any
	emptyStream func() streamReader

	ControlPredecessors map[string]dependencyState
	Values              map[string]any
	DataPredecessors    map[string]bool
	Skipped             bool

	mergeConfig FanInMergeConfig
}

func (ch *dagChannel) setMergeConfig(cfg FanInMergeConfig) { _ = "STUB: not implemented"; return }

func (ch *dagChannel) load(c channel) error { _ = "STUB: not implemented"; return nil }

func (ch *dagChannel) reportValues(ins map[string]any) error { _ = "STUB: not implemented"; return nil }

func (ch *dagChannel) reportDependencies(dependencies []string) { _ = "STUB: not implemented"; return }

func (ch *dagChannel) reportSkip(keys []string) bool { _ = "STUB: not implemented"; return false }

func (ch *dagChannel) get(isStream bool, name string, edgeHandler *edgeHandlerManager) (
	any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

func (ch *dagChannel) convertValues(fn func(map[string]any) error) error {
	_ = "STUB: not implemented"
	return nil
}
