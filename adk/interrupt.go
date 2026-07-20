package adk

import (
	"context"
	"encoding/gob"
	"sync"

	"github.com/cloudwego/eino/internal/core"
	"github.com/cloudwego/eino/schema"
)

type ResumeInfo struct {
	EnableStreaming bool

	*InterruptInfo

	WasInterrupted bool
	InterruptState any
	IsResumeTarget bool
	ResumeData     any
}

type InterruptInfo struct {
	Data any

	InterruptContexts []*InterruptCtx
}

func TypedInterrupt[M MessageType](ctx context.Context, info any) *TypedAgentEvent[M] {
	_ = "STUB: not implemented"
	return nil
}

func Interrupt(ctx context.Context, info any) *AgentEvent { _ = "STUB: not implemented"; return nil }

func TypedStatefulInterrupt[M MessageType](ctx context.Context, info any, state any) *TypedAgentEvent[M] {
	_ = "STUB: not implemented"
	return nil
}

func StatefulInterrupt(ctx context.Context, info any, state any) *AgentEvent {
	_ = "STUB: not implemented"
	return nil
}

func TypedCompositeInterrupt[M MessageType](ctx context.Context, info any, state any,
	subInterruptSignals ...*InterruptSignal) *TypedAgentEvent[M] {
	_ = "STUB: not implemented"
	return nil
}

func CompositeInterrupt(ctx context.Context, info any, state any,
	subInterruptSignals ...*InterruptSignal) *AgentEvent {
	_ = "STUB: not implemented"
	return nil
}

type Address = core.Address
type AddressSegment = core.AddressSegment
type AddressSegmentType = core.AddressSegmentType

const (
	AddressSegmentAgent AddressSegmentType = "agent"
	AddressSegmentTool  AddressSegmentType = "tool"
)

var allowedAddressSegmentTypes = []AddressSegmentType{AddressSegmentAgent, AddressSegmentTool}

func AppendAddressSegment(ctx context.Context, segType AddressSegmentType, segID string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type InterruptCtx = core.InterruptCtx
type InterruptSignal = core.InterruptSignal

func FromInterruptContexts(contexts []*InterruptCtx) *InterruptSignal {
	_ = "STUB: not implemented"
	return nil
}

func WithCheckPointID(id string) AgentRunOption {
	_ = "STUB: not implemented"
	return *new(AgentRunOption)
}

func init() {
	schema.RegisterName[*serialization]("_eino_adk_serialization")
	schema.RegisterName[*WorkflowInterruptInfo]("_eino_adk_workflow_interrupt_info")

	gob.Register([]byte{})
}

type serialization struct {
	RunCtx *runContext

	Info                *InterruptInfo
	EnableStreaming     bool
	InterruptID2Address map[string]Address
	InterruptID2State   map[string]core.InterruptState
}

func runnerLoadCheckPointImpl(store CheckPointStore, ctx context.Context, checkpointID string) (
	context.Context, *runContext, *ResumeInfo, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil, nil
}

func preprocessADKCheckpoint(data []byte) []byte { _ = "STUB: not implemented"; return nil }

func runnerSaveCheckPointImpl(
	enableStreaming bool,
	store CheckPointStore,
	ctx context.Context,
	key string,
	info *InterruptInfo,
	is *core.InterruptSignal,
) error {
	_ = "STUB: not implemented"
	return nil
}

const bridgeCheckpointID = "adk_react_mock_key"

func newBridgeStore() *bridgeStore { _ = "STUB: not implemented"; return nil }

func newResumeBridgeStore(checkPointID string, data []byte) *bridgeStore {
	_ = "STUB: not implemented"
	return nil
}

type bridgeStore struct {
	mu   sync.Mutex
	data map[string][]byte
}

func (m *bridgeStore) Get(_ context.Context, key string) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *bridgeStore) Set(_ context.Context, key string, checkPoint []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func getNextResumeAgent(ctx context.Context, _ *ResumeInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getNextResumeAgents(ctx context.Context, _ *ResumeInfo) (map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildResumeInfo(ctx context.Context, nextAgentID string, info *ResumeInfo) (
	context.Context, *ResumeInfo) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
