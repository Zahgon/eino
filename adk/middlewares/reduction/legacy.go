package reduction

import "github.com/cloudwego/eino/adk/middlewares/reduction/internal"

type (
	ClearToolResultConfig = internal.ClearToolResultConfig
	ToolResultConfig      = internal.ToolResultConfig
	Backend               = internal.Backend
)

var (
	NewClearToolResult = internal.NewClearToolResult

	NewToolResultMiddleware = internal.NewToolResultMiddleware
)
