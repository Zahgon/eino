package parser

type Options struct {
	URI string

	ExtraMeta map[string]any
}

type Option struct {
	apply func(opts *Options)

	implSpecificOptFn any
}

func WithURI(uri string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithExtraMeta(meta map[string]any) Option { _ = "STUB: not implemented"; return *new(Option) }

func GetCommonOptions(base *Options, opts ...Option) *Options {
	_ = "STUB: not implemented"
	return nil
}

func WrapImplSpecificOptFn[T any](optFn func(*T)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func GetImplSpecificOptions[T any](base *T, opts ...Option) *T {
	_ = "STUB: not implemented"
	return nil
}
