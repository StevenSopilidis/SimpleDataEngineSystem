package main

import (
	m "github.com/StevenSopilidis/SimpleDataEngineSystem/memtable"
	s "github.com/StevenSopilidis/SimpleDataEngineSystem/storage"
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
	data := table.RemoveAll()
	s := s.SegmentHandler{}
	s.AppendSegment(data)
	s.GetSegment("748b9787-1693-4c34-ad4d-54d4719bb0fb")
}
