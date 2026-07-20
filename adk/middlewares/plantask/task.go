package plantask

import (
	"context"
	"regexp"

	"github.com/cloudwego/eino/adk/middlewares/filesystem"
)

var validTaskIDRegex = regexp.MustCompile(`^\d+$`)

const highWatermarkFileName = ".highwatermark"

type task struct {
	ID          string         `json:"id"`
	Subject     string         `json:"subject"`
	Description string         `json:"description"`
	Status      string         `json:"status"`
	Blocks      []string       `json:"blocks"`
	BlockedBy   []string       `json:"blockedBy"`
	ActiveForm  string         `json:"activeForm,omitempty"`
	Owner       string         `json:"owner,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}

type taskOut struct {
	Result string `json:"result"`
}

const (
	taskStatusPending    = "pending"
	taskStatusInProgress = "in_progress"
	taskStatusCompleted  = "completed"
	taskStatusDeleted    = "deleted"
)

type FileInfo = filesystem.FileInfo
type LsInfoRequest = filesystem.LsInfoRequest
type ReadRequest = filesystem.ReadRequest
type WriteRequest = filesystem.WriteRequest

type DeleteRequest struct {
	FilePath string
}

type Backend interface {
	LsInfo(ctx context.Context, req *LsInfoRequest) ([]FileInfo, error)

	Read(ctx context.Context, req *ReadRequest) (*filesystem.FileContent, error)

	Write(ctx context.Context, req *WriteRequest) error

	Delete(ctx context.Context, req *DeleteRequest) error
}

func isValidTaskID(taskID string) bool { _ = "STUB: not implemented"; return false }

func appendUnique(slice []string, items ...string) []string { _ = "STUB: not implemented"; return nil }

func hasCyclicDependency(taskMap map[string]*task, blockerID, blockedID string) bool {
	_ = "STUB: not implemented"
	return false
}

func canReach(taskMap map[string]*task, fromID, toID string, visited map[string]bool) bool {
	_ = "STUB: not implemented"
	return false
}
