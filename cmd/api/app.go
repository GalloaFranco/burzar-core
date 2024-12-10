package app

import (
	"log"
	"net/http"

	drivenrest "github.com/GalloaFranco/burzar-core/internal/adapters/drivens/rest"
	driverrest "github.com/GalloaFranco/burzar-core/internal/adapters/drivers/rest"
	"github.com/GalloaFranco/burzar-core/internal/core/services"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

var glambda *ginadapter.GinLambda

func Run() {
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

	glambda = ginadapter.New(router)
	log.Printf("ENVIRONMENT: %v", env)
	if env == "local" {
		log.Fatal(router.Run(port))
	} else {
		lambda.Start(glambda.Proxy)
	}
}
