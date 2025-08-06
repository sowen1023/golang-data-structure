package btree

import (
	"encoding/binary"
)

const (
	HEADER             = 4
	BNODE_NODE         = 1 // internal nodes with pointers
	BNODE_LEAF         = 2 // leaf nodes with values
	BTREE_PAGE_SIZE    = 4096
	BTREE_MAX_KEY_SIZE = 1000
	BTREE_MAX_VAL_SIZE = 3000
)

//func init() {
//	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
//	assert(node1max > BTREE_PAGE_SIZE)
//}

type BTree struct {
	root uint64 // disk page number
	// callbacks to manage disk page references
	get func(uint64) BNode // to dereference pointer
	new func(BNode) uint64 // to allocate new page
	del func(uint64)       // to deallocate new page
}

type Node struct {
	keys [][]byte
	vals [][]byte
	kids []*Node
}

/*
| type | nkeys | pointers | offsets |            key-values           | unused |
|   2  |   2   | nil nil  |  8 19   | 2 2 "k1" "hi"  2 5 "k3" "hello" |        |
|  2B  |  2B   |   2×8B   |  2×2B   | 4B + 2B + 2B + 4B + 2B + 5B     |        |
*/

type BNode []byte // can be dumped to the disk

func (node BNode) btype() uint16 {
	return binary.LittleEndian.Uint16(node[0:2])
}

func (node BNode) nkeys() uint16 {
	return binary.LittleEndian.Uint16(node[2:4])
}

func (node BNode) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node[0:2], btype)
	binary.LittleEndian.PutUint16(node[2:4], nkeys)
}

func (node BNode) getPtr(idx uint16) uint64 {
	assert(idx < node.nkeys())
	pos := 4 + 8*idx
	return binary.LittleEndian.Uint64(node[pos:])
}

func (node BNode) setPtr(idx uint16, val uint64) {
	assert(idx < node.nkeys())
	pos := 4 + 8*idx
	binary.LittleEndian.PutUint64(node[pos:], val)
}

// offset functions
// returns the value of the offset i.e. the location of the kv-pair at given index
func offsetPos(node BNode, idx uint16) uint16 {
	assert(idx >= 1 && idx <= node.nkeys())

	return HEADER + 8*node.nkeys() + 2*(idx-1)
}

// Read KV pairs
// The start offset of the n-th KV pairs is read from the offset array.

// read the `offsets` array
func (node BNode) getOffset(idx uint16) uint16 {
	if idx == 0 {
		return 0
	}
	pos := 4 + 8*node.nkeys() + 2*(idx-1)
	return binary.LittleEndian.Uint16(node[pos:])
}

// updates offset for kv-pair at given index
func (node BNode) setOffset(idx uint16, offset uint16) {
	pos := offsetPos(node, idx)
	binary.LittleEndian.PutUint16(node[pos:], offset)
}

func (node BNode) kvPos(idx uint16) uint16 {
	assert(idx <= node.nkeys())
	return 4 + 8*node.nkeys() + 2*node.nkeys() + node.getOffset(idx)
}

func (node BNode) getKey(idx uint16) []byte {
	assert(idx < node.nkeys())
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node[pos:])
	return node[pos+4:][:klen]
}

func (node BNode) getVal(idx uint16) []byte {
	assert(idx < node.nkeys())
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node[pos+0:])
	vlen := binary.LittleEndian.Uint16(node[pos+2:])
	return node[pos+4+klen:][:vlen]
}

func nodeAppendKV(new BNode, idx uint16, ptr uint64, key []byte, val []byte) {
	new.setPtr(idx, ptr)
	// KVs
	pos := new.kvPos(idx) // uses the offset value of the previous key
	// 4-bytes KV sizes
	binary.LittleEndian.PutUint16(new[pos+0:], uint16(len(key)))
	binary.LittleEndian.PutUint16(new[pos+2:], uint16(len(val)))
	// KV data
	copy(new[pos+4:], key)
	copy(new[pos+4+uint16(len(key)):], val)
	// update the offset value for the next key
	new.setOffset(idx+1, new.getOffset(idx)+4+uint16(len(key)+len(val)))
}

// node size in bytes
func (node BNode) nbytes() uint16 {
	return node.kvPos(node.nkeys()) // uses the offset value of the last key
}
