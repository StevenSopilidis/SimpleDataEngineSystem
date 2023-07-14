package storage

import (
	"github.com/StevenSopilidis/SimpleDataEngineSystem/memtable"
)

type ItemNotFoundError struct {
	key string
}

func (e *ItemNotFoundError) Error() string {
	return "Item not found: " + e.key
}

type Storage struct {
	segmentHandler SegmentHandler
	memtable       memtable.Memtable
}

type StorageEntry struct {
	Data []byte
}

func CreateStorage() *Storage {
	memtable := memtable.CreateBinaryMemtable()
	return &Storage{
		segmentHandler: SegmentHandler{},
		memtable:       memtable,
	}
}

func (s *Storage) SetEntry(key []byte, data []byte) error {
	s.memtable.Insert(key, data)

	if s.memtable.IsFull() {
		data := s.memtable.RemoveAll()
		err := s.segmentHandler.AppendSegment(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) GetEntry(key []byte) (*StorageEntry, error) {
	memtableEntry := s.memtable.Get(key)
	if memtableEntry != nil {
		return &StorageEntry{
			Data: memtableEntry,
		}, nil
	}
	segmentEntry := s.segmentHandler.FindItem(key)
	if segmentEntry != nil {
		return &StorageEntry{
			Data: segmentEntry.Data,
		}, nil
	}
	return nil, &ItemNotFoundError{
		key: string(key),
	}
}
