package memtable

type Memtable interface {
	// inserts a value in to the memtable
	// if key already exists we just update the old value
	Insert(key []byte, data interface{})
	// gets a value from  the memtable
	// if it does not exist returns null
	Get(key []byte) interface{}
}
