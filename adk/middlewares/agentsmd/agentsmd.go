package agentsmd

import (
	"context"

	"github.com/cloudwego/eino/adk"
)

type Config struct {
	Backend Backend

	AgentsMDFiles []string

	AllAgentsMDMaxBytes int

	PerAgentsMDMaxBytes int

	OnLoadWarning func(filePath string, err error)
}

func NewTyped[M adk.MessageType](_ context.Context, cfg *Config) (adk.TypedChatModelAgentMiddleware[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(ctx context.Context, cfg *Config) (adk.ChatModelAgentMiddleware, error) {
	_ = "STUB: not implemented"
	return *new(adk.ChatModelAgentMiddleware), nil
}

type typedMiddleware[M adk.MessageType] struct {
	*adk.TypedBaseChatModelAgentMiddleware[M]
	loader *loaderConfig
}

const agentsMDCacheKey = "__agentsmd_content_cache__"
const agentsMDExtraKey = "__agentsmd_content__"

func (m *typedMiddleware[M]) BeforeModelRewriteState(ctx context.Context, state *adk.TypedChatModelAgentState[M], _ *adk.TypedModelContext[M]) (context.Context, *adk.TypedChatModelAgentState[M], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func hasAgentsMDExtra[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func typedInsertBeforeFirstUser[M adk.MessageType](msgs []M, content string) []M {
	_ = "STUB: not implemented"
	return nil
}

func isUserRole[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func makeUserMsgWithExtra[M adk.MessageType](content string) M {
	_ = "STUB: not implemented"
	return *new(M)
}

func (m *typedMiddleware[M]) loadContent(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Config) validate() error { _ = "STUB: not implemented"; return nil }
