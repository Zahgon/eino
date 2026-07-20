package prompt

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

type DefaultChatTemplate struct {
	templates []schema.MessagesTemplate

	formatType schema.FormatType
}

func FromMessages(formatType schema.FormatType, templates ...schema.MessagesTemplate) *DefaultChatTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultChatTemplate) Format(ctx context.Context,
	vs map[string]any, _ ...Option) (result []*schema.Message, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *DefaultChatTemplate) GetType() string { _ = "STUB: not implemented"; return "" }

func (t *DefaultChatTemplate) IsCallbacksEnabled() bool { _ = "STUB: not implemented"; return false }
