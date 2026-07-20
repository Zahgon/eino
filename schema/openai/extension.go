package openai

type ResponseMetaExtension struct {
	ID                   string               `json:"id,omitempty"`
	Status               ResponseStatus       `json:"status,omitempty"`
	Error                *ResponseError       `json:"error,omitempty"`
	IncompleteDetails    *IncompleteDetails   `json:"incomplete_details,omitempty"`
	PreviousResponseID   string               `json:"previous_response_id,omitempty"`
	Reasoning            *Reasoning           `json:"reasoning,omitempty"`
	ServiceTier          ServiceTier          `json:"service_tier,omitempty"`
	CreatedAt            int64                `json:"created_at,omitempty"`
	PromptCacheRetention PromptCacheRetention `json:"prompt_cache_retention,omitempty"`
}

type AssistantGenTextExtension struct {
	Refusal     *OutputRefusal    `json:"refusal,omitempty"`
	Annotations []*TextAnnotation `json:"annotations,omitempty"`
}

type ReasoningExtension struct {
	Content []*ReasoningContent `json:"content,omitempty"`
}

type ReasoningContent struct {
	Index *int `json:"index,omitempty"`

	Text string `json:"text,omitempty"`
}

type ResponseError struct {
	Code    ResponseErrorCode `json:"code,omitempty"`
	Message string            `json:"message,omitempty"`
}

type IncompleteDetails struct {
	Reason string `json:"reason,omitempty"`
}

type Reasoning struct {
	Effort  ReasoningEffort  `json:"effort,omitempty"`
	Summary ReasoningSummary `json:"summary,omitempty"`
}

type OutputRefusal struct {
	Reason string `json:"reason,omitempty"`
}

type TextAnnotation struct {
	Index int `json:"index,omitempty"`

	Type TextAnnotationType `json:"type,omitempty"`

	FileCitation          *TextAnnotationFileCitation          `json:"file_citation,omitempty"`
	URLCitation           *TextAnnotationURLCitation           `json:"url_citation,omitempty"`
	ContainerFileCitation *TextAnnotationContainerFileCitation `json:"container_file_citation,omitempty"`
	FilePath              *TextAnnotationFilePath              `json:"file_path,omitempty"`
}

type TextAnnotationFileCitation struct {
	FileID string `json:"file_id,omitempty"`

	Filename string `json:"filename,omitempty"`

	Index int `json:"index,omitempty"`
}

type TextAnnotationURLCitation struct {
	Title string `json:"title,omitempty"`

	URL string `json:"url,omitempty"`

	StartIndex int `json:"start_index,omitempty"`

	EndIndex int `json:"end_index,omitempty"`
}

type TextAnnotationContainerFileCitation struct {
	ContainerID string `json:"container_id,omitempty"`

	FileID string `json:"file_id,omitempty"`

	Filename string `json:"filename,omitempty"`

	StartIndex int `json:"start_index,omitempty"`

	EndIndex int `json:"end_index,omitempty"`
}

type TextAnnotationFilePath struct {
	FileID string `json:"file_id,omitempty"`

	Index int `json:"index,omitempty"`
}

func ConcatAssistantGenTextExtensions(chunks []*AssistantGenTextExtension) (*AssistantGenTextExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ConcatReasoningExtensions(chunks []*ReasoningExtension) (*ReasoningExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ConcatResponseMetaExtensions(chunks []*ResponseMetaExtension) (*ResponseMetaExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
