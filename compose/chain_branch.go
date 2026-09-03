package compose

import (
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/retriever"
	"github.com/cloudwego/eino/internal/generic"
)

type nodeOptionsPair generic.Pair[*graphNode, *graphAddNodeOpts]

type ChainBranch struct {
	internalBranch *GraphBranch
	key2BranchNode map[string]nodeOptionsPair
	err            error
}

func NewChainMultiBranch[T any](cond GraphMultiBranchCondition[T]) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func NewStreamChainMultiBranch[T any](cond StreamGraphMultiBranchCondition[T]) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func NewChainBranch[T any](cond GraphBranchCondition[T]) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func NewStreamChainBranch[T any](cond StreamGraphBranchCondition[T]) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddChatModel(key string, node model.BaseChatModel, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddAgenticModel(key string, node model.AgenticModel, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddChatTemplate(key string, node prompt.ChatTemplate, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddAgenticChatTemplate(key string, node prompt.AgenticChatTemplate, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddToolsNode(key string, node *ToolsNode, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddAgenticToolsNode(key string, node *AgenticToolsNode, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddLambda(key string, node *Lambda, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddEmbedding(key string, node embedding.Embedder, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddRetriever(key string, node retriever.Retriever, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddLoader(key string, node document.Loader, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddIndexer(key string, node indexer.Indexer, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddDocumentTransformer(key string, node document.Transformer, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddGraph(key string, node AnyGraph, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) AddPassthrough(key string, opts ...GraphAddNodeOpt) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ChainBranch) addNode(key string, node *graphNode, options *graphAddNodeOpts) *ChainBranch {
	_ = "STUB: not implemented"
	return nil
}
