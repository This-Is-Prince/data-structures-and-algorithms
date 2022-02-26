package binary_tree

import "errors"

type SNode struct {
	data *Node
	next *SNode
}

type Stack struct {
	top *SNode
}

func (stk *Stack) Push(value *Node) error {
	if newNode := (&SNode{data: value}); newNode == nil {
		return errors.New("overflow")
	} else {
		newNode.next = stk.top
		stk.top = newNode
	}
	return nil
}

func (stk *Stack) Pop() (*Node, error) {
	if stk.IsEmpty() {
		return nil, errors.New("underflow")
	} else {
		value := stk.top
		stk.top = stk.top.next
		value.next = nil
		return value.data, nil
	}
}

func (stk *Stack) Peek() (*Node, error) {
	if stk.IsEmpty() {
		return nil, errors.New("underflow")
	} else {
		return stk.top.data, nil
	}
}

func (stk *Stack) IsEmpty() bool {
	return stk.top == nil
}
