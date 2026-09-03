package filesystem

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

type FileInfo struct {
	Path string

	IsDir bool

	Size int64

	ModifiedAt string
}

type GrepMatch struct {
	Content string

	Path string

	Line int
}

type LsInfoRequest struct {
	Path string
}

type ReadRequest struct {
	FilePath string

	Offset int

	Limit int
}

type MultiModalReadRequest struct {
	ReadRequest

	Pages string
}

type GrepRequest struct {
	Pattern string

	Path string

	Glob string

	FileType string

	CaseInsensitive bool

	EnableMultiline bool

	AfterLines int

	BeforeLines int
}

type GlobInfoRequest struct {
	Pattern string

	Path string
}

type WriteRequest struct {
	FilePath string

	Content string
}

type EditRequest struct {
	FilePath string

	OldString string

	NewString string

	ReplaceAll bool
}

type FileContentPartType string

const (
	FileContentPartTypeImage FileContentPartType = "image"

	FileContentPartTypePDF FileContentPartType = "pdf"
)

type FileContentPart struct {
	Type FileContentPartType

	MIMEType string

	Data []byte
}

type FileContent struct {
	Content string
}

type MultiFileContent struct {
	*FileContent

	Parts []FileContentPart
}

type MultiModalReader interface {
	MultiModalRead(ctx context.Context, req *MultiModalReadRequest) (*MultiFileContent, error)
}

type Backend interface {
	LsInfo(ctx context.Context, req *LsInfoRequest) ([]FileInfo, error)

	Read(ctx context.Context, req *ReadRequest) (*FileContent, error)

	GrepRaw(ctx context.Context, req *GrepRequest) ([]GrepMatch, error)

	GlobInfo(ctx context.Context, req *GlobInfoRequest) ([]FileInfo, error)

	Write(ctx context.Context, req *WriteRequest) error

	Edit(ctx context.Context, req *EditRequest) error
}

type ExecuteRequest struct {
	Command            string
	RunInBackendGround bool
}

type ExecuteResponse struct {
	Output    string
	ExitCode  *int
	Truncated bool
}

type Shell interface {
	Execute(ctx context.Context, input *ExecuteRequest) (result *ExecuteResponse, err error)
}

type StreamingShell interface {
	ExecuteStreaming(ctx context.Context, input *ExecuteRequest) (result *schema.StreamReader[*ExecuteResponse], err error)
}
