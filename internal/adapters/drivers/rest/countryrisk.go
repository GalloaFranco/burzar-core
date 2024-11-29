package driver_rest

import "github.com/GalloaFranco/burzar-core/internal/core/ports"

type CountryRisk struct {
	countryRisk ports.CountryRisk
}

func NewCountryRisk(crs ports.CountryRisk) *CountryRisk {
	return &CountryRisk{countryRisk: crs}
}

func (cr *CountryRisk) Get() (*CountryRiskView, error) {
	r, err := cr.countryRisk.Get()
	if err != nil {
		return nil, err
	}
	return &CountryRiskView{
		Risk: r.Risk,
	}, err
}
