package plantask

import (
	"context"
	"sync"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

func newTaskUpdateTool(backend Backend, baseDir string, lock *sync.Mutex) *taskUpdateTool {
	_ = "STUB: not implemented"
	return nil
}

type taskUpdateTool struct {
	Backend Backend
	BaseDir string
	lock    *sync.Mutex
}

type taskUpdateArgs struct {
	TaskID       string         `json:"taskId"`
	Subject      string         `json:"subject,omitempty"`
	Description  string         `json:"description,omitempty"`
	ActiveForm   string         `json:"activeForm,omitempty"`
	Status       string         `json:"status,omitempty"`
	AddBlocks    []string       `json:"addBlocks,omitempty"`
	AddBlockedBy []string       `json:"addBlockedBy,omitempty"`
	Owner        string         `json:"owner,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`
}

func (t *taskUpdateTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *taskUpdateTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *taskUpdateTool) removeTaskFromDependencies(ctx context.Context, deletedTaskID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *taskUpdateTool) addBlockedByToTask(ctx context.Context, targetTaskID, blockerTaskID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *taskUpdateTool) addBlocksToTask(ctx context.Context, targetTaskID, blockedTaskID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *taskUpdateTool) checkIfNeedDeleteAllTasks(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

const TaskUpdateToolName = "TaskUpdate"
const taskUpdateToolDesc = `Use this tool to update a task in the task list.

## When to Use This Tool

**Mark tasks as resolved:**
- When you have completed the work described in a task
- When a task is no longer needed or has been superseded
- IMPORTANT: Always mark your assigned tasks as resolved when you finish them
- After resolving, call TaskList to find your next task

- ONLY mark a task as completed when you have FULLY accomplished it
- If you encounter errors, blockers, or cannot finish, keep the task as in_progress
- When blocked, create a new task describing what needs to be resolved
- Never mark a task as completed if:
  - Tests are failing
  - Implementation is partial
  - You encountered unresolved errors
  - You couldn't find necessary files or dependencies

**Delete tasks:**
- When a task is no longer relevant or was created in error
- Setting status to ` + "`deleted`" + ` permanently removes the task

**Update task details:**
- When requirements change or become clearer
- When establishing dependencies between tasks

## Fields You Can Update

- **status**: The task status (see Status Workflow below)
- **subject**: Change the task title (imperative form, e.g., "Run tests")
- **description**: Change the task description
- **activeForm**: Present continuous form shown in spinner when in_progress (e.g., "Running tests")
- **owner**: Change the task owner (agent name)
- **metadata**: Merge metadata keys into the task (set a key to null to delete it)
- **addBlocks**: Mark tasks that cannot start until this one completes
- **addBlockedBy**: Mark tasks that must complete before this one can start

## Status Workflow

Status progresses: ` + "`pending`" + ` → ` + "`in_progress`" + ` → ` + "`completed`" + `

Use ` + "`deleted`" + ` to permanently remove a task.

## Staleness

Make sure to read a task's latest state using ` + "`TaskGet`" + ` before updating it.

## Examples

Mark task as in progress when starting work:
` + "```json" + `
{"taskId": "1", "status": "in_progress"}
` + "```" + `

Mark task as completed after finishing work:
` + "```json" + `
{"taskId": "1", "status": "completed"}
` + "```" + `

Delete a task:
` + "```json" + `
{"taskId": "1", "status": "deleted"}
` + "```" + `

Claim a task by setting owner:
` + "```json" + `
{"taskId": "1", "owner": "my-name"}
` + "```" + `

Set up task dependencies:
` + "```json" + `
{"taskId": "2", "addBlockedBy": ["1"]}
` + "```" + `
`

const taskUpdateToolDescChinese = `使用此工具更新任务列表中的任务。

## 何时使用此工具

**将任务标记为已完成：**
- 当你完成了任务中描述的工作时
- 当任务不再需要或已被取代时
- 重要：完成分配给你的任务后，务必将其标记为已完成
- 完成后，调用 TaskList 查找下一个任务

- 只有在完全完成任务时才将其标记为已完成
- 如果遇到错误、阻塞或无法完成，请保持任务为 in_progress 状态
- 当被阻塞时，创建一个新任务描述需要解决的问题
- 在以下情况下不要将任务标记为已完成：
  - 测试失败
  - 实现不完整
  - 遇到未解决的错误
  - 找不到必要的文件或依赖项

**删除任务：**
- 当任务不再相关或创建错误时
- 将状态设置为 ` + "`deleted`" + ` 会永久删除任务

**更新任务详情：**
- 当需求变更或变得更清晰时
- 当建立任务之间的依赖关系时

## 可更新的字段

- **status**：任务状态（参见下方状态流程）
- **subject**：更改任务标题（使用祈使句形式，例如"运行测试"）
- **description**：更改任务描述
- **activeForm**：in_progress 状态时在加载动画中显示的现在进行时形式（例如"正在运行测试"）
- **owner**：更改任务所有者（代理名称）
- **metadata**：将元数据键合并到任务中（将键设置为 null 可删除它）
- **addBlocks**：标记在此任务完成之前无法开始的任务
- **addBlockedBy**：标记必须在此任务开始之前完成的任务

## 状态流程

状态进展：` + "`pending`" + ` → ` + "`in_progress`" + ` → ` + "`completed`" + `

使用 ` + "`deleted`" + ` 永久删除任务。

## 过期性

更新任务前，请确保使用 ` + "`TaskGet`" + ` 读取任务的最新状态。

## 示例

开始工作时将任务标记为进行中：
` + "```json" + `
{"taskId": "1", "status": "in_progress"}
` + "```" + `

完成工作后将任务标记为已完成：
` + "```json" + `
{"taskId": "1", "status": "completed"}
` + "```" + `

删除任务：
` + "```json" + `
{"taskId": "1", "status": "deleted"}
` + "```" + `

通过设置 owner 认领任务：
` + "```json" + `
{"taskId": "1", "owner": "my-name"}
` + "```" + `

设置任务依赖关系：
` + "```json" + `
{"taskId": "2", "addBlockedBy": ["1"]}
` + "```" + `
`
