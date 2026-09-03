package filesystem

import (
	"context"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/adk/filesystem"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
)

const (
	ToolNameLs        = "ls"
	ToolNameReadFile  = "read_file"
	ToolNameWriteFile = "write_file"
	ToolNameEditFile  = "edit_file"
	ToolNameGlob      = "glob"
	ToolNameGrep      = "grep"
	ToolNameExecute   = "execute"

	noFilesFound   = "No files found"
	noMatchesFound = "No matches found"
)

type ToolConfig struct {
	Name string

	Desc *string

	CustomTool tool.BaseTool

	Disable bool
}

type Config struct {
	Backend filesystem.Backend

	Shell filesystem.Shell

	StreamingShell filesystem.StreamingShell

	LsToolConfig *ToolConfig

	ReadFileToolConfig *ToolConfig

	WriteFileToolConfig *ToolConfig

	EditFileToolConfig *ToolConfig

	GlobToolConfig *ToolConfig

	GrepToolConfig *ToolConfig

	WithoutLargeToolResultOffloading bool

	LargeToolResultOffloadingTokenLimit int

	LargeToolResultOffloadingPathGen func(ctx context.Context, input *compose.ToolInput) (string, error)

	CustomSystemPrompt *string

	CustomLsToolDesc *string

	CustomReadFileToolDesc *string

	CustomGrepToolDesc *string

	CustomGlobToolDesc *string

	CustomWriteFileToolDesc *string

	CustomEditToolDesc *string
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func NewMiddleware(ctx context.Context, config *Config) (adk.AgentMiddleware, error) {
	_ = "STUB: not implemented"
	return *new(adk.AgentMiddleware), nil
}

type MiddlewareConfig struct {
	Backend filesystem.Backend

	Shell filesystem.Shell

	StreamingShell filesystem.StreamingShell

	LsToolConfig *ToolConfig

	ReadFileToolConfig *ToolConfig

	WriteFileToolConfig *ToolConfig

	EditFileToolConfig *ToolConfig

	GlobToolConfig *ToolConfig

	GrepToolConfig *ToolConfig

	UseMultiModalRead bool

	CustomSystemPrompt *string

	CustomLsToolDesc *string

	CustomReadFileToolDesc *string

	CustomGrepToolDesc *string

	CustomGlobToolDesc *string

	CustomWriteFileToolDesc *string

	CustomEditToolDesc *string
}

func (c *MiddlewareConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *MiddlewareConfig) mergeToolConfigWithDesc(
	toolConfig *ToolConfig,
	legacyDesc *string,
) *ToolConfig {
	_ = "STUB: not implemented"
	return nil
}

func NewTyped[M adk.MessageType](ctx context.Context, config *MiddlewareConfig) (adk.TypedChatModelAgentMiddleware[M], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(ctx context.Context, config *MiddlewareConfig) (adk.ChatModelAgentMiddleware, error) {
	_ = "STUB: not implemented"
	return *new(adk.ChatModelAgentMiddleware), nil
}

type typedFilesystemMiddleware[M adk.MessageType] struct {
	*adk.TypedBaseChatModelAgentMiddleware[M]
	additionalInstruction string
	additionalTools       []tool.BaseTool
}

func (m *typedFilesystemMiddleware[M]) BeforeAgent(ctx context.Context, runCtx *adk.ChatModelAgentContext) (context.Context, *adk.ChatModelAgentContext, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

type toolSpec struct {
	config     *ToolConfig
	legacyDesc *string
	createFunc func(name, desc string) (tool.BaseTool, error)
}

func getFilesystemTools(_ context.Context, middlewareConfig *MiddlewareConfig) ([]tool.BaseTool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createToolFromSpec(middlewareConfig *MiddlewareConfig, spec toolSpec) (tool.BaseTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool), nil
}

func getOrCreateTool(customTool tool.BaseTool, createFunc func() (tool.BaseTool, error)) (tool.BaseTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool), nil
}

type lsArgs struct {
	Path string `json:"path"`
}

func newLsTool(fs filesystem.Backend, name string, desc string) (tool.BaseTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool), nil
}

type readFileArgs struct {
	FilePath string `json:"file_path" jsonschema:"description=The path to the file to read"`

	Offset int `json:"offset" jsonschema:"description=The line number to start reading from. Only provide if the file is too large to read at once"`

	Limit int `json:"limit" jsonschema:"description=The number of lines to read. Only provide if the file is too large to read at once."`
}

type multiModalReadFileArgs struct {
	readFileArgs

	Pages string `json:"pages,omitempty" jsonschema:"description=Page range for PDF files (e.g.\\, \"1-5\"\\, \"3\"\\, \"10-20\"). Only applicable to PDF files. Maximum 20 pages per request."`
}

func newReadFileTool(fs filesystem.Backend, name string, desc string) (tool.BaseTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool), nil
}

func formatReadResult(content, filePath string, startLine int) string {
	_ = "STUB: not implemented"
	return ""
}

func formatLineNumbers(content string, startLine int) string { _ = "STUB: not implemented"; return "" }

const maxPagesPerRequest = 20

func validatePages(pages string) error { _ = "STUB: not implemented"; return nil }

func newMultiModalReadFileTool(fs filesystem.Backend, name string, desc string) (tool.BaseTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool), nil
}

type writeFileArgs struct {
	FilePath string `json:"file_path" jsonschema:"description=The path to the file to write"`

	Content string `json:"content" jsonschema:"description=The content to write to the file"`
}

func newWriteFileTool(fs filesystem.Backend, name string, desc string) (tool.BaseTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool), nil
}

type editFileArgs struct {
	FilePath string `json:"file_path" jsonschema:"description=The path to the file to modify"`

	OldString string `json:"old_string" jsonschema:"description=The text to replace"`

	NewString string `json:"new_string" jsonschema:"description=The text to replace it with (must be different from old_string)"`

	ReplaceAll bool `json:"replace_all" jsonschema:"description=Replace all occurrences of old_string (default false),default=false"`
}

func newEditFileTool(fs filesystem.Backend, name string, desc string) (tool.BaseTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool), nil
}

type globArgs struct {
	Pattern string `json:"pattern" jsonschema:"description=The glob pattern to match files against"`

	Path string `json:"path" jsonschema:"description=The directory to search in. If not specified\\, the current working directory will be used. IMPORTANT: Omit this field to use the default directory. DO NOT enter 'undefined' or 'null' - simply omit it for the default behavior. Must be a valid directory path if provided."`
}

func newGlobTool(fs filesystem.Backend, name string, desc string) (tool.BaseTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool), nil
}

type grepArgs struct {
	Pattern string `json:"pattern" jsonschema:"description=The regular expression pattern to search for in file contents"`

	Path *string `json:"path,omitempty" jsonschema:"description=File or directory to search in (rg PATH). Defaults to current working directory."`

	Glob *string `json:"glob,omitempty" jsonschema:"description=Glob pattern to filter files (e.g. '*.js'\\, '*.{ts\\,tsx}') - maps to rg --glob"`

	OutputMode string `json:"output_mode,omitempty" jsonschema:"description=Output mode: 'content' shows matching lines (supports -A/-B/-C context\\, -n line numbers\\, head_limit)\\, 'files_with_matches' shows file paths (supports head_limit)\\, 'count' shows match counts (supports head_limit). Defaults to 'files_with_matches'.,enum=content,enum=files_with_matches,enum=count"`

	Context *int `json:"-C,omitempty" jsonschema:"description=Number of lines to show before and after each match (rg -C). Requires output_mode: 'content'\\, ignored otherwise."`

	BeforeLines *int `json:"-B,omitempty" jsonschema:"description=Number of lines to show before each match (rg -B). Requires output_mode: 'content'\\, ignored otherwise."`

	AfterLines *int `json:"-A,omitempty" jsonschema:"description=Number of lines to show after each match (rg -A). Requires output_mode: 'content'\\, ignored otherwise."`

	ShowLineNumbers *bool `json:"-n,omitempty" jsonschema:"description=Show line numbers in output (rg -n). Requires output_mode: 'content'\\, ignored otherwise. Defaults to true."`

	CaseInsensitive *bool `json:"-i,omitempty" jsonschema:"description=Case insensitive search (rg -i)"`

	FileType *string `json:"type,omitempty" jsonschema:"description=File type to search (rg --type). Common types: js\\, py\\, rust\\, go\\, java\\, etc. More efficient than include for standard file types."`

	HeadLimit *int `json:"head_limit,omitempty" jsonschema:"description=Limit output to first N lines/entries\\, equivalent to '| head -N'. Works across all output modes: content (limits output lines)\\, files_with_matches (limits file paths)\\, count (limits count entries). Defaults to 0 (unlimited)."`

	Offset *int `json:"offset,omitempty" jsonschema:"description=Skip first N lines/entries before applying head_limit\\, equivalent to '| tail -n +N | head -N'. Works across all output modes. Defaults to 0."`

	Multiline *bool `json:"multiline,omitempty" jsonschema:"description=Enable multiline mode where . matches newlines and patterns can span lines (rg -U --multiline-dotall). Default: false."`
}

func newGrepTool(fs filesystem.Backend, name string, desc string) (tool.BaseTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool), nil
}

type executeArgs struct {
	Command string `json:"command"`
}

func newExecuteTool(sb filesystem.Shell, name string, desc string) (tool.BaseTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool), nil
}

func newStreamingExecuteTool(sb filesystem.StreamingShell, name string, desc string) (tool.BaseTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.BaseTool), nil
}

func convExecuteResponse(response *filesystem.ExecuteResponse) string {
	_ = "STUB: not implemented"
	return ""
}

func valueOrDefault[T any](ptr *T, defaultValue T) T { _ = "STUB: not implemented"; return *new(T) }

type base64Encoder struct {
	buf []byte
}

func (e *base64Encoder) encode(data []byte) string { _ = "STUB: not implemented"; return "" }

func applyPagination[T any](items []T, offset, headLimit int) []T {
	_ = "STUB: not implemented"
	return nil
}

func formatFileMatches(matches []filesystem.GrepMatch, offset, headLimit int) string {
	_ = "STUB: not implemented"
	return ""
}

func formatContentMatches(matches []filesystem.GrepMatch, showLineNum bool) string {
	_ = "STUB: not implemented"
	return ""
}

func formatCountMatches(matches []filesystem.GrepMatch, offset, headLimit int) string {
	_ = "STUB: not implemented"
	return ""
}

func selectToolDesc(customDesc string, defaultEnglish, defaultChinese string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func selectToolName(customName string, defaultName string) string {
	_ = "STUB: not implemented"
	return ""
}
