package planexecute

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func init() {
	schema.RegisterName[*defaultPlan]("_eino_adk_plan_execute_default_plan")
	schema.RegisterName[ExecutedStep]("_eino_adk_plan_execute_executed_step")
	schema.RegisterName[[]ExecutedStep]("_eino_adk_plan_execute_executed_steps")
}

type Plan interface {
	FirstStep() string

	json.Marshaler

	json.Unmarshaler
}

type NewPlan func(ctx context.Context) Plan

type defaultPlan struct {
	Steps []string `json:"steps"`
}

func (p *defaultPlan) FirstStep() string { _ = "STUB: not implemented"; return "" }

func (p *defaultPlan) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *defaultPlan) UnmarshalJSON(bytes []byte) error { _ = "STUB: not implemented"; return nil }

type Response struct {
	Response string `json:"response"`
}

var (
	PlanToolInfo = schema.ToolInfo{
		Name: "plan",
		Desc: "Plan with a list of steps to execute in order. Each step should be clear, actionable, and arranged in a logical sequence. The output will be used to guide the execution process.",
		ParamsOneOf: schema.NewParamsOneOfByParams(
			map[string]*schema.ParameterInfo{
				"steps": {
					Type:     schema.Array,
					ElemInfo: &schema.ParameterInfo{Type: schema.String},
					Desc:     "different steps to follow, should be in sorted order",
					Required: true,
				},
			},
		),
	}

	RespondToolInfo = schema.ToolInfo{
		Name: "respond",
		Desc: "Generate a direct response to the user. Use this tool when you have all the information needed to provide a final answer.",
		ParamsOneOf: schema.NewParamsOneOfByParams(
			map[string]*schema.ParameterInfo{
				"response": {
					Type:     schema.String,
					Desc:     "The complete response to provide to the user",
					Required: true,
				},
			},
		),
	}

	PlannerPrompt = prompt.FromMessages(schema.FString,
		schema.SystemMessage(`You are an expert planning agent. Given an objective, create a comprehensive step-by-step plan to achieve the objective.

## YOUR TASK
Analyze the objective and generate a strategic plan that breaks down the goal into manageable, executable steps.

## PLANNING REQUIREMENTS
Each step in your plan must be:
- **Specific and actionable**: Clear instructions that can be executed without ambiguity
- **Self-contained**: Include all necessary context, parameters, and requirements
- **Independently executable**: Can be performed by an agent without dependencies on other steps
- **Logically sequenced**: Arranged in optimal order for efficient execution
- **Objective-focused**: Directly contribute to achieving the main goal

## PLANNING GUIDELINES
- Eliminate redundant or unnecessary steps
- Include relevant constraints, parameters, and success criteria for each step
- Ensure the final step produces a complete answer or deliverable
- Anticipate potential challenges and include mitigation strategies
- Structure steps to build upon each other logically
- Provide sufficient detail for successful execution

## QUALITY CRITERIA
- Plan completeness: Does it address all aspects of the objective?
- Step clarity: Can each step be understood and executed independently?
- Logical flow: Do steps follow a sensible progression?
- Efficiency: Is this the most direct path to the objective?
- Adaptability: Can the plan handle unexpected results or changes?`),
		schema.MessagesPlaceholder("input", false),
	)

	ExecutorPrompt = prompt.FromMessages(schema.FString,
		schema.SystemMessage(`You are a diligent and meticulous executor agent. Follow the given plan and execute your tasks carefully and thoroughly.`),
		schema.UserMessage(`## OBJECTIVE
{input}
## Given the following plan:
{plan}
## COMPLETED STEPS & RESULTS
{executed_steps}
## Your task is to execute the first step, which is: 
{step}`))

	ReplannerPrompt = prompt.FromMessages(schema.FString,
		schema.SystemMessage(
			`You are going to review the progress toward an objective. Analyze the current state and determine the optimal next action.

## YOUR TASK
Based on the progress above, you MUST choose exactly ONE action:

### Option 1: COMPLETE (if objective is fully achieved)
Call '{respond_tool}' with:
- A comprehensive final answer
- Clear conclusion summarizing how the objective was met
- Key insights from the execution process

### Option 2: CONTINUE (if more work is needed)
Call '{plan_tool}' with a revised plan that:
- Contains ONLY remaining steps (exclude completed ones)
- Incorporates lessons learned from executed steps
- Addresses any gaps or issues discovered
- Maintains logical step sequence

## PLANNING REQUIREMENTS
Each step in your plan must be:
- **Specific and actionable**: Clear instructions that can be executed without ambiguity
- **Self-contained**: Include all necessary context, parameters, and requirements
- **Independently executable**: Can be performed by an agent without dependencies on other steps
- **Logically sequenced**: Arranged in optimal order for efficient execution
- **Objective-focused**: Directly contribute to achieving the main goal

## PLANNING GUIDELINES
- Eliminate redundant or unnecessary steps
- Adapt strategy based on new information
- Include relevant constraints, parameters, and success criteria for each step

## DECISION CRITERIA
- Has the original objective been completely satisfied?
- Are there any remaining requirements or sub-goals?
- Do the results suggest a need for strategy adjustment?
- What specific actions are still required?`),
		schema.UserMessage(`## OBJECTIVE
{input}

## ORIGINAL PLAN
{plan}

## COMPLETED STEPS & RESULTS
{executed_steps}`),
	)
)

const (
	UserInputSessionKey = "UserInput"

	PlanSessionKey = "Plan"

	ExecutedStepSessionKey = "ExecutedStep"

	ExecutedStepsSessionKey = "ExecutedSteps"
)

type PlannerConfig struct {
	ChatModelWithFormattedOutput model.BaseChatModel

	ToolCallingChatModel model.ToolCallingChatModel

	ToolInfo *schema.ToolInfo

	GenInputFn GenPlannerModelInputFn

	NewPlan NewPlan
}

type GenPlannerModelInputFn func(ctx context.Context, userInput []adk.Message) ([]adk.Message, error)

func defaultNewPlan(ctx context.Context) Plan { _ = "STUB: not implemented"; return *new(Plan) }

func defaultGenPlannerInputFn(ctx context.Context, userInput []adk.Message) ([]adk.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type planner struct {
	toolCall   bool
	chatModel  model.BaseChatModel
	genInputFn GenPlannerModelInputFn
	newPlan    NewPlan
}

func (p *planner) Name(_ context.Context) string { _ = "STUB: not implemented"; return "" }

func (p *planner) Description(_ context.Context) string { _ = "STUB: not implemented"; return "" }

func argToContent(msg adk.Message) (adk.Message, error) {
	_ = "STUB: not implemented"
	return *new(adk.Message), nil
}

func (p *planner) Run(ctx context.Context, input *adk.AgentInput,
	_ ...adk.AgentRunOption) *adk.AsyncIterator[*adk.AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

func NewPlanner(_ context.Context, cfg *PlannerConfig) (adk.Agent, error) {
	_ = "STUB: not implemented"
	return *new(adk.Agent), nil
}

type ExecutionContext struct {
	UserInput     []adk.Message
	Plan          Plan
	ExecutedSteps []ExecutedStep
}

type GenModelInputFn func(ctx context.Context, in *ExecutionContext) ([]adk.Message, error)

type ExecutorConfig struct {
	Model model.BaseChatModel

	ToolsConfig adk.ToolsConfig

	MaxIterations int

	GenInputFn GenModelInputFn
}

type ExecutedStep struct {
	Step   string
	Result string
}

func NewExecutor(ctx context.Context, cfg *ExecutorConfig) (adk.Agent, error) {
	_ = "STUB: not implemented"
	return *new(adk.Agent), nil
}

func defaultGenExecutorInputFn(ctx context.Context, in *ExecutionContext) ([]adk.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type replanner struct {
	chatModel   model.ToolCallingChatModel
	planTool    *schema.ToolInfo
	respondTool *schema.ToolInfo

	genInputFn GenModelInputFn
	newPlan    NewPlan
}

type ReplannerConfig struct {
	ChatModel model.ToolCallingChatModel

	PlanTool *schema.ToolInfo

	RespondTool *schema.ToolInfo

	GenInputFn GenModelInputFn

	NewPlan NewPlan
}

func formatInput(input []adk.Message) string { _ = "STUB: not implemented"; return "" }

func formatExecutedSteps(results []ExecutedStep) string { _ = "STUB: not implemented"; return "" }

func (r *replanner) Name(_ context.Context) string { _ = "STUB: not implemented"; return "" }

func (r *replanner) Description(_ context.Context) string { _ = "STUB: not implemented"; return "" }

func (r *replanner) genInput(ctx context.Context) ([]adk.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *replanner) Run(ctx context.Context, input *adk.AgentInput, _ ...adk.AgentRunOption) *adk.AsyncIterator[*adk.AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

func buildGenReplannerInputFn(planToolName, respondToolName string) GenModelInputFn {
	_ = "STUB: not implemented"
	return *new(GenModelInputFn)
}

func NewReplanner(_ context.Context, cfg *ReplannerConfig) (adk.Agent, error) {
	_ = "STUB: not implemented"
	return *new(adk.Agent), nil
}

type Config struct {
	Planner adk.Agent

	Executor adk.Agent

	Replanner adk.Agent

	MaxIterations int
}

func New(ctx context.Context, cfg *Config) (adk.ResumableAgent, error) {
	_ = "STUB: not implemented"
	return *new(adk.ResumableAgent), nil
}
