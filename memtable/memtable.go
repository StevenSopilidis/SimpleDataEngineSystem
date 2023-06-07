package memtable

const (
	// max size of memtable before copying it to a segment
	MaxMemTableSize = 10
)

type Memtable interface {
	// inserts a value in to the memtable
	// if key already exists we just update the old value
	Insert(key []byte, data interface{})
	// gets a value from  the memtable
	// if it does not exist returns null
	Get(key []byte) interface{}
	// returns if the memtable is empty
	IsEmpty() bool
	// removes all elements from memtable and returns them
	// in an ascending order based on the key
	RemoveAll() map[string]interface{}
}
