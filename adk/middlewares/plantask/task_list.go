package plantask

import (
	"context"
	"sync"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

func newTaskListTool(backend Backend, baseDir string, lock *sync.Mutex) *taskListTool {
	_ = "STUB: not implemented"
	return nil
}

type taskListTool struct {
	Backend Backend
	BaseDir string
	lock    *sync.Mutex
}

func (t *taskListTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listTasks(ctx context.Context, backend Backend, baseDir string) ([]*task, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *taskListTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

const TaskListToolName = "TaskList"
const taskListToolDesc = `Use this tool to list all tasks in the task list.

## When to Use This Tool

- To see what tasks are available to work on (status: 'pending', no owner, not blocked)
- To check overall progress on the project
- To find tasks that are blocked and need dependencies resolved
- After completing a task, to check for newly unblocked work or claim the next available task
- **Prefer working on tasks in ID order** (lowest ID first) when multiple tasks are available, as earlier tasks often set up context for later ones

## Output

Returns a summary of each task:
- **id**: Task identifier (use with TaskGet, TaskUpdate)
- **subject**: Brief description of the task
- **status**: 'pending', 'in_progress', or 'completed'
- **owner**: Agent ID if assigned, empty if available
- **blockedBy**: List of open task IDs that must be resolved first (tasks with blockedBy cannot be claimed until dependencies resolve)

Use TaskGet with a specific task ID to view full details including description and comments.
`

const taskListToolDescChinese = `使用此工具列出任务列表中的所有任务。

## 何时使用此工具

- 查看可以处理的任务（状态：'pending'，无所有者，未被阻塞）
- 检查项目的整体进度
- 查找被阻塞且需要解决依赖关系的任务
- 完成任务后，检查新解除阻塞的工作或认领下一个可用任务
- **优先按 ID 顺序处理任务**（最小 ID 优先），当有多个任务可用时，因为较早的任务通常为后续任务建立上下文

## 输出

返回每个任务的摘要：
- **id**：任务标识符（与 TaskGet、TaskUpdate 一起使用）
- **subject**：任务的简要描述
- **status**：'pending'、'in_progress' 或 'completed'
- **owner**：如果已分配则为代理 ID，如果可用则为空
- **blockedBy**：必须首先解决的开放任务 ID 列表（具有 blockedBy 的任务在依赖关系解决之前无法被认领）

使用 TaskGet 配合特定任务 ID 查看完整详情，包括描述和评论。
`
