package schema

import (
	"context"
	"sync"

	"github.com/nikolalohinski/gonja"

	"github.com/cloudwego/eino/internal"
)

func init() {
	internal.RegisterStreamChunkConcatFunc(ConcatMessages)
	internal.RegisterStreamChunkConcatFunc(ConcatMessageArray)

	internal.RegisterStreamChunkConcatFunc(ConcatAgenticMessages)
	internal.RegisterStreamChunkConcatFunc(ConcatAgenticMessagesArray)

	internal.RegisterStreamChunkConcatFunc(ConcatToolResults)
}

func buildConcatGenericArray[T any](f func([]*T) (*T, error)) func([][]*T) ([]*T, error) {
	_ = "STUB: not implemented"
	return nil
}

func ConcatMessageArray(mas [][]*Message) ([]*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type FormatType uint8

const (
	FString FormatType = 0

	GoTemplate FormatType = 1

	Jinja2 FormatType = 2
)

type RoleType string

const (
	Assistant RoleType = "assistant"

	User RoleType = "user"

	System RoleType = "system"

	Tool RoleType = "tool"
)

type FunctionCall struct {
	Name string `json:"name,omitempty"`

	Arguments string `json:"arguments,omitempty"`
}

type ToolCall struct {
	Index *int `json:"index,omitempty"`

	ID string `json:"id"`

	Type string `json:"type"`

	Function FunctionCall `json:"function"`

	Extra map[string]any `json:"extra,omitempty"`
}

type ImageURLDetail string

const (
	ImageURLDetailHigh ImageURLDetail = "high"

	ImageURLDetailLow ImageURLDetail = "low"

	ImageURLDetailAuto ImageURLDetail = "auto"
)

type MessagePartCommon struct {
	URL *string `json:"url,omitempty"`

	Base64Data *string `json:"base64data,omitempty"`

	MIMEType string `json:"mime_type,omitempty"`

	Extra map[string]any `json:"extra,omitempty"`
}

type MessageInputImage struct {
	MessagePartCommon

	Detail ImageURLDetail `json:"detail,omitempty"`
}

type MessageInputAudio struct {
	MessagePartCommon
}

type MessageInputVideo struct {
	MessagePartCommon
}

type MessageInputFile struct {
	MessagePartCommon

	Name string `json:"name,omitempty"`
}

type MessageInputPart struct {
	Type ChatMessagePartType `json:"type"`

	Text string `json:"text,omitempty"`

	Image *MessageInputImage `json:"image,omitempty"`

	Audio *MessageInputAudio `json:"audio,omitempty"`

	Video *MessageInputVideo `json:"video,omitempty"`

	File *MessageInputFile `json:"file,omitempty"`

	ToolSearchResult *ToolSearchResult `json:"tool_search_result,omitempty"`

	Extra map[string]any `json:"extra,omitempty"`
}

type MessageOutputImage struct {
	MessagePartCommon
}

type MessageOutputAudio struct {
	MessagePartCommon
}

type MessageOutputVideo struct {
	MessagePartCommon
}

type MessageOutputReasoning struct {
	Text string `json:"text,omitempty"`

	Signature string `json:"signature,omitempty"`
}

type MessageStreamingMeta struct {
	Index int `json:"index,omitempty"`
}

type MessageOutputPart struct {
	Type ChatMessagePartType `json:"type"`

	Text string `json:"text,omitempty"`

	Image *MessageOutputImage `json:"image,omitempty"`

	Audio *MessageOutputAudio `json:"audio,omitempty"`

	Video *MessageOutputVideo `json:"video,omitempty"`

	Reasoning *MessageOutputReasoning `json:"reasoning,omitempty"`

	Extra map[string]any `json:"extra,omitempty"`

	StreamingMeta *MessageStreamingMeta `json:"-"`
}

type ChatMessageImageURL struct {
	URL string `json:"url,omitempty"`

	URI string `json:"uri,omitempty"`

	Detail ImageURLDetail `json:"detail,omitempty"`

	MIMEType string `json:"mime_type,omitempty"`

	Extra map[string]any `json:"extra,omitempty"`
}

type ChatMessagePartType string

const (
	ChatMessagePartTypeText ChatMessagePartType = "text"

	ChatMessagePartTypeImageURL ChatMessagePartType = "image_url"

	ChatMessagePartTypeAudioURL ChatMessagePartType = "audio_url"

	ChatMessagePartTypeVideoURL ChatMessagePartType = "video_url"

	ChatMessagePartTypeFileURL ChatMessagePartType = "file_url"

	ChatMessagePartTypeReasoning ChatMessagePartType = "reasoning"

	ChatMessagePartTypeToolSearchResult ChatMessagePartType = "tool_search_result"
)

type ChatMessageAudioURL struct {
	URL string `json:"url,omitempty"`
	URI string `json:"uri,omitempty"`

	MIMEType string `json:"mime_type,omitempty"`

	Extra map[string]any `json:"extra,omitempty"`
}

type ChatMessageVideoURL struct {
	URL string `json:"url,omitempty"`
	URI string `json:"uri,omitempty"`

	MIMEType string `json:"mime_type,omitempty"`

	Extra map[string]any `json:"extra,omitempty"`
}

type ChatMessageFileURL struct {
	URL string `json:"url,omitempty"`
	URI string `json:"uri,omitempty"`

	MIMEType string `json:"mime_type,omitempty"`

	Name string `json:"name,omitempty"`

	Extra map[string]any `json:"extra,omitempty"`
}

type ChatMessagePart struct {
	Type ChatMessagePartType `json:"type,omitempty"`

	Text string `json:"text,omitempty"`

	ImageURL *ChatMessageImageURL `json:"image_url,omitempty"`

	AudioURL *ChatMessageAudioURL `json:"audio_url,omitempty"`

	VideoURL *ChatMessageVideoURL `json:"video_url,omitempty"`

	FileURL *ChatMessageFileURL `json:"file_url,omitempty"`
}

type LogProbs struct {
	Content []LogProb `json:"content"`
}

type LogProb struct {
	Token string `json:"token"`

	LogProb float64 `json:"logprob"`

	Bytes []int64 `json:"bytes,omitempty"`

	TopLogProbs []TopLogProb `json:"top_logprobs"`
}

type TopLogProb struct {
	Token string `json:"token"`

	LogProb float64 `json:"logprob"`

	Bytes []int64 `json:"bytes,omitempty"`
}

type ResponseMeta struct {
	FinishReason string `json:"finish_reason,omitempty"`

	Usage *TokenUsage `json:"usage,omitempty"`

	LogProbs *LogProbs `json:"logprobs,omitempty"`
}

type Message struct {
	Role RoleType `json:"role"`

	Content string `json:"content"`

	MultiContent []ChatMessagePart `json:"multi_content,omitempty"`

	UserInputMultiContent []MessageInputPart `json:"user_input_multi_content,omitempty"`

	AssistantGenMultiContent []MessageOutputPart `json:"assistant_output_multi_content,omitempty"`

	Name string `json:"name,omitempty"`

	ToolCalls []ToolCall `json:"tool_calls,omitempty"`

	ToolCallID string `json:"tool_call_id,omitempty"`

	ToolName string `json:"tool_name,omitempty"`

	ResponseMeta *ResponseMeta `json:"response_meta,omitempty"`

	ReasoningContent string `json:"reasoning_content,omitempty"`

	Extra map[string]any `json:"extra,omitempty"`
}

type TokenUsage struct {
	PromptTokens int `json:"prompt_tokens"`

	PromptTokenDetails PromptTokenDetails `json:"prompt_token_details"`

	CompletionTokens int `json:"completion_tokens"`

	TotalTokens int `json:"total_tokens"`

	CompletionTokensDetails CompletionTokensDetails `json:"completion_token_details"`
}

type CompletionTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens,omitempty"`
}

type PromptTokenDetails struct {
	CachedTokens int `json:"cached_tokens"`
}

var _ MessagesTemplate = &Message{}
var _ MessagesTemplate = MessagesPlaceholder("", false)

type MessagesTemplate interface {
	Format(ctx context.Context, vs map[string]any, formatType FormatType) ([]*Message, error)
}

type messagesPlaceholder struct {
	key      string
	optional bool
}

func MessagesPlaceholder(key string, optional bool) MessagesTemplate {
	_ = "STUB: not implemented"
	return *new(MessagesTemplate)
}

func (p *messagesPlaceholder) Format(_ context.Context, vs map[string]any, _ FormatType) ([]*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatContent(content string, vs map[string]any, formatType FormatType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *Message) Format(_ context.Context, vs map[string]any, formatType FormatType) ([]*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatMultiContent(multiContent []ChatMessagePart, vs map[string]any, formatType FormatType) ([]ChatMessagePart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatUserInputMultiContent(userInputMultiContent []MessageInputPart, vs map[string]any, formatType FormatType) ([]MessageInputPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Message) String() string { _ = "STUB: not implemented"; return "" }

func formatInputPart(part MessageInputPart) string { _ = "STUB: not implemented"; return "" }

func formatMessageInputMedia[T MessageInputImage | MessageInputAudio | MessageInputVideo](media *T) string {
	_ = "STUB: not implemented"
	return ""
}

func formatMessageInputFile(file *MessageInputFile) string { _ = "STUB: not implemented"; return "" }

func formatOutputPart(part MessageOutputPart) string { _ = "STUB: not implemented"; return "" }

func formatMessageOutputMedia[T MessageOutputImage | MessageOutputAudio | MessageOutputVideo](media *T) string {
	_ = "STUB: not implemented"
	return ""
}

func formatChatMessagePart(part ChatMessagePart) string { _ = "STUB: not implemented"; return "" }

func SystemMessage(content string) *Message { _ = "STUB: not implemented"; return nil }

func AssistantMessage(content string, toolCalls []ToolCall) *Message {
	_ = "STUB: not implemented"
	return nil
}

func UserMessage(content string) *Message { _ = "STUB: not implemented"; return nil }

type toolMessageOptions struct {
	toolName string
}

type ToolMessageOption func(*toolMessageOptions)

func WithToolName(name string) ToolMessageOption {
	_ = "STUB: not implemented"
	return *new(ToolMessageOption)
}

func ToolMessage(content string, toolCallID string, opts ...ToolMessageOption) *Message {
	_ = "STUB: not implemented"
	return nil
}

func ConcatToolResults(chunks []*ToolResult) (*ToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatToolOutputParts(parts []ToolOutputPart) ([]ToolOutputPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func groupToolOutputParts(parts []ToolOutputPart) [][]ToolOutputPart {
	_ = "STUB: not implemented"
	return nil
}

func mergeToolTextParts(group []ToolOutputPart) (ToolOutputPart, error) {
	_ = "STUB: not implemented"
	return *new(ToolOutputPart), nil
}

func concatToolCalls(chunks []ToolCall) ([]ToolCall, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatAssistantMultiContent(parts []MessageOutputPart) ([]MessageOutputPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func groupOutputParts(parts []MessageOutputPart) [][]MessageOutputPart {
	_ = "STUB: not implemented"
	return nil
}

func canMergeOutputParts(current, next MessageOutputPart) bool {
	_ = "STUB: not implemented"
	return false
}

func isMergeableOutputPartType(part MessageOutputPart) bool {
	_ = "STUB: not implemented"
	return false
}

func mergeOutputPartGroup(group []MessageOutputPart) (MessageOutputPart, error) {
	_ = "STUB: not implemented"
	return *new(MessageOutputPart), nil
}

func mergeTextParts(group []MessageOutputPart) (MessageOutputPart, error) {
	_ = "STUB: not implemented"
	return *new(MessageOutputPart), nil
}

func mergeReasoningParts(group []MessageOutputPart) (MessageOutputPart, error) {
	_ = "STUB: not implemented"
	return *new(MessageOutputPart), nil
}

func mergeAudioParts(group []MessageOutputPart) (MessageOutputPart, error) {
	_ = "STUB: not implemented"
	return *new(MessageOutputPart), nil
}

func isBase64MessageOutputAudioPart(part MessageOutputPart) bool {
	_ = "STUB: not implemented"
	return false
}

func concatUserMultiContent(parts []MessageInputPart) ([]MessageInputPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatExtra(extraList []map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ConcatMessages(msgs []*Message) (*Message, error) { _ = "STUB: not implemented"; return nil, nil }

func ConcatMessageStream(s *StreamReader[*Message]) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var jinjaEnvOnce sync.Once
var jinjaEnv *gonja.Environment
var envInitErr error

const (
	jinjaInclude = "include"
	jinjaExtends = "extends"
	jinjaImport  = "import"
	jinjaFrom    = "from"
	jinjaFile    = "file"
	jinjaFileSet = "fileset"
)

func getJinjaEnv() (*gonja.Environment, error) { _ = "STUB: not implemented"; return nil, nil }
