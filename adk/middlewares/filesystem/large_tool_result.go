package filesystem

import (
	"context"

	"github.com/cloudwego/eino/adk/filesystem"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

type toolResultOffloadingConfig struct {
	Backend       filesystem.Backend
	TokenLimit    int
	PathGenerator func(ctx context.Context, input *compose.ToolInput) (string, error)
}

func newToolResultOffloading(ctx context.Context, config *toolResultOffloadingConfig) compose.ToolMiddleware {
	_ = "STUB: not implemented"
	return *new(compose.ToolMiddleware)
}

type toolResultOffloading struct {
	backend       filesystem.Backend
	tokenLimit    int
	pathGenerator func(ctx context.Context, input *compose.ToolInput) (string, error)
}

func (t *toolResultOffloading) invoke(endpoint compose.InvokableToolEndpoint) compose.InvokableToolEndpoint {
	_ = "STUB: not implemented"
	return *new(compose.InvokableToolEndpoint)
}

func (t *toolResultOffloading) stream(endpoint compose.StreamableToolEndpoint) compose.StreamableToolEndpoint {
	_ = "STUB: not implemented"
	return *new(compose.StreamableToolEndpoint)
}

func (t *toolResultOffloading) handleResult(ctx context.Context, result string, input *compose.ToolInput) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func concatString(sr *schema.StreamReader[string]) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func formatToolMessage(s string) string { _ = "STUB: not implemented"; return "" }
