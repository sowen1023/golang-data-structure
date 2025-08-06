
### Summary of decoding the B+tree nodes
- btype(), nkeys(): Read the fixed-size header.
- setHeader(type, nkeys): Write the fixed-size header.
- getPtr(n), setPtr(n, ptr): Read and write the pointer array (for internal nodes).
- getOffset(n): Read the offsets array to locate the nth key in O(1).
- kvPos(n): Return the position of the nth key using getOffset().
- getKey(n): Get the nth key data as a slice.
- getVal(n): Get the nth value data as a slice (for leaf nodes).

### Summary of creating B+tree nodes
-  Nodes are immutable, we will only create new nodes from old nodes.
-  To create a new node:
-  Allocate a byte array.
-  Set the number of keys with setHeader().
-  Add each key in sort order with nodeAppendKV().
-  nbytes() returns the number of bytes used by the node.