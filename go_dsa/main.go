package main

import (
	"fmt"
	"go_dsa/trees/binary_tree"
)

func main() {
	fmt.Println("Data Structures and Algorithms")

	binaryTree := binary_tree.BinaryTree{}
	binaryTree.Create([]int{5, 8, 6, -1, 9, 3, -1, 4, 2}, -1)
	binaryTree.PreOrderUsingRecursion(",")
	binaryTree.PostOrderUsingRecursion(",")
	binaryTree.InOrderUsingRecursion(",")
	binaryTree.InOrderUsingLoop(",")
	binaryTree.LevelOrderUsingLoop(",")
}
