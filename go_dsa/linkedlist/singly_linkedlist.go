package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList struct {
	head *Node
}

func (list *SinglyLinkedList) Create(values ...int) {
	var last *Node
	for _, val := range values {
		if list.head == nil {
			list.head = &Node{val, nil}
			last = list.head
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
	tmp := list.head
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
	tmp := list.head
	var last *Node
	for tmp != last {
		for tmp.Next != last {
			tmp = tmp.Next
		}
		last = tmp
		if tmp == list.head {
			fmt.Println(tmp.Data)
		} else {
			fmt.Print(tmp.Data, sep)
			tmp = list.head
		}
	}
}

// Display function using recursion
func (list *SinglyLinkedList) DisplayUsingRecursion(sep string) {
	list.displayUsingRecursion(sep, list.head)
}
func (list *SinglyLinkedList) DisplayUsingRecursionInReverseOrder(sep string) {
	list.displayUsingRecursionInReverseOrder(sep, list.head)
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
		if node == list.head {
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
	tmp := list.head
	for tmp != nil {
		tmp = tmp.Next
		noOfNodes++
	}
	return noOfNodes
}

// count function using recursion
func (list *SinglyLinkedList) CountUsingRecursion() int {
	return list.countUsingRecursion(list.head)
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
	tmp := list.head
	for tmp != nil {
		sum += tmp.Data
		tmp = tmp.Next
	}
	return sum
}

// sum function using recursion
func (list *SinglyLinkedList) SumUsingRecursion() int {
	return list.sumUsingRecursion(list.head)
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
	tmp := list.head
	for tmp != nil {
		sum += tmp.Data
		tmp = tmp.Next
	}
	return float64(sum / list.CountUsingLoop())
}

// average function using recursion
func (list *SinglyLinkedList) AverageUsingRecursion() float64 {
	return float64(list.averageUsingRecursion(list.head) / list.CountUsingLoop())
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

func (list *SinglyLinkedList) Max() int {
	max := -1
	if tmp := list.head; tmp != nil {
		max = tmp.Data
		tmp = tmp.Next
		for tmp != nil {
			if max < tmp.Data {
				max = tmp.Data
			}
			tmp = tmp.Next
		}
	}
	return max
}
