package document

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

type Source struct {
	URI string
}

//go:generate  mockgen -destination ../../internal/mock/components/document/document_mock.go --package document -source interface.go

type Loader interface {
	Load(ctx context.Context, src Source, opts ...LoaderOption) ([]*schema.Document, error)
}

type Transformer interface {
	Transform(ctx context.Context, src []*schema.Document, opts ...TransformerOption) ([]*schema.Document, error)
}
