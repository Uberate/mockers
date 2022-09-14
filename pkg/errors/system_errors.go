package errors

import "mockers/pkg/i18n"

var SystemErrNamespace = NewErrMessage()
var SystemErrNamespaceStr = "system-err-namespace"

var SystemUnHeathError = SystemErrNamespace.
	NewError("system-un-health-error", SystemErrNamespaceStr, i18n.KeyEN, "The system is un-heath").
	NewLanguage(i18n.KeyZHCN, "系统不可用")

var ValueUnExpectValue = SystemErrNamespace.
	NewError("value-un-expect-value", SystemErrNamespaceStr, i18n.KeyEN, "The value is unexpect, is [%v], but want [%v]").
	NewLanguage(i18n.KeyZHCN, "非期望数据，实际 [%v], 期望 [%v]")
