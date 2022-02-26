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
	values := []int{}
	values = tree.preOrderUsingRecursion(values, sep, tree.root)
	tree.print(values, sep)
}

// Preorder Traversals using recursion, utils
func (tree *BinaryTree) preOrderUsingRecursion(values []int, sep string, node *Node) []int {
	if node != nil {
		values = append(values, node.Data)
		values = tree.preOrderUsingRecursion(values, sep, node.Left)
		values = tree.preOrderUsingRecursion(values, sep, node.Right)
	}
	return values
}

// InOrder Traversals using recursion
func (tree *BinaryTree) InOrderUsingRecursion(sep string) {
	values := []int{}
	values = tree.inOrderUsingRecursion(values, sep, tree.root)
	tree.print(values, sep)
}

// InOrder Traversals using recursion, utils
func (tree *BinaryTree) inOrderUsingRecursion(values []int, sep string, node *Node) []int {
	if node != nil {
		values = tree.inOrderUsingRecursion(values, sep, node.Left)
		values = append(values, node.Data)
		values = tree.inOrderUsingRecursion(values, sep, node.Right)
	}
	return values
}

// Postorder Traversals using recursion
func (tree *BinaryTree) PostOrderUsingRecursion(sep string) {
	values := []int{}
	values = tree.postOrderUsingRecursion(values, sep, tree.root)
	tree.print(values, sep)
}

// Postorder Traversals using recursion, utils
func (tree *BinaryTree) postOrderUsingRecursion(values []int, sep string, node *Node) []int {
	if node != nil {
		values = tree.postOrderUsingRecursion(values, sep, node.Left)
		values = tree.postOrderUsingRecursion(values, sep, node.Right)
		values = append(values, node.Data)
	}
	return values
}

// LevelOrder Traversals using Loop
func (tree *BinaryTree) LevelOrderUsingLoop(sep string) {
	if root := tree.root; root != nil {
		q := &Queue{}
		q.EnQueue(root)
		values := []int{}
		for !q.IsEmpty() {
			root, _ = q.DeQueue()
			values = append(values, root.Data)
			if root.Left != nil {
				q.EnQueue(root.Left)
			}
			if root.Right != nil {
				q.EnQueue(root.Right)
			}
		}
		tree.print(values, sep)
	}
}

// PreOrder Traversals using Loops
func (tree *BinaryTree) PreOrderUsingLoop(sep string) {
	stk := &Stack{}
	root := tree.root
	values := []int{}
	for !stk.IsEmpty() || root != nil {
		if root != nil {
			values = append(values, root.Data)
			stk.Push(root)
			root = root.Left
		} else {
			root, _ = stk.Pop()
			root = root.Right
		}
	}
	tree.print(values, sep)
}

// InOrder Traversals using Loops
func (tree *BinaryTree) InOrderUsingLoop(sep string) {
	stk := &Stack{}
	root := tree.root
	values := []int{}
	for !stk.IsEmpty() || root != nil {
		if root != nil {
			stk.Push(root)
			root = root.Left
		} else {
			root, _ = stk.Pop()
			values = append(values, root.Data)
			root = root.Right
		}
	}
	tree.print(values, sep)
}

// PostOrder Traversals using Loops
func (tree *BinaryTree) PostOrderUsingLoop(sep string) {
	stk := &Stack{}
	root := tree.root
	values := []int{}
	for !stk.IsEmpty() || root != nil {
		if root != nil {
			if root.Right != nil {
				stk.Push(root.Right)
			}
			stk.Push(root)
			root = root.Left
		} else {
			root, _ = stk.Pop()
			top, _ := stk.Peek()
			if root.Right != nil && root.Right == top {
				stk.Pop()
				stk.Push(root)
				root = root.Right
			} else {
				values = append(values, root.Data)
				root = nil
			}
		}
	}
	tree.print(values, sep)
}

// utility function for printing array of values
func (tree *BinaryTree) print(values []int, sep string) {
	n := len(values)
	index := 0
	for index < n {
		if index+1 == n {
			fmt.Println(values[index])
		} else {

			fmt.Print(values[index], sep)
		}
		index++
	}
}
