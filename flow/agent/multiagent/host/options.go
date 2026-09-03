package host

import "github.com/cloudwego/eino/flow/agent"

type options struct {
	agentCallbacks []MultiAgentCallback
}

func WithAgentCallbacks(agentCallbacks ...MultiAgentCallback) agent.AgentOption {
	_ = "STUB: not implemented"
	return *new(agent.AgentOption)
}
