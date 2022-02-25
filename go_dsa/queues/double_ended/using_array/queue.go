package using_array

import (
	"errors"
	"fmt"
)

type dequeue struct {
	data              []int
	front, rear, size int
}

func Create(size int) (*dequeue, error) {
	if size <= 0 {
		return nil, errors.New("size can't be negative or zero")
	} else {
		q := &dequeue{make([]int, size), 0, 0, size}
		return q, nil
	}
}

func (q *dequeue) AddLast(value int) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.rear = (q.rear + 1) % q.size
	q.data[q.rear] = value
	return nil
}

func (q *dequeue) AddFirst(value int) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.data[q.front] = value
	if q.front == 0 {
		q.front = q.size - 1
	} else {
		q.front--
	}
	return nil
}

func (q *dequeue) RemoveFirst() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	} else {
		value := q.data[q.front]
		q.data[q.front] = 0
		q.front = (q.front + 1) % q.size
		return value, nil
	}
}

func (q *dequeue) RemoveLast() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	} else {
		value := q.data[q.rear]
		q.data[q.rear] = 0
		if q.rear == 0 {
			q.rear = q.size - 1
		} else {
			q.rear--
		}
		return value, nil
	}
}

func (q *dequeue) IsEmpty() bool {
	if q.rear == q.front {
		return true
	} else {
		return false
	}
}

func (q *dequeue) IsFull() bool {
	if (q.rear+1)%q.size == q.front {
		return true
	} else {
		return false
	}
}

func (q *dequeue) Display(sep string) {
	if !q.IsEmpty() {
		index := (q.front + 1) % q.size
		for index != q.rear {
			fmt.Print(q.data[index], sep)
			index = (index + 1) % q.size
		}
		fmt.Println(q.data[index])
	}
}
