// Package i18n provides some function to quick deal the i18n problem. And for some quick deal with it.
//
// The i18n build the message info like it:
// - ln -> namespace -> code -> message.
//
// The ln is the language name. The namespace is the namespace of the message info. It like a middle layout in the
// i18n system. Like error, user, application and so on. In different namespace, the code can be same.
package i18n

import "fmt"

// MessageObject I18n quick struct.
type MessageObject struct {
	Language  LanguageKey `json:"language" yaml:"language" mapstructure:"language"`
	Namespace string      `json:"namespace" yaml:"namespace" mapstructure:"language"`
	Code      string      `json:"code" yaml:"code" mapstructure:"language"`
	Message   string      `json:"message" yaml:"message" mapstructure:"language"`
}

// I18n --------------------------------------------------

// I18n is the pool of i18n infos. The i18n message info build should complete before the application run time, it
// should build at application bootstrap age. And after build, the info should not be changed, else maybe has error.
// And you can't remove the i18n info by handler. But you can set emtpy info to remove a message. If specify namespace
// already has same code message, the old message will be replaced.
type I18n struct {
	Languages map[LanguageKey]*Language

	//--------------------------------------------------
	// configs, control the behavior of I18n.

	DefaultLanguage LanguageKey

	//--------------------------------------------------
	// inner settings, like index status.

}

// Len return the message count of current I18n.
func (i *I18n) Len() int {
	res := 0
	for _, languageItem := range i.Languages {
		res += languageItem.Len()
	}
	return res
}

// ToMessageObjects return a MessageObject array.
func (i *I18n) ToMessageObjects() []MessageObject {
	res := make([]MessageObject, 0, i.Len())
	i.WalkMessage(func(key LanguageKey, namespace, code, message string) {
		res = append(res, MessageObject{
			Language:  key,
			Namespace: namespace,
			Code:      code,
			Message:   message,
		})
	})
	return res
}

// WalkMessage will invoke the func for all message.
func (i *I18n) WalkMessage(f func(key LanguageKey, namespace, code, message string)) {
	for ln, namespaceItems := range i.Languages {
		for namespace, messageItems := range namespaceItems.Namespaces {
			for code, message := range messageItems.Messages {
				f(ln, namespace, code, message.Message)
			}
		}
	}
}

// MessageWithParam return the template message info with params. The base is invoked the fmt.Sprintf(). If not found,
// return string in format 'namespace-code' and false.
func (i *I18n) MessageWithParam(ln LanguageKey, namespace, code string, params ...any) (string, bool) {
	if languageItem, ok := i.Languages[ln]; ok {
		return languageItem.MessageWithParam(namespace, code, params...)
	}

	// if the default language has value, search in default language scope.
	if ln != i.DefaultLanguage && len(i.DefaultLanguage) != 0 && i.DefaultLanguage != EmptyLanguage {
		return i.MessageWithParam(i.DefaultLanguage, namespace, code, params...)
	}

	// if is default, and not found message, return "namespace-code"
	return fmt.Sprintf("%s-%s", namespace, code), false
}

// Message return the message info. Different from the MessageWithParam, it is return value directly. If not found
// specify value, return string in format 'namespace-code' and false.
func (i *I18n) Message(ln LanguageKey, namespace, code string) (string, bool) {
	if languageItem, ok := i.Languages[ln]; ok {
		return languageItem.Message(namespace, code)
	}

	// if the default language has value, search in default language scope.
	if ln != i.DefaultLanguage && len(i.DefaultLanguage) != 0 && i.DefaultLanguage != EmptyLanguage {
		return i.Message(i.DefaultLanguage, namespace, code)
	}

	// if is default, and not found message, return "namespace-code"
	return fmt.Sprintf("%s-%s", namespace, code), false
}

// RegisterMessage will register a language to the namespace. If the code already in a namespace, the old message info
// will be replaced. If input message info is empty, the message will be removed.
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

// RegisterMessageObject like RegisterMessage, but this function receive a MessageObject.
func (i *I18n) RegisterMessageObject(object MessageObject) {
	i.RegisterMessage(object.Language, object.Namespace, object.Code, object.Message)
}

// Language --------------------------------------------------

// Language save one language message info.
type Language struct {
	Language   LanguageKey
	Namespaces map[string]*Namespace
}

func (l *Language) Len() int {
	res := 0
	for _, namespaceItem := range l.Namespaces {
		res += namespaceItem.Len()
	}

	return res
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

func (n *Namespace) Len() int {
	return len(n.Messages)
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
	if len(message) == 0 {
		delete(n.Messages, code)
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
