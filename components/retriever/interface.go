package retriever

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

//go:generate mockgen -destination ../../internal/mock/components/retriever/retriever_mock.go --package retriever -source interface.go

type Retriever interface {
	Retrieve(ctx context.Context, query string, opts ...Option) ([]*schema.Document, error)
}
