package server

import (
	"github.com/gin-gonic/gin"
	"mockers/pkg/errors"
	"mockers/pkg/i18n"
	"net/http"
	"strings"
)

type GinConfig struct {
	Host                 string
	Port                 string
	EnableLivenessProbe  bool
	EnableReadinessProbe bool
}

// GinHandler save a handler value. Will be registered to the gin.
type GinHandler struct {
	HandlerPath string
	Method      []string
	Handler     []gin.HandlerFunc
}

func Start(config GinConfig, handlers ...GinHandler) error {
	ginInstance := gin.Default()

	// if enable the gin liveness probe, register all method function: gin-liveness-probe
	if config.EnableLivenessProbe {
		handlers = append(handlers,
			GinHandler{
				Method:      AnyMethod,
				HandlerPath: "gin-liveness-probe",
				Handler: []gin.HandlerFunc{
					func(context *gin.Context) {
						context.Writer.WriteHeader(http.StatusOK)
					},
				},
			},
		)
	}

	isReady := false
	// if enable the gin readiness probe, register all method function: gin-readiness-probe
	if config.EnableReadinessProbe {
		handlers = append(handlers,
			GinHandler{
				Method:      AnyMethod,
				HandlerPath: "gin-readiness-probe",
				Handler: []gin.HandlerFunc{
					func(context *gin.Context) {
						if isReady {
							context.Writer.WriteHeader(http.StatusServiceUnavailable)
						} else {
							context.Writer.WriteHeader(http.StatusOK)
						}
					},
				},
			},
		)
	}

	for _, handler := range handlers {
		for _, method := range handler.Method {
			switch method {
			case http.MethodGet:
				ginInstance.GET(handler.HandlerPath, handler.Handler...)
			case http.MethodHead:
				ginInstance.HEAD(handler.HandlerPath, handler.Handler...)
			case http.MethodPost:
				ginInstance.POST(handler.HandlerPath, handler.Handler...)
			case http.MethodPut:
				ginInstance.PUT(handler.HandlerPath, handler.Handler...)
			case http.MethodPatch:
				ginInstance.PATCH(handler.HandlerPath, handler.Handler...)
			case http.MethodDelete:
				ginInstance.DELETE(handler.HandlerPath, handler.Handler...)
			case http.MethodConnect:
				break
			case http.MethodOptions:
				ginInstance.OPTIONS(handler.HandlerPath, handler.Handler...)
			case http.MethodTrace:
				break
			}
		}
	}

	isReady = true

	return ginInstance.Run(strings.Join([]string{config.Host, config.Port}, ":"))
}

const (
	HeaderLanguage = "language"
)

func GetLanguage(ctx *gin.Context) string {
	return ctx.GetHeader(HeaderLanguage)
}

func ResSuccess(ctx *gin.Context, code *int, obj interface{}) {
	x := http.StatusOK
	if code == nil {
		code = &x
	}
	ctx.JSONP(*code, obj)
}

func AutoRes(ctx *gin.Context, code int, obj interface{}, err error) {
	x := &code
	if code == 0 {
		x = nil
	}
	if err != nil {
		ResError(ctx, x, err)
	} else {
		ResSuccess(ctx, x, obj)
	}
}

func ResError(ctx *gin.Context, code *int, err error) {
	x := http.StatusBadRequest
	if code == nil {
		code = &x
	}
	if e, ok := errors.IsHttpErrorItem(err); ok {
		ctx.JSONP(e.Code, e.ErrorWithLanguage(i18n.LanguageKey(GetLanguage(ctx))))
	} else if e, ok := errors.IsErrorItem(err); ok {
		ctx.JSONP(x, e.ErrorWithLanguage(i18n.LanguageKey(GetLanguage(ctx))))
	} else {
		ctx.JSONP(x, err)
	}
}
