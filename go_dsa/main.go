package main

import (
	"fmt"
	"go_dsa/trees/binary"
)

func main() {
	fmt.Println("Data Structures and Algorithms")

	binaryTree := binary.BinaryTree{}
	binaryTree.Create([]int{1, 2, 3, 4, 5})
	binaryTree.PreOrderUsingRecursion(",")
	binaryTree.PostOrderUsingRecursion(",")
	binaryTree.InOrderUsingRecursion(",")
	binaryTree.LevelOrderUsingLoop(",")
}
