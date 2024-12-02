package driven_rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockRoundTripper implements RoundTripper
type MockRoundTripper struct {
	MockResponse *http.Response
	MockError    error
}

func (m *MockRoundTripper) RoundTrip(_ *http.Request) (*http.Response, error) {
	return m.MockResponse, m.MockError
}

func TestCountryRiskRepository_NewCountryRisk(t *testing.T) {
	sut := NewCountryRiskRepository(&http.Client{}, "")

	assert.NotNil(t, sut)
	assert.Equal(t, "*driven_rest.CountryRiskRepository", fmt.Sprintf("%T", sut))
}

func TestCountryRiskRepository_Obtain_ShouldReturnCountryRisk(t *testing.T) {
	response := CountryRiskResponse{
		Risk:      "400",
		Value:     "",
		Variation: "",
	}
	bodyBytes, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(bodyBytes)),
	}

	// Using MockRoundTripper to intercept the request
	mockTransport := &MockRoundTripper{MockResponse: mockResponse}
	client := &http.Client{Transport: mockTransport}

	sut := NewCountryRiskRepository(client, "")
	cv, err := sut.Obtain()

	assert.Nil(t, err)
	assert.NotNil(t, cv)
	assert.Equal(t, "400", cv.Risk)
}

func TestCountryRiskRepository_Obtain_ShouldReturnError(t *testing.T) {
	mockResponse := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       nil,
	}

	// Using MockRoundTripper to intercept the request
	mockTransport := &MockRoundTripper{MockResponse: mockResponse}
	client := &http.Client{Transport: mockTransport}

	sut := NewCountryRiskRepository(client, "")
	cv, err := sut.Obtain()

	assert.Nil(t, cv)
	assert.NotNil(t, err)
	assert.Equal(t, "error: status 500", err.Error())
}
