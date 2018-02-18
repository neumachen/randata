package testdata

import (
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
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
	ok := assert.True(t, len(lat) == 2)
	if !ok {
		spew.Dump(address.Latitude)
	}
	assert.True(t, len(lat[1]) < 7)
	long := strings.Split(address.Longitude, ".")
	ok = assert.True(t, len(long) == 2)
	if !ok {
		spew.Dump(address.Longitude)
	}
	assert.True(t, len(long[1]) < 7)
}
