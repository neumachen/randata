package testdata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomEmailAddress_DefaultLength(t *testing.T) {
	emailAddress := strings.Split(RandomEmailAddress("something.com", 0), "@")
	assert.Equal(t, 10, len(emailAddress[0]))
}
