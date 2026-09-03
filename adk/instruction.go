package adk

import (
	"context"
)

const (
	TransferToAgentInstruction = `Available other agents: %s

Decision rule:
- If you're best suited for the question according to your description: ANSWER
- If another agent is better according its description: CALL '%s' function with their agent name

When transferring: OUTPUT ONLY THE FUNCTION CALL`

	TransferToAgentInstructionChinese = `可用的其他 agent：%s

决策规则：
- 如果根据你的职责描述，你最适合回答这个问题：ANSWER
- 如果根据其职责描述，另一个 agent 更适合：调用 %s 函数，并传入该 agent 的名称

当进行移交时：只输出函数调用，不要输出其他任何内容`

	agentDescriptionTpl        = "\n- Agent name: %s\n  Agent description: %s"
	agentDescriptionTplChinese = "\n- Agent 名字: %s\n  Agent 描述: %s"
)

func genTransferToAgentInstruction[M MessageType](ctx context.Context, agents []TypedAgent[M]) string {
	_ = "STUB: not implemented"
	return ""
}
