package document

import "github.com/cloudwego/eino/components/document/parser"

type LoaderOptions struct {
	ParserOptions []parser.Option
}

type LoaderOption struct {
	apply func(opts *LoaderOptions)

	implSpecificOptFn any
}

func WrapLoaderImplSpecificOptFn[T any](optFn func(*T)) LoaderOption {
	_ = "STUB: not implemented"
	return *new(LoaderOption)
}

func GetLoaderImplSpecificOptions[T any](base *T, opts ...LoaderOption) *T {
	_ = "STUB: not implemented"
	return nil
}

func GetLoaderCommonOptions(base *LoaderOptions, opts ...LoaderOption) *LoaderOptions {
	_ = "STUB: not implemented"
	return nil
}

func WithParserOptions(opts ...parser.Option) LoaderOption {
	_ = "STUB: not implemented"
	return *new(LoaderOption)
}

type TransformerOption struct {
	implSpecificOptFn any
}

func WrapTransformerImplSpecificOptFn[T any](optFn func(*T)) TransformerOption {
	_ = "STUB: not implemented"
	return *new(TransformerOption)
}

func GetTransformerImplSpecificOptions[T any](base *T, opts ...TransformerOption) *T {
	_ = "STUB: not implemented"
	return nil
}
