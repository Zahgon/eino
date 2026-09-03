package components

type Typer interface {
	GetType() string
}

func GetType(component any) (string, bool) { _ = "STUB: not implemented"; return "", false }

type Checker interface {
	IsCallbacksEnabled() bool
}

func IsCallbacksEnabled(i any) bool { _ = "STUB: not implemented"; return false }

type Component string

const (
	ComponentOfPrompt Component = "ChatTemplate"

	ComponentOfAgenticPrompt Component = "AgenticChatTemplate"

	ComponentOfChatModel Component = "ChatModel"

	ComponentOfAgenticModel Component = "AgenticModel"

	ComponentOfEmbedding Component = "Embedding"

	ComponentOfIndexer Component = "Indexer"

	ComponentOfRetriever Component = "Retriever"

	ComponentOfLoader Component = "Loader"

	ComponentOfTransformer Component = "DocumentTransformer"

	ComponentOfTool Component = "Tool"
)
