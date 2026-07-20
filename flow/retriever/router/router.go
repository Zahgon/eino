package router

import (
	"context"
	"fmt"
	"sort"

	"github.com/cloudwego/eino/components/retriever"
	"github.com/cloudwego/eino/schema"
)

var rrf = func(ctx context.Context, result map[string][]*schema.Document) ([]*schema.Document, error) {
	if len(result) < 1 {
		return nil, fmt.Errorf("no documents")
	}
	if len(result) == 1 {
		for _, docs := range result {
			return docs, nil
		}
	}

	docRankMap := make(map[string]float64)
	docMap := make(map[string]*schema.Document)
	for _, v := range result {
		for i := range v {
			docMap[v[i].ID] = v[i]
			if _, ok := docRankMap[v[i].ID]; !ok {
				docRankMap[v[i].ID] = 1.0 / float64(i+60)
			} else {
				docRankMap[v[i].ID] += 1.0 / float64(i+60)
			}
		}
	}
	docList := make([]*schema.Document, 0, len(docMap))
	for id := range docMap {
		docList = append(docList, docMap[id])
	}

	sort.Slice(docList, func(i, j int) bool {
		return docRankMap[docList[i].ID] > docRankMap[docList[j].ID]
	})

	return docList, nil
}

func NewRetriever(ctx context.Context, config *Config) (retriever.Retriever, error) {
	_ = "STUB: not implemented"
	return *new(retriever.Retriever), nil
}

type Config struct {
	Retrievers map[string]retriever.Retriever

	Router func(ctx context.Context, query string) ([]string, error)

	FusionFunc func(ctx context.Context, result map[string][]*schema.Document) ([]*schema.Document, error)
}

type routerRetriever struct {
	retrievers map[string]retriever.Retriever
	router     func(ctx context.Context, query string) ([]string, error)
	fusionFunc func(ctx context.Context, result map[string][]*schema.Document) ([]*schema.Document, error)
}

func (e *routerRetriever) Retrieve(ctx context.Context, query string, opts ...retriever.Option) ([]*schema.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *routerRetriever) GetType() string { _ = "STUB: not implemented"; return "" }

func ctxWithRouterRunInfo(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ctxWithFusionRunInfo(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
