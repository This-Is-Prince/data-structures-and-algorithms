package singly

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

func (list *LinkedList) GetRoot() *Node {
	return list.root
}

/*
All Display functions
*/

// Display function using loops
func (list *LinkedList) DisplayUsingLoop(sep string) {
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

func (list *LinkedList) DisplayUsingLoopInReverseOrder(sep string) {
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
func (list *LinkedList) DisplayUsingRecursion(sep string) {
	list.displayUsingRecursion(sep, list.root)
}
func (list *LinkedList) DisplayUsingRecursionInReverseOrder(sep string) {
	list.displayUsingRecursionInReverseOrder(sep, list.root)
}

// Display function using recursions, utility function
func (list *LinkedList) displayUsingRecursion(sep string, node *Node) {
	if node == nil {
		return
	} else if node.Next == nil {
		fmt.Println(node.Data)
	} else {
		fmt.Print(node.Data, sep)
		list.displayUsingRecursion(sep, node.Next)
	}
}
func (list *LinkedList) displayUsingRecursionInReverseOrder(sep string, node *Node) {
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
func (list *LinkedList) CountUsingLoop() int {
	noOfNodes := 0
	tmp := list.root
	for tmp != nil {
		tmp = tmp.Next
		noOfNodes++
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
	tmp := list.root
	for tmp != nil {
		sum += tmp.Data
		tmp = tmp.Next
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
func (list *LinkedList) MaxUsingRecursion() (int, error) {
	if list.root == nil {
		return -1, errors.New("linkedlist is empty")
	} else {
		return list.maxUsingRecursion(list.root), nil
	}
}

// maximum function using recursion, utils
func (list *LinkedList) maxUsingRecursion(node *Node) int {
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
func (list *LinkedList) MinUsingLoop() (int, error) {
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
func (list *LinkedList) MinUsingRecursion() (int, error) {
	if list.root == nil {
		return 0, errors.New("linkedlist is empty")
	} else {
		return list.minUsingRecursion(list.root), nil
	}
}

// minimum function using recursion, utils
func (list *LinkedList) minUsingRecursion(node *Node) int {
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
func (list *LinkedList) LinearSearchUsingLoop(key int) (*Node, error) {
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
func (list *LinkedList) LinearSearchUsingRecursion(key int) (*Node, error) {
	if list.root == nil {
		return nil, errors.New("linkedlist is empty")
	} else {
		return list.linearSearchUsingRecursion(list.root, key)
	}
}

// linearSearch function using Recursion, utils
func (list *LinkedList) linearSearchUsingRecursion(node *Node, key int) (*Node, error) {
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
func (list *LinkedList) InsertUsingLoop(index int, key int) error {
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
func (list *LinkedList) InsertUsingRecursion(index int, key int) error {
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
func (list *LinkedList) insertUsingRecursion(node *Node, index int, key int) error {
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
func (list *LinkedList) AddFirst(key int) {
	tmp := &Node{key, nil}
	if list.root == nil {
		list.root = tmp
	} else {
		tmp.Next = list.root
		list.root = tmp
	}
}

// addLast function using loop
func (list *LinkedList) AddLastUsingLoop(key int) {
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
func (list *LinkedList) AddLastUsingRecursion(key int) {
	list.root = list.addLastUsingRecursion(list.root, key)
}

// addLast function using Recursion, utils
func (list *LinkedList) addLastUsingRecursion(node *Node, key int) *Node {
	if node == nil {
		return &Node{key, nil}
	} else {
		node.Next = list.addLastUsingRecursion(node.Next, key)
		return node
	}
}

/*
	insert in sorted linkedlist Function
*/

// insertInSortedList function using Loop
func (list *LinkedList) InsertInSortedListUsingLoop(key int) {
	newNode := &Node{key, nil}
	if list.root == nil {
		list.root = newNode
	} else if list.root.Data >= key {
		newNode.Next = list.root
		list.root = newNode
	} else {
		temp := list.root
		for temp.Next != nil && temp.Next.Data < key {
			temp = temp.Next
		}
		newNode.Next = temp.Next
		temp.Next = newNode
	}
}

// insertInSortedList function using Recursion
func (list *LinkedList) InsertInSortedListUsingRecursion(key int) {
	list.root = list.insertInSortedListUsingRecursion(list.root, key)
}

// insertInSortedList function using Recursion, utils
func (list *LinkedList) insertInSortedListUsingRecursion(node *Node, key int) *Node {
	if list.root == nil {
		return &Node{key, nil}
	} else {
		if node.Data < key {
			node.Next = list.insertInSortedListUsingRecursion(node.Next, key)
			return node
		} else {
			newNode := &Node{key, nil}
			newNode.Next = node
			return newNode
		}
	}
}

/*
	deleting node Function
*/
// deleteFirst function using Loop
func (list *LinkedList) DeleteFirst() error {
	if list.root != nil {
		list.root = list.root.Next
		return nil
	} else {
		return errors.New("list is empty")
	}
}

// deleteLast function using Loop
func (list *LinkedList) DeleteLastUsingLoop() error {
	if tmp := list.root; tmp != nil {
		var last *Node
		for tmp.Next != nil {
			last = tmp
			tmp = tmp.Next
		}
		if last == nil {
			list.root = nil
		} else {
			last.Next = nil
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
	} else {
		list.root = list.deleteLastUsingRecursion(list.root)
		return nil
	}
}

// deleteLast function using Recursion, utils
func (list *LinkedList) deleteLastUsingRecursion(node *Node) *Node {
	if node.Next == nil {
		return nil
	} else {
		node.Next = list.deleteLastUsingRecursion(node.Next)
		return node
	}
}

// delete function using Loop
func (list *LinkedList) DeleteUsingLoop(index int) error {
	if index < 0 {
		return errors.New("index is invalid")
	} else if tmp := list.root; tmp != nil {
		if index == 0 {
			list.root = list.root.Next
			return nil
		} else {
			for tmp.Next != nil && index != 1 {
				tmp = tmp.Next
				index--
			}
			if tmp.Next == nil {
				return errors.New("index is invalid")
			} else {
				tmp.Next = tmp.Next.Next
				return nil
			}
		}
	} else {
		return errors.New("list is empty")
	}
}

// delete function using Recursion
func (list *LinkedList) DeleteUsingRecursion(index int) error {
	if index < 0 {
		return errors.New("index is invalid")
	}
	if list.root == nil {
		return errors.New("list is empty")
	} else {
		node, err := list.deleteUsingRecursion(list.root, index)
		list.root = node
		return err
	}
}

// delete function using Recursion, utils
func (list *LinkedList) deleteUsingRecursion(node *Node, index int) (*Node, error) {
	if node == nil {
		return nil, errors.New("index is invalid")
	} else if index == 0 {
		return node.Next, nil
	} else {
		temp, err := list.deleteUsingRecursion(node.Next, index-1)
		node.Next = temp
		return node, err
	}
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
		for node.Next != nil {
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
	if node.Next == nil {
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
		for node.Next != nil {
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
	if node.Next == nil {
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

/*
	reversing linkedlist Function
*/
// reversing linkedlist function using loop
func (list *LinkedList) ReverseListUsingLoop() error {
	if curr := list.root; curr == nil {
		return errors.New("list is empty")
	} else {
		var prev, next *Node
		for curr != nil {
			next = curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		list.root = prev
		return nil
	}
}

// reversing linkedlist function using Recursion
func (list *LinkedList) ReverseListUsingRecursion() error {
	if list.root == nil {
		return errors.New("list is empty")
	} else {
		list.reverseListUsingRecursion(list.root)
		return nil
	}
}

// reversing linkedlist function using Recursion, utils
func (list *LinkedList) reverseListUsingRecursion(node *Node) *Node {
	if node.Next == nil {
		list.root = node
		return node
	} else {
		tmp := list.reverseListUsingRecursion(node.Next)
		node.Next = nil
		tmp.Next = node
		return node
	}
}

/*
	concating linkedlist Function
*/
// concating linkedlist function using loop
func (list *LinkedList) ConcatUsingLoop(otherList *LinkedList) {
	if node := list.root; node == nil {
		list.root = otherList.root
	} else {
		for node.Next != nil {
			node = node.Next
		}
		node.Next = otherList.root
	}
	otherList.root = nil
}

// concating linkedlist function using Recursion
func (list *LinkedList) ConcatUsingRecursion(otherList *LinkedList) {
	list.root = list.concatUsingRecursion(otherList, list.root)
	otherList.root = nil
}

// concating linkedlist function using Recursion, utils
func (list *LinkedList) concatUsingRecursion(otherList *LinkedList, node *Node) *Node {
	if node == nil {
		return otherList.root
	} else {
		node.Next = list.concatUsingRecursion(otherList, node.Next)
		return node
	}
}

/*
	merging linkedlist Function
*/
// merging linkedlist function using loop
func (list *LinkedList) MergeUsingLoop(otherList *LinkedList) {
	second := otherList.root
	if first := list.root; first == nil {
		list.root = second
	} else if second != nil {
		var last *Node
		if first.Data <= second.Data {
			list.root = first
			last = first
			first = first.Next
		} else {
			list.root = second
			last = second
			second = second.Next
		}
		for first != nil && second != nil {
			if first.Data <= second.Data {
				last.Next = first
				last = first
				first = first.Next
			} else {
				last.Next = second
				last = second
				second = second.Next
			}
		}
		if first == nil {
			last.Next = second
		} else {
			last.Next = first
		}
	}
	otherList.root = nil
}

// merging linkedlist function using Recursion
func (list *LinkedList) MergeUsingRecursion(otherList *LinkedList) {
	if list.root == nil {
		list.root = otherList.root
	} else if otherList.root != nil {
		list.root = list.mergeUsingRecursion(list.root, otherList.root)
	}
	otherList.root = nil
}

// merging linkedlist function using Recursion, utils
func (list *LinkedList) mergeUsingRecursion(first *Node, second *Node) *Node {
	if first == nil {
		return second
	} else if second == nil {
		return first
	} else if first.Data <= second.Data {
		first.Next = list.mergeUsingRecursion(first.Next, second)
		return first
	} else {
		second.Next = list.mergeUsingRecursion(first, second.Next)
		return second
	}
}

/*
	check loop in linkedlist Function
*/
// check loop in linkedlist Function using loop
func (list *LinkedList) CheckLoopUsingLoop() (bool, error) {
	if list.root == nil {
		return false, errors.New("list is empty")
	} else {
		p := list.root
		q := list.root
		for q != nil && q.Next != nil {
			p = p.Next
			q = q.Next.Next
			if p == q {
				return true, nil
			}
		}
		return false, nil
	}
}

// check loop in linkedlist Function using Recursion
func (list *LinkedList) CheckLoopUsingRecursion() (bool, error) {
	if list.root == nil {
		return false, errors.New("list is empty")
	} else {
		return list.checkLoopUsingRecursion(list.root, list.root)
	}
}

// check loop in linkedlist Function using Recursion, utils
func (list *LinkedList) checkLoopUsingRecursion(first *Node, second *Node) (bool, error) {
	if first == nil || second == nil || second.Next == nil {
		return false, nil
	} else {
		first = first.Next
		second = second.Next.Next
		if first == second {
			return true, nil
		}
		return list.checkLoopUsingRecursion(first, second)
	}
}
