package services

import (
	"github.com/GalloaFranco/burzar-core/internal/core/domain"
	"github.com/GalloaFranco/burzar-core/internal/core/ports"
)

type CountryRisk struct {
	crr ports.CountryRiskRepository
}

func NewCountryRisk(crr ports.CountryRiskRepository) *CountryRisk {
	return &CountryRisk{crr: crr}
}

func (c *CountryRisk) Get() (*domain.CountryRisk, error) {
	return c.crr.Obtain()
}
