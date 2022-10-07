package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uberate/mockers/cmd/i18n/server/cfg"
	"github.com/uberate/mockers/cmd/i18n/server/model"
	"github.com/uberate/mockers/cmd/utils"
	"github.com/uberate/mockers/pkg/i18n"
	"net/http"
)

type I18nHandlers struct {
	I18nCenter i18n.I18n
	WebConfig  cfg.I18nWebConfig
}

// Message is a handler in gin. It will return a message by language, namespace and code.
//
// return model.I18nMessage
func (i *I18nHandlers) Message(context *gin.Context) {
	ln := context.Param("ln")
	namespace := context.Param("namespace")
	code := context.Param("code")
	value, ok := i.I18nCenter.Message(i18n.GetLanguageKey(ln), namespace, code)
	if !ok {
		// if specify message not found, and set return 404 when target specify message not found. return nil with
		// 404 value.
		if i.WebConfig.NotFoundWith404 {
			// TODO : quick return 404 status with nil object.
			utils.JSON(context, http.StatusNotFound, nil)
			return
		}
	}

	message := model.I18nMessage{
		Message: value,
	}

	utils.Success(context, message)
}

// Languages is a handler in gin. It will return all language key value and describe.
//
// return map[string]string as a json string.
func (i *I18nHandlers) Languages(context *gin.Context) {
	context.JSON(http.StatusOK, i18n.GetLanguageDescribe())
}

// Create is a handler in gin. It will register a message to i18n center.
func (i *I18nHandlers) Create(context *gin.Context) {
	if !i.WebConfig.EnableChange {
		utils.Error(context, fmt.Errorf("The application was disbale change i18n value "))
		return
	}
	ln := context.Param("ln")
	namespace := context.Param("namespace")
	code := context.Param("code")

	messageObj := model.I18nMessage{}
	if err := context.BindJSON(&messageObj); err != nil {
		utils.Error(context, err)
		return
	}

	i.I18nCenter.RegisterMessage(i18n.GetLanguageKey(ln), namespace, code, messageObj.Message)
	utils.Success(context, "success")
}
