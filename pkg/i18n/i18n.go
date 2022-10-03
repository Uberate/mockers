// Package i18n provides some function to quick deal the i18n problem. And for some quick deal with it.
//
// The i18n build the message info like it:
// - ln -> namespace -> code -> message.
//
// The ln is the language name. The namespace is the namespace of the message info. It like a middle layout in the
// i18n system. Like error, user, application and so on. In different namespace, the code can be same.
package i18n

import "fmt"

// I18n --------------------------------------------------

// I18n is the pool of i18n infos.
type I18n struct {
	Languages       map[LanguageKey]*Language
	DefaultLanguage LanguageKey
}

// MessageWithParam return the template message info with params. The base is invoked the fmt.Sprintf(). If not found,
// return string in format 'namespace-code' and false.
func (i *I18n) MessageWithParam(ln LanguageKey, namespace, code string, params ...any) (string, bool) {
	if languageItem, ok := i.Languages[ln]; ok {
		return languageItem.MessageWithParam(namespace, code, params...)
	}

	return "", false
}

// Message return the message info. Different from the MessageWithParam, it is return value directly. If not found
// specify value, return string in format 'namespace-code' and false.
func (i *I18n) Message(ln LanguageKey, namespace, code string) (string, bool) {
	if languageItem, ok := i.Languages[ln]; ok {
		return languageItem.Message(namespace, code)
	}
	return "", false
}

// RegisterMessage will register a language to the namespace.
func (i *I18n) RegisterMessage(ln LanguageKey, namespace, code, message string) {
	if i.Languages == nil {
		i.Languages = map[LanguageKey]*Language{}
	}
	if _, ok := i.Languages[ln]; !ok {
		i.Languages[ln] = &Language{
			Language:   ln,
			Namespaces: map[string]*Namespace{},
		}
	}
	i.Languages[ln].RegisterMessage(namespace, code, message)
}

// Language --------------------------------------------------

// Language save one language message info.
type Language struct {
	Language   LanguageKey
	Namespaces map[string]*Namespace
}

func (l *Language) MessageWithParam(namespace, code string, param ...any) (string, bool) {
	if namespaceItem, ok := l.Namespaces[namespace]; ok {
		return namespaceItem.MessageWithParam(code, param...)
	}
	return "", false
}

func (l *Language) Message(namespace, code string) (string, bool) {
	if namespaceItem, ok := l.Namespaces[namespace]; ok {
		return namespaceItem.Message(code)
	}
	return "", false
}

func (l *Language) RegisterMessage(namespace, code, message string) {
	if l.Namespaces == nil {
		l.Namespaces = map[string]*Namespace{}
	}
	if _, ok := l.Namespaces[namespace]; !ok {
		l.Namespaces[namespace] = &Namespace{
			Namespace: namespace,
			Messages:  map[string]*Message{},
		}
	}
	l.Namespaces[namespace].RegisterMessage(code, message)
}

// Namespace --------------------------------------------------

// Namespace is a middle layout to help deal the complex system.
type Namespace struct {
	Namespace string
	Messages  map[string]*Message
}

// MessageWithParam return the message of the specify code in current namespace. If not found the code, return empty
// string and false. Else return the value of the code by param. The param is used fmt.Sprintf(). The message should a
// template string value.
func (n *Namespace) MessageWithParam(code string, param ...any) (string, bool) {
	if message, ok := n.Message(code); ok {
		value := fmt.Sprintf(message, param...)
		return value, true
	}
	return "", false
}

// Message return the message of the specify code in current namespace. If not found the code, return the empty string
// and false. Else return the message. It will return the value directly.
func (n *Namespace) Message(code string) (string, bool) {
	if value, ok := n.Messages[code]; ok {
		return value.Message, true
	}

	return "", false
}

// RegisterMessage will append one message to current namespace.
func (n *Namespace) RegisterMessage(code, message string) {
	newMessage := Message{
		Code:    code,
		Message: message,
	}

	if n.Messages == nil {
		n.Messages = map[string]*Message{}
	}
	n.Messages[code] = &newMessage
}

// Message --------------------------------------------------

// Message contain the message and code of an i18n message info. The message is i18n the smallest cell. It can receive
// the template value. But the value can't auto translate.
type Message struct {
	Message string `mapstructure:"message"`
	Code    string `mapstructure:"code"`
}
