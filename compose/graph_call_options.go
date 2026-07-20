package compose

import (
	"context"
	"time"

	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/retriever"
)

type graphCancelChanKey struct{}
type graphCancelChanVal struct {
	ch chan *time.Duration
}

type graphInterruptOptions struct {
	timeout *time.Duration
}

type GraphInterruptOption func(o *graphInterruptOptions)

func WithGraphInterruptTimeout(timeout time.Duration) GraphInterruptOption {
	_ = "STUB: not implemented"
	return *new(GraphInterruptOption)
}

func WithGraphInterrupt(parent context.Context) (ctx context.Context, interrupt func(opts ...GraphInterruptOption)) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func getGraphCancel(ctx context.Context) *graphCancelChanVal { _ = "STUB: not implemented"; return nil }

type Option struct {
	options []any
	handler []callbacks.Handler

	paths []*NodePath

	maxRunSteps         int
	checkPointID        *string
	writeToCheckPointID *string
	forceNewRun         bool
	stateModifier       StateModifier
}

func (o Option) deepCopy() Option { _ = "STUB: not implemented"; return *new(Option) }

func (o Option) DesignateNode(nodeKey ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func (o Option) DesignateNodeWithPath(path ...*NodePath) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithEmbeddingOption(opts ...embedding.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithRetrieverOption(opts ...retriever.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLoaderOption(opts ...document.LoaderOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithDocumentTransformerOption(opts ...document.TransformerOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithIndexerOption(opts ...indexer.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithChatModelOption(opts ...model.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithChatTemplateOption(opts ...prompt.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithToolsNodeOption(opts ...ToolsNodeOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLambdaOption(opts ...any) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCallbacks(cbs ...callbacks.Handler) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRuntimeMaxSteps(maxSteps int) Option { _ = "STUB: not implemented"; return *new(Option) }

func withComponentOption[TOption any](opts ...TOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func convertOption[TOption any](opts ...any) ([]TOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
