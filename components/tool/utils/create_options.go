package utils

import (
	"context"
	"reflect"

	"github.com/eino-contrib/jsonschema"
)

type UnmarshalArguments func(ctx context.Context, arguments string) (any, error)

type MarshalOutput func(ctx context.Context, output any) (string, error)

type toolOptions struct {
	um         UnmarshalArguments
	m          MarshalOutput
	scModifier SchemaModifierFn
}

type Option func(o *toolOptions)

func WithUnmarshalArguments(um UnmarshalArguments) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMarshalOutput(m MarshalOutput) Option { _ = "STUB: not implemented"; return *new(Option) }

type SchemaModifierFn func(jsonTagName string, t reflect.Type, tag reflect.StructTag, schema *jsonschema.Schema)

func WithSchemaModifier(modifier SchemaModifierFn) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func getToolOptions(opt ...Option) *toolOptions { _ = "STUB: not implemented"; return nil }
