package schema

import (
	"reflect"
)

func init() {
	RegisterName[*Message]("_eino_message")
	RegisterName[[]*Message]("_eino_message_slice")
	RegisterName[*AgenticMessage]("_eino_agentic_message")
	RegisterName[[]*AgenticMessage]("_eino_agentic_message_slice")
	RegisterName[Document]("_eino_document")
	RegisterName[RoleType]("_eino_role_type")
	RegisterName[ToolCall]("_eino_tool_call")
	RegisterName[FunctionCall]("_eino_function_call")
	RegisterName[ResponseMeta]("_eino_response_meta")
	RegisterName[TokenUsage]("_eino_token_usage")
	RegisterName[LogProbs]("_eino_log_probs")
	RegisterName[ChatMessagePart]("_eino_chat_message_part")
	RegisterName[ChatMessagePartType]("_eino_chat_message_type")
	RegisterName[ChatMessageImageURL]("_eino_chat_message_image_url")
	RegisterName[ChatMessageAudioURL]("_eino_chat_message_audio_url")
	RegisterName[ChatMessageVideoURL]("_eino_chat_message_video_url")
	RegisterName[ChatMessageFileURL]("_eino_chat_message_file_url")
	RegisterName[MessageInputPart]("_eino_message_input_part")
	RegisterName[MessageInputImage]("_eino_message_input_image")
	RegisterName[MessageInputAudio]("_eino_message_input_audio")
	RegisterName[MessageInputVideo]("_eino_message_input_video")
	RegisterName[MessageInputFile]("_eino_message_input_file")
	RegisterName[MessageOutputPart]("_eino_message_output_part")
	RegisterName[MessageOutputImage]("_eino_message_output_image")
	RegisterName[MessageOutputAudio]("_eino_message_output_audio")
	RegisterName[MessageOutputVideo]("_eino_message_output_video")
	RegisterName[MessagePartCommon]("_eino_message_part_common")
	RegisterName[ImageURLDetail]("_eino_image_url_detail")
	RegisterName[PromptTokenDetails]("_eino_prompt_token_details")
}

func RegisterName[T any](name string) { _ = "STUB: not implemented"; return }

func getTypeName(rt reflect.Type) string { _ = "STUB: not implemented"; return "" }

func Register[T any]() { _ = "STUB: not implemented"; return }
