package parser

import (
	"context"
	"io"

	"github.com/cloudwego/eino/schema"
)

type ExtParserConfig struct {
	Parsers map[string]Parser

	FallbackParser Parser
}

type ExtParser struct {
	parsers map[string]Parser

	fallbackParser Parser
}

func NewExtParser(ctx context.Context, conf *ExtParserConfig) (*ExtParser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ExtParser) GetParsers() map[string]Parser { _ = "STUB: not implemented"; return nil }

func (p *ExtParser) Parse(ctx context.Context, reader io.Reader, opts ...Option) ([]*schema.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
