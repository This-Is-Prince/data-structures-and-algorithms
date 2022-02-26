package binary_tree

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
		q := &Queue{}
		if values[0] == nilValue {
			return errors.New("invalid tree root")
		}
		index, value := 1, 0
		tree.root = (&Node{values[0], nil, nil})
		q.EnQueue(tree.root)

		for index < n && !q.IsEmpty() {
			root, _ := q.DeQueue()
			value = values[index]

			if value != nilValue {
				root.Left = &Node{Data: value}
				q.EnQueue(root.Left)
			}

			if index+1 < n {
				value = values[index+1]
				if value != nilValue {
					root.Right = &Node{Data: value}
					q.EnQueue(root.Right)
				}
			}
			index += 2
		}
		return nil
	} else {
		return errors.New("tree is empty")
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
	if tree.root != nil {
		q := &Queue{}
		q.EnQueue(tree.root)
		for !q.IsEmpty() {
			root, _ := q.DeQueue()
			if root.Left != nil {
				q.EnQueue(root.Left)
			}
			if root.Right != nil {
				q.EnQueue(root.Right)
			}
			if q.IsEmpty() {
				fmt.Println(root.Data)
			} else {
				fmt.Print(root.Data, sep)
			}
		}
	}
}

// InOrder Traversals using Loops
func (tree *BinaryTree) InOrderUsingLoop(sep string) {
	stk := &Stack{}
	root := tree.root
	for !stk.IsEmpty() || root != nil {
		for root != nil {
			stk.Push(root)
			root = root.Left
		}
		root, _ = stk.Pop()
		fmt.Print(root.Data, sep)
		root = root.Right
	}
	fmt.Println()
}

// PreOrder Traversals using Loops
func (tree *BinaryTree) PreOrderUsingLoop(sep string) {
	if tree.root != nil {
		stk := &Stack{}
		stk.Push(tree.root)
		for !stk.IsEmpty() {
			root, _ := stk.Pop()
			if root.Right != nil {
				stk.Push(root.Right)
			}
			if root.Left != nil {
				stk.Push(root.Left)
			}
			if stk.IsEmpty() {
				fmt.Println(root.Data)
			} else {
				fmt.Print(root.Data, sep)
			}
		}
	}
}
