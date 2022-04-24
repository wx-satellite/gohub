package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRoutes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"name": "wx"})
		})
	}
}
