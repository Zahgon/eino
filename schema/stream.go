package schema

import (
	"errors"
	"reflect"
	"sync"
)

var ErrNoValue = errors.New("no value")

var ErrRecvAfterClosed = errors.New("recv after stream closed")

type SourceEOF struct {
	sourceName string
}

func (e *SourceEOF) Error() string { _ = "STUB: not implemented"; return "" }

func GetSourceName(err error) (string, bool) { _ = "STUB: not implemented"; return "", false }

func Pipe[T any](cap int) (*StreamReader[T], *StreamWriter[T]) {
	_ = "STUB: not implemented"
	return nil, nil
}

type StreamWriter[T any] struct {
	stm *stream[T]
}

func (sw *StreamWriter[T]) Send(chunk T, err error) (closed bool) {
	_ = "STUB: not implemented"
	return false
}

func (sw *StreamWriter[T]) Close() { _ = "STUB: not implemented"; return }

type StreamReader[T any] struct {
	typ readerType

	st *stream[T]

	ar *arrayReader[T]

	msr *multiStreamReader[T]

	srw *streamReaderWithConvert[T]

	csr *childStreamReader[T]
}

func (sr *StreamReader[T]) Recv() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (sr *StreamReader[T]) Close() { _ = "STUB: not implemented"; return }

func (sr *StreamReader[T]) Copy(n int) []*StreamReader[T] { _ = "STUB: not implemented"; return nil }

func (sr *StreamReader[T]) SetAutomaticClose() { _ = "STUB: not implemented"; return }

func (sr *StreamReader[T]) recvAny() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (sr *StreamReader[T]) copyAny(n int) []iStreamReader { _ = "STUB: not implemented"; return nil }

func arrToStream[T any](arr []T) *stream[T] { _ = "STUB: not implemented"; return nil }

func (sr *StreamReader[T]) toStream() *stream[T] { _ = "STUB: not implemented"; return nil }

type readerType int

const (
	readerTypeStream readerType = iota
	readerTypeArray
	readerTypeMultiStream
	readerTypeWithConvert
	readerTypeChild
)

type iStreamReader interface {
	recvAny() (any, error)
	copyAny(int) []iStreamReader
	Close()
	SetAutomaticClose()
}

type stream[T any] struct {
	items chan streamItem[T]

	closed chan struct{}

	automaticClose bool
	closedFlag     *uint32
}

type streamItem[T any] struct {
	chunk T
	err   error
}

func newStream[T any](cap int) *stream[T] { _ = "STUB: not implemented"; return nil }

func (s *stream[T]) asReader() *StreamReader[T] { _ = "STUB: not implemented"; return nil }

func (s *stream[T]) recv() (chunk T, err error) { _ = "STUB: not implemented"; return *new(T), nil }

func (s *stream[T]) send(chunk T, err error) (closed bool) { _ = "STUB: not implemented"; return false }

func (s *stream[T]) closeSend() { _ = "STUB: not implemented"; return }

func (s *stream[T]) closeRecv() { _ = "STUB: not implemented"; return }

func StreamReaderFromArray[T any](arr []T) *StreamReader[T] { _ = "STUB: not implemented"; return nil }

type arrayReader[T any] struct {
	arr   []T
	index int
}

func (ar *arrayReader[T]) recv() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (ar *arrayReader[T]) copy(n int) []*arrayReader[T] { _ = "STUB: not implemented"; return nil }

func (ar *arrayReader[T]) toStream() *stream[T] { _ = "STUB: not implemented"; return nil }

type multiArrayReader[T any] struct {
	ars   []*arrayReader[T]
	index int
}

type multiStreamReader[T any] struct {
	sts []*stream[T]

	itemsCases []reflect.SelectCase

	nonClosed []int

	sourceReaderNames []string
}

func newMultiStreamReader[T any](sts []*stream[T]) *multiStreamReader[T] {
	_ = "STUB: not implemented"
	return nil
}

func (msr *multiStreamReader[T]) recv() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (msr *multiStreamReader[T]) nonClosedStreams() []*stream[T] {
	_ = "STUB: not implemented"
	return nil
}

func (msr *multiStreamReader[T]) close() { _ = "STUB: not implemented"; return }

func (msr *multiStreamReader[T]) toStream() *stream[T] { _ = "STUB: not implemented"; return nil }

type streamReaderWithConvert[T any] struct {
	sr iStreamReader

	convert func(any) (T, error)

	errWrapper func(error) error
	onEOF      func() (T, error)
	eofDone    bool
}

func newStreamReaderWithConvert[T any](origin iStreamReader, convert func(any) (T, error), opts ...ConvertOption) *StreamReader[T] {
	_ = "STUB: not implemented"
	return nil
}

type convertOptions struct {
	ErrWrapper func(error) error
	OnEOF      func() (any, error)
}

type ConvertOption func(*convertOptions)

func WithErrWrapper(wrapper func(error) error) ConvertOption {
	_ = "STUB: not implemented"
	return *new(ConvertOption)
}

func WithOnEOF(fn func() (any, error)) ConvertOption {
	_ = "STUB: not implemented"
	return *new(ConvertOption)
}

func StreamReaderWithConvert[T, D any](sr *StreamReader[T], convert func(T) (D, error), opts ...ConvertOption) *StreamReader[D] {
	_ = "STUB: not implemented"
	return nil
}

func (srw *streamReaderWithConvert[T]) recv() (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (srw *streamReaderWithConvert[T]) close() { _ = "STUB: not implemented"; return }

type reader[T any] interface {
	recv() (T, error)
	close()
}

func toStream[T any, Reader reader[T]](r Reader) *stream[T] { _ = "STUB: not implemented"; return nil }

func (srw *streamReaderWithConvert[T]) toStream() *stream[T] { _ = "STUB: not implemented"; return nil }

type cpStreamElement[T any] struct {
	once sync.Once
	next *cpStreamElement[T]
	item streamItem[T]
}

func copyStreamReaders[T any](sr *StreamReader[T], n int) []*StreamReader[T] {
	_ = "STUB: not implemented"
	return nil
}

type parentStreamReader[T any] struct {
	sr *StreamReader[T]

	subStreamList []*cpStreamElement[T]

	closedNum uint32
}

func (p *parentStreamReader[T]) peek(idx int) (t T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (p *parentStreamReader[T]) close(idx int) { _ = "STUB: not implemented"; return }

type childStreamReader[T any] struct {
	parent *parentStreamReader[T]
	index  int
}

func (csr *childStreamReader[T]) recv() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (csr *childStreamReader[T]) toStream() *stream[T] { _ = "STUB: not implemented"; return nil }

func (csr *childStreamReader[T]) close() { _ = "STUB: not implemented"; return }

func MergeStreamReaders[T any](srs []*StreamReader[T]) *StreamReader[T] {
	_ = "STUB: not implemented"
	return nil
}

func MergeNamedStreamReaders[T any](srs map[string]*StreamReader[T]) *StreamReader[T] {
	_ = "STUB: not implemented"
	return nil
}

func InternalMergeNamedStreamReaders[T any](srs []*StreamReader[T], names []string) *StreamReader[T] {
	_ = "STUB: not implemented"
	return nil
}
