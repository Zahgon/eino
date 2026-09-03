package indexer

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

//go:generate  mockgen -destination ../../internal/mock/components/indexer/indexer_mock.go --package indexer -source interface.go
type Indexer interface {
	Store(ctx context.Context, docs []*schema.Document, opts ...Option) (ids []string, err error)
}
