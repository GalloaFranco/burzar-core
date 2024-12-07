package main

import (
	"log"
	"net/http"

	drivenrest "github.com/GalloaFranco/burzar-core/internal/adapters/drivens/rest"
	driverrest "github.com/GalloaFranco/burzar-core/internal/adapters/drivers/rest"
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
	countryRiskRepository := drivenrest.NewCountryRiskRepository(httpClient, countryRiskBaseURL)
	countryRiskService := services.NewCountryRisk(countryRiskRepository)
	countryRiskDriver := driverrest.NewCountryRisk(countryRiskService)
	// TODO: wrap all services that router will use in one big service (DI)

	router := gin.Default()
	driverrest.RegisterRoutes(router, countryRiskDriver)

	log.Fatal(router.Run(port))
}
