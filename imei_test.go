package imei

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewIMEI(t *testing.T) {
	t.Parallel()

	t.Run("IMEI", func(t *testing.T) {
		t.Parallel()

		v, err := NewIMEI("35-209900-176148-1")
		require.NoError(t, err)
		require.NotZero(t, v)

		require.Equal(t, "1", string(v.GetCD()))
		require.True(t, v.IsIMEI())
		require.False(t, v.IsIMEISV())
		require.Equal(t, "176148", string(v.GetSNR()))
		require.Equal(t, "35 209900 176148 1", v.String())
		require.Zero(t, v.GetSVN())
		require.Equal(t, "209900", string(v.GetTAC().GetID()))
		require.Equal(t, "35", string(v.GetTAC().GetRBI()))
	})

	t.Run("IMEISV", func(t *testing.T) {
		t.Parallel()

		v, err := NewIMEI("35-209900-176148-23")
		require.NoError(t, err)
		require.NotZero(t, v)

		require.Zero(t, v.GetCD())
		require.False(t, v.IsIMEI())
		require.True(t, v.IsIMEISV())
		require.Equal(t, "176148", string(v.GetSNR()))
		require.Equal(t, "35 209900 176148 23", v.String())
		require.Equal(t, "23", string(v.GetSVN()))
		require.Equal(t, "209900", string(v.GetTAC().GetID()))
		require.Equal(t, "35", string(v.GetTAC().GetRBI()))
	})
}

func TestNewIMEI_Err(t *testing.T) {
	t.Parallel()

	v, err := NewIMEI("")
	require.Error(t, err)
	require.Zero(t, v)
}
