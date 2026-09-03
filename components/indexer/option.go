package indexer

import "github.com/cloudwego/eino/components/embedding"

type Options struct {
	Index *string

	SubIndexes []string

	Embedding embedding.Embedder
}

func WithIndex(index string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSubIndexes(subIndexes []string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEmbedding(emb embedding.Embedder) Option { _ = "STUB: not implemented"; return *new(Option) }

type Option struct {
	apply func(opts *Options)

	implSpecificOptFn any
}

func GetCommonOptions(base *Options, opts ...Option) *Options {
	_ = "STUB: not implemented"
	return nil
}

func WrapImplSpecificOptFn[T any](optFn func(*T)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func GetImplSpecificOptions[T any](base *T, opts ...Option) *T {
	_ = "STUB: not implemented"
	return nil
}
