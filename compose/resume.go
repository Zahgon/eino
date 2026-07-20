package compose

import (
	"context"
)

func GetInterruptState[T any](ctx context.Context) (wasInterrupted bool, hasState bool, state T) {
	_ = "STUB: not implemented"
	return false, false, *new(T)
}

func GetResumeContext[T any](ctx context.Context) (isResumeFlow bool, hasData bool, data T) {
	_ = "STUB: not implemented"
	return false, false, *new(T)
}

func GetCurrentAddress(ctx context.Context) Address {
	_ = "STUB: not implemented"
	return *new(Address)
}

func Resume(ctx context.Context, interruptIDs ...string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ResumeWithData(ctx context.Context, interruptID string, data any) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func BatchResumeWithData(ctx context.Context, resumeData map[string]any) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getNodePath(ctx context.Context) (*NodePath, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func AppendAddressSegment(ctx context.Context, segType AddressSegmentType, segID string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func appendToolAddressSegment(ctx context.Context, segID string, subID string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
