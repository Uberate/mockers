package i18n

import "strings"

// LanguageKey is an alias of string to save the language name.
type LanguageKey string

// ToString will return the LanguageKey and auto turn to an upper word. The language key follow the `ISO 639-1`.
func (lk LanguageKey) ToString() string {
	return strings.ToUpper(string(lk))
}

const (
	EN            LanguageKey = "EN"
	ZHCN          LanguageKey = "ZHCN"
	EmptyLanguage LanguageKey = "__empty__"
)

func GetAllLanguages() map[string]LanguageKey {
	return map[string]LanguageKey{
		EN.ToString():   EN,
		ZHCN.ToString(): ZHCN,
	}
}

// GetLanguageDescribe may be should an i18n base.
func GetLanguageDescribe() map[string]string {
	return map[string]string{
		EN.ToString():   "English",
		ZHCN.ToString(): "中文",
	}
}

func GetLanguageKey(ln string) LanguageKey {
	if value, ok := GetAllLanguages()[strings.ToUpper(ln)]; ok {
		return value
	} else {
		return EmptyLanguage
	}
}
