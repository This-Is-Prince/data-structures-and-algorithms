package using_linkedlist

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
	prev *node
}

type Queue struct {
	front *node
	rear  *node
}

func (q *Queue) AddFirst(value int) error {
	if newNode := (&node{data: value}); newNode == nil {
		return errors.New("queue is full")
	} else {
		if q.front == nil {
			q.rear = newNode
		} else {
			newNode.next = q.front
			q.front.prev = newNode
		}
		q.front = newNode
		return nil
	}
}

func (q *Queue) AddLast(value int) error {
	if newNode := (&node{data: value}); newNode == nil {
		return errors.New("queue is full")
	} else {
		if q.front == nil {
			q.front = newNode
		} else {
			q.rear.next = newNode
			newNode.prev = q.rear
		}
		q.rear = newNode
		return nil
	}
}

func (q *Queue) RemoveFirst() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	} else {
		value := q.front
		q.front = q.front.next
		if q.front == nil {
			q.rear = nil
		} else {
			q.front.prev = nil
			value.next = nil
		}
		return value.data, nil
	}
}

func (q *Queue) RemoveLast() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	} else {
		value := q.rear
		q.rear = q.rear.prev
		if q.rear == nil {
			q.front = nil
		} else {
			q.rear.next = nil
			value.prev = nil
		}
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

func (q *Queue) Display(sep string) {
	if tmp := q.front; tmp != nil {
		for tmp.next != nil {
			fmt.Print(tmp.data, sep)
			tmp = tmp.next
		}
		fmt.Println(tmp.data)
	}
}
