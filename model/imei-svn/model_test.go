package imei_svn_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSVN(t *testing.T) {
	t.Parallel()

	v, err := NewSVN("23")
	require.NoError(t, err)
	require.NotZero(t, v)
}

func TestNewSVN_Err(t *testing.T) {
	t.Parallel()

	v, err := NewSVN("")
	require.Error(t, err)
	require.Zero(t, v)
}

func TestSVN_IsReserved(t *testing.T) {
	t.Parallel()

	t.Run("false", func(t *testing.T) {
		t.Parallel()

		v, err := NewSVN("23")
		require.NoError(t, err)
		require.False(t, v.IsReserved())
	})

	t.Run("true", func(t *testing.T) {
		t.Parallel()

		v, err := NewSVN("99")
		require.NoError(t, err)
		require.True(t, v.IsReserved())
	})
}
