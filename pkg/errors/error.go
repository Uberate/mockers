package errors

import (
	"fmt"
	"mockers/pkg/i18n"
)

type ErrMessage struct {
	languagePool        i18n.LanguagePool
	unknownErrorMessage string
}

func NewErrMessage() ErrMessage {
	return ErrMessage{
		languagePool:        i18n.LanguagePool{},
		unknownErrorMessage: "Unknown error message: [Code: %s], [Namespace: %s], [Language: %s] with value: [%v]",
	}
}

func (e *ErrMessage) Message(code, namespace string, ln i18n.LanguageKey, vars ...interface{}) (string, bool) {
	return e.languagePool.Message(code, namespace, ln, vars...)
}

func (e *ErrMessage) NewError(code, namespace string, ln i18n.LanguageKey, info string) *ErrorGenerator {
	e.languagePool.RegistryMessage(code, namespace, ln, info)
	return &ErrorGenerator{
		code:      code,
		namespace: namespace,
		e:         *e,
	}
}

func (e *ErrMessage) NewHttpError(code, namespace string, ln i18n.LanguageKey, info string, httpCode int) *HttpErrorGenerator {
	newErr := e.NewError(code, namespace, ln, info)
	return &HttpErrorGenerator{
		ErrorGenerator: newErr,
		httpCode:       httpCode,
	}
}

type ErrorGenerator struct {
	code      string
	namespace string

	e ErrMessage
}

func (eg *ErrorGenerator) Param(vars ...interface{}) *ErrorItem {
	ei := ErrorItem{
		eg:     eg,
		ln:     eg.e.languagePool.DefaultLanguage,
		params: vars,
	}
	return &ei
}

func (eg *ErrorGenerator) NewLanguage(ln i18n.LanguageKey, info string) *ErrorGenerator {
	eg.e.NewError(eg.code, eg.namespace, ln, info)
	return eg
}

func (eg *ErrorGenerator) Message(ln i18n.LanguageKey, vars ...interface{}) string {
	res, ok := eg.e.Message(eg.code, eg.namespace, ln, vars...)
	if !ok {
		return fmt.Sprintf(eg.e.unknownErrorMessage, eg.code, eg.namespace, ln, vars)
	}
	return res
}

type HttpErrorGenerator struct {
	*ErrorGenerator

	httpCode int
}

func (h *HttpErrorGenerator) Code() int {
	return h.httpCode
}

func (h *HttpErrorGenerator) Param(vars ...interface{}) HttpErrorItem {
	ei := HttpErrorItem{
		ErrorItem: ErrorItem{
			eg:     h.ErrorGenerator,
			ln:     h.e.languagePool.DefaultLanguage,
			params: vars,
		},
		Code: h.Code(),
	}
	return ei
}

type ErrorItem struct {
	eg *ErrorGenerator

	ln     i18n.LanguageKey
	params []interface{}
}

type HttpErrorItem struct {
	ErrorItem
	Code int
}

func (ei *ErrorItem) Error() string {
	return ei.eg.Message(ei.ln, ei.params...)
}

func (ei *ErrorItem) ErrorWithLanguage(ln i18n.LanguageKey) string {
	return ei.eg.Message(ln, ei.params...)
}

func IsHttpErrorItem(err error) (*HttpErrorItem, bool) {
	if value, ok := err.(*HttpErrorItem); ok {
		return value, true
	} else {
		return nil, false
	}
}
func IsErrorItem(err error) (*ErrorItem, bool) {
	if value, ok := err.(*ErrorItem); ok {
		return value, true
	} else {
		return nil, false
	}
}
