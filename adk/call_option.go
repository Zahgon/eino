package adk

import "github.com/cloudwego/eino/callbacks"

type options struct {
	sharedParentSession  bool
	sessionValues        map[string]any
	checkPointID         *string
	skipTransferMessages bool
	handlers             []callbacks.Handler
	cancelCtx            *cancelContext
}

type AgentRunOption struct {
	implSpecificOptFn any

	agentNames []string
}

func (o AgentRunOption) DesignateAgent(name ...string) AgentRunOption {
	_ = "STUB: not implemented"
	return *new(AgentRunOption)
}

func getCommonOptions(base *options, opts ...AgentRunOption) *options {
	_ = "STUB: not implemented"
	return nil
}

func WithSessionValues(v map[string]any) AgentRunOption {
	_ = "STUB: not implemented"
	return *new(AgentRunOption)
}

func WithSkipTransferMessages() AgentRunOption {
	_ = "STUB: not implemented"
	return *new(AgentRunOption)
}

func withSharedParentSession() AgentRunOption {
	_ = "STUB: not implemented"
	return *new(AgentRunOption)
}

func WithCallbacks(handlers ...callbacks.Handler) AgentRunOption {
	_ = "STUB: not implemented"
	return *new(AgentRunOption)
}

func WrapImplSpecificOptFn[T any](optFn func(*T)) AgentRunOption {
	_ = "STUB: not implemented"
	return *new(AgentRunOption)
}

func GetImplSpecificOptions[T any](base *T, opts ...AgentRunOption) *T {
	_ = "STUB: not implemented"
	return nil
}

func filterCallbackHandlersForNestedAgents(currentAgentName string, opts []AgentRunOption) []AgentRunOption {
	_ = "STUB: not implemented"
	return nil
}

func filterCancelOption(opts []AgentRunOption) []AgentRunOption {
	_ = "STUB: not implemented"
	return nil
}

func filterOptions(agentName string, opts []AgentRunOption) []AgentRunOption {
	_ = "STUB: not implemented"
	return nil
}
