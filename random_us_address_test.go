package testdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomUSAddress_Success(t *testing.T) {
	address := RandomUSAddress()
	assert.NotNil(t, address)
	assert.NotEmpty(t, address.StreetNumber)
	assert.NotEmpty(t, address.Route)
	assert.NotEmpty(t, address.Locality)
	assert.NotEmpty(t, address.AdministrativeAreaLevel1)
	assert.NotEmpty(t, address.PostalCode)
	assert.NotEmpty(t, address.Country)
	assert.NotEmpty(t, address.Latitude)
	assert.NotEmpty(t, address.Longitude)
}
