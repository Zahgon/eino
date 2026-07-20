package deep

import (
	"context"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

func typedTaskToolMiddleware[M adk.MessageType](
	ctx context.Context,
	taskToolDescriptionGenerator func(ctx context.Context, subAgents []adk.TypedAgent[M]) (string, error),
	subAgents []adk.TypedAgent[M],

	withoutGeneralSubAgent bool,
	cm model.BaseModel[M],
	instruction string,
	toolsConfig adk.ToolsConfig,
	maxIteration int,
	middlewares []adk.AgentMiddleware,
	handlers []adk.TypedChatModelAgentMiddleware[M],
	modelFailoverConfig *adk.ModelFailoverConfig[M],
) (adk.TypedChatModelAgentMiddleware[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func typedNewTaskTool[M adk.MessageType](
	ctx context.Context,
	taskToolDescriptionGenerator func(ctx context.Context, subAgents []adk.TypedAgent[M]) (string, error),
	subAgents []adk.TypedAgent[M],

	withoutGeneralSubAgent bool,
	cm model.BaseModel[M],
	instruction string,
	toolsConfig adk.ToolsConfig,
	maxIteration int,
	middlewares []adk.AgentMiddleware,
	handlers []adk.TypedChatModelAgentMiddleware[M],
	modelFailoverConfig *adk.ModelFailoverConfig[M],
) (tool.InvokableTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.InvokableTool), nil
}

type typedTaskTool[M adk.MessageType] struct {
	subAgents     map[string]tool.InvokableTool
	subAgentSlice []adk.TypedAgent[M]
	descGen       func(ctx context.Context, subAgents []adk.TypedAgent[M]) (string, error)
}

func (t *typedTaskTool[M]) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type taskToolArgument struct {
	SubagentType string `json:"subagent_type"`
	Description  string `json:"description"`
}

func (t *typedTaskTool[M]) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func typedDefaultTaskToolDescription[M adk.MessageType](ctx context.Context, subAgents []adk.TypedAgent[M]) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
