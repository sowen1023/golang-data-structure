/* Copyright © INFINI LTD. All rights reserved.
 * Web: https://infinilabs.com
 * Email: hello#infini.ltd */

package btree

//func test() {
//	// Example 1, node {"k1":"hi", "k2":"a", "k3":"hello"} is updated with "k2":"b":
//	node := BNode(make([]byte, BTREE_PAGE_SIZE))
//	node.setHeader(BNODE_LEAF, 3)
//	nodeAppendKV(node, 0, 0, old.getKey(0), old.getVal(0))
//	nodeAppendKV(node, 1, 0, []byte("k2"), []byte("b"))
//	nodeAppendKV(node, 2, 0, old.getKey(2), old.getVal(2))
//
//	// Example 2, remove "k2" from that node:
//	new := BNode(make([]byte, BTREE_PAGE_SIZE))
//	new.setHeader(BNODE_LEAF, 2)
//	nodeAppendKV(new, 0, 0, old.getKey(0), old.getVal(0))
//	nodeAppendKV(new, 1, 0, old.getKey(2), old.getVal(2))
//
//	// Example 3, insert "a":"b" into that node:
//	new := BNode(make([]byte, 2*BTREE_PAGE_SIZE)) // larger
//	new.setHeader(BNODE_LEAF, 4)
//	nodeAppendKV(new, 0, 0, []byte("a"), []byte("b"))
//	nodeAppendKV(new, 1, 0, old.getKey(0), old.getVal(0))
//	nodeAppendKV(new, 2, 0, old.getKey(1), old.getVal(1))
//	nodeAppendKV(new, 3, 0, old.getKey(2), old.getVal(2))
//}
