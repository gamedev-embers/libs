package testhelper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testObject struct {
	I    int
	J    string
	K    map[string]int
	L    []int
	Zero uint32
}

func TestSnapshot(t *testing.T) {
	makeData := func() *testObject {
		return &testObject{
			I: 1,
			J: "2",
			K: map[string]int{
				"3": 3,
			},
			L:    []int{4, 5, 6},
			Zero: 0,
		}
	}
	data := makeData()

	snap, err := Snapshot(data)
	require.NoError(t, err)
	require.NotNil(t, snap)

	data.I = 100
	data.Zero = 9527

	err = snap.Rollback()
	require.NoError(t, err)

	cloned := makeData()
	require.Equal(t, cloned, data)
}
