package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uberate/mockers/cmd/common/server"
	cfg2 "github.com/uberate/mockers/cmd/common/server/cfg"
	"github.com/uberate/mockers/cmd/i18n/server/cfg"
	"github.com/uberate/mockers/cmd/i18n/server/handler"
	"github.com/uberate/mockers/cmd/utils"
	"github.com/uberate/mockers/pkg/i18n"
)

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

	var i18nInstanceOptions []i18n.Option = make([]i18n.Option, 0, 2)

	if webConfig.EnableChange {
		i18nInstanceOptions = append(i18nInstanceOptions, i18n.EnableI18nChange())
	}
	i18nInstanceOptions = append(i18nInstanceOptions, i18n.DefaultLanguage(webConfig.DefaultLanguage))

	i18nInstance := handler.I18nHandlers{
		I18nCenter: *i18n.NewI18nInstance(i18nInstanceOptions...),
		WebConfig:  webConfig,
	}

	engine.GET("config", func(context *gin.Context) {
		utils.Success(context, webConfig)
	})

	engine.GET("languages", i18nInstance.Languages)

	// get one message info
	engine.GET("language/:ln/namespace/:namespace/code/:code", i18nInstance.Message)
	engine.POST("language/:ln/namespace/:namespace/code/:code", i18nInstance.Create)

	if err := server.GinStart(engine, webConfig.WebCfg); err != nil {
		panic(err)
	}
}
