package adk

import "github.com/cloudwego/eino/adk/internal"

type Language = internal.Language

const (
	LanguageEnglish Language = internal.LanguageEnglish

	LanguageChinese Language = internal.LanguageChinese
)

func SetLanguage(lang Language) error { _ = "STUB: not implemented"; return nil }
