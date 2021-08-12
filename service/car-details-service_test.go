package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	carDetailsService = NewCarDetailsService()
)

func TestGetCarDetails(t *testing.T) {
	carDetails := carDetailsService.GetDetails()

	assert.NotNil(t, carDetails)
	assert.Equal(t, 1, carDetails.ID)
	assert.Equal(t, "Mitsubishi", carDetails.Brand)
	assert.Equal(t, 2002, carDetails.Year)
	assert.Equal(t, "Alyosha", carDetails.FirstName)
	assert.Equal(t, "Caldero", carDetails.LastName)
	assert.Equal(t, "acaldero0@behance.net", carDetails.Email)
}
