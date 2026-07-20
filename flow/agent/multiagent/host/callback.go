package host

import (
	"context"

	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/flow/agent"
)

type MultiAgentCallback interface {
	OnHandOff(ctx context.Context, info *HandOffInfo) context.Context
}

type HandOffInfo struct {
	ToAgentName string
	Argument    string
}

func ConvertCallbackHandlers(handlers ...MultiAgentCallback) callbacks.Handler {
	_ = "STUB: not implemented"
	return *new(callbacks.Handler)
}

func convertCallbacks(opts ...agent.AgentOption) callbacks.Handler {
	_ = "STUB: not implemented"
	return *new(callbacks.Handler)
}
