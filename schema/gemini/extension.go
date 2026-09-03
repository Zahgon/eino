package gemini

type ResponseMetaExtension struct {
	ID            string             `json:"id,omitempty"`
	FinishReason  string             `json:"finish_reason,omitempty"`
	GroundingMeta *GroundingMetadata `json:"grounding_meta,omitempty"`
}

type GroundingMetadata struct {
	GroundingChunks []*GroundingChunk `json:"grounding_chunks,omitempty"`

	GroundingSupports []*GroundingSupport `json:"grounding_supports,omitempty"`

	SearchEntryPoint *SearchEntryPoint `json:"search_entry_point,omitempty"`

	WebSearchQueries []string `json:"web_search_queries,omitempty"`
}

type GroundingChunk struct {
	Web *GroundingChunkWeb `json:"web,omitempty"`
}

type GroundingChunkWeb struct {
	Domain string `json:"domain,omitempty"`

	Title string `json:"title,omitempty"`

	URI string `json:"uri,omitempty"`
}

type GroundingSupport struct {
	ConfidenceScores []float32 `json:"confidence_scores,omitempty"`

	GroundingChunkIndices []int `json:"grounding_chunk_indices,omitempty"`

	Segment *Segment `json:"segment,omitempty"`
}

type Segment struct {
	EndIndex int `json:"end_index,omitempty"`

	PartIndex int `json:"part_index,omitempty"`

	StartIndex int `json:"start_index,omitempty"`

	Text string `json:"text,omitempty"`
}

type SearchEntryPoint struct {
	RenderedContent string `json:"rendered_content,omitempty"`

	SDKBlob []byte `json:"sdk_blob,omitempty"`
}

func ConcatResponseMetaExtensions(chunks []*ResponseMetaExtension) (*ResponseMetaExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
