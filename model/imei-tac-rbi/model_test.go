package imei_tac_rbi_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRBI(t *testing.T) {
	t.Parallel()

	v, err := NewRBI("35")
	require.NoError(t, err)
	require.NotZero(t, v)
}

func TestNewRBI_Err(t *testing.T) {
	t.Parallel()

	v, err := NewRBI("")
	require.Error(t, err)
	require.Zero(t, v)
}

func TestRBI_IsTestIMEI(t *testing.T) {
	t.Parallel()

	for k := range test {
		t.Run(string(k), func(t *testing.T) {
			t.Parallel()

			v, err := NewRBI(string(k))
			require.NoError(t, err)
			require.True(t, v.IsTestIMEI())
		})
	}
}

func TestRBI_Validate(t *testing.T) {
	t.Parallel()

	for k := range whitelist {
		t.Run(string(k), func(t *testing.T) {
			t.Parallel()

			v, err := NewRBI(string(k))
			require.NoError(t, err)
			require.NotZero(t, v)
		})
	}
}

func TestRBI_Validate_Err(t *testing.T) {
	t.Parallel()

	v, err := NewRBI("11")
	require.Error(t, err)
	require.Zero(t, v)
}
