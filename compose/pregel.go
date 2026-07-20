package compose

func pregelChannelBuilder(_ []string, _ []string, _ func() any, _ func() streamReader) channel {
	_ = "STUB: not implemented"
	return *new(channel)
}

type pregelChannel struct {
	Values map[string]any

	mergeConfig FanInMergeConfig
}

func (ch *pregelChannel) setMergeConfig(cfg FanInMergeConfig) { _ = "STUB: not implemented"; return }

func (ch *pregelChannel) load(c channel) error { _ = "STUB: not implemented"; return nil }

func (ch *pregelChannel) convertValues(fn func(map[string]any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *pregelChannel) reportValues(ins map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *pregelChannel) get(isStream bool, name string, edgeHandler *edgeHandlerManager) (
	any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

func (ch *pregelChannel) reportSkip(_ []string) bool { _ = "STUB: not implemented"; return false }

func (ch *pregelChannel) reportDependencies(_ []string) { _ = "STUB: not implemented"; return }
