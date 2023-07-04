package randata

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSSN(t *testing.T) {
	t.Parallel()

	ssn := SSN(false, 100)
	require.NotEmpty(t, ssn)
}
