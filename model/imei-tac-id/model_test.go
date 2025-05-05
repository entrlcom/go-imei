package imei_tac_id_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewID(t *testing.T) {
	t.Parallel()

	v, err := NewID("209900")
	require.NoError(t, err)
	require.NotZero(t, v)
}

func TestNewID_Err(t *testing.T) {
	t.Parallel()

	v, err := NewID("")
	require.Error(t, err)
	require.Zero(t, v)
}
