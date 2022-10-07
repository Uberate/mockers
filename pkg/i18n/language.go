package i18n

import "strings"

// LanguageKey is an alias of string to save the language name.
type LanguageKey string

// ToString will return the LanguageKey and auto turn to an upper word.
func (lk LanguageKey) ToString() string {
	return strings.ToUpper(string(lk))
}

const (
	EN            LanguageKey = "EN"
	ZHCN          LanguageKey = "ZHCN"
	EmptyLanguage LanguageKey = "__empty__"
)

func GetLanguageKey(ln string) LanguageKey {
	switch strings.ToUpper(ln) {
	case EN.ToString():
		return EN
	case ZHCN.ToString():
		return ZHCN
	default:
		return EmptyLanguage
	}
}
