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
		m.Insert(value, []byte("asdwetiowet"))
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
		"TestIsEmpty":    testIsEmpty,
		"TestRemoveAll":  testRemoveAll,
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

func testIsEmpty(t *testing.T, m *BinaryMemtable) {
	require.True(t, m.IsEmpty())
	populateMemTable(t, m)
	require.False(t, m.IsEmpty())
}

func testRemoveAll(t *testing.T, m *BinaryMemtable) {
	data := populateMemTable(t, m)
	memtable_data := m.RemoveAll()
	keys := make([]string, 0)
	for k := range memtable_data {
		keys = append(keys, k)
	}
	require.Equal(t, len(data), len(keys))
	require.True(t, m.IsEmpty())
	require.Zero(t, m.currentSize)
	// check that the data is in ascending order
	for i := 1; i < len(data); i++ {
		require.True(t, keys[i-1] <= keys[i])
	}
}
