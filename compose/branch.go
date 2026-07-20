package compose

import (
	"context"
	"reflect"

	"github.com/cloudwego/eino/schema"
)

type GraphBranchCondition[T any] func(ctx context.Context, in T) (endNode string, err error)

type StreamGraphBranchCondition[T any] func(ctx context.Context, in *schema.StreamReader[T]) (endNode string, err error)

type GraphMultiBranchCondition[T any] func(ctx context.Context, in T) (endNode map[string]bool, err error)

type StreamGraphMultiBranchCondition[T any] func(ctx context.Context, in *schema.StreamReader[T]) (endNodes map[string]bool, err error)

type GraphBranch struct {
	invoke    func(ctx context.Context, input any) (output []string, err error)
	collect   func(ctx context.Context, input streamReader) (output []string, err error)
	inputType reflect.Type
	*genericHelper
	endNodes   map[string]bool
	idx        int
	noDataFlow bool
}

func (gb *GraphBranch) GetEndNode() map[string]bool { _ = "STUB: not implemented"; return nil }

func newGraphBranch[T any](r *runnablePacker[T, []string, any], endNodes map[string]bool) *GraphBranch {
	_ = "STUB: not implemented"
	return nil
}

func NewGraphMultiBranch[T any](condition GraphMultiBranchCondition[T], endNodes map[string]bool) *GraphBranch {
	_ = "STUB: not implemented"
	return nil
}

func NewStreamGraphMultiBranch[T any](condition StreamGraphMultiBranchCondition[T],
	endNodes map[string]bool) *GraphBranch {
	_ = "STUB: not implemented"
	return nil
}

func NewGraphBranch[T any](condition GraphBranchCondition[T], endNodes map[string]bool) *GraphBranch {
	_ = "STUB: not implemented"
	return nil
}

func NewStreamGraphBranch[T any](condition StreamGraphBranchCondition[T], endNodes map[string]bool) *GraphBranch {
	_ = "STUB: not implemented"
	return nil
}
