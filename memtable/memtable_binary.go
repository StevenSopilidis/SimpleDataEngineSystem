package memtable

import (
	hk "github.com/StevenSopilidis/SimpleDataEngineSystem/internal"
)

// implementation of memtable using binary search tree
type BinaryMemtable struct {
	head *node
	// current size of memtable
	currentSize int
}

type node struct {
	key   string
	data  []byte
	left  *node
	right *node
}

func CreateBinaryMemtable() *BinaryMemtable {
	return &BinaryMemtable{
		head: nil,
	}
}

func (b *BinaryMemtable) IsEmpty() bool {
	return b.head == nil
}

func (b *BinaryMemtable) Insert(key []byte, data []byte) {
	if b.head == nil {
		b.head = &node{
			key:  hk.HashStringToSHA256(key),
			data: data,
		}
		return
	}
	recersiveInsert(&b.head, hk.HashStringToSHA256(key), data)
	b.currentSize++
}

func recersiveInsert(head **node, key string, data []byte) {
	if *head == nil {
		*head = &node{
			key:  key,
			data: data,
		}
	} else if (*head).key < key {
		recersiveInsert(&(*head).right, key, data)
	} else if (*head).key > key {
		recersiveInsert(&(*head).left, key, data)
	}
}

func (b *BinaryMemtable) Get(key []byte) []byte {
	return recersiveGet(b.head, hk.HashStringToSHA256(key))
}

func (b *BinaryMemtable) RemoveAll() map[string][]byte {
	data := make(map[string][]byte)
	recersiveRemove(b.head, &data)
	b.head = nil
	b.currentSize = 0
	return data
}

func recersiveRemove(head *node, buffer *map[string][]byte) {
	if head == nil {
		return

	} else {
		recersiveRemove(head.left, buffer)
		(*buffer)[head.key] = head.data
		recersiveRemove(head.right, buffer)
	}
}

func recersiveGet(head *node, key string) []byte {
	if head == nil {
		return nil
	}
	if head.key == key {
		return head.data
	} else if head.key < key {
		return recersiveGet(head.right, key)
	} else {
		return recersiveGet(head.left, key)
	}
}

func (b *BinaryMemtable) IsFull() bool {
	return b.currentSize >= MaxMemTableSize
}
