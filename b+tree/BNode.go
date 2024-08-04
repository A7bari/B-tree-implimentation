package btree

type Value struct {
	Value string
}

// BNode represents a node in a B+ tree
type BNode struct {
	isLeaf   bool
	keys     []int
	children []*BNode
	values   []Value
}

func search(node *BNode, key int) *BNode {
	nkeys := len(node.keys)
	found := -1
	// the first key is a copy from the parent node,
	// thus it's always less than or equal to the key.
	for i := 0; i < nkeys && node.keys[i] <= key; i++ {
		found = i
	}

	if node.isLeaf {
		if found >= 0 && node.keys[found] == key {
			return node
		}
		return nil
	}

	return search(node.children[found+1], key)
}

func (node *BNode) leafInsert(key int, value Value) *BNode {
	nkeys := len(node.keys)
	for i := 0; i < nkeys; i++ {
		if node.keys[i] == key {
			node.values[i] = value
			return nil
		}
		if node.keys[i] > key {
			node.keys = append(node.keys[:i], append([]int{key}, node.keys[i:]...)...)
			node.values = append(node.values[:i], append([]Value{value}, node.values[i:]...)...)
			return nil
		}
	}
	node.keys = append(node.keys, key)
	node.values = append(node.values, value)
	return nil
}

func (node *BNode) insertNonFull(key int, value Value) *BNode {
	if node.isLeaf {
		// insert the key and value into the leaf node
		return node.leafInsert(key, value)
	}
	// find the child node to insert the key
	nkeys := len(node.keys)
	found := 0
	for i := 0; i < nkeys && node.keys[i] < key; i++ {
		found = i + 1
	}

	child := node.children[found]
	if len(child.keys) == MAX_DEGREE {
		// split the child node
		splitChild(node, found)
		return node.insertNonFull(key, value)
	}
	return child.insertNonFull(key, value)
}

func splitChild(node *BNode, idx int) {
	child := node.children[idx]
	newChild := BNode{isLeaf: child.isLeaf}

	mid := (MAX_DEGREE / 2)

	// move the last (MAX_DEGREE-1)/2 keys and children to the new child
	//keys
	newChild.keys = child.keys[mid:]
	child.keys = child.keys[:mid]

	// values if it's a leaf node
	if newChild.isLeaf {
		newChild.values = child.values[mid:]
		child.values = child.values[:mid]
	} else {
		//children
		newChild.children = child.children[mid:]
		child.children = child.children[:mid]
	}

	// insert the middle key of the child node into the parent node
	if len(node.keys) == 0 {
		node.keys = []int{newChild.keys[0]}
	} else if idx == 0 {
		node.keys = append([]int{newChild.keys[0]}, node.keys[1:]...)
	} else {
		node.keys = append(node.keys[:idx], append([]int{newChild.keys[0]}, node.keys[idx:]...)...)
	}
	node.children = append(node.children[:idx+1], append([]*BNode{&newChild}, node.children[idx+1:]...)...)
}
