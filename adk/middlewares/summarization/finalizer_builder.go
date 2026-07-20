package summarization

import (
	"context"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

func DefaultFinalize[M adk.MessageType](ctx context.Context, originalMessages []M, summary M) ([]M, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TypedFinalizerBuilder[M adk.MessageType] struct {
	handlers []TypedFinalizeFunc[M]
	errs     []error
}

type FinalizerBuilder = TypedFinalizerBuilder[*schema.Message]

func NewTypedFinalizer[M adk.MessageType]() *TypedFinalizerBuilder[M] {
	_ = "STUB: not implemented"
	return nil
}

func NewFinalizer() *FinalizerBuilder { _ = "STUB: not implemented"; return nil }

func (b *TypedFinalizerBuilder[M]) Build() (TypedFinalizeFunc[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PreserveSkillsConfig struct {
	SkillToolName string

	MaxSkills *int

	MaxTokensPerSkill *int

	SkillsTokenBudget *int
}

func (b *TypedFinalizerBuilder[M]) PreserveSkills(config *PreserveSkillsConfig) *TypedFinalizerBuilder[M] {
	_ = "STUB: not implemented"
	return nil
}

func (c *PreserveSkillsConfig) check() error { _ = "STUB: not implemented"; return nil }

type skillInfo struct {
	Name    string
	Content string
}

func extractSkillInfos[M adk.MessageType](messages []M, skillTool string) ([]*skillInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildPreservedSkillsText[M adk.MessageType](_ context.Context, messages []M, config *PreserveSkillsConfig) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func truncateSkillContent(content string, maxTokens int) string {
	_ = "STUB: not implemented"
	return ""
}
