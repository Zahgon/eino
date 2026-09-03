package compose

import (
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/retriever"
)

func NewParallel() *Parallel { _ = "STUB: not implemented"; return nil }

type Parallel struct {
	nodes      []nodeOptionsPair
	outputKeys map[string]bool
	err        error
}

func (p *Parallel) AddChatModel(outputKey string, node model.BaseChatModel, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddAgenticModel(outputKey string, node model.AgenticModel, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddChatTemplate(outputKey string, node prompt.ChatTemplate, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddAgenticChatTemplate(outputKey string, node prompt.AgenticChatTemplate, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddToolsNode(outputKey string, node *ToolsNode, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddAgenticToolsNode(outputKey string, node *AgenticToolsNode, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddLambda(outputKey string, node *Lambda, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddEmbedding(outputKey string, node embedding.Embedder, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddRetriever(outputKey string, node retriever.Retriever, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddLoader(outputKey string, node document.Loader, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddIndexer(outputKey string, node indexer.Indexer, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddDocumentTransformer(outputKey string, node document.Transformer, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddGraph(outputKey string, node AnyGraph, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) AddPassthrough(outputKey string, opts ...GraphAddNodeOpt) *Parallel {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parallel) addNode(outputKey string, node *graphNode, options *graphAddNodeOpts) *Parallel {
	_ = "STUB: not implemented"
	return nil
}
