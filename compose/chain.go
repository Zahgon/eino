package compose

import (
	"context"
	"errors"
	"reflect"

	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/retriever"
)

func NewChain[I, O any](opts ...NewGraphOption) *Chain[I, O] { _ = "STUB: not implemented"; return nil }

type Chain[I, O any] struct {
	err error

	gg *Graph[I, O]

	nodeIdx int

	preNodeKeys []string

	hasEnd bool
}

var ErrChainCompiled = errors.New("chain has been compiled, cannot be modified")

func (c *Chain[I, O]) compile(ctx context.Context, option *graphCompileOptions) (*composableRunnable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Chain[I, O]) addEndIfNeeded() error { _ = "STUB: not implemented"; return nil }

func (c *Chain[I, O]) getGenericHelper() *genericHelper { _ = "STUB: not implemented"; return nil }

func (c *Chain[I, O]) inputType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (c *Chain[I, O]) outputType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (c *Chain[I, O]) component() component { _ = "STUB: not implemented"; return *new(component) }

func (c *Chain[I, O]) Compile(ctx context.Context, opts ...GraphCompileOption) (Runnable[I, O], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Chain[I, O]) AppendChatModel(node model.BaseChatModel, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendAgenticModel(node model.AgenticModel, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendChatTemplate(node prompt.ChatTemplate, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendAgenticChatTemplate(node prompt.AgenticChatTemplate, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendToolsNode(node *ToolsNode, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendAgenticToolsNode(node *AgenticToolsNode, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendDocumentTransformer(node document.Transformer, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendLambda(node *Lambda, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendEmbedding(node embedding.Embedder, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendRetriever(node retriever.Retriever, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendLoader(node document.Loader, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendIndexer(node indexer.Indexer, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendBranch(b *ChainBranch) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendParallel(p *Parallel) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendGraph(node AnyGraph, opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) AppendPassthrough(opts ...GraphAddNodeOpt) *Chain[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chain[I, O]) nextNodeKey() string { _ = "STUB: not implemented"; return "" }

func (c *Chain[I, O]) reportError(err error) { _ = "STUB: not implemented"; return }

func (c *Chain[I, O]) addNode(node *graphNode, options *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return
}
