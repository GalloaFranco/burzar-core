package driven_rest

import "github.com/GalloaFranco/burzar-core/internal/core/domain"

/*
	{
	  "ultimo": "755",
	  "fecha": "28-11-2024 19:55:02",
	  "valor": "755",
	  "varpesos": "-3",
	  "variacion-nombre": "Var. puntos"
	}
*/
type CountryRiskResponse struct {
	Risk      string `json:"ultimo"`
	Value     string `json:"valor"`
	Variation string `json:"varpesos"`
}

func (c *CountryRiskResponse) ToCountryRisk() *domain.CountryRisk {
	return &domain.CountryRisk{Risk: c.Risk}
}
