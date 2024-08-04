package btree

import (
	"fmt"
	"strings"
)

type BTree struct {
	root       *BNode
	value_type string
}

func NewBTree(key int, value Value) *BTree {
	return &BTree{
		root: &BNode{
			isLeaf: true,
			keys:   []int{key},
			values: []Value{value},
		},
		value_type: "int",
	}
}

func (b *BTree) Insert(key int, value Value) {
	root := b.root
	if len(root.keys) == MAX_DEGREE {
		// split the root node
		newRoot := &BNode{isLeaf: false}
		newRoot.children = []*BNode{root}
		b.root = newRoot
		splitChild(newRoot, 0)
		newRoot.insertNonFull(key, value)
	} else {
		root.insertNonFull(key, value)
	}
}

func (t *BTree) Get(key int) *Value {
	node := search(t.root, key)
	if node == nil {
		return nil
	}
	nkeys := len(node.keys)
	for i := 0; i < nkeys; i++ {
		if node.keys[i] == key {
			return &node.values[i]
		}
	}
	return nil
}

func (t *BTree) PrintTree() {
	if t.root != nil {
		t.printNode(t.root, 0)
	}
}

func (t *BTree) printNode(node *BNode, level int) {
	fmt.Println(strings.Repeat("     ", level), node.keys)
	if !node.isLeaf {
		for _, child := range node.children {
			if child.isLeaf {
				fmt.Println(strings.Repeat("     ", level+1), child.keys, child.values)
			} else {
				t.printNode(child, level+1)
			}
		}
	}

}
