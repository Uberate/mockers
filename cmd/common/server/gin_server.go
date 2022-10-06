package server

import (
	"github.com/gin-gonic/gin"
	"github.com/uberate/mockers/cmd/common/server/cfg"
)

func GinStart(engine *gin.Engine, webConfig cfg.WebConfig) error {
	return engine.Run(webConfig.Port)
}
