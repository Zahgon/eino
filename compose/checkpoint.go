package compose

import (
	"context"

	"github.com/cloudwego/eino/internal/core"
	"github.com/cloudwego/eino/internal/serialization"
	"github.com/cloudwego/eino/schema"
)

func init() {
	schema.RegisterName[*checkpoint]("_eino_checkpoint")
	schema.RegisterName[*dagChannel]("_eino_dag_channel")
	schema.RegisterName[*pregelChannel]("_eino_pregel_channel")
	schema.RegisterName[dependencyState]("_eino_dependency_state")
	_ = serialization.GenericRegister[channel]("_eino_channel")
}

func RegisterSerializableType[T any](name string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type CheckPointStore = core.CheckPointStore

type Serializer interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

func WithCheckPointStore(store CheckPointStore) GraphCompileOption {
	_ = "STUB: not implemented"
	return *new(GraphCompileOption)
}

func WithSerializer(serializer Serializer) GraphCompileOption {
	_ = "STUB: not implemented"
	return *new(GraphCompileOption)
}

func WithCheckPointID(checkPointID string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWriteToCheckPointID(checkPointID string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithForceNewRun() Option { _ = "STUB: not implemented"; return *new(Option) }

type StateModifier func(ctx context.Context, path NodePath, state any) error

func WithStateModifier(sm StateModifier) Option { _ = "STUB: not implemented"; return *new(Option) }

type checkpoint struct {
	Channels       map[string]channel
	Inputs         map[string]any
	State          any
	SkipPreHandler map[string]bool
	RerunNodes     []string

	SubGraphs map[string]*checkpoint

	InterruptID2Addr  map[string]Address
	InterruptID2State map[string]core.InterruptState
}

type stateModifierKey struct{}
type checkPointKey struct{}

func getStateModifier(ctx context.Context) StateModifier {
	_ = "STUB: not implemented"
	return *new(StateModifier)
}

func setStateModifier(ctx context.Context, modifier StateModifier) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getCheckPointFromStore(ctx context.Context, id string, cpr *checkPointer) (cp *checkpoint, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setCheckPointToCtx(ctx context.Context, cp *checkpoint) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getCheckPointFromCtx(ctx context.Context) *checkpoint { _ = "STUB: not implemented"; return nil }

func forwardCheckPoint(ctx context.Context, nodeKey string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func newCheckPointer(
	inputPairs, outputPairs map[string]streamConvertPair,
	store CheckPointStore,
	serializer Serializer,
) *checkPointer {
	_ = "STUB: not implemented"
	return nil
}

type checkPointer struct {
	sc         *streamConverter
	store      CheckPointStore
	serializer Serializer
}

func (c *checkPointer) get(ctx context.Context, id string) (*checkpoint, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (c *checkPointer) set(ctx context.Context, id string, cp *checkpoint) error {
	_ = "STUB: not implemented"
	return nil
}

func normalizeCheckpointTypedNilInputs(cp *checkpoint) { _ = "STUB: not implemented"; return }

func isTypedNil(v any) bool { _ = "STUB: not implemented"; return false }

func MigrateCheckpointState(data []byte, serializer Serializer, migrate func(state any) (any, bool, error)) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func migrateCheckpoint(cp *checkpoint, migrate func(state any) (any, bool, error)) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *checkPointer) convertCheckPoint(cp *checkpoint, isStream bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *checkPointer) restoreCheckPoint(cp *checkpoint, isStream bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func newStreamConverter(inputPairs, outputPairs map[string]streamConvertPair) *streamConverter {
	_ = "STUB: not implemented"
	return nil
}

type streamConverter struct {
	inputPairs, outputPairs map[string]streamConvertPair
}

func (s *streamConverter) convertInputs(isStream bool, values map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *streamConverter) restoreInputs(isStream bool, values map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *streamConverter) convertOutputs(isStream bool, values map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *streamConverter) restoreOutputs(isStream bool, values map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func convert(values map[string]any, convPairs map[string]streamConvertPair, isStream bool) error {
	_ = "STUB: not implemented"
	return nil
}

func restore(values map[string]any, convPairs map[string]streamConvertPair, isStream bool) error {
	_ = "STUB: not implemented"
	return nil
}
