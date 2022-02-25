package using_linkedlist

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
}

type Queue struct {
	front *node
	rear  *node
}

func (q *Queue) EnQueue(value int) error {
	if newNode := (&node{data: value}); newNode == nil {
		return errors.New("queue is full")
	} else {
		if q.front == nil {
			q.front, q.rear = newNode, newNode
		} else {
			q.rear.next = newNode
			q.rear = newNode
		}
		q.rear.next = q.front
		return nil
	}
}

func (q *Queue) DeQueue() (int, error) {
	if q.front == nil {
		return 0, errors.New("queue is empty")
	} else {
		value := q.front
		if q.front == q.front.next {
			q.front = nil
			q.rear = nil
		} else {
			q.front = q.front.next
			q.rear.next = q.front
		}
		value.next = nil
		return value.data, nil
	}
}

func (q *Queue) IsEmpty() bool {
	if q.front == nil {
		return true
	} else {
		return false
	}
}

func (q *Queue) IsFull() bool {
	if tmp := (&node{}); tmp == nil {
		return true
	} else {
		return false
	}
}

func (q *Queue) First() (int, error) {
	if q.front == nil {
		return 0, errors.New("queue is empty")
	} else {
		return q.front.data, nil
	}
}

func (q *Queue) Last() (int, error) {
	if q.rear == nil {
		return 0, errors.New("queue is empty")
	} else {
		return q.rear.data, nil
	}
}

func (q *Queue) Display(sep string) {
	if tmp := q.front; tmp != nil {
		for tmp.next != q.front {
			fmt.Print(tmp.data, sep)
			tmp = tmp.next
		}
		fmt.Println(tmp.data)
	}
}
