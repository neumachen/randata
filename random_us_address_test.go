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

	ok := assert.True(t, len(lat) == 2)
	if !ok {
		assert.FailNow("latitude:", address.Latitude)
	}
	assert.True(t, len(lat[1]) < 7)
	ok = assert.True(t, len(long) == 2)
	if !ok {
		assert.FailNow("longitude:", address.Longitude)
	}

	ok = assert.True(t, len(lat[1]) == 6)
	if !ok {
		assert.FailNow("latitude:", lat[1])
	}
	ok = assert.True(t, len(long[1]) == 6)
	if !ok {
		assert.FailNow("Longitude:", long[1])
	}
	assert.True(t, len(long[1]) < 7)
}
