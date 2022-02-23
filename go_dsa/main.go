package main

import (
	"fmt"
	"go_dsa/trees/binary"
)

func main() {
	fmt.Println("Data Structures and Algorithms")

	binaryTree := binary.BinaryTree{}
	binaryTree.Create([]int{5, 8, 6, -1, 9, 3, -1, 4, 2}, -1)
	binaryTree.PreOrderUsingRecursion(",")
	binaryTree.PostOrderUsingRecursion(",")
	binaryTree.InOrderUsingRecursion(",")
	binaryTree.InOrderUsingLoop(",")
	binaryTree.LevelOrderUsingLoop(",")
}
