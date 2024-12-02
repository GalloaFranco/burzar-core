package driver_rest

import (
	"errors"
	"fmt"
	"testing"

	"github.com/GalloaFranco/burzar-core/internal/core/domain"

	"github.com/GalloaFranco/burzar-core/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCountryRiskDriver_NewCountryRisk(t *testing.T) {
	crs := mocks.NewCountryRisk(t)
	sut := NewCountryRisk(crs)

	assert.NotNil(t, sut)
	assert.Equal(t, "*driver_rest.CountryRisk", fmt.Sprintf("%T", sut))
}

func TestCountryRiskDriver_Obtain_ShouldReturnCountryRiskView(t *testing.T) {
	crs := mocks.NewCountryRisk(t)
	crs.On("Get").Return(&domain.CountryRisk{Risk: "400"}, nil)

	sut := NewCountryRisk(crs)
	cv, err := sut.Get()

	assert.Nil(t, err)
	assert.NotNil(t, cv)
	assert.Equal(t, "400", cv.Risk)
	crs.AssertNumberOfCalls(t, "Get", 1)
	crs.AssertCalled(t, "Get")
}

func TestCountryRiskDriver_Obtain_ShouldReturnError(t *testing.T) {
	crs := mocks.NewCountryRisk(t)
	crs.On("Get").Return(nil, errors.New("test_error"))

	sut := NewCountryRisk(crs)
	cv, err := sut.Get()

	assert.Nil(t, cv)
	assert.NotNil(t, err)
	assert.Equal(t, "test_error", err.Error())
	crs.AssertNumberOfCalls(t, "Get", 1)
	crs.AssertCalled(t, "Get")
}
