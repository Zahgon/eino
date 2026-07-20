package claude

type ResponseMetaExtension struct {
	ID           string       `json:"id,omitempty"`
	StopReason   string       `json:"stop_reason,omitempty"`
	StopSequence string       `json:"stop_sequence,omitempty"`
	StopDetails  *StopDetails `json:"stop_details,omitempty"`
}

type StopDetails struct {
	Category    string `json:"category,omitempty"`
	Explanation string `json:"explanation,omitempty"`
}

type AssistantGenTextExtension struct {
	Citations []*TextCitation `json:"citations,omitempty"`
}

type TextCitation struct {
	Type TextCitationType `json:"type,omitempty"`

	CharLocation            *CitationCharLocation            `json:"char_location,omitempty"`
	PageLocation            *CitationPageLocation            `json:"page_location,omitempty"`
	ContentBlockLocation    *CitationContentBlockLocation    `json:"content_block_location,omitempty"`
	WebSearchResultLocation *CitationWebSearchResultLocation `json:"web_search_result_location,omitempty"`
}

type CitationCharLocation struct {
	CitedText string `json:"cited_text,omitempty"`

	DocumentTitle string `json:"document_title,omitempty"`
	DocumentIndex int    `json:"document_index,omitempty"`

	StartCharIndex int `json:"start_char_index,omitempty"`
	EndCharIndex   int `json:"end_char_index,omitempty"`
}

type CitationPageLocation struct {
	CitedText string `json:"cited_text,omitempty"`

	DocumentTitle string `json:"document_title,omitempty"`
	DocumentIndex int    `json:"document_index,omitempty"`

	StartPageNumber int `json:"start_page_number,omitempty"`
	EndPageNumber   int `json:"end_page_number,omitempty"`
}

type CitationContentBlockLocation struct {
	CitedText string `json:"cited_text,omitempty"`

	DocumentTitle string `json:"document_title,omitempty"`
	DocumentIndex int    `json:"document_index,omitempty"`

	StartBlockIndex int `json:"start_block_index,omitempty"`
	EndBlockIndex   int `json:"end_block_index,omitempty"`
}

type CitationWebSearchResultLocation struct {
	CitedText string `json:"cited_text,omitempty"`

	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`

	EncryptedIndex string `json:"encrypted_index,omitempty"`
}

func ConcatAssistantGenTextExtensions(chunks []*AssistantGenTextExtension) (*AssistantGenTextExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ConcatResponseMetaExtensions(chunks []*ResponseMetaExtension) (*ResponseMetaExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
