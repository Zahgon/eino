package plantask

import (
	"context"
	"sync"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

func newTaskGetTool(backend Backend, baseDir string, lock *sync.Mutex) *taskGetTool {
	_ = "STUB: not implemented"
	return nil
}

type taskGetTool struct {
	Backend Backend
	BaseDir string
	lock    *sync.Mutex
}

func (t *taskGetTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type taskGetArgs struct {
	TaskID string `json:"taskId"`
}

func (t *taskGetTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

const TaskGetToolName = "TaskGet"
const taskGetToolDesc = `Use this tool to retrieve a task by its ID from the task list.

## When to Use This Tool

- When you need the full description and context before starting work on a task
- To understand task dependencies (what it blocks, what blocks it)
- After being assigned a task, to get complete requirements

## Output

Returns full task details:
- **subject**: Task title
- **description**: Detailed requirements and context
- **status**: 'pending', 'in_progress', or 'completed'
- **blocks**: Tasks waiting on this one to complete
- **blockedBy**: Tasks that must complete before this one can start

## Tips

- After fetching a task, verify its blockedBy list is empty before beginning work.
- Use TaskList to see all tasks in summary form.
`

const taskGetToolDescChinese = `使用此工具通过任务 ID 从任务列表中获取任务。

## 何时使用此工具

- 当你需要在开始处理任务之前获取完整的描述和上下文时
- 了解任务依赖关系（它阻塞什么，什么阻塞它）
- 被分配任务后，获取完整的需求

## 输出

返回完整的任务详情：
- **subject**：任务标题
- **description**：详细的需求和上下文
- **status**：'pending'、'in_progress' 或 'completed'
- **blocks**：等待此任务完成的任务
- **blockedBy**：必须在此任务开始之前完成的任务

## 提示

- 获取任务后，在开始工作之前验证其 blockedBy 列表是否为空。
- 使用 TaskList 查看所有任务的摘要形式。
`
