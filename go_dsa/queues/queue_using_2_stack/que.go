package queue_using_2_stack

import (
	"errors"
	"fmt"
)

type Queue struct {
	s1 *stack
	s2 *stack
}

func (q *Queue) EnQueue(value int) error {
	if q.s1 == nil {
		q.s1 = &stack{}
		q.s2 = &stack{}
	}
	if q.s1.Push(value) != nil {
		return errors.New("queue is full")
	}
	return nil
}

func (q *Queue) DeQueue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	} else {
		if q.s2.IsEmpty() {
			for !q.s1.IsEmpty() {
				value, _ := q.s1.Pop()
				q.s2.Push(value)
			}
		}
		value, _ := q.s2.Pop()
		return value, nil
	}
}

func (q *Queue) IsEmpty() bool {
	return q.s1 == nil || (q.s1.IsEmpty() && q.s2.IsEmpty())
}

func (q *Queue) Display(sep string) error {
	if q.IsEmpty() {
		return errors.New("queue is empty")
	}
	err := q.s2.Display(sep)

	if !q.s1.IsEmpty() {
		if err == nil {
			fmt.Print(sep)
		}
		q.s1.DisplayReverse(sep)
	}
	fmt.Println()
	return nil

}
