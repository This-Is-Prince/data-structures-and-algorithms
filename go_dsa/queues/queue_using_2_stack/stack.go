package queue_using_2_stack

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
}

type stack struct {
	top *node
}

func (stk *stack) Create(values ...int) error {
	for _, value := range values {
		newNode := &node{data: value}
		if newNode == nil {
			return errors.New("overflow")
		}
		if stk.top != nil {
			newNode.next = stk.top
		}
		stk.top = newNode
	}
	return nil
}

func (stk *stack) Push(value int) error {
	newNode := &node{data: value}
	if newNode == nil {
		return errors.New("overflow")
	}
	if stk.top != nil {
		newNode.next = stk.top
	}
	stk.top = newNode
	return nil
}

func (stk *stack) Pop() (int, error) {
	if stk.IsEmpty() {
		return 0, errors.New("underflow")
	} else {
		value := stk.top
		stk.top = stk.top.next
		value.next = nil
		return value.data, nil
	}
}

func (stk *stack) IsEmpty() bool {
	return stk.top == nil
}

func (stk *stack) Display(sep string) error {
	if stk.IsEmpty() {
		return errors.New("underflow")
	}
	tmp := stk.top
	for tmp.next != nil {
		fmt.Print(tmp.data, sep)
		tmp = tmp.next
	}
	fmt.Print(tmp.data)
	return nil
}

func (stk *stack) DisplayReverse(sep string) error {
	if stk.IsEmpty() {
		return errors.New("underflow")
	} else {
		stk.displayReverse(stk.top.next, sep)
		fmt.Print(stk.top.data)
	}
	return nil
}

func (stk *stack) displayReverse(top *node, sep string) {
	if top != nil {
		stk.displayReverse(top.next, sep)
		fmt.Print(top.data, sep)
	}
}
