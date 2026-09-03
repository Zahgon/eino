package react

import (
	"context"

	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent"
	"github.com/cloudwego/eino/internal"
	"github.com/cloudwego/eino/schema"
)

func WithToolOptions(opts ...tool.Option) agent.AgentOption {
	_ = "STUB: not implemented"
	return *new(agent.AgentOption)
}

func WithChatModelOptions(opts ...model.Option) agent.AgentOption {
	_ = "STUB: not implemented"
	return *new(agent.AgentOption)
}

func WithToolList(tools ...tool.BaseTool) agent.AgentOption {
	_ = "STUB: not implemented"
	return *new(agent.AgentOption)
}

func WithTools(ctx context.Context, tools ...tool.BaseTool) ([]agent.AgentOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Iterator[T any] struct {
	ch *internal.UnboundedChan[item[T]]
}

func (iter *Iterator[T]) Next() (T, bool, error) {
	_ = "STUB: not implemented"
	return *new(T), false, nil
}

type MessageFuture interface {
	GetMessages() *Iterator[*schema.Message]

	GetMessageStreams() *Iterator[*schema.StreamReader[*schema.Message]]
}

func WithMessageFuture() (agent.AgentOption, MessageFuture) {
	_ = "STUB: not implemented"
	return *new(agent.AgentOption), *new(MessageFuture)
}

type item[T any] struct {
	v   T
	err error
}

type cbHandler struct {
	graphName string

	ownAddress compose.Address
	ownClaimed bool

	msgs  *internal.UnboundedChan[item[*schema.Message]]
	sMsgs *internal.UnboundedChan[item[*schema.StreamReader[*schema.Message]]]

	started chan struct{}
}

func (h *cbHandler) GetMessages() *Iterator[*schema.Message] { _ = "STUB: not implemented"; return nil }

func (h *cbHandler) GetMessageStreams() *Iterator[*schema.StreamReader[*schema.Message]] {
	_ = "STUB: not implemented"
	return nil
}

func (h *cbHandler) isOwnGraph(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (h *cbHandler) claimOwnership(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *cbHandler) onChatModelEnd(ctx context.Context,
	_ *callbacks.RunInfo, input *model.CallbackOutput) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *cbHandler) onChatModelEndWithStreamOutput(ctx context.Context,
	_ *callbacks.RunInfo, input *schema.StreamReader[*model.CallbackOutput]) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *cbHandler) onGraphError(ctx context.Context,
	_ *callbacks.RunInfo, err error) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *cbHandler) onGraphEnd(ctx context.Context,
	_ *callbacks.RunInfo, _ callbacks.CallbackOutput) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *cbHandler) onGraphEndWithStreamOutput(ctx context.Context,
	_ *callbacks.RunInfo, _ *schema.StreamReader[callbacks.CallbackOutput]) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *cbHandler) onGraphStart(ctx context.Context,
	_ *callbacks.RunInfo, _ callbacks.CallbackInput) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *cbHandler) onGraphStartWithStreamInput(ctx context.Context, _ *callbacks.RunInfo,
	input *schema.StreamReader[callbacks.CallbackInput]) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *cbHandler) sendMessage(msg *schema.Message) { _ = "STUB: not implemented"; return }

func (h *cbHandler) sendMessageStream(sMsg *schema.StreamReader[*schema.Message]) {
	_ = "STUB: not implemented"
	return
}
