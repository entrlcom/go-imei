package imei_tac_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTAC(t *testing.T) {
	t.Parallel()

	v, err := ParseTAC("35209900")
	require.NoError(t, err)
	require.NotZero(t, v)
}

func TestParseTAC_Err(t *testing.T) {
	t.Parallel()

	v, err := ParseTAC("")
	require.Error(t, err)
	require.Zero(t, v)
}

func TestTAC_GetID(t *testing.T) {
	t.Parallel()

	v, err := ParseTAC("35209900")
	require.NoError(t, err)
	require.Equal(t, "209900", string(v.GetID()))
}

func TestTAC_GetRBI(t *testing.T) {
	t.Parallel()

	v, err := ParseTAC("35209900")
	require.NoError(t, err)
	require.Equal(t, "35", string(v.GetRBI()))
}
