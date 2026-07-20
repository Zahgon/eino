package reduction

const (
	truncFmt = `<persisted-output>
Output too large ({original_size}). Full output saved to: {file_path}
Preview (first {preview_size}):
{preview_first}

Preview (last {preview_size}):
{preview_last}

Use {read_tool_name} to view
</persisted-output>`
	truncFmtZh = `<persisted-output>
输出结果过大 ({original_size}). 完整输出保存到: {file_path}
预览 (前 {preview_size}):
{preview_first}

预览 (后 {preview_size}):
{preview_last}

使用 {read_tool_name} 进行查看
</persisted-output>`
)

const (
	streamTruncFmt = `<persisted-output>
Output truncated after {preview_size} bytes were streamed. {offload_notify} {error_msg_notify}
</persisted-output>`
	streamTruncFmtZh = `<persisted-output>
输出结果在流式传输 {preview_size} 字节后被截断。{offload_notify} {error_msg_notify}
</persisted-output>`
)

const (
	clearWithOffloadingFmt = `<persisted-output>Tool result saved to: {file_path}
Use {read_tool_name} to view</persisted-output>`
	clearWithOffloadingFmtZh = `<persisted-output>工具结果已保存至: {file_path}
使用 {read_tool_name} 进行查看</persisted-output>`

	clearWithoutOffloadingFmt   = `[Old tool result content cleared]`
	clearWithoutOffloadingFmtZh = `[工具输出结果已清理]`
)

const (
	msgClearedFlag = "_reduction_mw_processed"
)

func getTruncFmt() string { _ = "STUB: not implemented"; return "" }

func getStreamTruncFmt() string { _ = "STUB: not implemented"; return "" }

func formatStreamOffloadSavedNotify(filePath, readFileToolName string) string {
	_ = "STUB: not implemented"
	return ""
}

func formatStreamOffloadFailedNotify(err error) string { _ = "STUB: not implemented"; return "" }

func getClearWithOffloadingFmt() string { _ = "STUB: not implemented"; return "" }

func getClearWithoutOffloadingFmt() string { _ = "STUB: not implemented"; return "" }

type scene int

const (
	sceneTruncation scene = 1
	sceneClear      scene = 2
)
