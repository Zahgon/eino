package deep

import (
	"context"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/adk/filesystem"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

func init() {
	schema.RegisterName[TODO]("_eino_adk_prebuilt_deep_todo")
	schema.RegisterName[[]TODO]("_eino_adk_prebuilt_deep_todo_slice")
}

type TypedConfig[M adk.MessageType] struct {
	Name string

	Description string

	ChatModel model.BaseModel[M]

	Instruction string

	SubAgents []adk.TypedAgent[M]

	ToolsConfig adk.ToolsConfig

	MaxIteration int

	Backend filesystem.Backend

	Shell filesystem.Shell

	StreamingShell filesystem.StreamingShell

	WithoutWriteTodos bool

	WithoutGeneralSubAgent bool

	TaskToolDescriptionGenerator func(ctx context.Context, availableAgents []adk.TypedAgent[M]) (string, error)

	Middlewares []adk.AgentMiddleware

	Handlers []adk.TypedChatModelAgentMiddleware[M]

	ModelRetryConfig *adk.TypedModelRetryConfig[M]

	ModelFailoverConfig *adk.ModelFailoverConfig[M]

	OutputKey string
}

type Config = TypedConfig[*schema.Message]

func NewTyped[M adk.MessageType](ctx context.Context, cfg *TypedConfig[M]) (adk.TypedResumableAgent[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(ctx context.Context, cfg *Config) (adk.ResumableAgent, error) {
	_ = "STUB: not implemented"
	return *new(adk.ResumableAgent), nil
}

func typedGenModelInput[M adk.MessageType](_ context.Context, instruction string, input *adk.TypedAgentInput[M]) ([]M, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildTypedBuiltinAgentMiddlewares[M adk.MessageType](ctx context.Context, cfg *TypedConfig[M]) ([]adk.TypedChatModelAgentMiddleware[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TODO struct {
	Content    string `json:"content"`
	ActiveForm string `json:"activeForm"`
	Status     string `json:"status" jsonschema:"enum=pending,enum=in_progress,enum=completed"`
}

type writeTodosArguments struct {
	Todos []TODO `json:"todos"`
}

func typedNewWriteTodos[M adk.MessageType]() (adk.TypedChatModelAgentMiddleware[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
