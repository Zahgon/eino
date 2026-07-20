package prompt

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

func FromAgenticMessages(formatType schema.FormatType, templates ...schema.AgenticMessagesTemplate) *DefaultAgenticChatTemplate {
	_ = "STUB: not implemented"
	return nil
}

type DefaultAgenticChatTemplate struct {
	templates  []schema.AgenticMessagesTemplate
	formatType schema.FormatType
}

func (t *DefaultAgenticChatTemplate) Format(ctx context.Context, vs map[string]any, opts ...Option) (result []*schema.AgenticMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *DefaultAgenticChatTemplate) GetType() string { _ = "STUB: not implemented"; return "" }

func (t *DefaultAgenticChatTemplate) IsCallbacksEnabled() bool {
	_ = "STUB: not implemented"
	return false
}
