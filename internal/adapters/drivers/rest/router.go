package driver_rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, crs *CountryRisk) {
	rg := router.Group("")
	{
		rg.GET("/hello", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello")})
		})
		rg.GET("/country_risk", func(context *gin.Context) {
			risk, err := crs.Get()
			if err != nil {
				context.JSON(
					http.StatusInternalServerError,
					gin.H{"message": fmt.Sprintf("Error: %+v", err)},
				)
			}
			context.JSON(http.StatusOK, risk)
		})
	}
}
