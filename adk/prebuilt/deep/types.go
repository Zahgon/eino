package deep

import (
	"context"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
)

const (
	generalAgentName = "general-purpose"
	taskToolName     = "task"
)

const (
	SessionKeyTodos = "deep_agent_session_key_todos"
)

func assertAgentTool(t tool.BaseTool) (tool.InvokableTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.InvokableTool), nil
}

func typedBuildAppendPromptTool[M adk.MessageType](prompt string, t tool.BaseTool) adk.TypedChatModelAgentMiddleware[M] {
	_ = "STUB: not implemented"
	return nil
}

type typedAppendPromptTool[M adk.MessageType] struct {
	*adk.TypedBaseChatModelAgentMiddleware[M]
	t      tool.BaseTool
	prompt string
}

func (w *typedAppendPromptTool[M]) BeforeAgent(ctx context.Context, runCtx *adk.ChatModelAgentContext) (context.Context, *adk.ChatModelAgentContext, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}
