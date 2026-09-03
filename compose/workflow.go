package compose

import (
	"context"
	"reflect"

	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/retriever"
)

type WorkflowNode struct {
	g                *graph
	key              string
	addInputs        []func() error
	staticValues     map[string]any
	dependencySetter func(fromNodeKey string, typ dependencyType)
	mappedFieldPath  map[string]any
}

type Workflow[I, O any] struct {
	g                *graph
	workflowNodes    map[string]*WorkflowNode
	workflowBranches []*WorkflowBranch
	dependencies     map[string]map[string]dependencyType
}

type dependencyType int

const (
	normalDependency dependencyType = iota
	noDirectDependency
	branchDependency
)

func NewWorkflow[I, O any](opts ...NewGraphOption) *Workflow[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) Compile(ctx context.Context, opts ...GraphCompileOption) (Runnable[I, O], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (wf *Workflow[I, O]) AddChatModelNode(key string, chatModel model.BaseChatModel, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddAgenticModelNode(key string, agenticModel model.AgenticModel, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddChatTemplateNode(key string, chatTemplate prompt.ChatTemplate, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddAgenticChatTemplateNode(key string, chatTemplate prompt.AgenticChatTemplate, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddToolsNode(key string, tools *ToolsNode, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddAgenticToolsNode(key string, tools *AgenticToolsNode, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddRetrieverNode(key string, retriever retriever.Retriever, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddEmbeddingNode(key string, embedding embedding.Embedder, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddIndexerNode(key string, indexer indexer.Indexer, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddLoaderNode(key string, loader document.Loader, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddDocumentTransformerNode(key string, transformer document.Transformer, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddGraphNode(key string, graph AnyGraph, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddLambdaNode(key string, lambda *Lambda, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) End() *WorkflowNode { _ = "STUB: not implemented"; return nil }

func (wf *Workflow[I, O]) AddPassthroughNode(key string, opts ...GraphAddNodeOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (n *WorkflowNode) AddInput(fromNodeKey string, inputs ...*FieldMapping) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

type workflowAddInputOpts struct {
	noDirectDependency bool

	dependencyWithoutInput bool
}

type WorkflowAddInputOpt func(*workflowAddInputOpts)

func getAddInputOpts(opts []WorkflowAddInputOpt) *workflowAddInputOpts {
	_ = "STUB: not implemented"
	return nil
}

func WithNoDirectDependency() WorkflowAddInputOpt {
	_ = "STUB: not implemented"
	return *new(WorkflowAddInputOpt)
}

func (n *WorkflowNode) AddInputWithOptions(fromNodeKey string, inputs []*FieldMapping, opts ...WorkflowAddInputOpt) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (n *WorkflowNode) AddDependency(fromNodeKey string) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (n *WorkflowNode) SetStaticValue(path FieldPath, value any) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (n *WorkflowNode) addDependencyRelation(fromNodeKey string, inputs []*FieldMapping, options *workflowAddInputOpts) *WorkflowNode {
	_ = "STUB: not implemented"
	return nil
}

func (n *WorkflowNode) checkAndAddMappedPath(paths []FieldPath) error {
	_ = "STUB: not implemented"
	return nil
}

type WorkflowBranch struct {
	fromNodeKey string
	*GraphBranch
}

func (wf *Workflow[I, O]) AddBranch(fromNodeKey string, branch *GraphBranch) *WorkflowBranch {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) AddEnd(fromNodeKey string, inputs ...*FieldMapping) *Workflow[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (wf *Workflow[I, O]) compile(ctx context.Context, options *graphCompileOptions) (*composableRunnable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (wf *Workflow[I, O]) initNode(key string) *WorkflowNode { _ = "STUB: not implemented"; return nil }

func (wf *Workflow[I, O]) getGenericHelper() *genericHelper { _ = "STUB: not implemented"; return nil }

func (wf *Workflow[I, O]) inputType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (wf *Workflow[I, O]) outputType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (wf *Workflow[I, O]) component() component { _ = "STUB: not implemented"; return *new(component) }
