package compose

func RegisterValuesMergeFunc[T any](fn func([]T) (T, error)) { _ = "STUB: not implemented"; return }

type mergeOptions struct {
	streamMergeWithSourceEOF bool
	names                    []string
}

func mergeValues(vs []any, opts *mergeOptions) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
