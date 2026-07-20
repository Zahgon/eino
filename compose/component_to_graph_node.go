package compose

import (
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/retriever"
)

func toComponentNode[I, O, TOption any](
	node any,
	componentType component,
	invoke Invoke[I, O, TOption],
	stream Stream[I, O, TOption],
	collect Collect[I, O, TOption],
	transform Transform[I, O, TOption],
	opts ...GraphAddNodeOpt,
) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toEmbeddingNode(node embedding.Embedder, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toRetrieverNode(node retriever.Retriever, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toLoaderNode(node document.Loader, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toIndexerNode(node indexer.Indexer, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toChatModelNode(node model.BaseChatModel, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toAgenticModelNode(node model.AgenticModel, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toChatTemplateNode(node prompt.ChatTemplate, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toAgenticChatTemplateNode(node prompt.AgenticChatTemplate, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toDocumentTransformerNode(node document.Transformer, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toToolsNode(node *ToolsNode, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toAgenticToolsNode(node *AgenticToolsNode, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toLambdaNode(node *Lambda, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toAnyGraphNode(node AnyGraph, opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toPassthroughNode(opts ...GraphAddNodeOpt) (*graphNode, *graphAddNodeOpts) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toNode(nodeInfo *nodeInfo, executor *composableRunnable, graph AnyGraph,
	meta *executorMeta, instance any, opts ...GraphAddNodeOpt) *graphNode {
	_ = "STUB: not implemented"
	return nil
}
