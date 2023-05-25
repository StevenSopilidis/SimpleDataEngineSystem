package memtable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func setUpTest(t *testing.T) *BinaryMemtable {
	t.Helper()
	return CreateBinaryMemtable()
}

func populateMemTable(t *testing.T, m *BinaryMemtable) [][]byte {
	t.Helper()
	values := [][]byte{
		[]byte("asdnjiowehwetuioh"),
		[]byte("ascdshuiweruitgefwrd"),
		[]byte("czxsnjoksidioufewst"),
		[]byte("cnsuiwehd8rt94325r98h"),
		[]byte("scnmviuoehdsnfuihghbewrtg"),
		[]byte("cnsiudoghbfn8394egy5rh4fc"),
		[]byte("r43f3w35345rt"),
	}
	for _, value := range values {
		m.Insert(value, "Data......")
	}
	return values
}

func TestBinaryMemtable(t *testing.T) {
	for scenario, fn := range map[string]func(
		t *testing.T,
		m *BinaryMemtable,
	){
		"TestValidGet":   testValidGet,
		"TestInvalidGet": testInvalidGet,
	} {
		t.Run(scenario, func(t *testing.T) {
			table := setUpTest(t)
			fn(t, table)
		})
	}
}

func testValidGet(t *testing.T, m *BinaryMemtable) {
	values := populateMemTable(t, m)
	for _, value := range values {
		require.NotNil(t, m.Get(value))
	}
}

func testInvalidGet(t *testing.T, m *BinaryMemtable) {
	require.Nil(t, m.Get([]byte("InvalidKey")))
}
