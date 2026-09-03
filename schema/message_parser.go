package schema

import (
	"context"
)

type MessageParser[T any] interface {
	Parse(ctx context.Context, m *Message) (T, error)
}

type MessageParseFrom string

const (
	MessageParseFromContent  MessageParseFrom = "content"
	MessageParseFromToolCall MessageParseFrom = "tool_call"
)

type MessageJSONParseConfig struct {
	ParseFrom MessageParseFrom `json:"parse_from,omitempty"`

	ParseKeyPath string `json:"parse_key_path,omitempty"`
}

func NewMessageJSONParser[T any](config *MessageJSONParseConfig) MessageParser[T] {
	_ = "STUB: not implemented"
	return nil
}

type MessageJSONParser[T any] struct {
	ParseFrom    MessageParseFrom
	ParseKeyPath string
}

func (p *MessageJSONParser[T]) Parse(ctx context.Context, m *Message) (parsed T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (p *MessageJSONParser[T]) extractData(data string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *MessageJSONParser[T]) parse(data string) (parsed T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
