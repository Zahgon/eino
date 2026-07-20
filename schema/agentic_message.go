package schema

import (
	"context"
	"reflect"

	"github.com/eino-contrib/jsonschema"

	"github.com/cloudwego/eino/schema/claude"
	"github.com/cloudwego/eino/schema/gemini"
	"github.com/cloudwego/eino/schema/openai"
)

type ContentBlockType string

const (
	ContentBlockTypeReasoning               ContentBlockType = "reasoning"
	ContentBlockTypeUserInputText           ContentBlockType = "user_input_text"
	ContentBlockTypeUserInputImage          ContentBlockType = "user_input_image"
	ContentBlockTypeUserInputAudio          ContentBlockType = "user_input_audio"
	ContentBlockTypeUserInputVideo          ContentBlockType = "user_input_video"
	ContentBlockTypeUserInputFile           ContentBlockType = "user_input_file"
	ContentBlockTypeToolSearchResult        ContentBlockType = "tool_search_result"
	ContentBlockTypeAssistantGenText        ContentBlockType = "assistant_gen_text"
	ContentBlockTypeAssistantGenImage       ContentBlockType = "assistant_gen_image"
	ContentBlockTypeAssistantGenAudio       ContentBlockType = "assistant_gen_audio"
	ContentBlockTypeAssistantGenVideo       ContentBlockType = "assistant_gen_video"
	ContentBlockTypeFunctionToolCall        ContentBlockType = "function_tool_call"
	ContentBlockTypeFunctionToolResult      ContentBlockType = "function_tool_result"
	ContentBlockTypeServerToolCall          ContentBlockType = "server_tool_call"
	ContentBlockTypeServerToolResult        ContentBlockType = "server_tool_result"
	ContentBlockTypeMCPToolCall             ContentBlockType = "mcp_tool_call"
	ContentBlockTypeMCPToolResult           ContentBlockType = "mcp_tool_result"
	ContentBlockTypeMCPListToolsResult      ContentBlockType = "mcp_list_tools_result"
	ContentBlockTypeMCPToolApprovalRequest  ContentBlockType = "mcp_tool_approval_request"
	ContentBlockTypeMCPToolApprovalResponse ContentBlockType = "mcp_tool_approval_response"
)

type AgenticRoleType string

const (
	AgenticRoleTypeSystem    AgenticRoleType = "system"
	AgenticRoleTypeUser      AgenticRoleType = "user"
	AgenticRoleTypeAssistant AgenticRoleType = "assistant"
)

type AgenticMessage struct {
	Role AgenticRoleType `json:"role"`

	ContentBlocks []*ContentBlock `json:"content_blocks,omitempty"`

	ResponseMeta *AgenticResponseMeta `json:"response_meta,omitempty"`

	Extra map[string]any `json:"extra,omitempty"`
}

type AgenticResponseMeta struct {
	TokenUsage *TokenUsage `json:"token_usage,omitempty"`

	OpenAIExtension *openai.ResponseMetaExtension `json:"openai_extension,omitempty"`

	GeminiExtension *gemini.ResponseMetaExtension `json:"gemini_extension,omitempty"`

	ClaudeExtension *claude.ResponseMetaExtension `json:"claude_extension,omitempty"`

	Extension any `json:"extension,omitempty"`
}

type ContentBlock struct {
	Type ContentBlockType `json:"type"`

	Reasoning *Reasoning `json:"reasoning,omitempty"`

	UserInputText *UserInputText `json:"user_input_text,omitempty"`

	UserInputImage *UserInputImage `json:"user_input_image,omitempty"`

	UserInputAudio *UserInputAudio `json:"user_input_audio,omitempty"`

	UserInputVideo *UserInputVideo `json:"user_input_video,omitempty"`

	UserInputFile *UserInputFile `json:"user_input_file,omitempty"`

	AssistantGenText *AssistantGenText `json:"assistant_gen_text,omitempty"`

	AssistantGenImage *AssistantGenImage `json:"assistant_gen_image,omitempty"`

	AssistantGenAudio *AssistantGenAudio `json:"assistant_gen_audio,omitempty"`

	AssistantGenVideo *AssistantGenVideo `json:"assistant_gen_video,omitempty"`

	FunctionToolCall *FunctionToolCall `json:"function_tool_call,omitempty"`

	FunctionToolResult *FunctionToolResult `json:"function_tool_result,omitempty"`

	ToolSearchFunctionToolResult *ToolSearchFunctionToolResult `json:"tool_search_function_tool_result,omitempty"`

	ServerToolCall *ServerToolCall `json:"server_tool_call,omitempty"`

	ServerToolResult *ServerToolResult `json:"server_tool_result,omitempty"`

	MCPToolCall *MCPToolCall `json:"mcp_tool_call,omitempty"`

	MCPToolResult *MCPToolResult `json:"mcp_tool_result,omitempty"`

	MCPListToolsResult *MCPListToolsResult `json:"mcp_list_tools_result,omitempty"`

	MCPToolApprovalRequest *MCPToolApprovalRequest `json:"mcp_tool_approval_request,omitempty"`

	MCPToolApprovalResponse *MCPToolApprovalResponse `json:"mcp_tool_approval_response,omitempty"`

	StreamingMeta *StreamingMeta `json:"streaming_meta,omitempty"`

	Extra map[string]any `json:"extra,omitempty"`
}

type StreamingMeta struct {
	Index int `json:"index"`
}

type UserInputText struct {
	Text string `json:"text,omitempty"`
}

type UserInputImage struct {
	URL string `json:"url,omitempty"`

	Base64Data string `json:"base64_data,omitempty"`

	MIMEType string `json:"mime_type,omitempty"`

	Detail ImageURLDetail `json:"detail,omitempty"`
}

type UserInputAudio struct {
	URL string `json:"url,omitempty"`

	Base64Data string `json:"base64_data,omitempty"`

	MIMEType string `json:"mime_type,omitempty"`
}

type UserInputVideo struct {
	URL string `json:"url,omitempty"`

	Base64Data string `json:"base64_data,omitempty"`

	MIMEType string `json:"mime_type,omitempty"`
}

type UserInputFile struct {
	URL string `json:"url,omitempty"`

	Name string `json:"name,omitempty"`

	Base64Data string `json:"base64_data,omitempty"`

	MIMEType string `json:"mime_type,omitempty"`
}

type AssistantGenText struct {
	Text string `json:"text,omitempty"`

	OpenAIExtension *openai.AssistantGenTextExtension `json:"openai_extension,omitempty"`

	ClaudeExtension *claude.AssistantGenTextExtension `json:"claude_extension,omitempty"`

	Extension any `json:"extension,omitempty"`
}

type AssistantGenImage struct {
	URL string `json:"url,omitempty"`

	Base64Data string `json:"base64_data,omitempty"`

	MIMEType string `json:"mime_type,omitempty"`
}

type AssistantGenAudio struct {
	URL string `json:"url,omitempty"`

	Base64Data string `json:"base64_data,omitempty"`

	MIMEType string `json:"mime_type,omitempty"`
}

type AssistantGenVideo struct {
	URL string `json:"url,omitempty"`

	Base64Data string `json:"base64_data,omitempty"`

	MIMEType string `json:"mime_type,omitempty"`
}

type Reasoning struct {
	Text string `json:"text,omitempty"`

	Signature string `json:"signature,omitempty"`

	OpenAIExtension *openai.ReasoningExtension `json:"openai_extension,omitempty"`
}

type FunctionToolCall struct {
	CallID string `json:"call_id,omitempty"`

	Name string `json:"name"`

	Arguments string `json:"arguments,omitempty"`
}

type FunctionToolResultContentBlockType string

const (
	FunctionToolResultContentBlockTypeText  FunctionToolResultContentBlockType = "text"
	FunctionToolResultContentBlockTypeImage FunctionToolResultContentBlockType = "image"
	FunctionToolResultContentBlockTypeAudio FunctionToolResultContentBlockType = "audio"
	FunctionToolResultContentBlockTypeVideo FunctionToolResultContentBlockType = "video"
	FunctionToolResultContentBlockTypeFile  FunctionToolResultContentBlockType = "file"
)

type FunctionToolResultContentBlock struct {
	Type FunctionToolResultContentBlockType `json:"type"`

	Text *UserInputText `json:"text,omitempty"`

	Image *UserInputImage `json:"image,omitempty"`

	Audio *UserInputAudio `json:"audio,omitempty"`

	Video *UserInputVideo `json:"video,omitempty"`

	File *UserInputFile `json:"file,omitempty"`

	Extra map[string]any `json:"extra,omitempty"`
}

func (b *FunctionToolResultContentBlock) String() string { _ = "STUB: not implemented"; return "" }

type FunctionToolResult struct {
	CallID string `json:"call_id,omitempty"`

	Name string `json:"name"`

	Content []*FunctionToolResultContentBlock `json:"content,omitempty"`
}

type ToolSearchFunctionToolResult struct {
	CallID string `json:"call_id,omitempty"`

	Name string `json:"name"`

	Result *ToolSearchResult `json:"result,omitempty"`
}

func (t *ToolSearchFunctionToolResult) String() string { _ = "STUB: not implemented"; return "" }

type ServerToolCall struct {
	Name string `json:"name"`

	CallID string `json:"call_id,omitempty"`

	Arguments any `json:"arguments,omitempty"`
}

type ServerToolResult struct {
	Name string `json:"name"`

	CallID string `json:"call_id,omitempty"`

	Content any `json:"content,omitempty"`
}

type MCPToolCall struct {
	ServerLabel string `json:"server_label,omitempty"`

	ApprovalRequestID string `json:"approval_request_id,omitempty"`

	CallID string `json:"call_id,omitempty"`

	Name string `json:"name"`

	Arguments string `json:"arguments,omitempty"`
}

type MCPToolResult struct {
	ServerLabel string `json:"server_label,omitempty"`

	CallID string `json:"call_id,omitempty"`

	Name string `json:"name"`

	Content string `json:"content,omitempty"`

	Error *MCPToolCallError `json:"error,omitempty"`
}

type MCPToolCallError struct {
	Code *int64 `json:"code,omitempty"`

	Message string `json:"message,omitempty"`
}

type MCPListToolsResult struct {
	ServerLabel string `json:"server_label,omitempty"`

	Tools []*MCPListToolsItem `json:"tools,omitempty"`

	Error string `json:"error,omitempty"`
}

type MCPListToolsItem struct {
	Name string `json:"name"`

	Description string `json:"description"`

	InputSchema *jsonschema.Schema `json:"input_schema,omitempty"`
}

type mcpListToolsItemGob struct {
	Name            string
	Description     string
	InputSchemaJSON []byte
}

func (m *MCPListToolsItem) GobEncode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *MCPListToolsItem) GobDecode(data []byte) error { _ = "STUB: not implemented"; return nil }

type MCPToolApprovalRequest struct {
	ID string `json:"id,omitempty"`

	Name string `json:"name"`

	Arguments string `json:"arguments,omitempty"`

	ServerLabel string `json:"server_label,omitempty"`
}

type MCPToolApprovalResponse struct {
	ApprovalRequestID string `json:"approval_request_id,omitempty"`

	Approve bool `json:"approve"`

	Reason string `json:"reason,omitempty"`
}

func SystemAgenticMessage(text string) *AgenticMessage { _ = "STUB: not implemented"; return nil }

func UserAgenticMessage(text string) *AgenticMessage { _ = "STUB: not implemented"; return nil }

type contentBlockVariant interface {
	Reasoning | userInputVariant | assistantGenVariant | functionToolCallVariant | serverToolCallVariant | mcpToolCallVariant
}

type userInputVariant interface {
	UserInputText | UserInputImage | UserInputAudio | UserInputVideo | UserInputFile
}

type assistantGenVariant interface {
	AssistantGenText | AssistantGenImage | AssistantGenAudio | AssistantGenVideo
}

type functionToolCallVariant interface {
	FunctionToolCall | FunctionToolResult | ToolSearchFunctionToolResult
}

type serverToolCallVariant interface {
	ServerToolCall | ServerToolResult
}

type mcpToolCallVariant interface {
	MCPToolCall | MCPToolResult | MCPListToolsResult | MCPToolApprovalRequest | MCPToolApprovalResponse
}

func NewContentBlock[T contentBlockVariant](content *T) *ContentBlock {
	_ = "STUB: not implemented"
	return nil
}

func NewContentBlockChunk[T contentBlockVariant](content *T, meta *StreamingMeta) *ContentBlock {
	_ = "STUB: not implemented"
	return nil
}

type AgenticMessagesTemplate interface {
	Format(ctx context.Context, vs map[string]any, formatType FormatType) ([]*AgenticMessage, error)
}

var _ AgenticMessagesTemplate = &AgenticMessage{}
var _ AgenticMessagesTemplate = AgenticMessagesPlaceholder("", false)

type agenticMessagesPlaceholder struct {
	key      string
	optional bool
}

func AgenticMessagesPlaceholder(key string, optional bool) AgenticMessagesTemplate {
	_ = "STUB: not implemented"
	return *new(AgenticMessagesTemplate)
}

func (p *agenticMessagesPlaceholder) Format(_ context.Context, vs map[string]any, _ FormatType) ([]*AgenticMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *AgenticMessage) Format(_ context.Context, vs map[string]any, formatType FormatType) ([]*AgenticMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatUserInputText(uit *UserInputText, vs map[string]any, formatType FormatType) (*UserInputText, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatUserInputImage(uii *UserInputImage, vs map[string]any, formatType FormatType) (*UserInputImage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatUserInputAudio(uia *UserInputAudio, vs map[string]any, formatType FormatType) (*UserInputAudio, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatUserInputVideo(uiv *UserInputVideo, vs map[string]any, formatType FormatType) (*UserInputVideo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatUserInputFile(uif *UserInputFile, vs map[string]any, formatType FormatType) (*UserInputFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ConcatAgenticMessagesArray(mas [][]*AgenticMessage) ([]*AgenticMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ConcatAgenticMessages(msgs []*AgenticMessage) (*AgenticMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatAgenticResponseMeta(metas []*AgenticResponseMeta) (ret *AgenticResponseMeta, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatTokenUsage(usages []*TokenUsage) *TokenUsage { _ = "STUB: not implemented"; return nil }

func concatChunksOfSameContentBlock(blocks []*ContentBlock) (*ContentBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatContentBlockHelper[T contentBlockVariant](
	blocks []*ContentBlock,
	expectedType ContentBlockType,
	getter func(*ContentBlock) *T,
	concatFunc func([]*T) (*T, error),
) (*ContentBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func genericGetTFromContentBlocks[T any](blocks []*ContentBlock, checkAndGetter func(block *ContentBlock) (T, error)) ([]T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatReasoning(reasons []*Reasoning) (ret *Reasoning, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatUserInputTexts(texts []*UserInputText) (*UserInputText, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatUserInputImages(images []*UserInputImage) (*UserInputImage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatUserInputAudios(audios []*UserInputAudio) (*UserInputAudio, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatUserInputVideos(videos []*UserInputVideo) (*UserInputVideo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatUserInputFiles(files []*UserInputFile) (*UserInputFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatToolSearchFunctionToolResult(results []*ToolSearchFunctionToolResult) (*ToolSearchFunctionToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatAssistantGenTexts(texts []*AssistantGenText) (ret *AssistantGenText, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatAssistantGenImages(images []*AssistantGenImage) (*AssistantGenImage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatAssistantGenAudios(audios []*AssistantGenAudio) (*AssistantGenAudio, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatAssistantGenVideos(videos []*AssistantGenVideo) (*AssistantGenVideo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatFunctionToolCalls(calls []*FunctionToolCall) (*FunctionToolCall, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatFunctionToolResults(results []*FunctionToolResult) (*FunctionToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatServerToolCalls(calls []*ServerToolCall) (ret *ServerToolCall, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatServerToolResults(results []*ServerToolResult) (ret *ServerToolResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatMCPToolCalls(calls []*MCPToolCall) (*MCPToolCall, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatMCPToolResults(results []*MCPToolResult) (*MCPToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatMCPListToolsResults(results []*MCPListToolsResult) (*MCPListToolsResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatMCPToolApprovalRequests(requests []*MCPToolApprovalRequest) (*MCPToolApprovalRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concatMCPToolApprovalResponses(responses []*MCPToolApprovalResponse) (*MCPToolApprovalResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *AgenticMessage) String() string { _ = "STUB: not implemented"; return "" }

func (b *ContentBlock) String() string { _ = "STUB: not implemented"; return "" }

func (r *Reasoning) String() string { _ = "STUB: not implemented"; return "" }

func (u *UserInputText) String() string { _ = "STUB: not implemented"; return "" }

func (u *UserInputImage) String() string { _ = "STUB: not implemented"; return "" }

func (u *UserInputAudio) String() string { _ = "STUB: not implemented"; return "" }

func (u *UserInputVideo) String() string { _ = "STUB: not implemented"; return "" }

func (u *UserInputFile) String() string { _ = "STUB: not implemented"; return "" }

func (a *AssistantGenText) String() string { _ = "STUB: not implemented"; return "" }

func (a *AssistantGenImage) String() string { _ = "STUB: not implemented"; return "" }

func (a *AssistantGenAudio) String() string { _ = "STUB: not implemented"; return "" }

func (a *AssistantGenVideo) String() string { _ = "STUB: not implemented"; return "" }

func (f *FunctionToolCall) String() string { _ = "STUB: not implemented"; return "" }

func (f *FunctionToolResult) String() string { _ = "STUB: not implemented"; return "" }

func (s *ServerToolCall) String() string { _ = "STUB: not implemented"; return "" }

func (s *ServerToolResult) String() string { _ = "STUB: not implemented"; return "" }

func (m *MCPToolCall) String() string { _ = "STUB: not implemented"; return "" }

func (m *MCPToolResult) String() string { _ = "STUB: not implemented"; return "" }

func (m *MCPListToolsResult) String() string { _ = "STUB: not implemented"; return "" }

func (m *MCPToolApprovalRequest) String() string { _ = "STUB: not implemented"; return "" }

func (m *MCPToolApprovalResponse) String() string { _ = "STUB: not implemented"; return "" }

func (a *AgenticResponseMeta) String() string { _ = "STUB: not implemented"; return "" }

func truncateString(s string, maxLen int) string { _ = "STUB: not implemented"; return "" }

func formatMediaString(url, base64Data string, mimeType string, detail string) string {
	_ = "STUB: not implemented"
	return ""
}

func validateExtensionType(expected reflect.Type, actual any) (reflect.Type, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Type), false
}

func printAny(a any) string { _ = "STUB: not implemented"; return "" }
