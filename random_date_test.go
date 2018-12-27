package randata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRandomDate(t *testing.T) {
	minDate, err := time.Parse(time.RFC822, "01 Jan 90 00:00 UTC")
	assert.NoError(t, err)
	maxDate, err := time.Parse(time.RFC822, "31 Dec 91 11:59 UTC")
	assert.NoError(t, err)

	date := RandomDate(1990, 1991)
	assert.True(t, date.After(minDate) && date.Before(maxDate))
}
