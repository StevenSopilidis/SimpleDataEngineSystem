package main

import (
	m "github.com/StevenSopilidis/SimpleDataEngineSystem/memtable"
)

func main() {
	table := m.CreateBinaryMemtable()
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
		table.Insert(value, []byte("asdwetiowet"))
	}
}
