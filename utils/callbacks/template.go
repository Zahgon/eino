package callbacks

import (
	"context"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components"
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/retriever"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

func NewHandlerHelper() *HandlerHelper { _ = "STUB: not implemented"; return nil }

type HandlerHelper struct {
	promptHandler           *PromptCallbackHandler
	chatModelHandler        *ModelCallbackHandler
	embeddingHandler        *EmbeddingCallbackHandler
	indexerHandler          *IndexerCallbackHandler
	retrieverHandler        *RetrieverCallbackHandler
	loaderHandler           *LoaderCallbackHandler
	transformerHandler      *TransformerCallbackHandler
	toolHandler             *ToolCallbackHandler
	toolsNodeHandler        *ToolsNodeCallbackHandlers
	agentHandler            *AgentCallbackHandler
	agenticAgentHandler     *AgenticAgentCallbackHandler
	agenticPromptHandler    *AgenticPromptCallbackHandler
	agenticModelHandler     *AgenticModelCallbackHandler
	agenticToolsNodeHandler *AgenticToolsNodeCallbackHandlers
	composeTemplates        map[components.Component]callbacks.Handler
}

func (c *HandlerHelper) Handler() callbacks.Handler {
	_ = "STUB: not implemented"
	return *new(callbacks.Handler)
}

func (c *HandlerHelper) Prompt(handler *PromptCallbackHandler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) ChatModel(handler *ModelCallbackHandler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) Embedding(handler *EmbeddingCallbackHandler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) Indexer(handler *IndexerCallbackHandler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) Retriever(handler *RetrieverCallbackHandler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) Loader(handler *LoaderCallbackHandler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) Transformer(handler *TransformerCallbackHandler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) Tool(handler *ToolCallbackHandler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) ToolsNode(handler *ToolsNodeCallbackHandlers) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) AgenticPrompt(handler *AgenticPromptCallbackHandler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) AgenticModel(handler *AgenticModelCallbackHandler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) AgenticToolsNode(handler *AgenticToolsNodeCallbackHandlers) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) Agent(handler *AgentCallbackHandler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) AgenticAgent(handler *AgenticAgentCallbackHandler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) Graph(handler callbacks.Handler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) Chain(handler callbacks.Handler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

func (c *HandlerHelper) Lambda(handler callbacks.Handler) *HandlerHelper {
	_ = "STUB: not implemented"
	return nil
}

type handlerTemplate struct {
	*HandlerHelper
}

func (c *handlerTemplate) OnStart(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *handlerTemplate) OnEnd(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *handlerTemplate) OnError(ctx context.Context, info *callbacks.RunInfo, err error) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *handlerTemplate) OnStartWithStreamInput(ctx context.Context, info *callbacks.RunInfo, input *schema.StreamReader[callbacks.CallbackInput]) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *handlerTemplate) OnEndWithStreamOutput(ctx context.Context, info *callbacks.RunInfo, output *schema.StreamReader[callbacks.CallbackOutput]) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

//nolint:cyclop
func (c *handlerTemplate) Needed(ctx context.Context, info *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type LoaderCallbackHandler struct {
	OnStart func(ctx context.Context, runInfo *callbacks.RunInfo, input *document.LoaderCallbackInput) context.Context
	OnEnd   func(ctx context.Context, runInfo *callbacks.RunInfo, output *document.LoaderCallbackOutput) context.Context
	OnError func(ctx context.Context, runInfo *callbacks.RunInfo, err error) context.Context
}

func (ch *LoaderCallbackHandler) Needed(ctx context.Context, runInfo *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type TransformerCallbackHandler struct {
	OnStart func(ctx context.Context, runInfo *callbacks.RunInfo, input *document.TransformerCallbackInput) context.Context
	OnEnd   func(ctx context.Context, runInfo *callbacks.RunInfo, output *document.TransformerCallbackOutput) context.Context
	OnError func(ctx context.Context, runInfo *callbacks.RunInfo, err error) context.Context
}

func (ch *TransformerCallbackHandler) Needed(ctx context.Context, runInfo *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type EmbeddingCallbackHandler struct {
	OnStart func(ctx context.Context, runInfo *callbacks.RunInfo, input *embedding.CallbackInput) context.Context
	OnEnd   func(ctx context.Context, runInfo *callbacks.RunInfo, output *embedding.CallbackOutput) context.Context
	OnError func(ctx context.Context, runInfo *callbacks.RunInfo, err error) context.Context
}

func (ch *EmbeddingCallbackHandler) Needed(ctx context.Context, runInfo *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type IndexerCallbackHandler struct {
	OnStart func(ctx context.Context, runInfo *callbacks.RunInfo, input *indexer.CallbackInput) context.Context
	OnEnd   func(ctx context.Context, runInfo *callbacks.RunInfo, output *indexer.CallbackOutput) context.Context
	OnError func(ctx context.Context, runInfo *callbacks.RunInfo, err error) context.Context
}

func (ch *IndexerCallbackHandler) Needed(ctx context.Context, runInfo *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type ModelCallbackHandler struct {
	OnStart               func(ctx context.Context, runInfo *callbacks.RunInfo, input *model.CallbackInput) context.Context
	OnEnd                 func(ctx context.Context, runInfo *callbacks.RunInfo, output *model.CallbackOutput) context.Context
	OnEndWithStreamOutput func(ctx context.Context, runInfo *callbacks.RunInfo, output *schema.StreamReader[*model.CallbackOutput]) context.Context
	OnError               func(ctx context.Context, runInfo *callbacks.RunInfo, err error) context.Context
}

func (ch *ModelCallbackHandler) Needed(ctx context.Context, runInfo *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type PromptCallbackHandler struct {
	OnStart func(ctx context.Context, runInfo *callbacks.RunInfo, input *prompt.CallbackInput) context.Context

	OnEnd func(ctx context.Context, runInfo *callbacks.RunInfo, output *prompt.CallbackOutput) context.Context

	OnError func(ctx context.Context, runInfo *callbacks.RunInfo, err error) context.Context
}

func (ch *PromptCallbackHandler) Needed(ctx context.Context, runInfo *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type RetrieverCallbackHandler struct {
	OnStart func(ctx context.Context, runInfo *callbacks.RunInfo, input *retriever.CallbackInput) context.Context

	OnEnd func(ctx context.Context, runInfo *callbacks.RunInfo, output *retriever.CallbackOutput) context.Context

	OnError func(ctx context.Context, runInfo *callbacks.RunInfo, err error) context.Context
}

func (ch *RetrieverCallbackHandler) Needed(ctx context.Context, runInfo *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type ToolCallbackHandler struct {
	OnStart               func(ctx context.Context, info *callbacks.RunInfo, input *tool.CallbackInput) context.Context
	OnEnd                 func(ctx context.Context, info *callbacks.RunInfo, output *tool.CallbackOutput) context.Context
	OnEndWithStreamOutput func(ctx context.Context, info *callbacks.RunInfo, output *schema.StreamReader[*tool.CallbackOutput]) context.Context
	OnError               func(ctx context.Context, info *callbacks.RunInfo, err error) context.Context
}

func (ch *ToolCallbackHandler) Needed(ctx context.Context, runInfo *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type ToolsNodeCallbackHandlers struct {
	OnStart               func(ctx context.Context, info *callbacks.RunInfo, input *schema.Message) context.Context
	OnEnd                 func(ctx context.Context, info *callbacks.RunInfo, input []*schema.Message) context.Context
	OnEndWithStreamOutput func(ctx context.Context, info *callbacks.RunInfo, output *schema.StreamReader[[]*schema.Message]) context.Context
	OnError               func(ctx context.Context, info *callbacks.RunInfo, err error) context.Context
}

func (ch *ToolsNodeCallbackHandlers) Needed(ctx context.Context, runInfo *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

func convToolsNodeCallbackInput(src callbacks.CallbackInput) *schema.Message {
	_ = "STUB: not implemented"
	return nil
}

func convToolsNodeCallbackOutput(src callbacks.CallbackInput) []*schema.Message {
	_ = "STUB: not implemented"
	return nil
}

type AgentCallbackHandler struct {
	OnStart func(ctx context.Context, info *callbacks.RunInfo, input *adk.AgentCallbackInput) context.Context

	OnEnd func(ctx context.Context, info *callbacks.RunInfo, output *adk.AgentCallbackOutput) context.Context
}

func (ch *AgentCallbackHandler) Needed(ctx context.Context, info *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type AgenticAgentCallbackHandler struct {
	OnStart func(ctx context.Context, info *callbacks.RunInfo, input *adk.TypedAgentCallbackInput[*schema.AgenticMessage]) context.Context

	OnEnd func(ctx context.Context, info *callbacks.RunInfo, output *adk.TypedAgentCallbackOutput[*schema.AgenticMessage]) context.Context
}

func (ch *AgenticAgentCallbackHandler) Needed(ctx context.Context, info *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type AgenticPromptCallbackHandler struct {
	OnStart func(ctx context.Context, runInfo *callbacks.RunInfo, input *prompt.CallbackInput) context.Context

	OnEnd func(ctx context.Context, runInfo *callbacks.RunInfo, output *prompt.CallbackOutput) context.Context

	OnError func(ctx context.Context, runInfo *callbacks.RunInfo, err error) context.Context
}

func (ch *AgenticPromptCallbackHandler) Needed(ctx context.Context, runInfo *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type AgenticModelCallbackHandler struct {
	OnStart               func(ctx context.Context, runInfo *callbacks.RunInfo, input *model.AgenticCallbackInput) context.Context
	OnEnd                 func(ctx context.Context, runInfo *callbacks.RunInfo, output *model.AgenticCallbackOutput) context.Context
	OnEndWithStreamOutput func(ctx context.Context, runInfo *callbacks.RunInfo, output *schema.StreamReader[*model.AgenticCallbackOutput]) context.Context
	OnError               func(ctx context.Context, runInfo *callbacks.RunInfo, err error) context.Context
}

func (ch *AgenticModelCallbackHandler) Needed(ctx context.Context, runInfo *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

type AgenticToolsNodeCallbackHandlers struct {
	OnStart               func(ctx context.Context, info *callbacks.RunInfo, input *schema.AgenticMessage) context.Context
	OnEnd                 func(ctx context.Context, info *callbacks.RunInfo, input []*schema.AgenticMessage) context.Context
	OnEndWithStreamOutput func(ctx context.Context, info *callbacks.RunInfo, output *schema.StreamReader[[]*schema.AgenticMessage]) context.Context
	OnError               func(ctx context.Context, info *callbacks.RunInfo, err error) context.Context
}

func (ch *AgenticToolsNodeCallbackHandlers) Needed(ctx context.Context, runInfo *callbacks.RunInfo, timing callbacks.CallbackTiming) bool {
	_ = "STUB: not implemented"
	return false
}

func convAgenticToolsNodeCallbackInput(src callbacks.CallbackInput) *schema.AgenticMessage {
	_ = "STUB: not implemented"
	return nil
}

func convAgenticToolsNodeCallbackOutput(src callbacks.CallbackInput) []*schema.AgenticMessage {
	_ = "STUB: not implemented"
	return nil
}
