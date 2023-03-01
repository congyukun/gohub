package bootstrap

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetRoute(router *gin.Engine) {
	registerGlobalMiddleWare(router)
	setup404Handler(router)
}

// 注册全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(gin.Logger(), gin.Recovery())
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		stringAccept := ctx.Request.Header.Get("Accept")
		if strings.Contains(stringAccept, "text/html") {
			ctx.String(http.StatusNotFound, "404 Not Found")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

}
