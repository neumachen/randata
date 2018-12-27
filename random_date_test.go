package randata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomDate(t *testing.T) {
	minDate := "1989-12-31 00:00:01 +0000 UTC"
	maxDate := "1992-01-01 00:00:01 +0000 UTC"

	date := RandomDate(1990, 1991)
	assert.True(t, date.After(minDate) && date.Before(maxDate))
}
