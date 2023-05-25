package memtable

import (
	hk "github.com/StevenSopilidis/SimpleDataEngineSystem/memtable/internal"
)

// implementation of memtable using binary search tree
type BinaryMemtable struct {
	head *node
}

type node struct {
	key   string
	data  interface{}
	left  *node
	right *node
}

func CreateBinaryMemtable() *BinaryMemtable {
	return &BinaryMemtable{
		head: nil,
	}
}

func (b *BinaryMemtable) Insert(key []byte, data interface{}) {
	if b.head == nil {
		b.head = &node{
			key:  hk.HashStringToSHA256(key),
			data: data,
		}
		return
	}
	recersiveInsert(&b.head, hk.HashStringToSHA256(key), data)
}

func recersiveInsert(head **node, key string, data interface{}) {
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

func (b *BinaryMemtable) Get(key []byte) interface{} {
	return recersiveGet(b.head, hk.HashStringToSHA256(key))
}

func recersiveGet(head *node, key string) interface{} {
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
