package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	rg := router.Group("")
	{
		rg.GET("/hello", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello")})
		})
	}
}
