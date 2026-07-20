package parser

import (
	"context"
	"io"

	"github.com/cloudwego/eino/schema"
)

const (
	MetaKeySource = "_source"
)

type TextParser struct{}

func (dp TextParser) Parse(ctx context.Context, reader io.Reader, opts ...Option) ([]*schema.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
