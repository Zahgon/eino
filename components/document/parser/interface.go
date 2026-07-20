package parser

import (
	"context"
	"io"

	"github.com/cloudwego/eino/schema"
)

type Parser interface {
	Parse(ctx context.Context, reader io.Reader, opts ...Option) ([]*schema.Document, error)
}
