package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uberate/mockers/cmd/common/server"
	cfg2 "github.com/uberate/mockers/cmd/common/server/cfg"
	"github.com/uberate/mockers/cmd/i18n/server/cfg"
	"github.com/uberate/mockers/cmd/utils"
	"github.com/uberate/mockers/pkg/i18n"
	"net/http"
)

var i18nCenter i18n.I18n

func main() {

	engine := gin.Default()

	webConfig := cfg.I18nWebConfig{
		WebCfg: cfg2.WebConfig{
			Port: "3000",
		},
	}

	if err := utils.ReadConfig("", "", "", true,
		"", nil, &webConfig); err != nil {
		panic(err)
	}

	// get one message info
	engine.GET("message/:ln/:namespace/:code", func(context *gin.Context) {
		ln := context.Param("ln")
		namespace := context.Param("namespace")
		code := context.Param("code")
		value, ok := i18nCenter.Message(i18n.GetLanguageKey(ln), namespace, code)
		if !ok {
			// if specify message not found, and set return 404 when target specify message not found. return nil with
			// 404 value.
			if webConfig.NotFoundWith404 {
				// TODO : quick return 404 status with nil object.
				context.JSON(http.StatusNotFound, nil)
				return
			}
		}

		context.JSON(http.StatusOK, map[string]string{"value": value})
	})

	if err := server.GinStart(engine, webConfig.WebCfg); err != nil {
		panic(err)
	}
}

func init() {
	// read the config

}
