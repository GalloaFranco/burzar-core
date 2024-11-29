package main

import (
	"log"
	"net/http"

	"github.com/GalloaFranco/burzar-core/internal/adapters/drivens/rest"
	driver_rest "github.com/GalloaFranco/burzar-core/internal/adapters/drivers/rest"
	"github.com/GalloaFranco/burzar-core/internal/core/services"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func main() {
	// ENV & CONFIGS
	env := GetEnvironment()
	if err := LoadConfigs(env); err != nil {
		log.Fatalf("Error loading configuration: %+v", err)
	}
	port := viper.GetString("server.port")
	countryRiskBaseURL := viper.GetString("country_risk.baseURL")

	// SERVICES & ADAPTERS
	httpClient := &http.Client{}
	countryRiskRepository := driven_rest.NewCountryRiskRepository(httpClient, countryRiskBaseURL)
	countryRiskService := services.NewCountryRisk(countryRiskRepository)
	countryRiskDriver := driver_rest.NewCountryRisk(countryRiskService)

	router := gin.Default()
	driver_rest.RegisterRoutes(router, countryRiskDriver)

	log.Fatal(router.Run(port))
}
