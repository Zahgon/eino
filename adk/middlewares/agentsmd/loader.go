package agentsmd

import (
	"context"
	"regexp"

	"github.com/cloudwego/eino/adk/filesystem"
)

var importRegex = regexp.MustCompile(`@([a-zA-Z0-9_.~/][a-zA-Z0-9_.~/\-]*)`)

var allowedImportExts = map[string]bool{
	".md":   true,
	".txt":  true,
	".mdx":  true,
	".yaml": true,
	".yml":  true,
	".json": true,
	".toml": true,
}

const maxImportDepth = 5

const minTruncatableBytes = 500

type ReadRequest = filesystem.ReadRequest
type FileContent = filesystem.FileContent

type Backend interface {
	Read(ctx context.Context, req *ReadRequest) (*FileContent, error)
}

type loaderConfig struct {
	backend      Backend
	files        []string
	maxBytes     int
	perFileBytes int
	onWarning    func(filePath string, err error)
}

func newLoaderConfig(backend Backend, files []string, maxBytes, perFileBytes int, onWarning func(filePath string, err error)) *loaderConfig {
	_ = "STUB: not implemented"
	return nil
}

type loader struct {
	*loaderConfig
	totalBytes int
	stopped    bool
}

func (cfg *loaderConfig) newLoader() *loader { _ = "STUB: not implemented"; return nil }

func (cfg *loaderConfig) load(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (l *loader) loadFile(ctx context.Context, filePath string, depth int, visited map[string]bool, seen map[string]bool) ([]loadedFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *loader) applyBudget(filePath, content string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (l *loader) collectImports(ctx context.Context, hostPath, content string, depth int, visited map[string]bool, seen map[string]bool) ([]loadedFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type loadedFile struct {
	path    string
	content string
}

func truncateBytes(s string, max int) string { _ = "STUB: not implemented"; return "" }

func truncationNotice(cut, total int) string { _ = "STUB: not implemented"; return "" }

const formatHeaderEn = `<system-reminder>
As you answer the user's questions, you can use the following context:
Codebase and user instructions are shown below. Be sure to adhere to these instructions. IMPORTANT: These instructions OVERRIDE any default behavior and you MUST follow them exactly as written.
`

const formatHeaderCn = `<system-reminder>
在回答用户问题时，你可以使用以下上下文：
代码库和用户指令如下。请务必遵守这些指令。重要提示：这些指令会覆盖任何默认行为，你必须严格按照要求执行。
`

const formatFileHeaderEn = "\nContents of "

const formatFileHeaderCn = "\n文件内容："

const formatFileLabelEn = " (instructions):\n\n"

const formatFileLabelCn = "（指令）：\n\n"

const formatFooterEn = `IMPORTANT: this context may or may not be relevant to your tasks. You should not respond to this context unless it is highly relevant to your task.
</system-reminder>`

const formatFooterCn = `重要提示：此上下文可能与你的任务相关，也可能不相关。除非此上下文与你的任务高度相关，否则不要响应此上下文。
</system-reminder>`

func omittedNotice(paths []string) string { _ = "STUB: not implemented"; return "" }

func formatContent(files []loadedFile, omitted []string) string {
	_ = "STUB: not implemented"
	return ""
}
