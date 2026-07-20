package summarization

import (
	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

type TypedCustomizedAction[M adk.MessageType] struct {
	Type ActionType `json:"type"`

	Before *TypedBeforeSummarizeAction[M] `json:"before,omitempty"`

	After *TypedAfterSummarizeAction[M] `json:"after,omitempty"`

	GenerateSummary *TypedGenerateSummaryAction[M] `json:"generate_summary,omitempty"`
}

type CustomizedAction = TypedCustomizedAction[*schema.Message]

type TypedBeforeSummarizeAction[M adk.MessageType] struct {
	Messages []M `json:"messages,omitempty"`
}

type BeforeSummarizeAction = TypedBeforeSummarizeAction[*schema.Message]

type TypedAfterSummarizeAction[M adk.MessageType] struct {
	Messages []M `json:"messages,omitempty"`
}

type AfterSummarizeAction = TypedAfterSummarizeAction[*schema.Message]

type GenerateSummaryPhase string

const (
	GenerateSummaryPhasePrimary GenerateSummaryPhase = "primary"

	GenerateSummaryPhaseFailover GenerateSummaryPhase = "failover"
)

type TypedGenerateSummaryAction[M adk.MessageType] struct {
	Attempt int `json:"attempt"`

	Phase GenerateSummaryPhase `json:"phase"`

	ModelResponse M `json:"model_response,omitempty"`

	err error
}

type GenerateSummaryAction = TypedGenerateSummaryAction[*schema.Message]

func (a *TypedGenerateSummaryAction[M]) GetError() error { _ = "STUB: not implemented"; return nil }
