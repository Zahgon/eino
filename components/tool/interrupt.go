package tool

import (
	"context"
)

func Interrupt(ctx context.Context, info any) error { _ = "STUB: not implemented"; return nil }

func StatefulInterrupt(ctx context.Context, info any, state any) error {
	_ = "STUB: not implemented"
	return nil
}

func CompositeInterrupt(ctx context.Context, info any, state any, errs ...error) error {
	_ = "STUB: not implemented"
	return nil
}

func GetInterruptState[T any](ctx context.Context) (wasInterrupted bool, hasState bool, state T) {
	_ = "STUB: not implemented"
	return false, false, *new(T)
}

func GetResumeContext[T any](ctx context.Context) (isResumeTarget bool, hasData bool, data T) {
	_ = "STUB: not implemented"
	return false, false, *new(T)
}
