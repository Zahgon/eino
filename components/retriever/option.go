package retriever

import "github.com/cloudwego/eino/components/embedding"

type Options struct {
	Index *string

	SubIndex *string

	TopK *int

	ScoreThreshold *float64

	Embedding embedding.Embedder

	DSLInfo map[string]any
}

func WithIndex(index string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSubIndex(subIndex string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTopK(topK int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithScoreThreshold(threshold float64) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEmbedding(emb embedding.Embedder) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDSLInfo(dsl map[string]any) Option { _ = "STUB: not implemented"; return *new(Option) }

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
