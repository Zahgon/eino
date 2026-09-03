package skill

import (
	"context"

	"github.com/cloudwego/eino/adk/filesystem"
)

const skillFileName = "SKILL.md"

type filesystemBackend struct {
	backend filesystem.Backend
	baseDir string
}

type BackendFromFilesystemConfig struct {
	Backend filesystem.Backend

	BaseDir string
}

func NewBackendFromFilesystem(_ context.Context, config *BackendFromFilesystemConfig) (Backend, error) {
	_ = "STUB: not implemented"
	return *new(Backend), nil
}

func (b *filesystemBackend) List(ctx context.Context) ([]FrontMatter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *filesystemBackend) Get(ctx context.Context, name string) (Skill, error) {
	_ = "STUB: not implemented"
	return *new(Skill), nil
}

func (b *filesystemBackend) list(ctx context.Context) ([]Skill, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *filesystemBackend) loadSkillFromFile(ctx context.Context, path string) (Skill, error) {
	_ = "STUB: not implemented"
	return *new(Skill), nil
}

func stripLineNumbers(data string) string { _ = "STUB: not implemented"; return "" }

func parseFrontmatter(data string) (frontmatter string, content string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
