package imei_cd_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCD(t *testing.T) {
	t.Parallel()

	v, err := NewCD("1")
	require.NoError(t, err)
	require.NotZero(t, v)
}

func TestNewCD_Err(t *testing.T) {
	t.Parallel()

	v, err := NewCD("")
	require.Error(t, err)
	require.Zero(t, v)
}

func TestComputeCD(t *testing.T) {
	t.Parallel()

	v, err := ComputeCD("35209900176148")
	require.NoError(t, err)
	require.NotZero(t, v)
}

func TestComputeCD_Err(t *testing.T) {
	t.Parallel()

	v, err := ComputeCD("-")
	require.Error(t, err)
	require.Zero(t, v)
}
