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

/*
	insert Function
*/

// insert function using loop
func (list *LinkedList) InsertUsingLoop(index int, key int) error {
	if index < 0 || (list.root == nil && index > 0) {
		return errors.New("index is invalid")
	}
	if index == 0 {
		list.AddFirst(key)
	} else {
		node := list.root
		newNode := &Node{key, nil}
		for i := 0; i < index-1; i++ {
			node = node.Next
			if node == list.root {
				return errors.New("index is invalid")
			}
		}
		newNode.Next = node.Next
		node.Next = newNode
	}
	return nil
}

// insert function using recursion
func (list *LinkedList) InsertUsingRecursion(index int, key int) error {
	if index < 0 || (list.root == nil && index > 0) {
		return errors.New("index is invalid")
	} else if index == 0 {
		list.AddFirst(key)
		return nil
	} else {
		return list.insertUsingRecursion(list.root, index, key)
	}
}

// insert function using recursion, utils
func (list *LinkedList) insertUsingRecursion(node *Node, index int, key int) error {
	if index == 1 {
		newNode := &Node{key, nil}
		newNode.Next = node.Next
		node.Next = newNode
		return nil
	} else if node.Next == list.root {
		return errors.New("index is invalid")
	} else {
		return list.insertUsingRecursion(node.Next, index-1, key)
	}
}

/*
	addFirst/addLast Function
*/
// addFirst function using loop
func (list *LinkedList) AddFirst(key int) {
	newNode := &Node{key, nil}
	if list.root == nil {
		newNode.Next = newNode
	} else {
		node := list.root
		for node.Next != list.root {
			node = node.Next
		}
		node.Next = newNode
		newNode.Next = list.root
	}
	list.root = newNode
}

// addLast function using loop
func (list *LinkedList) AddLastUsingLoop(key int) {
	newNode := &Node{key, nil}
	if list.root == nil {
		newNode.Next = newNode
		list.root = newNode
	} else {
		node := list.root
		for node.Next != list.root {
			node = node.Next
		}
		newNode.Next = node.Next
		node.Next = newNode
	}
}

// addLast function using Recursion
func (list *LinkedList) AddLastUsingRecursion(key int) {
	if list.root == nil {
		newNode := &Node{key, nil}
		newNode.Next = newNode
		list.root = newNode
	} else {
		list.addLastUsingRecursion(list.root, key)
	}
}

// addLast function using Recursion, utils
func (list *LinkedList) addLastUsingRecursion(node *Node, key int) {
	if node.Next == list.root {
		newNode := &Node{key, nil}
		newNode.Next = node.Next
		node.Next = newNode
	} else {
		list.addLastUsingRecursion(node.Next, key)
	}
}

/*
	insert in sorted linkedlist Function
*/

// insertInSortedList function using Loop
func (list *LinkedList) InsertInSortedListUsingLoop(key int) {
	newNode := &Node{key, nil}
	if list.root == nil {
		newNode.Next = newNode
		list.root = newNode
	} else if list.root.Data >= key {
		list.AddFirst(key)
	} else {
		node := list.root
		for node.Next != list.root && node.Next.Data < key {
			node = node.Next
		}
		newNode.Next = node.Next
		node.Next = newNode
	}
}

// insertInSortedList function using Recursion
func (list *LinkedList) InsertInSortedListUsingRecursion(key int) {
	if list.root == nil {
		newNode := &Node{key, nil}
		newNode.Next = newNode
		list.root = newNode
	} else if list.root.Data >= key {
		list.AddFirst(key)
	} else {
		list.insertInSortedListUsingRecursion(list.root, key)
	}
}

// insertInSortedList function using Recursion, utils
func (list *LinkedList) insertInSortedListUsingRecursion(node *Node, key int) {
	if node.Next == list.root || node.Next.Data >= key {
		newNode := &Node{key, nil}
		newNode.Next = node.Next
		node.Next = newNode
	} else {
		list.insertInSortedListUsingRecursion(node.Next, key)
	}
}

/*
	deleting node Function
*/
// deleteFirst function using Loop
func (list *LinkedList) DeleteFirst() error {
	if node := list.root; node != nil {
		if node == node.Next {
			list.root = nil
		} else {
			for node.Next != list.root {
				node = node.Next
			}
			root := list.root
			node.Next = root.Next
			list.root = root.Next
			root.Next = nil
		}
		return nil
	} else {
		return errors.New("list is empty")
	}
}

// deleteLast function using Loop
func (list *LinkedList) DeleteLastUsingLoop() error {
	if node := list.root; node != nil {
		var last *Node
		for node.Next != list.root {
			last = node
			node = node.Next
		}
		if last == nil {
			list.root = nil
		} else {
			last.Next = list.root
		}
		return nil
	} else {
		return errors.New("list is empty")
	}
}

// deleteLast function using Recursion
func (list *LinkedList) DeleteLastUsingRecursion() error {
	if list.root == nil {
		return errors.New("list is empty")
	}
	if list.root == list.root.Next {
		list.root = nil
	} else {
		list.root = list.deleteLastUsingRecursion(list.root)
	}
	return nil
}

// deleteLast function using Recursion, utils
func (list *LinkedList) deleteLastUsingRecursion(node *Node) *Node {
	if node.Next == list.root {
		return list.root
	} else {
		node.Next = list.deleteLastUsingRecursion(node.Next)
		return node
	}
}

// delete function using Loop
func (list *LinkedList) DeleteUsingLoop(index int) error {
	if index < 0 {
		return errors.New("index is invalid")
	}
	if index == 0 {
		return list.DeleteFirst()
	}
	if node := list.root; node == nil {
		return errors.New("list is empty")
	} else {
		var last *Node
		for i := 0; i < index; i++ {
			last = node
			node = node.Next
			if node == list.root {
				return errors.New("index is invalid")
			}
		}
		last.Next = node.Next
		return nil
	}
}

// delete function using Recursion
func (list *LinkedList) DeleteUsingRecursion(index int) error {
	if index < 0 {
		return errors.New("index is invalid")
	}
	if index == 0 {
		return list.DeleteFirst()
	}
	if list.root == nil {
		return errors.New("list is empty")
	} else {
		return list.deleteUsingRecursion(list.root, index)
	}
}

// delete function using Recursion, utils
func (list *LinkedList) deleteUsingRecursion(node *Node, index int) error {
	if node.Next == list.root {
		return errors.New("index is invalid")
	} else if index == 1 {
		node.Next = node.Next.Next
	} else {
		list.deleteUsingRecursion(node.Next, index-1)
	}
	return nil
}

/*
	isSorted Function
*/
// sorted function using loop
func (list *LinkedList) IsSortedUsingLoop() (bool, error) {
	isSorted := true
	if node := list.root; node == nil {
		return !isSorted, errors.New("list is empty")
	} else {
		for node.Next != list.root {
			if node.Data > node.Next.Data {
				isSorted = false
				break
			}
			node = node.Next
		}
		return isSorted, nil
	}
}

// sorted function using Recursion
func (list *LinkedList) IsSortedUsingRecursion() (bool, error) {
	if list.root == nil {
		return false, errors.New("list is empty")
	} else {
		return list.isSortedUsingRecursion(list.root)
	}
}

// sorted function using Recursion, utils
func (list *LinkedList) isSortedUsingRecursion(node *Node) (bool, error) {
	if node.Next == list.root {
		return true, nil
	} else {
		if node.Data < node.Next.Data {
			return list.isSortedUsingRecursion(node.Next)
		} else {
			return false, nil
		}
	}
}

/*
	remove duplicate from sorted linkedlist Function
*/
// remove duplicate from sorted linkedlist function using loop
func (list *LinkedList) RemoveDuplicateFromSortedListUsingLoop() error {
	if node := list.root; node == nil {
		return errors.New("list is empty")
	} else {
		for node.Next != list.root {
			if node.Data == node.Next.Data {
				node.Next = node.Next.Next
			} else {
				node = node.Next
			}
		}
		return nil
	}
}

// remove duplicate from sorted linkedlist function using Recursion
func (list *LinkedList) RemoveDuplicateFromSortedListUsingRecursion() error {
	if list.root == nil {
		return errors.New("list is empty")
	} else {
		list.root = list.removeDuplicateFromSortedListUsingRecursion(list.root)
		return nil
	}
}

// remove duplicate from sorted linkedlist function using Recursion, utils
func (list *LinkedList) removeDuplicateFromSortedListUsingRecursion(node *Node) *Node {
	if node.Next == list.root {
		return node
	} else {
		tmp := list.removeDuplicateFromSortedListUsingRecursion(node.Next)
		if tmp.Data == node.Data {
			return tmp
		} else {
			node.Next = tmp
			return node
		}
	}
}
