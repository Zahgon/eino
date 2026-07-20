package adk

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

var (
	ErrExceedMaxRetries = errors.New("exceeds max retries")
)

type RetryExhaustedError struct {
	LastErr      error
	TotalRetries int
}

func (e *RetryExhaustedError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *RetryExhaustedError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type WillRetryError struct {
	ErrStr       string
	RetryAttempt int
	rejectReason any
	err          error
}

func (e *WillRetryError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *WillRetryError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *WillRetryError) RejectReason() any { _ = "STUB: not implemented"; return *new(any) }

func init() {
	schema.RegisterName[*WillRetryError]("eino_adk_chatmodel_will_retry_error")
}

type TypedRetryContext[M MessageType] struct {
	RetryAttempt int

	InputMessages []M

	Options []model.Option

	OutputMessage M

	Err error
}

type RetryContext = TypedRetryContext[*schema.Message]

type TypedRetryDecision[M MessageType] struct {
	Retry bool

	RewriteError error

	ModifiedInputMessages []M

	PersistModifiedInputMessages bool

	AdditionalOptions []model.Option

	Backoff time.Duration

	RejectReason any
}

type RetryDecision = TypedRetryDecision[*schema.Message]

type TypedModelRetryConfig[M MessageType] struct {
	MaxRetries int

	ShouldRetry func(ctx context.Context, retryCtx *TypedRetryContext[M]) *TypedRetryDecision[M]

	IsRetryAble func(ctx context.Context, err error) bool

	BackoffFunc func(ctx context.Context, attempt int) time.Duration
}

type ModelRetryConfig = TypedModelRetryConfig[*schema.Message]

func defaultIsRetryAble(_ context.Context, err error) bool { _ = "STUB: not implemented"; return false }

func defaultBackoff(_ context.Context, attempt int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func genErrWrapper(ctx context.Context, maxRetries, attempt int, isRetryAbleFunc func(ctx context.Context, err error) bool) func(error) error {
	_ = "STUB: not implemented"
	return nil
}

func consumeStreamForError[M any](stream *schema.StreamReader[M]) error {
	_ = "STUB: not implemented"
	return nil
}

type retryVerdictSignal struct {
	ch chan retryVerdict
}

type retryVerdict struct {
	WillRetry    bool
	RetryAttempt int
	Err          error
	RejectReason any
}

type typedRetryModelWrapper[M MessageType] struct {
	inner  model.BaseModel[M]
	config *TypedModelRetryConfig[M]
}

func newTypedRetryModelWrapper[M MessageType](inner model.BaseModel[M], config *TypedModelRetryConfig[M]) *typedRetryModelWrapper[M] {
	_ = "STUB: not implemented"
	return nil
}

func (r *typedRetryModelWrapper[M]) Generate(ctx context.Context, input []M, opts ...model.Option) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func (r *typedRetryModelWrapper[M]) generateLegacy(ctx context.Context, input []M, opts ...model.Option) (zero M, _ error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func generateWithShouldRetry[M MessageType](r *typedRetryModelWrapper[M], ctx context.Context, input []M, opts ...model.Option) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func (r *typedRetryModelWrapper[M]) contextAwareSleep(ctx context.Context, delay time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func streamWithShouldRetry[M MessageType](r *typedRetryModelWrapper[M], ctx context.Context, input []M, opts ...model.Option) (
	*schema.StreamReader[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyDecisionForRetry[M MessageType](currentInput *[]M, currentOpts *[]model.Option, ctx context.Context, decision *TypedRetryDecision[M]) {
	_ = "STUB: not implemented"
	return
}

func (r *typedRetryModelWrapper[M]) Stream(ctx context.Context, input []M, opts ...model.Option) (
	*schema.StreamReader[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *typedRetryModelWrapper[M]) streamLegacy(ctx context.Context, input []M, opts ...model.Option) (
	*schema.StreamReader[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
