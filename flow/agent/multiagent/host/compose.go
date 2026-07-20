package host

import (
	"context"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

const (
	defaultHostNodeKey                 = "host"
	defaultHostPrompt                  = "decide which tool is best for the task and call only the best tool."
	specialistsAnswersCollectorNodeKey = "specialist_answers_collect"
	singleIntentAnswerNodeKey          = "single_intent_answer"
	multiIntentSummarizeNodeKey        = "multi_intents_summarize"
	defaultSummarizerPrompt            = "summarize the answers from the specialists into a single answer."
	map2ListConverterNodeKey           = "map_to_list"
)

type state struct {
	msgs              []*schema.Message
	isMultipleIntents bool
}

func NewMultiAgent(ctx context.Context, config *MultiAgentConfig) (*MultiAgent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addSpecialistAgent(specialist *Specialist, g *compose.Graph[[]*schema.Message, *schema.Message]) error {
	_ = "STUB: not implemented"
	return nil
}

func addHostAgent(model model.BaseChatModel, prompt string, g *compose.Graph[[]*schema.Message, *schema.Message], hostNodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

func addDirectAnswerBranch(convertorName string, g *compose.Graph[[]*schema.Message, *schema.Message],
	toolCallChecker func(ctx context.Context, modelOutput *schema.StreamReader[*schema.Message]) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func addMultiSpecialistsBranch(convertorName string, agentMap map[string]bool, g *compose.Graph[[]*schema.Message, *schema.Message]) error {
	_ = "STUB: not implemented"
	return nil
}

func addSingleIntentAnswerNode(g *compose.Graph[[]*schema.Message, *schema.Message]) error {
	_ = "STUB: not implemented"
	return nil
}

func addAfterSpecialistsBranch(g *compose.Graph[[]*schema.Message, *schema.Message]) error {
	_ = "STUB: not implemented"
	return nil
}

func addMultiIntentsSummarizeNode(summarizer *Summarizer, g *compose.Graph[[]*schema.Message, *schema.Message]) error {
	_ = "STUB: not implemented"
	return nil
}
