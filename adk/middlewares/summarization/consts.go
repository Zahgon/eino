package summarization

type ActionType string

const (
	ActionTypeBeforeSummarize ActionType = "before_summarize"
	ActionTypeAfterSummarize  ActionType = "after_summarize"
	ActionTypeGenerateSummary ActionType = "generate_summary"
)

const (
	extraKeyContentType       = "_eino_summarization_content_type"
	preserveUserMsgsMaxTokens = 30000
)

type summarizationContentType string

const (
	contentTypeSummary summarizationContentType = "summary"
	contentTypeSkills  summarizationContentType = "skills"
)

type ctxKeyModelInput struct{}
