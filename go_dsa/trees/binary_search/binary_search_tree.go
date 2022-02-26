package binary_search

import (
	"errors"
	"fmt"
	"math"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	root *Node
}

/*
	Generate BST Using Traversals
*/
// generate bst using preorder
func (tree *BinarySearchTree) GenerateUsingPreOrder(values ...int) error {
	if tree.root == nil {
		if length := len(values); length > 0 {
			root, stk, index, value := &Node{Data: values[0]}, &Stack{}, 1, 0
			tree.root = root
			var newNode *Node

			for index < length {
				value = values[index]
				if value == root.Data {
					return errors.New("duplicate values")
				} else {
					newNode = &Node{Data: value}
					if value < root.Data {
						stk.Push(root)
						root.Left = newNode
					} else {
						peek, _ := stk.Peek()
						for !stk.IsEmpty() && peek.Data <= value {
							if value == peek.Data {
								return errors.New("duplicate values")
							}
							root, _ = stk.Pop()
							peek, _ = stk.Peek()
						}
						root.Right = newNode
					}
					root = newNode
					index++
				}
			}
			return nil
		} else {
			return errors.New("can't generate tree, there is no values")
		}
	} else {
		return errors.New("tree is already present")
	}
}

/*
	Create
*/
// Create using loop
func (tree *BinarySearchTree) CreateUsingLoop(values ...int) error {
	if tree.root == nil {
		var err error
		for _, value := range values {
			err = tree.InsertUsingLoop(value)
			if err != nil {
				tree.root = nil
				return err
			}
		}
		return nil
	} else {
		return errors.New("tree already exists")
	}
}

// Create using recursion
func (tree *BinarySearchTree) CreateUsingRecursion(values ...int) error {
	if tree.root == nil {
		var err error
		for _, value := range values {
			err = tree.InsertUsingRecursion(value)
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return errors.New("tree already exists")
	}
}

/*
	Deletion
*/

// delete function
func (tree *BinarySearchTree) Delete(key int) error {
	if tree.root == nil {
		return errors.New("tree is empty")
	} else {
		var err error
		tree.root, err = tree.delete(tree.root, key)
		return err
	}
}

// delete function, utils
func (tree *BinarySearchTree) delete(root *Node, key int) (*Node, error) {
	if root == nil {
		return nil, errors.New("can't delete, key is not found")
	} else if root.Left == nil && root.Right == nil {
		if root.Data == key {
			return nil, nil
		} else {
			return root, errors.New("can't delete, key is not found")
		}
	} else {
		var err error
		if key < root.Data {
			root.Left, err = tree.delete(root.Left, key)
		} else if key > root.Data {
			root.Right, err = tree.delete(root.Right, key)
		} else {
			if tree.height(root.Left) > tree.height(root.Right) {
				inPre := tree.inOrderPredecessor(root.Left)
				root.Data = inPre.Data
				root.Left, err = tree.delete(root.Left, inPre.Data)
			} else {
				inSuc := tree.inOrderSuccessor(root.Right)
				root.Data = inSuc.Data
				root.Right, err = tree.delete(root.Right, inSuc.Data)
			}
		}
		return root, err
	}
}

/*
	InOrder Predecessor
*/
// InOrder Predecessor function
func (tree *BinarySearchTree) InOrderPredecessor() *Node {
	return tree.inOrderPredecessor(tree.root)
}

// InOrder Predecessor function, utils
func (tree *BinarySearchTree) inOrderPredecessor(root *Node) *Node {
	for root != nil && root.Right != nil {
		root = root.Right
	}
	return root
}

/*
	InOrder Successor
*/
// InOrder Successor function
func (tree *BinarySearchTree) InOrderSuccessor() *Node {
	return tree.inOrderPredecessor(tree.root)
}

// InOrder Successor function, utils
func (tree *BinarySearchTree) inOrderSuccessor(root *Node) *Node {
	for root != nil && root.Left != nil {
		root = root.Left
	}
	return root
}

/*
	Height
*/
// height function
func (tree *BinarySearchTree) Height() int {
	return int(tree.height(tree.root))
}

// height function, utils
func (tree *BinarySearchTree) height(root *Node) float64 {
	if root == nil {
		return 0
	} else {
		leftHeight := tree.height(root.Left)
		rightHeight := tree.height(root.Right)
		return math.Max(leftHeight, rightHeight) + 1
	}
}

/*
	Insertion
*/
// insert using loop
func (tree *BinarySearchTree) InsertUsingLoop(value int) error {
	if newNode := (&Node{Data: value}); newNode == nil {
		return errors.New("tree is full")
	} else {
		if root := tree.root; root == nil {
			tree.root = newNode
		} else {
			var prev *Node
			for root != nil {
				prev = root
				if root.Data == value {
					return errors.New("value already exist, duplicate insertion")
				} else if value < root.Data {
					root = root.Left
				} else {
					root = root.Right
				}
			}
			if value < prev.Data {
				prev.Left = newNode
			} else {
				prev.Right = newNode
			}
		}
		return nil
	}
}

// insert using recursion
func (tree *BinarySearchTree) InsertUsingRecursion(value int) error {
	var err error
	tree.root, err = tree.insertUsingRecursion(value, tree.root)
	return err
}

// insert using recursion, utils
func (tree *BinarySearchTree) insertUsingRecursion(value int, root *Node) (*Node, error) {
	if root == nil {
		return &Node{Data: value}, nil
	} else if root.Data == value {
		return nil, errors.New("value already exist, duplicate insertion")
	} else if value < root.Data {
		left, err := tree.insertUsingRecursion(value, root.Left)
		if err != nil {
			return nil, err
		}
		root.Left = left
	} else {
		right, err := tree.insertUsingRecursion(value, root.Right)
		if err != nil {
			return nil, err
		}
		root.Right = right
	}
	return root, nil
}

/*
	Searching
*/
// searching using loop
func (tree *BinarySearchTree) SearchUsingLoop(key int) (*Node, error) {
	root := tree.root
	for root != nil {
		if root.Data == key {
			return root, nil
		} else if key < root.Data {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil, errors.New("key not found")
}

// searching using recursion
func (tree *BinarySearchTree) SearchUsingRecursion(key int) (*Node, error) {
	return tree.searchUsingRecursion(key, tree.root)
}

// searching using recursion, utils
func (tree *BinarySearchTree) searchUsingRecursion(key int, root *Node) (*Node, error) {
	if root == nil {
		return nil, errors.New("key not found")
	} else {
		if root.Data == key {
			return root, nil
		} else if key < root.Data {
			return tree.searchUsingRecursion(key, root.Left)
		} else {
			return tree.searchUsingRecursion(key, root.Right)
		}
	}
}

/*
	Traversals
*/

// preorder using recursion
func (tree *BinarySearchTree) PreOrderUsingRecursion(sep string) {
	values := []int{}
	values = tree.preOrderUsingRecursion(values, sep, tree.root)
	tree.print(values, sep)
}

// preorder using recursion, utils
func (tree *BinarySearchTree) preOrderUsingRecursion(values []int, sep string, root *Node) []int {
	if root != nil {
		values = append(values, root.Data)
		values = tree.preOrderUsingRecursion(values, sep, root.Left)
		values = tree.preOrderUsingRecursion(values, sep, root.Right)
	}
	return values
}

// preorder using recursion
func (tree *BinarySearchTree) InOrderUsingRecursion(sep string) {
	values := []int{}
	values = tree.inOrderUsingRecursion(values, sep, tree.root)
	tree.print(values, sep)
}

// preorder using recursion, utils
func (tree *BinarySearchTree) inOrderUsingRecursion(values []int, sep string, root *Node) []int {
	if root != nil {
		values = tree.inOrderUsingRecursion(values, sep, root.Left)
		values = append(values, root.Data)
		values = tree.inOrderUsingRecursion(values, sep, root.Right)
	}
	return values
}

// postorder using recursion
func (tree *BinarySearchTree) PostOrderUsingRecursion(sep string) {
	values := []int{}
	values = tree.postOrderUsingRecursion(values, sep, tree.root)
	tree.print(values, sep)
}

// postorder using recursion, utils
func (tree *BinarySearchTree) postOrderUsingRecursion(values []int, sep string, root *Node) []int {
	if root != nil {
		values = tree.postOrderUsingRecursion(values, sep, root.Left)
		values = tree.postOrderUsingRecursion(values, sep, root.Right)
		values = append(values, root.Data)
	}
	return values
}

// print utility function
func (tree *BinarySearchTree) print(values []int, sep string) {
	index, n := 0, len(values)
	for index < n {
		if index+1 == n {
			fmt.Println(values[index])
		} else {
			fmt.Print(values[index], sep)
		}
		index++
	}
}
