package storage

import "github.com/StevenSopilidis/SimpleDataEngineSystem/memtable"

type Storage struct {
	segmentHandler *SegmentHandler
	memtable       *memtable.Memtable
}

type StorageEntry struct {
	Data []byte
}

func (s *Storage) SetEntry(key []byte, data []byte) error {
	return nil
}

func (s *Storage) GetEntry(key []byte) StorageEntry {
	return StorageEntry{}
}
