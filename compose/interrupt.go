package compose

import (
	"context"
	"errors"

	"github.com/cloudwego/eino/internal/core"
	"github.com/cloudwego/eino/schema"
)

func WithInterruptBeforeNodes(nodes []string) GraphCompileOption {
	_ = "STUB: not implemented"
	return *new(GraphCompileOption)
}

func WithInterruptAfterNodes(nodes []string) GraphCompileOption {
	_ = "STUB: not implemented"
	return *new(GraphCompileOption)
}

var InterruptAndRerun = deprecatedInterruptAndRerun
var deprecatedInterruptAndRerun = errors.New("interrupt and rerun")

func NewInterruptAndRerunErr(extra any) error { _ = "STUB: not implemented"; return nil }

func deprecatedInterruptAndRerunErr(extra any) error { _ = "STUB: not implemented"; return nil }

type wrappedInterruptAndRerun struct {
	ps    Address
	inner error
}

func (w *wrappedInterruptAndRerun) Error() string { _ = "STUB: not implemented"; return "" }

func (w *wrappedInterruptAndRerun) Unwrap() error { _ = "STUB: not implemented"; return nil }

func WrapInterruptAndRerunIfNeeded(ctx context.Context, step AddressSegment, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func Interrupt(ctx context.Context, info any) error { _ = "STUB: not implemented"; return nil }

func StatefulInterrupt(ctx context.Context, info any, state any) error {
	_ = "STUB: not implemented"
	return nil
}

func CompositeInterrupt(ctx context.Context, info any, state any, errs ...error) error {
	_ = "STUB: not implemented"
	return nil
}

func IsInterruptRerunError(err error) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func isInterruptRerunError(err error) (info any, state any, ok bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(any), false
}

type InterruptInfo struct {
	State             any
	BeforeNodes       []string
	AfterNodes        []string
	RerunNodes        []string
	RerunNodesExtra   map[string]any
	SubGraphs         map[string]*InterruptInfo
	InterruptContexts []*InterruptCtx
}

func init() {
	schema.RegisterName[*InterruptInfo]("_eino_compose_interrupt_info")
}

type AddressSegmentType = core.AddressSegmentType

const (
	AddressSegmentNode AddressSegmentType = "node"

	AddressSegmentTool AddressSegmentType = "tool"

	AddressSegmentRunnable AddressSegmentType = "runnable"
)

type Address = core.Address

type AddressSegment = core.AddressSegment

type InterruptCtx = core.InterruptCtx

func ExtractInterruptInfo(err error) (info *InterruptInfo, existed bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type interruptError struct {
	Info *InterruptInfo
}

func (e *interruptError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *interruptError) GetInterruptContexts() []*InterruptCtx {
	_ = "STUB: not implemented"
	return nil
}

func isSubGraphInterrupt(err error) *subGraphInterruptError { _ = "STUB: not implemented"; return nil }

type subGraphInterruptError struct {
	Info       *InterruptInfo
	CheckPoint *checkpoint

	signal *core.InterruptSignal
}

func (e *subGraphInterruptError) Error() string { _ = "STUB: not implemented"; return "" }

func isInterruptError(err error) bool { _ = "STUB: not implemented"; return false }
