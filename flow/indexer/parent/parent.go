package parent

import (
	"context"

	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/schema"
)

type Config struct {
	Indexer indexer.Indexer

	Transformer document.Transformer

	ParentIDKey string

	SubIDGenerator func(ctx context.Context, parentID string, num int) ([]string, error)
}

func NewIndexer(ctx context.Context, config *Config) (indexer.Indexer, error) {
	_ = "STUB: not implemented"
	return *new(indexer.Indexer), nil
}

type parentIndexer struct {
	indexer        indexer.Indexer
	transformer    document.Transformer
	parentIDKey    string
	subIDGenerator func(ctx context.Context, parentID string, num int) ([]string, error)
}

func (p *parentIndexer) Store(ctx context.Context, docs []*schema.Document, opts ...indexer.Option) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
