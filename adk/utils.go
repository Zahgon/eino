package adk

import (
	"context"

	"github.com/cloudwego/eino/internal"
)

type AsyncIterator[T any] struct {
	ch *internal.UnboundedChan[T]
}

func (ai *AsyncIterator[T]) Next() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

type AsyncGenerator[T any] struct {
	ch *internal.UnboundedChan[T]
}

func (ag *AsyncGenerator[T]) Send(v T) { _ = "STUB: not implemented"; return }

func (ag *AsyncGenerator[T]) trySend(v T) bool { _ = "STUB: not implemented"; return false }

func (ag *AsyncGenerator[T]) Close() { _ = "STUB: not implemented"; return }

func NewAsyncIteratorPair[T any]() (*AsyncIterator[T], *AsyncGenerator[T]) {
	_ = "STUB: not implemented"
	return nil, nil
}

func copyMap[K comparable, V any](m map[K]V) map[K]V { _ = "STUB: not implemented"; return nil }

func cloneSlice[T any](s []T) []T { _ = "STUB: not implemented"; return nil }

func concatInstructions(instructions ...string) string { _ = "STUB: not implemented"; return "" }

func GenTransferMessages(_ context.Context, destAgentName string) (Message, Message) {
	_ = "STUB: not implemented"
	return *new(Message), *new(Message)
}

func typedSetAutomaticClose[M MessageType](e *TypedAgentEvent[M]) {
	_ = "STUB: not implemented"
	return
}

func setAutomaticClose(e *AgentEvent) { _ = "STUB: not implemented"; return }

func getMessageFromTypedWrappedEvent[M MessageType](e *typedAgentEventWrapper[M]) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}

func getMessageFromWrappedEvent(e *agentEventWrapper) (Message, error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}

func (e *agentEventWrapper) consumeStream() { _ = "STUB: not implemented"; return }

func copyTypedAgentEvent[M MessageType](ae *TypedAgentEvent[M]) *TypedAgentEvent[M] {
	_ = "STUB: not implemented"
	return nil
}

func TypedGetMessage[M MessageType](e *TypedAgentEvent[M]) (M, *TypedAgentEvent[M], error) {
	_ = "STUB: not implemented"
	return *new(M), nil, nil
}

func GetMessage(e *AgentEvent) (Message, *AgentEvent, error) {
	_ = "STUB: not implemented"
	return *new(Message), nil, nil
}

func typedErrorIter[M MessageType](err error) *AsyncIterator[*TypedAgentEvent[M]] {
	_ = "STUB: not implemented"
	return nil
}

func genErrorIter(err error) *AsyncIterator[*AgentEvent] { _ = "STUB: not implemented"; return nil }
