package using_linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	top *Node
}

func (stk *Stack) Push(value int) error {
	newNode := &Node{Data: value}
	if newNode == nil {
		return errors.New("stack overflow")
	}
	if stk.top == nil {
		stk.top = newNode
	} else {
		newNode.Next = stk.top
		stk.top = newNode
	}
	return nil
}

func (stk *Stack) Pop() (int, error) {
	if stk.IsEmpty() {
		return 0, errors.New("stack underflow")
	} else {
		top := stk.top
		stk.top = stk.top.Next
		top.Next = nil
		return top.Data, nil
	}
}

func (stk *Stack) Peek() (int, error) {
	if stk.IsEmpty() {
		return 0, errors.New("stack underflow")
	} else {
		return stk.top.Data, nil
	}
}

func (stk *Stack) IsEmpty() bool {
	if stk.top == nil {
		return true
	} else {
		return false
	}
}

func (stk *Stack) Display(sep string) {
	top := stk.top
	for top != nil {
		if top.Next == nil {
			fmt.Println(top.Data)
		} else {
			fmt.Print(top.Data, sep)
		}
		top = top.Next
	}
}
