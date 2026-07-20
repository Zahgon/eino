package supervisor

import (
	"context"

	"github.com/cloudwego/eino/adk"
)

type Config struct {
	Supervisor adk.Agent

	SubAgents []adk.Agent
}

type supervisorContainer struct {
	name  string
	inner adk.ResumableAgent
}

func (s *supervisorContainer) Name(_ context.Context) string { _ = "STUB: not implemented"; return "" }

func (s *supervisorContainer) Description(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *supervisorContainer) GetType() string { _ = "STUB: not implemented"; return "" }

func (s *supervisorContainer) Run(ctx context.Context, input *adk.AgentInput, opts ...adk.AgentRunOption) *adk.AsyncIterator[*adk.AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

func (s *supervisorContainer) Resume(ctx context.Context, info *adk.ResumeInfo, opts ...adk.AgentRunOption) *adk.AsyncIterator[*adk.AgentEvent] {
	_ = "STUB: not implemented"
	return nil
}

func New(ctx context.Context, conf *Config) (adk.ResumableAgent, error) {
	_ = "STUB: not implemented"
	return *new(adk.ResumableAgent), nil
}
