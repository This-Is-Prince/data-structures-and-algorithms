package binary_tree

import "errors"

type QNode struct {
	data       *Node
	prev, next *QNode
}

type Queue struct {
	front, rear *QNode
}

func (q *Queue) EnQueue(value *Node) error {
	if newNode := (&QNode{data: value}); newNode == nil {
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

func (q *Queue) DeQueue() (*Node, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	} else {
		value := q.front
		q.front = q.front.next
		if q.front != nil {
			q.front.prev = nil
		} else {
			q.rear = nil
		}
		value.next = nil
		return value.data, nil
	}
}

func (q *Queue) First() (*Node, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	} else {
		return q.front.data, nil
	}
}

func (q *Queue) Last() (*Node, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	} else {
		return q.rear.data, nil
	}
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}
