package using_array

import (
	"errors"
	"fmt"
)

type stack struct {
	size int
	data []int
	top  int
}

func Create(size int) (*stack, error) {
	if size <= 0 {
		return nil, errors.New("size can't be negative or zero")
	}
	data := make([]int, size)
	return &stack{size, data, -1}, nil
}

func (stk *stack) Push(value int) error {
	if stk.IsFull() {
		return errors.New("stack overflow")
	} else {
		stk.top++
		stk.data[stk.top] = value
		return nil
	}
}

func (stk *stack) Pop() (int, error) {
	if stk.IsEmpty() {
		return 0, errors.New("stack underflow")
	} else {
		value := stk.data[stk.top]
		stk.data[stk.top] = 0
		stk.top--
		return value, nil
	}
}

func (stk *stack) Peek() (int, error) {
	if stk.IsEmpty() {
		return 0, errors.New("stack underflow")
	} else {
		return stk.data[stk.top], nil
	}
}

func (stk *stack) IsEmpty() bool {
	if stk.top == -1 {
		return true
	} else {
		return false
	}
}

func (stk *stack) IsFull() bool {
	if stk.top == stk.size-1 {
		return true
	} else {
		return false
	}
}

func (stk *stack) Display(sep string) {
	for _, value := range stk.data {
		fmt.Print(value, sep)
	}
	fmt.Println()
}
