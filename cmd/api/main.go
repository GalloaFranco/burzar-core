package main

import (
	"log"

	burzarhttp "github.com/GalloaFranco/burzar-core/internal/adapters/drivers/http"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func main() {
	env := GetEnvironment()
	if err := LoadConfigs(env); err != nil {
		log.Fatalf("Error loading configuration: %+v", err)
	}
	port := viper.GetString("server.port")

	router := gin.Default()
	burzarhttp.RegisterRoutes(router)

	log.Fatal(router.Run(port))
}
