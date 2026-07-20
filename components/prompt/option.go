package prompt

type Option struct {
	implSpecificOptFn any
}

func WrapImplSpecificOptFn[T any](optFn func(*T)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func GetImplSpecificOptions[T any](base *T, opts ...Option) *T {
	_ = "STUB: not implemented"
	return nil
}
