package agent

import "github.com/cloudwego/eino/compose"

type AgentOption struct {
	implSpecificOptFn any
	composeOptions    []compose.Option
}

func GetComposeOptions(opts ...AgentOption) []compose.Option { _ = "STUB: not implemented"; return nil }

func WithComposeOptions(opts ...compose.Option) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

func WrapImplSpecificOptFn[T any](optFn func(*T)) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

func GetImplSpecificOptions[T any](base *T, opts ...AgentOption) *T {
	_ = "STUB: not implemented"
	return nil
}
