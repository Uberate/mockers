package i18n

import "fmt"

// MessageObject I18n quick struct.
type MessageObject struct {
	Language  LanguageKey `json:"language" yaml:"language" mapstructure:"language"`
	Namespace string      `json:"namespace" yaml:"namespace" mapstructure:"language"`
	Code      string      `json:"code" yaml:"code" mapstructure:"language"`
	Message   string      `json:"message" yaml:"message" mapstructure:"language"`
}

type I18n struct {
	messages map[string]map[string]map[LanguageKey]string

	DefaultLanguage LanguageKey

	EnableChange bool
}

func (i *I18n) Len() int {
	var res = 0

	i.WalkMessage(func(ln LanguageKey, namespace, code, message string) {
		res++
	})

	return res
}

func (i *I18n) ToMessageObjects() []MessageObject {
	var res []MessageObject

	i.WalkMessage(func(ln LanguageKey, namespace, code, message string) {
		obj := MessageObject{
			Language:  ln,
			Namespace: namespace,
			Code:      code,
			Message:   message,
		}
		res = append(res, obj)
	})
	return res
}

func (i *I18n) RegisterMessage(ln LanguageKey, namespace, code, message string) {
	if i.EnableChange {
		if _, ok := i.messages[namespace]; !ok {
			i.messages[namespace] = map[string]map[LanguageKey]string{}
		}
		if _, ok := i.messages[namespace][code]; !ok {
			i.messages[namespace][code] = map[LanguageKey]string{}
		}
		i.messages[namespace][code][ln] = message
	}
}

func (i *I18n) RegisterMessageObject(object MessageObject) {
	i.RegisterMessage(object.Language, object.Namespace, object.Code, object.Message)
}

func (i *I18n) WalkMessage(f func(ln LanguageKey, namespace, code, message string)) {
	for namespace, messages := range i.messages {
		for code, languageMessages := range messages {
			for language, message := range languageMessages {
				f(language, namespace, code, message)
			}
		}
	}
}

func (i *I18n) MessageWithParam(ln LanguageKey, namespace, code string, param ...any) (string, bool) {
	if res, ok := i.Message(ln, namespace, code); ok {
		return fmt.Sprintf(res, param...), true
	} else {
		return res, ok
	}
}

func (i *I18n) Message(ln LanguageKey, namespace, code string) (string, bool) {
	if _, ok := i.messages[namespace]; !ok {
		return fmt.Sprintf("%s-%s", namespace, code), false
	}
	if _, ok := i.messages[namespace][code]; !ok {
		return fmt.Sprintf("%s-%s", namespace, code), false
	}

	if message, ok := i.messages[namespace][code][ln]; ok {
		return message, true
	}

	if i.DefaultLanguage != EmptyLanguage && i.DefaultLanguage != ln {
		if message, ok := i.messages[namespace][code][i.DefaultLanguage]; ok {
			return message, true
		}
	}
	return fmt.Sprintf("%s-%s", namespace, code), false

}
