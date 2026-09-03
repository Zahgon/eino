package plantask

import (
	"context"

	"github.com/cloudwego/eino/adk"
)

type Config struct {
	Backend Backend
	BaseDir string
}

func NewTyped[M adk.MessageType](_ context.Context, config *Config) (adk.TypedChatModelAgentMiddleware[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(ctx context.Context, config *Config) (adk.ChatModelAgentMiddleware, error) {
	_ = "STUB: not implemented"
	return *new(adk.ChatModelAgentMiddleware), nil
}

type typedMiddleware[M adk.MessageType] struct {
	*adk.TypedBaseChatModelAgentMiddleware[M]
	backend Backend
	baseDir string
}

func (m *typedMiddleware[M]) BeforeAgent(ctx context.Context, runCtx *adk.ChatModelAgentContext) (context.Context, *adk.ChatModelAgentContext, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}
