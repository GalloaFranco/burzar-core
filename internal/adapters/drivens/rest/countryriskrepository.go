package driven_rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GalloaFranco/burzar-core/internal/core/domain"
)

type CountryRiskRepository struct {
	client  *http.Client
	baseURL string
}

func NewCountryRiskRepository(client *http.Client, baseURL string) *CountryRiskRepository {
	return &CountryRiskRepository{client: client, baseURL: baseURL}
}

// Obtain TODO: Error handling
func (c *CountryRiskRepository) Obtain() (*domain.CountryRisk, error) {
	response, err := c.client.Get(c.baseURL)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status %d", response.StatusCode)
	}

	var crr CountryRiskResponse
	if err := json.NewDecoder(response.Body).Decode(&crr); err != nil {
		return nil, err
	}
	return crr.ToCountryRisk(), nil
}
