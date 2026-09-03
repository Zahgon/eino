package compose

import (
	"reflect"

	"github.com/cloudwego/eino/schema"
)

type streamReader interface {
	copy(n int) []streamReader
	getType() reflect.Type
	getChunkType() reflect.Type
	merge([]streamReader) streamReader
	withKey(string) streamReader
	close()
	toAnyStreamReader() *schema.StreamReader[any]
	mergeWithNames([]streamReader, []string) streamReader
}

type streamReaderPacker[T any] struct {
	sr *schema.StreamReader[T]
}

func (srp streamReaderPacker[T]) close() { _ = "STUB: not implemented"; return }

func (srp streamReaderPacker[T]) copy(n int) []streamReader { _ = "STUB: not implemented"; return nil }

func (srp streamReaderPacker[T]) getType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (srp streamReaderPacker[T]) getChunkType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (srp streamReaderPacker[T]) toStreamReaders(srs []streamReader) []*schema.StreamReader[T] {
	_ = "STUB: not implemented"
	return nil
}

func (srp streamReaderPacker[T]) merge(isrs []streamReader) streamReader {
	_ = "STUB: not implemented"
	return *new(streamReader)
}

func (srp streamReaderPacker[T]) mergeWithNames(isrs []streamReader, names []string) streamReader {
	_ = "STUB: not implemented"
	return *new(streamReader)
}

func (srp streamReaderPacker[T]) withKey(key string) streamReader {
	_ = "STUB: not implemented"
	return *new(streamReader)
}

func (srp streamReaderPacker[T]) toAnyStreamReader() *schema.StreamReader[any] {
	_ = "STUB: not implemented"
	return nil
}

func packStreamReader[T any](sr *schema.StreamReader[T]) streamReader {
	_ = "STUB: not implemented"
	return *new(streamReader)
}

func unpackStreamReader[T any](isr streamReader) (*schema.StreamReader[T], bool) {
	_ = "STUB: not implemented"
	return nil, false
}
