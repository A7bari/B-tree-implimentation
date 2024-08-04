package main

import (
	btree "btree/b+tree"
	"fmt"
)

func main() {

	// example usage
	b := btree.NewBTree(1, btree.Value{Value: "value1"})
	b.Insert(2, btree.Value{Value: "value2"})
	b.Insert(3, btree.Value{Value: "value3"})
	b.Insert(4, btree.Value{Value: "value4"})
	b.Insert(5, btree.Value{Value: "value5"})

	b.PrintTree()
	fmt.Println("Get value for key 3: ", b.Get(3))
}
