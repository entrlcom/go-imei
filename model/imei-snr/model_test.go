package imei_snr_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSNR(t *testing.T) {
	t.Parallel()

	v, err := NewSNR("176148")
	require.NoError(t, err)
	require.NotZero(t, v)
}

func TestNewSNR_Err(t *testing.T) {
	t.Parallel()

	v, err := NewSNR("")
	require.Error(t, err)
	require.Zero(t, v)
}
