package utils

import (
	"context"

	"github.com/cloudwego/eino/components/retriever"
	"github.com/cloudwego/eino/schema"
)

type RetrieveTask struct {
	Name            string
	Retriever       retriever.Retriever
	Query           string
	RetrieveOptions []retriever.Option
	Result          []*schema.Document
	Err             error
}

func ConcurrentRetrieveWithCallback(ctx context.Context, tasks []*RetrieveTask) {
	_ = "STUB: not implemented"
	return
}

func ctxWithRetrieverRunInfo(ctx context.Context, r retriever.Retriever) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
