package randata

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandomUSAddress_Success(t *testing.T) {
	t.Parallel()
	address, err := RandomUSAddress()
	require.NoError(t, err)
	require.NotNil(t, address)
	assert.NotEmpty(t, address.StreetNumber)
	assert.NotEmpty(t, address.Route)
	assert.NotEmpty(t, address.Locality)
	assert.NotEmpty(t, address.AdministrativeAreaLevel1)
	assert.NotEmpty(t, address.PostalCode)
	assert.NotEmpty(t, address.Country)
	assert.NotEmpty(t, address.LatitudeStr)
	assert.NotEmpty(t, address.LongitudeStr)

	lat := strings.Split(address.LatitudeStr, ".")
	long := strings.Split(address.LongitudeStr, ".")

	assert.True(t, len(lat) == 2)
	assert.True(t, len(lat[1]) < 7)
	assert.True(t, len(long) == 2)
	assert.True(t, len(lat[1]) == 6)
	assert.True(t, len(long[1]) == 6)
	assert.True(t, len(long[1]) < 7)
}

func TestRandomUSStateAddress_Success(t *testing.T) {
	t.Parallel()
	state := "IL"
	address, err := RandomUSStateAddress(context.Background(), state, 10)
	require.NoError(t, err)
	require.NotNil(t, address)
	assert.Equal(t, state, address.AdministrativeAreaLevel1)
}
