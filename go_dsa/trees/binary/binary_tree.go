package binary

import (
	"errors"
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	root *Node
}

/*
	Creating tree using given values array
*/
func (tree *BinaryTree) Create(values []int, nilValue int) error {
	n := len(values)
	if n > 0 {
		if values[0] == nilValue {
			return errors.New("invalid tree nodes")
		}
		root := &Node{values[0], nil, nil}
		tree.root = root
		nodes := []*Node{root}
		index, value := 1, 0

		for index < n && len(nodes) != 0 {
			root = nodes[0]
			nodes = nodes[1:]

			value = values[index]
			if value != nilValue {
				root.Left = &Node{value, nil, nil}
				nodes = append(nodes, root.Left)
			}
			index++

			if index < n {
				value = values[index]
				if value != nilValue {
					root.Right = &Node{value, nil, nil}
					nodes = append(nodes, root.Right)
				}
				index++
			}
		}
		return nil
	} else {
		return errors.New("there is no values")
	}
}

/*
	Tree Traversals
*/

// Preorder Traversals using recursion
func (tree *BinaryTree) PreOrderUsingRecursion(sep string) {
	tree.preOrderUsingRecursion(sep, tree.root)
	fmt.Println()
}

// Preorder Traversals using recursion, utils
func (tree *BinaryTree) preOrderUsingRecursion(sep string, node *Node) {
	if node == nil {
		return
	} else {
		fmt.Print(node.Data, sep)
		tree.preOrderUsingRecursion(sep, node.Left)
		tree.preOrderUsingRecursion(sep, node.Right)
	}
}

// Postorder Traversals using recursion
func (tree *BinaryTree) PostOrderUsingRecursion(sep string) {
	tree.postOrderUsingRecursion(sep, tree.root)
	fmt.Println()
}

// Postorder Traversals using recursion, utils
func (tree *BinaryTree) postOrderUsingRecursion(sep string, node *Node) {
	if node == nil {
		return
	} else {
		tree.postOrderUsingRecursion(sep, node.Left)
		tree.postOrderUsingRecursion(sep, node.Right)
		fmt.Print(node.Data, sep)
	}
}

// InOrder Traversals using recursion
func (tree *BinaryTree) InOrderUsingRecursion(sep string) {
	tree.inOrderUsingRecursion(sep, tree.root)
	fmt.Println()
}

// InOrder Traversals using recursion, utils
func (tree *BinaryTree) inOrderUsingRecursion(sep string, node *Node) {
	if node == nil {
		return
	} else {
		tree.inOrderUsingRecursion(sep, node.Left)
		fmt.Print(node.Data, sep)
		tree.inOrderUsingRecursion(sep, node.Right)
	}
}

// LevelOrder Traversals using Loop
func (tree *BinaryTree) LevelOrderUsingLoop(sep string) {
	nodes := []*Node{}
	nodes = append(nodes, tree.root)
	n := len(nodes)
	for n > 0 {
		root := nodes[0]
		nodes = nodes[1:]
		if root.Left != nil {
			nodes = append(nodes, root.Left)
		}
		if root.Right != nil {
			nodes = append(nodes, root.Right)
		}

		n = len(nodes)
		if n == 0 {
			fmt.Println(root.Data)
		} else {
			fmt.Print(root.Data, sep)
		}
	}
}

// InOrder Traversals using Loops
func (tree *BinaryTree) InOrderUsingLoop(sep string) {
	stk := []*Node{}
	curr := tree.root
	for len(stk) != 0 || curr != nil {
		for curr != nil {
			stk = append(stk, curr)
			curr = curr.Left
		}

		n := len(stk)
		curr = stk[n-1]
		stk = stk[:n-1]
		fmt.Print(curr.Data, sep)
		curr = curr.Right
	}
	fmt.Println()
}
