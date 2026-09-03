package adk

import (
	"context"

	"github.com/cloudwego/eino/components"
	"github.com/cloudwego/eino/internal/core"
	"github.com/cloudwego/eino/schema"
)

const ComponentOfAgent components.Component = "Agent"

const ComponentOfAgenticAgent components.Component = "AgenticAgent"

type MessageType interface {
	*schema.Message | *schema.AgenticMessage
}

type Message = *schema.Message
type MessageStream = *schema.StreamReader[Message]

type AgenticMessage = *schema.AgenticMessage
type AgenticMessageStream = *schema.StreamReader[AgenticMessage]

func isNilMessage[M MessageType](msg M) bool { _ = "STUB: not implemented"; return false }

type TypedMessageVariant[M MessageType] struct {
	IsStreaming bool

	Message       M
	MessageStream *schema.StreamReader[M]

	Role schema.RoleType

	AgenticRole schema.AgenticRoleType

	ToolName string
}

func (mv *TypedMessageVariant[M]) GetMessage() (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

type MessageVariant = TypedMessageVariant[*schema.Message]

type messageVariantSerialization struct {
	IsStreaming   bool
	Message       Message
	MessageStream Message
	Role          schema.RoleType
	ToolName      string
}

type agenticMessageVariantSerialization struct {
	IsStreaming   bool
	Message       *schema.AgenticMessage
	MessageStream *schema.AgenticMessage
	Role          schema.RoleType
	AgenticRole   schema.AgenticRoleType
	ToolName      string
}

func (mv *TypedMessageVariant[M]) GobEncode() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mv *TypedMessageVariant[M]) GobDecode(b []byte) error { _ = "STUB: not implemented"; return nil }

func gobEncodeMessageVariant(mv *TypedMessageVariant[*schema.Message]) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func gobDecodeMessageVariant(mv *TypedMessageVariant[*schema.Message], b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func gobEncodeAgenticMessageVariant(mv *TypedMessageVariant[*schema.AgenticMessage]) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func gobDecodeAgenticMessageVariant(mv *TypedMessageVariant[*schema.AgenticMessage], b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func typedEventFromMessage[M MessageType](msg M, msgStream *schema.StreamReader[M],
	role schema.RoleType, toolName string) *TypedAgentEvent[M] {
	_ = "STUB: not implemented"
	return nil
}

func typedModelOutputEvent[M MessageType](msg M, msgStream *schema.StreamReader[M]) *TypedAgentEvent[M] {
	_ = "STUB: not implemented"
	return nil
}

func EventFromMessage(msg Message, msgStream *schema.StreamReader[Message],
	role schema.RoleType, toolName string) *AgentEvent {
	_ = "STUB: not implemented"
	return nil
}

func EventFromAgenticMessage(msg AgenticMessage, msgStream AgenticMessageStream, agenticRole schema.AgenticRoleType) *TypedAgentEvent[AgenticMessage] {
	_ = "STUB: not implemented"
	return nil
}

type TransferToAgentAction struct {
	DestAgentName string
}

type TypedAgentOutput[M MessageType] struct {
	MessageOutput *TypedMessageVariant[M]

	CustomizedOutput any
}

type AgentOutput = TypedAgentOutput[*schema.Message]

func NewTransferToAgentAction(destAgentName string) *AgentAction {
	_ = "STUB: not implemented"
	return nil
}

func NewExitAction() *AgentAction { _ = "STUB: not implemented"; return nil }

type AgentAction struct {
	Exit bool

	Interrupted *InterruptInfo

	TransferToAgent *TransferToAgentAction

	BreakLoop *BreakLoopAction

	CustomizedAction any

	internalInterrupted *core.InterruptSignal
}

type RunStep struct {
	agentName string
}

func init() {
	schema.RegisterName[[]RunStep]("eino_run_step_list")
}

func (r *RunStep) String() string { _ = "STUB: not implemented"; return "" }

func (r *RunStep) Equals(r1 RunStep) bool { _ = "STUB: not implemented"; return false }

func (r *RunStep) GobEncode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *RunStep) GobDecode(b []byte) error { _ = "STUB: not implemented"; return nil }

type runStepSerialization struct {
	AgentName string
}

type TypedAgentEvent[M MessageType] struct {
	AgentName string

	RunPath []RunStep

	Output *TypedAgentOutput[M]

	Action *AgentAction

	Err error
}

type AgentEvent = TypedAgentEvent[*schema.Message]

type TypedAgentInput[M MessageType] struct {
	Messages        []M
	EnableStreaming bool
}

type AgentInput = TypedAgentInput[*schema.Message]

type TypedAgent[M MessageType] interface {
	Name(ctx context.Context) string
	Description(ctx context.Context) string

	Run(ctx context.Context, input *TypedAgentInput[M], options ...AgentRunOption) *AsyncIterator[*TypedAgentEvent[M]]
}

//go:generate  mockgen -destination ../internal/mock/adk/Agent_mock.go --package adk github.com/cloudwego/eino/adk Agent,ResumableAgent
type Agent = TypedAgent[*schema.Message]

type OnSubAgents interface {
	OnSetSubAgents(ctx context.Context, subAgents []Agent) error
	OnSetAsSubAgent(ctx context.Context, parent Agent) error

	OnDisallowTransferToParent(ctx context.Context) error
}

type TypedResumableAgent[M MessageType] interface {
	TypedAgent[M]

	Resume(ctx context.Context, info *ResumeInfo, opts ...AgentRunOption) *AsyncIterator[*TypedAgentEvent[M]]
}

type ResumableAgent = TypedResumableAgent[*schema.Message]

func concatMessageStream[M MessageType](stream *schema.StreamReader[M]) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}
