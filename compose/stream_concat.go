package compose

import (
	"errors"

	"github.com/cloudwego/eino/schema"
)

func RegisterStreamChunkConcatFunc[T any](fn func([]T) (T, error)) {
	_ = "STUB: not implemented"
	return
}

var emptyStreamConcatErr = errors.New("stream reader is empty, concat fail")

func concatStreamReader[T any](sr *schema.StreamReader[T]) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
