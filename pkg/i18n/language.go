package i18n

import (
	"fmt"
)

type AbsI18N interface {
	Message(code, namespace string, ln LanguageKey, vars ...interface{}) (string, bool)
}

type NamespaceMessage map[string]string

// Message will return the message with specify vars, and if specify code is not exist, return "", false.
func (n NamespaceMessage) Message(code string, vars ...interface{}) (string, bool) {
	if _, ok := n[code]; !ok {
		return "", false
	}
	return fmt.Sprintf(n[code], vars...), true
}

// AppendMessage will insert a message to specify namespace. If specify key is already in, keep old value, ignore new
// value
// true.
func (n NamespaceMessage) AppendMessage(code, value string) NamespaceMessage {
	if _, ok := n[code]; !ok {
		n[code] = value
		return n
	}

	return n
}

// CoverMessage will insert a message to specify namespace. If specify key is already in, will cover old value.
func (n NamespaceMessage) CoverMessage(code, value string) NamespaceMessage {
	n[code] = value
	return n
}

func (n NamespaceMessage) DeleteMessage(code string) NamespaceMessage {
	delete(n, code)
	return n
}

type LanguageMessage map[string]NamespaceMessage

func (l LanguageMessage) Message(code, namespace string, vars ...interface{}) (string, bool) {
	if namespace, ok := l[namespace]; ok {
		return namespace.Message(code, vars...)
	}

	return "", false
}

func (l LanguageMessage) CreateNamespace(namespace string) LanguageMessage {
	_ = l.GetNamespace(namespace)
	return l
}

// GetNamespace will return a NamespaceMessage, and if not exist, create one.
func (l LanguageMessage) GetNamespace(namespace string) NamespaceMessage {
	if _, ok := l[namespace]; !ok {
		l[namespace] = NamespaceMessage{}
	}

	return l[namespace]
}

func (l LanguageMessage) DeleteNamespace(namespace string) LanguageMessage {
	if _, ok := l[namespace]; !ok {
		return l
	}
	delete(l, namespace)
	return l
}

type LanguagePool struct {
	Languages       map[LanguageKey]LanguageMessage
	DefaultLanguage LanguageKey
}

func (l LanguagePool) Message(code, namespace string, ln LanguageKey, vars ...interface{}) (string, bool) {
	if res, ok := l.GetLanguage(ln).GetNamespace(namespace).Message(code, vars...); ok {
		return res, true
	}
	return l.GetLanguage(l.DefaultLanguage).GetNamespace(namespace).Message(code, vars...)
}

func (l *LanguagePool) GetLanguage(ln LanguageKey) LanguageMessage {
	if l.Languages == nil {
		l.Languages = map[LanguageKey]LanguageMessage{}
	}
	if _, ok := l.Languages[ln]; !ok {
		l.Languages[ln] = LanguageMessage{}
	}

	return l.Languages[ln]
}

func (l *LanguagePool) RegistryMessage(code, namespace string, ln LanguageKey, value string) *LanguagePool {
	l.GetLanguage(ln).GetNamespace(namespace).AppendMessage(code, value)
	return l
}
