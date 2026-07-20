package parent

import (
	"context"

	"github.com/cloudwego/eino/components/retriever"
	"github.com/cloudwego/eino/schema"
)

type Config struct {
	Retriever retriever.Retriever

	ParentIDKey string

	OrigDocGetter func(ctx context.Context, ids []string) ([]*schema.Document, error)
}

func NewRetriever(ctx context.Context, config *Config) (retriever.Retriever, error) {
	_ = "STUB: not implemented"
	return *new(retriever.Retriever), nil
}

type parentRetriever struct {
	retriever     retriever.Retriever
	parentIDKey   string
	origDocGetter func(ctx context.Context, ids []string) ([]*schema.Document, error)
}

func (p *parentRetriever) Retrieve(ctx context.Context, query string, opts ...retriever.Option) ([]*schema.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inList(elem string, list []string) bool { _ = "STUB: not implemented"; return false }
