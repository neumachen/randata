package randata

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomSSN(t *testing.T) {
	ssn := RandomSSN(false, 100)
	require.NotEmpty(t, ssn)
}
