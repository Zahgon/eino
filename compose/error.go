package compose

import (
	"errors"
	"reflect"
)

var ErrExceedMaxSteps = errors.New("exceeds max steps")

func newUnexpectedInputTypeErr(expected reflect.Type, got reflect.Type) error {
	_ = "STUB: not implemented"
	return nil
}

type defaultImplAction string

const (
	actionInvokeByStream     defaultImplAction = "InvokeByStream"
	actionInvokeByCollect    defaultImplAction = "InvokeByCollect"
	actionInvokeByTransform  defaultImplAction = "InvokeByTransform"
	actionStreamByInvoke     defaultImplAction = "StreamByInvoke"
	actionStreamByTransform  defaultImplAction = "StreamByTransform"
	actionStreamByCollect    defaultImplAction = "StreamByCollect"
	actionCollectByTransform defaultImplAction = "CollectByTransform"
	actionCollectByInvoke    defaultImplAction = "CollectByInvoke"
	actionCollectByStream    defaultImplAction = "CollectByStream"
	actionTransformByStream  defaultImplAction = "TransformByStream"
	actionTransformByCollect defaultImplAction = "TransformByCollect"
	actionTransformByInvoke  defaultImplAction = "TransformByInvoke"
)

func newStreamReadError(err error) error { _ = "STUB: not implemented"; return nil }

func newGraphRunError(err error) error { _ = "STUB: not implemented"; return nil }

func wrapGraphNodeError(nodeKey string, err error) error { _ = "STUB: not implemented"; return nil }

type internalErrorType string

const (
	internalErrorTypeNodeRun  = "NodeRunError"
	internalErrorTypeGraphRun = "GraphRunError"
)

type internalError struct {
	typ       internalErrorType
	nodePath  NodePath
	origError error
}

func (i *internalError) Error() string { _ = "STUB: not implemented"; return "" }

func (i *internalError) Unwrap() error { _ = "STUB: not implemented"; return nil }
