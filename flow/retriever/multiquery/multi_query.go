package multiquery

import (
	"context"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/retriever"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

const (
	defaultRewritePrompt = `You are an helpful assistant.
	Your role is to create three different versions of the user query to retrieve relevant documents from store.
    Your goal is to improve the performance of similarity search by generating text from different perspectives based on the user query.
	Only provide the generated queries and separate them by newlines. 
	user query: {{query}}`
	defaultQueryVariable = "query"
	defaultMaxQueriesNum = 5
)

var deduplicateFusion = func(ctx context.Context, docs [][]*schema.Document) ([]*schema.Document, error) {
	m := map[string]bool{}
	var ret []*schema.Document
	for i := range docs {
		for j := range docs[i] {
			if _, ok := m[docs[i][j].ID]; !ok {
				m[docs[i][j].ID] = true
				ret = append(ret, docs[i][j])
			}
		}
	}
	return ret, nil
}

func NewRetriever(ctx context.Context, config *Config) (retriever.Retriever, error) {
	_ = "STUB: not implemented"
	return *new(retriever.Retriever), nil
}

type Config struct {
	RewriteLLM model.ChatModel

	RewriteTemplate prompt.ChatTemplate

	QueryVar string

	LLMOutputParser func(context.Context, *schema.Message) ([]string, error)

	RewriteHandler func(ctx context.Context, query string) ([]string, error)

	MaxQueriesNum int

	OrigRetriever retriever.Retriever

	FusionFunc func(ctx context.Context, docs [][]*schema.Document) ([]*schema.Document, error)
}

type multiQueryRetriever struct {
	queryRunner   compose.Runnable[string, []string]
	maxQueriesNum int
	origRetriever retriever.Retriever
	fusionFunc    func(ctx context.Context, docs [][]*schema.Document) ([]*schema.Document, error)
}

func (m *multiQueryRetriever) Retrieve(ctx context.Context, query string, opts ...retriever.Option) ([]*schema.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *multiQueryRetriever) GetType() string { _ = "STUB: not implemented"; return "" }

func ctxWithFusionRunInfo(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
