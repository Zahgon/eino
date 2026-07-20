package adk

import "sync"

type turnBuffer[T any] struct {
	buf      []T
	mu       sync.Mutex
	notEmpty *sync.Cond
	closed   bool
	woken    bool
}

func newTurnBuffer[T any]() *turnBuffer[T] { _ = "STUB: not implemented"; return nil }

func (tb *turnBuffer[T]) Send(value T) { _ = "STUB: not implemented"; return }

func (tb *turnBuffer[T]) TrySend(value T) bool { _ = "STUB: not implemented"; return false }

func (tb *turnBuffer[T]) Receive() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (tb *turnBuffer[T]) Close() { _ = "STUB: not implemented"; return }

func (tb *turnBuffer[T]) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (tb *turnBuffer[T]) TakeAll() []T { _ = "STUB: not implemented"; return nil }

func (tb *turnBuffer[T]) PushFront(values []T) { _ = "STUB: not implemented"; return }

func (tb *turnBuffer[T]) Wakeup() { _ = "STUB: not implemented"; return }

func (tb *turnBuffer[T]) ClearWakeup() { _ = "STUB: not implemented"; return }
