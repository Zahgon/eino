package embedding

import "context"

//go:generate  mockgen -destination ../../internal/mock/components/embedding/Embedding_mock.go --package embedding -source interface.go
type Embedder interface {
	EmbedStrings(ctx context.Context, texts []string, opts ...Option) ([][]float64, error)
}
