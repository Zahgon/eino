package skill

import (
	"context"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type ContextMode string

const (
	ContextModeFork ContextMode = "fork"

	ContextModeForkWithContext ContextMode = "fork_with_context"
)

type FrontMatter struct {
	Name        string      `yaml:"name"`
	Description string      `yaml:"description"`
	Context     ContextMode `yaml:"context"`
	Agent       string      `yaml:"agent"`
	Model       string      `yaml:"model"`
}

type Skill struct {
	FrontMatter

	Content string

	BaseDirectory string
}

type Backend interface {
	List(ctx context.Context) ([]FrontMatter, error)
	Get(ctx context.Context, name string) (Skill, error)
}

type TypedAgentHubOptions[M adk.MessageType] struct {
	Model model.BaseModel[M]
}

type AgentHubOptions = TypedAgentHubOptions[*schema.Message]

type TypedAgentHub[M adk.MessageType] interface {
	Get(ctx context.Context, name string, opts *TypedAgentHubOptions[M]) (adk.TypedAgent[M], error)
}

type AgentHub = TypedAgentHub[*schema.Message]

type TypedModelHub[M adk.MessageType] interface {
	Get(ctx context.Context, name string) (model.BaseModel[M], error)
}

type ModelHub = TypedModelHub[*schema.Message]

type SystemPromptFunc func(ctx context.Context, toolName string) string

type ToolDescriptionFunc func(ctx context.Context, skills []FrontMatter) string

type TypedSubAgentInput[M adk.MessageType] struct {
	Skill        Skill
	Mode         ContextMode
	RawArguments string
	SkillContent string
	History      []M
	ToolCallID   string
}

type SubAgentInput = TypedSubAgentInput[*schema.Message]

type TypedSubAgentOutput[M adk.MessageType] struct {
	Skill        Skill
	Mode         ContextMode
	RawArguments string
	Messages     []M
	Results      []string
}

type SubAgentOutput = TypedSubAgentOutput[*schema.Message]

type TypedConfig[M adk.MessageType] struct {
	Backend Backend

	SkillToolName *string

	UseChinese bool

	AgentHub TypedAgentHub[M]

	ModelHub TypedModelHub[M]

	CustomSystemPrompt SystemPromptFunc

	CustomToolDescription ToolDescriptionFunc

	CustomToolParams func(ctx context.Context, defaults map[string]*schema.ParameterInfo) (map[string]*schema.ParameterInfo, error)

	BuildContent func(ctx context.Context, skill Skill, rawArgs string) (string, error)

	BuildForkMessages func(ctx context.Context, in TypedSubAgentInput[M]) ([]M, error)

	FormatForkResult func(ctx context.Context, in TypedSubAgentOutput[M]) (string, error)
}

type Config = TypedConfig[*schema.Message]

func NewTyped[M adk.MessageType](ctx context.Context, config *TypedConfig[M]) (adk.TypedChatModelAgentMiddleware[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewMiddleware(ctx context.Context, config *Config) (adk.ChatModelAgentMiddleware, error) {
	_ = "STUB: not implemented"
	return *new(adk.ChatModelAgentMiddleware), nil
}

type typedSkillHandler[M adk.MessageType] struct {
	*adk.TypedBaseChatModelAgentMiddleware[M]
	instruction string
	tool        *typedSkillTool[M]
}

func (h *typedSkillHandler[M]) BeforeAgent(ctx context.Context, runCtx *adk.ChatModelAgentContext) (context.Context, *adk.ChatModelAgentContext, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func (h *typedSkillHandler[M]) WrapModel(ctx context.Context, m model.BaseModel[M], _ *adk.TypedModelContext[M]) (model.BaseModel[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const activeModelKey = "__skill_active_model__"

func New(ctx context.Context, config *Config) (adk.AgentMiddleware, error) {
	_ = "STUB: not implemented"
	return *new(adk.AgentMiddleware), nil
}

func buildSystemPrompt(skillToolName string, useChinese bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type typedSkillTool[M adk.MessageType] struct {
	b        Backend
	toolName string

	useChinese bool
	agentHub   TypedAgentHub[M]
	modelHub   TypedModelHub[M]

	customToolDesc ToolDescriptionFunc

	customToolParams func(ctx context.Context, defaults map[string]*schema.ParameterInfo) (map[string]*schema.ParameterInfo, error)
	buildContent     func(ctx context.Context, skill Skill, rawArgs string) (string, error)

	buildForkMessages func(ctx context.Context, in TypedSubAgentInput[M]) ([]M, error)
	formatForkResult  func(ctx context.Context, in TypedSubAgentOutput[M]) (string, error)
}

type descriptionTemplateHelper struct {
	Matters []FrontMatter
}

func (s *typedSkillTool[M]) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type inputArguments struct {
	Skill string `json:"skill"`
}

func (s *typedSkillTool[M]) InvokableRun(ctx context.Context, argumentsInJSON string, _ ...tool.Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *typedSkillTool[M]) setActiveModel(ctx context.Context, modelName string) {
	_ = "STUB: not implemented"
	return
}

func defaultToolParams() map[string]*schema.ParameterInfo { _ = "STUB: not implemented"; return nil }

func (s *typedSkillTool[M]) buildParamsOneOf(ctx context.Context) (*schema.ParamsOneOf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *typedSkillTool[M]) buildSkillResult(ctx context.Context, skill Skill, rawArguments string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *typedSkillTool[M]) defaultSkillContent(skill Skill) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *typedSkillTool[M]) runAgentMode(ctx context.Context, skill Skill, forkHistory bool, rawArguments string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func isNilMessage[M adk.MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

func (s *typedSkillTool[M]) getMessagesFromState(ctx context.Context) ([]M, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func renderToolDescription(matters []FrontMatter) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
