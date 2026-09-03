package schema

import (
	"github.com/eino-contrib/jsonschema"
)

type DataType string

const (
	Object  DataType = "object"
	Number  DataType = "number"
	Integer DataType = "integer"
	String  DataType = "string"
	Array   DataType = "array"
	Null    DataType = "null"
	Boolean DataType = "boolean"
)

type ToolChoice string

const (
	ToolChoiceForbidden ToolChoice = "forbidden"

	ToolChoiceAllowed ToolChoice = "allowed"

	ToolChoiceForced ToolChoice = "forced"
)

type AgenticToolChoice struct {
	Type ToolChoice

	Allowed *AgenticAllowedToolChoice

	Forced *AgenticForcedToolChoice
}

type AgenticAllowedToolChoice struct {
	Tools []*AllowedTool
}

type AgenticForcedToolChoice struct {
	Tools []*AllowedTool
}

type AllowedTool struct {
	FunctionName string

	MCPTool *AllowedMCPTool

	ServerTool *AllowedServerTool
}

type AllowedMCPTool struct {
	ServerLabel string

	Name string
}

type AllowedServerTool struct {
	Name string
}

type ToolInfo struct {
	Name string

	Desc string

	Extra map[string]any

	*ParamsOneOf
}

type toolInfoForJSON struct {
	Name           string                    `json:"name,omitempty"`
	Desc           string                    `json:"desc,omitempty"`
	Extra          map[string]any            `json:"extra,omitempty"`
	HasParamsOneOf bool                      `json:"has_params_one_of,omitempty"`
	Params         map[string]*ParameterInfo `json:"params,omitempty"`
	JSONSchema     *jsonschema.Schema        `json:"json_schema,omitempty"`
}

type toolInfoForGob struct {
	Name           string
	Desc           string
	Extra          map[string]any
	HasParamsOneOf bool
	Params         map[string]*ParameterInfo
	JSONSchema     *string
}

func (t *ToolInfo) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *ToolInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (t *ToolInfo) GobEncode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *ToolInfo) GobDecode(b []byte) error { _ = "STUB: not implemented"; return nil }

type ParameterInfo struct {
	Type DataType

	ElemInfo *ParameterInfo

	SubParams map[string]*ParameterInfo

	Desc string

	Enum []string

	Required bool
}

type ParamsOneOf struct {
	params map[string]*ParameterInfo

	jsonschema *jsonschema.Schema
}

func NewParamsOneOfByParams(params map[string]*ParameterInfo) *ParamsOneOf {
	_ = "STUB: not implemented"
	return nil
}

func NewParamsOneOfByJSONSchema(s *jsonschema.Schema) *ParamsOneOf {
	_ = "STUB: not implemented"
	return nil
}

func (p *ParamsOneOf) ToJSONSchema() (*jsonschema.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func paramInfoToJSONSchema(paramInfo *ParameterInfo) *jsonschema.Schema {
	_ = "STUB: not implemented"
	return nil
}

type ToolPartType string

const (
	ToolPartTypeText ToolPartType = "text"

	ToolPartTypeImage ToolPartType = "image"

	ToolPartTypeAudio ToolPartType = "audio"

	ToolPartTypeVideo ToolPartType = "video"

	ToolPartTypeFile ToolPartType = "file"

	ToolPartTypeToolSearchResult ToolPartType = "tool_search_result"
)

type ToolOutputImage struct {
	MessagePartCommon
}

type ToolOutputAudio struct {
	MessagePartCommon
}

type ToolOutputVideo struct {
	MessagePartCommon
}

type ToolOutputFile struct {
	MessagePartCommon
}

type ToolSearchResult struct {
	Tools []*ToolInfo
}

func (t *ToolSearchResult) String() string { _ = "STUB: not implemented"; return "" }

type ToolOutputPart struct {
	Type ToolPartType `json:"type"`

	Text string `json:"text,omitempty"`

	Image *ToolOutputImage `json:"image,omitempty"`

	Audio *ToolOutputAudio `json:"audio,omitempty"`

	Video *ToolOutputVideo `json:"video,omitempty"`

	File *ToolOutputFile `json:"file,omitempty"`

	ToolSearchResult *ToolSearchResult `json:"tool_search_result,omitempty"`

	Extra map[string]any `json:"extra,omitempty"`
}

type ToolArgument struct {
	Text string `json:"text,omitempty"`
}

type ToolResult struct {
	Parts []ToolOutputPart `json:"parts,omitempty"`
}

func convToolOutputPartToMessageInputPart(toolPart ToolOutputPart) (MessageInputPart, error) {
	_ = "STUB: not implemented"
	return *new(MessageInputPart), nil
}

func (tr *ToolResult) ToMessageInputParts() ([]MessageInputPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
