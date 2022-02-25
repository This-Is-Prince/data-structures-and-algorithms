package using_array

import (
	"errors"
	"fmt"
)

type queue struct {
	data              []int
	front, rear, size int
}

func Create(size int) (*queue, error) {
	if size <= 0 {
		return nil, errors.New("size can't be negative or zero")
	} else {
		q := &queue{make([]int, size), 0, 0, size}
		return q, nil
	}
}

func (q *queue) EnQueue(value int) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.rear = (q.rear + 1) % q.size
	q.data[q.rear] = value
	return nil
}

func (q *queue) DeQueue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	q.front = (q.front + 1) % q.size
	value := q.data[q.front]
	q.data[q.front] = 0
	return value, nil
}

func (q *queue) IsEmpty() bool {
	if q.front == q.rear {
		return true
	} else {
		return false
	}
}

func (q *queue) IsFull() bool {
	if (q.rear+1)%q.size == q.front {
		return true
	} else {
		return false
	}
}

func (q *queue) First() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	} else {
		return q.data[(q.front+1)%q.size], nil
	}
}

func (q *queue) Last() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	} else {
		return q.data[q.rear], nil
	}
}

func (q *queue) Display(sep string) {
	if !q.IsEmpty() {
		index := (q.front + 1) % q.size
		for index != q.rear {
			fmt.Print(q.data[index], sep)
			index = (index + 1) % q.size
		}
		fmt.Println(q.data[index])
	}
}
