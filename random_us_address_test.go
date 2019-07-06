package randata

import (
	"strings"
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

	lat := strings.Split(address.Latitude, ".")
	long := strings.Split(address.Longitude, ".")

	assert.True(t, len(lat) == 2)
	assert.True(t, len(lat[1]) < 7)
	assert.True(t, len(long) == 2)
	assert.True(t, len(lat[1]) == 6)
	assert.True(t, len(long[1]) == 6)
	assert.True(t, len(long[1]) < 7)
}

func TestRandomUSStateAddress_Success(t *testing.T) {
	state := "IL"
	address := RandomUSStateAddress(state, 10)
	assert.NotNil(t, address)
	assert.Equal(t, state, address.AdministrativeAreaLevel1)
}
