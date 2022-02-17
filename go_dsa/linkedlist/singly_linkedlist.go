package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList struct {
	root *Node
}

func (list *SinglyLinkedList) Create(values ...int) {
	var last *Node
	for _, val := range values {
		if list.root == nil {
			list.root = &Node{val, nil}
			last = list.root
		} else {
			last.Next = &Node{val, nil}
			last = last.Next
		}
	}
}

/*
All Display functions
*/

// Display function using loops
func (list *SinglyLinkedList) DisplayUsingLoop(sep string) {
	tmp := list.root
	for tmp != nil {
		if tmp.Next == nil {
			fmt.Println(tmp.Data)
		} else {
			fmt.Print(tmp.Data, sep)
		}
		tmp = tmp.Next
	}
}

func (list *SinglyLinkedList) DisplayUsingLoopInReverseOrder(sep string) {
	tmp := list.root
	var last *Node
	for tmp != last {
		for tmp.Next != last {
			tmp = tmp.Next
		}
		last = tmp
		if tmp == list.root {
			fmt.Println(tmp.Data)
		} else {
			fmt.Print(tmp.Data, sep)
			tmp = list.root
		}
	}
}

// Display function using recursion
func (list *SinglyLinkedList) DisplayUsingRecursion(sep string) {
	list.displayUsingRecursion(sep, list.root)
}
func (list *SinglyLinkedList) DisplayUsingRecursionInReverseOrder(sep string) {
	list.displayUsingRecursionInReverseOrder(sep, list.root)
}

// Display function using recursions, utility function
func (list *SinglyLinkedList) displayUsingRecursion(sep string, node *Node) {
	if node == nil {
		return
	} else if node.Next == nil {
		fmt.Println(node.Data)
	} else {
		fmt.Print(node.Data, sep)
		list.displayUsingRecursion(sep, node.Next)
	}
}
func (list *SinglyLinkedList) displayUsingRecursionInReverseOrder(sep string, node *Node) {
	if node == nil {
		return
	} else {
		list.displayUsingRecursionInReverseOrder(sep, node.Next)
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
func (list *SinglyLinkedList) CountUsingLoop() int {
	noOfNodes := 0
	tmp := list.root
	for tmp != nil {
		tmp = tmp.Next
		noOfNodes++
	}
	return noOfNodes
}

// count function using recursion
func (list *SinglyLinkedList) CountUsingRecursion() int {
	return list.countUsingRecursion(list.root)
}

// count function using recursion, utils
func (list *SinglyLinkedList) countUsingRecursion(node *Node) int {
	if node == nil {
		return 0
	} else {
		return list.countUsingRecursion(node.Next) + 1
	}
}

/*
	sum Function
*/
// sum function using loops
func (list *SinglyLinkedList) SumUsingLoop() int {
	sum := 0
	tmp := list.root
	for tmp != nil {
		sum += tmp.Data
		tmp = tmp.Next
	}
	return sum
}

// sum function using recursion
func (list *SinglyLinkedList) SumUsingRecursion() int {
	return list.sumUsingRecursion(list.root)
}

// sum function using recursion, utils
func (list *SinglyLinkedList) sumUsingRecursion(node *Node) int {
	if node == nil {
		return 0
	} else {
		return list.sumUsingRecursion(node.Next) + node.Data
	}
}

/*
	average Function
*/
// average function using loops
func (list *SinglyLinkedList) AverageUsingLoop() float64 {
	sum := 0
	tmp := list.root
	for tmp != nil {
		sum += tmp.Data
		tmp = tmp.Next
	}
	return float64(sum / list.CountUsingLoop())
}

// average function using recursion
func (list *SinglyLinkedList) AverageUsingRecursion() float64 {
	return float64(list.averageUsingRecursion(list.root) / list.CountUsingLoop())
}

// average function using recursion, utils
func (list *SinglyLinkedList) averageUsingRecursion(node *Node) int {
	if node == nil {
		return 0
	} else {
		return list.averageUsingRecursion(node.Next) + node.Data
	}
}

/*
	maximum Function
*/
// maximum function using loop
func (list *SinglyLinkedList) MaxUsingLoop() (int, error) {
	max := -1
	if tmp := list.root; tmp != nil {
		max = tmp.Data
		tmp = tmp.Next
		for tmp != nil {
			if max < tmp.Data {
				max = tmp.Data
			}
			tmp = tmp.Next
		}
		return max, nil
	} else {
		return max, errors.New("linkedlist is empty")
	}
}

// maximum function using recursion
func (list *SinglyLinkedList) MaxUsingRecursion() (int, error) {
	if list.root == nil {
		return -1, errors.New("linkedlist is empty")
	} else {
		return list.maxUsingRecursion(list.root), nil
	}
}

// maximum function using recursion, utils
func (list *SinglyLinkedList) maxUsingRecursion(node *Node) int {
	if node.Next == nil {
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
func (list *SinglyLinkedList) MinUsingLoop() (int, error) {
	min := 0
	if tmp := list.root; tmp != nil {
		min = tmp.Data
		tmp = tmp.Next
		for tmp != nil {
			if min > tmp.Data {
				min = tmp.Data
			}
			tmp = tmp.Next
		}
		return min, nil
	} else {
		return min, errors.New("linkedlist is empty")
	}
}

// minimum function using recursion
func (list *SinglyLinkedList) MinUsingRecursion() (int, error) {
	if list.root == nil {
		return 0, errors.New("linkedlist is empty")
	} else {
		return list.minUsingRecursion(list.root), nil
	}
}

// minimum function using recursion, utils
func (list *SinglyLinkedList) minUsingRecursion(node *Node) int {
	if node.Next == nil {
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
func (list *SinglyLinkedList) LinearSearchUsingLoop(key int) (*Node, error) {
	if tmp := list.root; tmp != nil {
		for tmp != nil {
			if tmp.Data == key {
				return tmp, nil
			} else {
				tmp = tmp.Next
			}
		}
		return nil, errors.New("no element found in linkedlist")
	} else {
		return nil, errors.New("linkedlist is empty")
	}
}

// linearSearch function using Recursion
func (list *SinglyLinkedList) LinearSearchUsingRecursion(key int) (*Node, error) {
	if list.root == nil {
		return nil, errors.New("linkedlist is empty")
	} else {
		return list.linearSearchUsingRecursion(list.root, key)
	}
}

// linearSearch function using Recursion, utils
func (list *SinglyLinkedList) linearSearchUsingRecursion(node *Node, key int) (*Node, error) {
	if node == nil {
		return nil, errors.New("no element found in linkedlist")
	} else {
		if node.Data == key {
			return node, nil
		}
		return list.linearSearchUsingRecursion(node.Next, key)
	}
}

/*
	insert Function
*/

// insert function using loop
func (list *SinglyLinkedList) InsertUsingLoop(index int, key int) error {
	if index < 0 {
		return errors.New("index is invalid")
	}

	if tmp := list.root; index == 0 {
		list.root = &Node{key, nil}
		list.root.Next = tmp
		return nil
	} else if tmp != nil {
		for index != 1 && tmp != nil {
			index--
			tmp = tmp.Next
		}
		if tmp == nil {
			return errors.New("index is invalid")
		} else {
			temp := tmp.Next
			tmp.Next = &Node{key, nil}
			tmp.Next.Next = temp
			return nil
		}
	} else {
		return errors.New("index is invalid")
	}
}

// insert function using recursion
func (list *SinglyLinkedList) InsertUsingRecursion(index int, key int) error {
	if index < 0 {
		return errors.New("index is invalid")
	} else if index == 0 {
		tmp := list.root
		list.root = &Node{key, nil}
		list.root.Next = tmp
		return nil
	} else {
		return list.insertUsingRecursion(list.root, index, key)
	}
}

// insert function using recursion, utils
func (list *SinglyLinkedList) insertUsingRecursion(node *Node, index int, key int) error {
	if node == nil {
		return errors.New("index is invalid")
	} else {
		if index == 1 {
			tmp := node.Next
			node.Next = &Node{key, nil}
			node.Next.Next = tmp
			return nil
		}
		return list.insertUsingRecursion(node.Next, index-1, key)
	}
}

/*
	addFirst/addLast Function
*/
// addFirst function using loop
func (list *SinglyLinkedList) AddFirst(key int) {
	tmp := &Node{key, nil}
	if list.root == nil {
		list.root = tmp
	} else {
		tmp.Next = list.root
		list.root = tmp
	}
}

// addLast function using loop
func (list *SinglyLinkedList) AddLastUsingLoop(key int) {
	tmp := &Node{key, nil}
	if list.root == nil {
		list.root = tmp
	} else {
		temp := list.root
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = tmp
	}
}

// addLast function using Recursion
func (list *SinglyLinkedList) AddLastUsingRecursion(key int) {
	list.root = list.addLastUsingRecursion(list.root, key)
}

// addLast function using Recursion, utils
func (list *SinglyLinkedList) addLastUsingRecursion(node *Node, key int) *Node {
	if node == nil {
		return &Node{key, nil}
	} else {
		node.Next = list.addLastUsingRecursion(node.Next, key)
		return node
	}
}
