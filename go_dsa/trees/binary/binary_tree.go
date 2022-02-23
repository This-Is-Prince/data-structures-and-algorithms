package binary

import "fmt"

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
func (tree *BinaryTree) Create(values []int) {
	n := len(values)
	nodes := make([]*Node, n)

	for index, value := range values {
		nodes[index] = &Node{value, nil, nil}
	}

	index, leftIndex, rightIndex := 0, 0, 0

	for index < n {
		if index == 0 {
			tree.root = nodes[index]
		}
		leftIndex = index*2 + 1
		rightIndex = index*2 + 2
		if leftIndex < n {
			nodes[index].Left = nodes[leftIndex]
		}
		if rightIndex < n {
			nodes[index].Right = nodes[rightIndex]
		}
		if leftIndex >= n && rightIndex >= n {
			break
		}
		index++
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
