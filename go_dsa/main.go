package main

import (
	"fmt"
	// "go_dsa/trees/binary_tree"
	"go_dsa/stack/using_array"
)

func main() {
	fmt.Println("Data Structures and Algorithms")

	// binaryTree := binary_tree.BinaryTree{}
	// binaryTree.Create([]int{5, 8, 6, -1, 9, 3, -1, 4, 2}, -1)
	// binaryTree.PreOrderUsingRecursion(",")
	// binaryTree.PostOrderUsingRecursion(",")
	// binaryTree.InOrderUsingRecursion(",")
	// binaryTree.InOrderUsingLoop(",")
	// binaryTree.LevelOrderUsingLoop(",")

	stk, err := using_array.Create(10)
	if err != nil {
		fmt.Println(err)
	} else {
		stk.Display(",")
		stk.Push(1)
		stk.Push(2)
		stk.Push(3)
		stk.Push(4)
		stk.Push(5)
		stk.Display(",")
	}

}
