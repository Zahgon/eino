package adk

import (
	"context"
	"encoding/gob"
	"errors"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

var ErrExceedMaxIterations = errors.New("exceeds max iterations")

type typedState[M MessageType] struct {
	Messages []M
	Extra    map[string]any

	ToolInfos []*schema.ToolInfo

	DeferredToolInfos []*schema.ToolInfo

	HasReturnDirectly        bool
	ReturnDirectlyToolCallID string
	ToolGenActions           map[string]*AgentAction
	AgentName                string
	RemainingIterations      int
	ReturnDirectlyEvent      *TypedAgentEvent[M]
	RetryAttempt             int
	ToolMsgIDs               map[string]map[string]string
}

type State = typedState[*schema.Message]

type agenticState = typedState[*schema.AgenticMessage]

const (
	stateGobNameV07 = "_eino_adk_react_state"

	stateGobNameV080 = "_eino_adk_state_v080_"
)

func init() {

	schema.RegisterName[*State](stateGobNameV07)
	schema.RegisterName[*stateV080](stateGobNameV080)

	schema.RegisterName[*typedState[*schema.AgenticMessage]]("_eino_adk_agentic_state")
	schema.RegisterName[*TypedAgentEvent[*schema.AgenticMessage]]("_eino_adk_agentic_event")

	gob.Register(&AgentEvent{})
	gob.Register(0)

	schema.RegisterName[*TypedAgentInput[*schema.AgenticMessage]]("_eino_adk_agentic_agent_input")
	schema.RegisterName[*typedAgentEventWrapper[*schema.AgenticMessage]]("_eino_adk_agentic_event_wrapper")
	schema.RegisterName[*[]*typedAgentEventWrapper[*schema.AgenticMessage]]("_eino_adk_agentic_event_wrapper_slice")
	schema.RegisterName[*reactInput]("_eino_adk_react_input")
	schema.RegisterName[*agenticReactInput]("_eino_adk_agentic_react_input")
}

func (s *typedState[M]) getReturnDirectlyEvent() *TypedAgentEvent[M] {
	_ = "STUB: not implemented"
	return nil
}

func (s *typedState[M]) setReturnDirectlyEvent(event *TypedAgentEvent[M]) {
	_ = "STUB: not implemented"
	return
}

func (s *typedState[M]) getRetryAttempt() int { _ = "STUB: not implemented"; return 0 }

func (s *typedState[M]) setRetryAttempt(attempt int) { _ = "STUB: not implemented"; return }

func (s *typedState[M]) getReturnDirectlyToolCallID() string { _ = "STUB: not implemented"; return "" }

func (s *typedState[M]) setReturnDirectlyToolCallID(id string) { _ = "STUB: not implemented"; return }

func (s *typedState[M]) getToolGenActions() map[string]*AgentAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *typedState[M]) setToolGenAction(key string, action *AgentAction) {
	_ = "STUB: not implemented"
	return
}

func (s *typedState[M]) popToolGenAction(key string) *AgentAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *typedState[M]) setToolMsgID(toolName, callID, msgID string) {
	_ = "STUB: not implemented"
	return
}

func (s *typedState[M]) popToolMsgID(toolName, callID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *typedState[M]) getRemainingIterations() int { _ = "STUB: not implemented"; return 0 }

func (s *typedState[M]) setRemainingIterations(iterations int) { _ = "STUB: not implemented"; return }

func (s *typedState[M]) decrementRemainingIterations() { _ = "STUB: not implemented"; return }

type stateV080 struct {
	Messages                 []Message
	HasReturnDirectly        bool
	ReturnDirectlyToolCallID string
	ToolGenActions           map[string]*AgentAction
	AgentName                string
	RemainingIterations      int
	RetryAttempt             int
	ReturnDirectlyEvent      *AgentEvent
	Extra                    map[string]any
	Internals                map[string]any
}

type stateV080Serialization stateV080

func (sc *stateV080) GobDecode(b []byte) error { _ = "STUB: not implemented"; return nil }

func stateV080ToState(sc *stateV080) *State { _ = "STUB: not implemented"; return nil }

func SendToolGenAction(ctx context.Context, toolName string, action *AgentAction) error {
	_ = "STUB: not implemented"
	return nil
}

type reactInput struct {
	Messages []Message
}

type typedReactConfig[M MessageType] struct {
	model model.BaseModel[M]

	toolsConfig      *compose.ToolsNodeConfig
	modelWrapperConf *typedModelWrapperConfig[M]

	toolsReturnDirectly map[string]bool

	agentName string

	maxIterations int

	cancelCtx *cancelContext

	afterAgentFunc func(ctx context.Context, msg M) (M, error)
}

type reactConfig = typedReactConfig[*schema.Message]

func genToolInfos(ctx context.Context, config *compose.ToolsNodeConfig) ([]*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type reactGraph = *compose.Graph[*reactInput, Message]

func getReturnDirectlyToolCallID(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func genReactState(config *reactConfig) func(ctx context.Context) *State {
	_ = "STUB: not implemented"
	return nil
}

func newReact(ctx context.Context, config *reactConfig) (reactGraph, error) {
	_ = "STUB: not implemented"
	return *new(reactGraph), nil
}

type agenticReactInput struct {
	Messages []*schema.AgenticMessage
}

type agenticReactConfig = typedReactConfig[*schema.AgenticMessage]

type agenticReactGraph = *compose.Graph[*agenticReactInput, *schema.AgenticMessage]

func getAgenticReturnDirectlyToolCallID(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func genAgenticReactState(config *agenticReactConfig) func(ctx context.Context) *agenticState {
	_ = "STUB: not implemented"
	return nil
}

func agenticMessageHasToolCalls(msg *schema.AgenticMessage) bool {
	_ = "STUB: not implemented"
	return false
}

func newAgenticReact(ctx context.Context, config *agenticReactConfig) (agenticReactGraph, error) {
	_ = "STUB: not implemented"
	return *new(agenticReactGraph), nil
}

func extractToolIdentifiers(msg *schema.AgenticMessage) (toolName, callID string) {
	_ = "STUB: not implemented"
	return "", ""
}
