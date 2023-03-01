// Package routers
package routers

import (
	"gohub/app/http/controllers/api/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/", func(c *gin.Context) {
		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "api!",
		})
	})
	authGroup := api.Group("auth")
	suc := new(auth.SignupController)
	authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
}
