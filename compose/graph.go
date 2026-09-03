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

const START = "start"

const END = "end"

type graphRunType string

const (
	runTypePregel graphRunType = "Pregel"

	runTypeDAG graphRunType = "DAG"
)

func (g graphRunType) String() string { _ = "STUB: not implemented"; return "" }

type graph struct {
	nodes        map[string]*graphNode
	controlEdges map[string][]string
	dataEdges    map[string][]string
	branches     map[string][]*GraphBranch
	startNodes   []string
	endNodes     []string

	toValidateMap map[string][]struct {
		endNode  string
		mappings []*FieldMapping
	}

	stateType      reflect.Type
	stateGenerator func(ctx context.Context) any
	newOpts        []NewGraphOption

	expectedInputType, expectedOutputType reflect.Type

	*genericHelper

	fieldMappingRecords map[string][]*FieldMapping

	buildError error

	cmp component

	compiled bool

	handlerOnEdges   map[string]map[string][]handlerPair
	handlerPreNode   map[string][]handlerPair
	handlerPreBranch map[string][][]handlerPair
}

type newGraphConfig struct {
	inputType, outputType reflect.Type
	gh                    *genericHelper
	cmp                   component
	stateType             reflect.Type
	stateGenerator        func(ctx context.Context) any
	newOpts               []NewGraphOption
}

func newGraphFromGeneric[I, O any](
	cmp component,
	stateGenerator func(ctx context.Context) any,
	stateType reflect.Type,
	opts []NewGraphOption,
) *graph {
	_ = "STUB: not implemented"
	return nil
}

func newGraph(cfg *newGraphConfig) *graph { _ = "STUB: not implemented"; return nil }

func (g *graph) component() component { _ = "STUB: not implemented"; return *new(component) }

func isChain(cmp component) bool { _ = "STUB: not implemented"; return false }

func isWorkflow(cmp component) bool { _ = "STUB: not implemented"; return false }

var ErrGraphCompiled = errors.New("graph has been compiled, cannot be modified")

func (g *graph) addNode(key string, node *graphNode, options *graphAddNodeOpts) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) addEdgeWithMappings(startNode, endNode string, noControl bool, noData bool, mappings ...*FieldMapping) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddEmbeddingNode(key string, node embedding.Embedder, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddRetrieverNode(key string, node retriever.Retriever, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddLoaderNode(key string, node document.Loader, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddIndexerNode(key string, node indexer.Indexer, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddChatModelNode(key string, node model.BaseChatModel, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddAgenticModelNode(key string, node model.AgenticModel, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddChatTemplateNode(key string, node prompt.ChatTemplate, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddAgenticChatTemplateNode(key string, node prompt.AgenticChatTemplate, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddToolsNode(key string, node *ToolsNode, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddAgenticToolsNode(key string, node *AgenticToolsNode, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddDocumentTransformerNode(key string, node document.Transformer, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddLambdaNode(key string, node *Lambda, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddGraphNode(key string, node AnyGraph, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddPassthroughNode(key string, opts ...GraphAddNodeOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) AddBranch(startNode string, branch *GraphBranch) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) addBranch(startNode string, branch *GraphBranch, skipData bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) addToValidateMap(startNode, endNode string, mapping []*FieldMapping) {
	_ = "STUB: not implemented"
	return
}

func (g *graph) updateToValidateMap() error { _ = "STUB: not implemented"; return nil }

func (g *graph) getNodeGenericHelper(name string) *genericHelper {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) getNodeInputType(name string) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (g *graph) getNodeOutputType(name string) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (g *graph) inputType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (g *graph) outputType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (g *graph) compile(ctx context.Context, opt *graphCompileOptions) (*composableRunnable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSuccessors(c *chanCall) []string { _ = "STUB: not implemented"; return nil }

func uniqueSlice(s []string) []string { _ = "STUB: not implemented"; return nil }

type subGraphCompileCallback struct {
	closure func(ctx context.Context, info *GraphInfo)
}

func (s *subGraphCompileCallback) OnFinish(ctx context.Context, info *GraphInfo) {
	_ = "STUB: not implemented"
	return
}

func (g *graph) beforeChildGraphsCompile(opt *graphCompileOptions) map[string]*GraphInfo {
	_ = "STUB: not implemented"
	return nil
}

func (gn *graphNode) beforeChildGraphCompile(nodeKey string, key2SubGraphs map[string]*GraphInfo) {
	_ = "STUB: not implemented"
	return
}

func (g *graph) toGraphInfo(opt *graphCompileOptions, key2SubGraphs map[string]*GraphInfo) *GraphInfo {
	_ = "STUB: not implemented"
	return nil
}

func (g *graph) onCompileFinish(ctx context.Context, opt *graphCompileOptions, key2SubGraphs map[string]*GraphInfo) {
	_ = "STUB: not implemented"
	return
}

func (g *graph) getGenericHelper() *genericHelper { _ = "STUB: not implemented"; return nil }

func (g *graph) GetType() string { _ = "STUB: not implemented"; return "" }

func transferTask(script [][]string, invertedEdges map[string][]string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

func validateDAG(chanSubscribeTo map[string]*chanCall, controlPredecessors map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}

var DAGInvalidLoopErr = errors.New("DAG is invalid, has loop")

func findLoops(startNodes []string, chanCalls map[string]*chanCall) [][]string {
	_ = "STUB: not implemented"
	return nil
}

func formatLoops(loops [][]string) string { _ = "STUB: not implemented"; return "" }

func NewNodePath(nodeKeyPath ...string) *NodePath { _ = "STUB: not implemented"; return nil }

type NodePath struct {
	path []string
}

func (p *NodePath) GetPath() []string { _ = "STUB: not implemented"; return nil }
