package circular

import (
	"errors"
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
	display in sequential order functions
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

/*
	display in reverse order functions
*/
// Display function using loops
func (list *LinkedList) DisplayInReverseOrderUsingLoop(sep string) {
	node := list.root
	last := list.root
	for {
		for node.Next != last {
			node = node.Next
		}
		last = node
		if node == list.root {
			fmt.Println(node.Data)
			break
		} else {
			fmt.Print(node.Data, sep)
			node = list.root
		}
	}
}

// Display function using recursion
func (list *LinkedList) DisplayInReverseOrderUsingRecursion(sep string) {
	list.displayInReverseOrderUsingRecursion(sep, list.root)
}

// Display function using recursion, utility
func (list *LinkedList) displayInReverseOrderUsingRecursion(sep string, node *Node) {
	if node.Next == list.root {
		fmt.Print(node.Data, sep)
	} else {
		list.displayInReverseOrderUsingRecursion(sep, node.Next)
		if node == list.root {
			fmt.Println(node.Data)
		} else {
			fmt.Print(node.Data, sep)
		}
	}
}

/*
	count Function
*/

// count function using loops
func (list *LinkedList) CountUsingLoop() int {
	noOfNodes := 0
	node := list.root
	for node != nil {
		noOfNodes++
		node = node.Next
		if node == list.root {
			break
		}
	}
	return noOfNodes
}

// count function using recursion
func (list *LinkedList) CountUsingRecursion() int {
	return list.countUsingRecursion(list.root)
}

// count function using recursion, utils
func (list *LinkedList) countUsingRecursion(node *Node) int {
	if node == nil {
		return 0
	} else if node.Next == list.root {
		return 1
	} else {
		return list.countUsingRecursion(node.Next) + 1
	}
}

/*
	sum Function
*/
// sum function using loops
func (list *LinkedList) SumUsingLoop() int {
	sum := 0
	node := list.root
	for node != nil {
		sum += node.Data
		node = node.Next
		if node == list.root {
			break
		}
	}
	return sum
}

// sum function using recursion
func (list *LinkedList) SumUsingRecursion() int {
	return list.sumUsingRecursion(list.root)
}

// sum function using recursion, utils
func (list *LinkedList) sumUsingRecursion(node *Node) int {
	if node == nil {
		return 0
	} else if node.Next == list.root {
		return node.Data
	} else {
		return list.sumUsingRecursion(node.Next) + node.Data
	}
}

/*
	average Function
*/
// average function using loops
func (list *LinkedList) AverageUsingLoop() float64 {
	count := list.CountUsingLoop()
	if count == 0 {
		return 0.0
	}
	return float64(list.SumUsingLoop()) / float64(count)
}

// average function using recursion
func (list *LinkedList) AverageUsingRecursion() float64 {
	count := list.CountUsingRecursion()
	if count == 0 {
		return 0.0
	}
	return float64(list.SumUsingRecursion()) / float64(count)
}

/*
	maximum Function
*/
// maximum function using loop
func (list *LinkedList) MaxUsingLoop() (int, error) {
	max := -1
	if node := list.root; node != nil {
		max = node.Data
		node = node.Next
		for node != list.root {
			if max < node.Data {
				max = node.Data
			}
			node = node.Next
		}
		return max, nil
	} else {
		return max, errors.New("linkedlist is empty")
	}
}

// maximum function using recursion
func (list *LinkedList) MaxUsingRecursion() (int, error) {
	if list.root == nil {
		return -1, errors.New("linkedlist is empty")
	} else {
		return list.maxUsingRecursion(list.root), nil
	}
}

// maximum function using recursion, utils
func (list *LinkedList) maxUsingRecursion(node *Node) int {
	if node.Next == list.root {
		return node.Data
	} else {
		data := list.maxUsingRecursion(node.Next)
		if data > node.Data {
			return data
		} else {
			return node.Data
		}
	}
}

/*
	minimum Function
*/
// minimum function using loop
func (list *LinkedList) MinUsingLoop() (int, error) {
	min := 0
	if node := list.root; node != nil {
		min = node.Data
		node = node.Next
		for node != list.root {
			if min > node.Data {
				min = node.Data
			}
			node = node.Next
		}
		return min, nil
	} else {
		return min, errors.New("linkedlist is empty")
	}
}

// minimum function using recursion
func (list *LinkedList) MinUsingRecursion() (int, error) {
	if list.root == nil {
		return 0, errors.New("linkedlist is empty")
	} else {
		return list.minUsingRecursion(list.root), nil
	}
}

// minimum function using recursion, utils
func (list *LinkedList) minUsingRecursion(node *Node) int {
	if node.Next == list.root {
		return node.Data
	} else {
		data := list.minUsingRecursion(node.Next)
		if data < node.Data {
			return data
		} else {
			return node.Data
		}
	}
}

/*
	search Function
*/

// linearSearch function using loop
func (list *LinkedList) LinearSearchUsingLoop(key int) (*Node, error) {
	if node := list.root; node != nil {
		for {
			if node.Data == key {
				return node, nil
			} else {
				node = node.Next
				if node == list.root {
					return nil, errors.New("no element found in linkedlist")
				}
			}
		}
	} else {
		return nil, errors.New("linkedlist is empty")
	}
}

// linearSearch function using Recursion
func (list *LinkedList) LinearSearchUsingRecursion(key int) (*Node, error) {
	if list.root == nil {
		return nil, errors.New("linkedlist is empty")
	} else {
		return list.linearSearchUsingRecursion(list.root, key)
	}
}

// linearSearch function using Recursion, utils
func (list *LinkedList) linearSearchUsingRecursion(node *Node, key int) (*Node, error) {
	if node.Data == key {
		return node, nil
	} else if node.Next == list.root {
		return nil, errors.New("no element found in linkedlist")
	} else {
		return list.linearSearchUsingRecursion(node.Next, key)
	}
}
