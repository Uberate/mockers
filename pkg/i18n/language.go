package i18n

import "strings"

// LanguageKey is an alias of string to save the language name.
type LanguageKey string

const (
	EN   LanguageKey = "EN"
	ZHCN LanguageKey = "ZHCN"
)

func GetLanguageKey(ln string) LanguageKey {
	switch strings.ToLower(ln) {
	case "en":
		return EN
	case "zhcn":
		return ZHCN
	default:
		return EN
	}
}
