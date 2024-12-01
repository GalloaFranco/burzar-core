package services

import (
	"errors"
	"fmt"
	"github.com/GalloaFranco/burzar-core/internal/core/domain"
	"github.com/GalloaFranco/burzar-core/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountryRiskService_NewCountryRisk(t *testing.T) {
	crs := mocks.NewCountryRiskRepository(t)
	sut := NewCountryRisk(crs)

	assert.NotNil(t, sut)
	assert.Equal(t, "*services.CountryRisk", fmt.Sprintf("%T", sut))
}

func TestCountryRiskService_Get_ShouldReturnCountryRisk(t *testing.T) {
	crs := mocks.NewCountryRiskRepository(t)
	crs.On("Obtain").Return(&domain.CountryRisk{Risk: "400"}, nil)

	sut := NewCountryRisk(crs)
	cv, err := sut.Get()

	assert.Nil(t, err)
	assert.NotNil(t, cv)
	assert.Equal(t, "400", cv.Risk)
	crs.AssertNumberOfCalls(t, "Obtain", 1)
	crs.AssertCalled(t, "Obtain")
}

func TestCountryRiskDriver_Obtain_ShouldReturnError(t *testing.T) {
	crs := mocks.NewCountryRiskRepository(t)
	crs.On("Obtain").Return(nil, errors.New("test_error"))

	sut := NewCountryRisk(crs)
	cv, err := sut.Get()

	assert.Nil(t, cv)
	assert.NotNil(t, err)
	assert.Equal(t, "test_error", err.Error())
	crs.AssertNumberOfCalls(t, "Obtain", 1)
	crs.AssertCalled(t, "Obtain")
}
