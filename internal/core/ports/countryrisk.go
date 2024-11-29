package ports

import (
	"github.com/GalloaFranco/burzar-core/internal/core/domain"
)

type CountryRisk interface {
	Get() (*domain.CountryRisk, error)
}

type CountryRiskRepository interface {
	Obtain() (*domain.CountryRisk, error)
}
