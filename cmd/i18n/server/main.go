package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uberate/mockers/cmd/common/server"
	cfg2 "github.com/uberate/mockers/cmd/common/server/cfg"
	"github.com/uberate/mockers/cmd/i18n/server/cfg"
	"github.com/uberate/mockers/cmd/utils"
	"github.com/uberate/mockers/pkg/i18n"
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

	})

	if err := server.GinStart(engine, webConfig.WebCfg); err != nil {
		panic(err)
	}
}

func init() {
	// read the config

}
