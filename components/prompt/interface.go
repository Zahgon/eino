package prompt

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

var _ ChatTemplate = &DefaultChatTemplate{}
var _ AgenticChatTemplate = &DefaultAgenticChatTemplate{}

type ChatTemplate interface {
	Format(ctx context.Context, vs map[string]any, opts ...Option) ([]*schema.Message, error)
}

type AgenticChatTemplate interface {
	Format(ctx context.Context, vs map[string]any, opts ...Option) ([]*schema.AgenticMessage, error)
}
