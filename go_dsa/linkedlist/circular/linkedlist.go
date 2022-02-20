package circular

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	root *Node
}

func (list *LinkedList) Create(values ...int) {
	var last, root *Node
	for _, val := range values {
		if list.root == nil {
			root = &Node{val, nil}
			list.root = root
			root.Next = root
			last = root
		} else {
			last.Next = &Node{val, root}
			last = last.Next
		}
	}
}

/*
All Display functions
*/
// Display function using loops
func (list *LinkedList) DisplayUsingLoop(sep string) {
	node := list.root
	for node != nil {
		if node.Next == list.root {
			fmt.Println(node.Data)
			break
		} else {
			fmt.Print(node.Data, sep)
		}
		node = node.Next
	}
}

// Display function using recursion
func (list *LinkedList) DisplayUsingRecursion(sep string) {
	list.displayUsingRecursion(sep, list.root)
}

// Display function using recursion, utility
func (list *LinkedList) displayUsingRecursion(sep string, node *Node) {
	if node.Next == list.root {
		fmt.Println(node.Data)
	} else {
		fmt.Print(node.Data, sep)
		list.displayUsingRecursion(sep, node.Next)
	}
}
