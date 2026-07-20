package callbacks

import (
	"github.com/cloudwego/eino/internal/callbacks"
)

type RunInfo = callbacks.RunInfo

type CallbackInput = callbacks.CallbackInput

type CallbackOutput = callbacks.CallbackOutput

type Handler = callbacks.Handler

func InitCallbackHandlers(handlers []Handler) { _ = "STUB: not implemented"; return }

func AppendGlobalHandlers(handlers ...Handler) { _ = "STUB: not implemented"; return }

type CallbackTiming = callbacks.CallbackTiming

const (
	TimingOnStart CallbackTiming = iota

	TimingOnEnd

	TimingOnError

	TimingOnStartWithStreamInput

	TimingOnEndWithStreamOutput
)

type TimingChecker = callbacks.TimingChecker
