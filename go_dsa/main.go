package main

import (
	"fmt"

	circular_queue_using_array "go_dsa/queues/circular/using_array"
	circular_queue_using_linkedlist "go_dsa/queues/circular/using_linkedlist"
	dequeue_using_array "go_dsa/queues/double_ended/using_array"
	dequeue_using_linkedlist "go_dsa/queues/double_ended/using_linkedlist"
	linear_queue_using_array "go_dsa/queues/linear/using_array"
	linear_queue_using_linkedlist "go_dsa/queues/linear/using_linkedlist"
	queue_using_2_stack "go_dsa/queues/queue_using_2_stack"
	stk_questions "go_dsa/stack/questions/parenthesis_matching"
	stk_using_array "go_dsa/stack/using_array"
	stk_using_linkedlist "go_dsa/stack/using_linkedlist"
	"go_dsa/trees/binary_tree"
)

func main() {
	fmt.Println("Data Structures and Algorithms")

	// debugStackUsingArray()
	// debugStackUsingLinkedList()
	// debugStackQuestions()
	// debugLinearQueueUsingArray()
	// debugLinearQueueUsingLinkedlist()
	// fmt.Println("Circular queue")
	// debugCircularQueueUsingArray()
	// debugCircularQueueUsingLinkedlist()
	// debugDoubleEndedQueueUsingArray()
	// debugDoubleEndedQueueUsingLinkedlist()
	// debugQueueUsing2Stack()
	debugBinaryTree()
}

func debugStackUsingArray() {
	stk, err := stk_using_array.Create(10)
	if err != nil {
		fmt.Println(err)
	} else {
		stk.Display(",")
		stk.Push(1)
		stk.Push(2)
		stk.Push(3)
		stk.Push(4)
		stk.Push(5)
		stk.Display(",")
	}
}

func debugStackUsingLinkedList() {

	stk1 := stk_using_linkedlist.Stack{}
	stk1.Display(",")
	stk1.Push(1)
	stk1.Push(2)
	stk1.Push(3)
	stk1.Push(4)
	stk1.Push(5)
	stk1.Display(",")
}

func debugStackQuestions() {
	fmt.Println("Is '((a+b)*(c-d))' Balance", stk_questions.IsBalance("((a+b)*(c-d))"))
	fmt.Println("Is '((a+b)*(c-d)' Balance", stk_questions.IsBalance("((a+b)*(c-d)"))
	fmt.Println("Is '((a+b)*(c-d)))' Balance", stk_questions.IsBalance("((a+b)*(c-d)))"))
	fmt.Println("Is '{([a+b]*[c-d])/e}' Balance", stk_questions.IsBalance("{([a+b]*[c-d])/e}"))
	fmt.Println("Is '{([a+b]*[c-d]/e}' Balance", stk_questions.IsBalance("{([a+b]*[c-d]/e}"))
	fmt.Println("Is '{([a+b]*[c-d)/e}' Balance", stk_questions.IsBalance("{([a+b]*[c-d)/e}"))
	fmt.Println("Is '{([a+b)*[c-d]]/e}' Balance", stk_questions.IsBalance("{([a+b)*[c-d]]/e}"))
}

func debugLinearQueueUsingArray() {
	q, err := linear_queue_using_array.Create(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	q.Display(",")
	fmt.Println(q.EnQueue(10))
	fmt.Println(q.EnQueue(20))
	fmt.Println(q.EnQueue(30))
	fmt.Println(q.EnQueue(40))
	fmt.Println(q.EnQueue(50))
	fmt.Println(q.EnQueue(60))
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.EnQueue(50))
	fmt.Println(q.EnQueue(60))
	q.Display(",")
}

func debugLinearQueueUsingLinkedlist() {
	q := linear_queue_using_linkedlist.Queue{}
	fmt.Println(q.EnQueue(10))
	q.Display(",")
	fmt.Println(q.EnQueue(20))
	q.Display(",")
	fmt.Println(q.EnQueue(30))
	q.Display(",")
	fmt.Println(q.EnQueue(40))
	q.Display(",")
	fmt.Println(q.EnQueue(50))
	q.Display(",")

	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
}

func debugCircularQueueUsingArray() {
	q, err := circular_queue_using_array.Create(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	q.Display(",")
	fmt.Println(q.EnQueue(10))
	fmt.Println(q.EnQueue(20))
	fmt.Println(q.EnQueue(30))
	fmt.Println(q.EnQueue(40))
	fmt.Println(q.EnQueue(50))
	fmt.Println(q.EnQueue(60))
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.EnQueue(50))
	fmt.Println(q.EnQueue(60))
	q.Display(",")
}

func debugCircularQueueUsingLinkedlist() {
	q := circular_queue_using_linkedlist.Queue{}
	fmt.Println(q.EnQueue(10))
	q.Display(",")
	fmt.Println(q.EnQueue(20))
	q.Display(",")
	fmt.Println(q.EnQueue(30))
	q.Display(",")
	fmt.Println(q.EnQueue(40))
	q.Display(",")
	fmt.Println(q.EnQueue(50))
	q.Display(",")

	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
}

func debugDoubleEndedQueueUsingArray() {
	q, err := dequeue_using_array.Create(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(q.AddLast(10))
	q.Display(",")
	fmt.Println(q.AddLast(20))
	q.Display(",")
	fmt.Println(q.AddLast(30))
	q.Display(",")
	fmt.Println(q.AddLast(40))
	q.Display(",")
	fmt.Println(q.AddLast(50))
	q.Display(",")

	fmt.Println(q.RemoveLast())
	q.Display(",")
	fmt.Println(q.RemoveLast())
	q.Display(",")
	fmt.Println(q.RemoveLast())
	q.Display(",")
	fmt.Println(q.RemoveFirst())
	q.Display(",")

	fmt.Println(q.AddFirst(10))
	q.Display(",")
	fmt.Println(q.AddFirst(20))
	q.Display(",")
	fmt.Println(q.AddFirst(30))
	q.Display(",")
	fmt.Println(q.AddFirst(40))
	q.Display(",")
	fmt.Println(q.AddFirst(50))
	q.Display(",")
}

func debugDoubleEndedQueueUsingLinkedlist() {
	q := dequeue_using_linkedlist.Queue{}

	fmt.Println(q.AddLast(10))
	q.Display(",")
	fmt.Println(q.AddLast(20))
	q.Display(",")
	fmt.Println(q.AddLast(30))
	q.Display(",")
	fmt.Println(q.AddLast(40))
	q.Display(",")
	fmt.Println(q.AddLast(50))
	q.Display(",")

	fmt.Println(q.RemoveLast())
	q.Display(",")
	fmt.Println(q.RemoveLast())
	q.Display(",")
	fmt.Println(q.RemoveLast())
	q.Display(",")
	fmt.Println(q.RemoveFirst())
	q.Display(",")
	fmt.Println(q.RemoveLast())
	q.Display(",")

	fmt.Println(q.AddFirst(10))
	q.Display(",")
	fmt.Println(q.AddFirst(20))
	q.Display(",")
	fmt.Println(q.AddFirst(30))
	q.Display(",")
	fmt.Println(q.AddFirst(40))
	q.Display(",")
	fmt.Println(q.AddFirst(50))
	q.Display(",")
}

func debugQueueUsing2Stack() {
	q := queue_using_2_stack.Queue{}
	fmt.Println(q.EnQueue(6))
	q.Display(",")
	fmt.Println(q.EnQueue(3))
	q.Display(",")
	fmt.Println(q.EnQueue(9))
	q.Display(",")
	fmt.Println(q.EnQueue(5))
	q.Display(",")
	fmt.Println(q.EnQueue(4))
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.DeQueue())
	q.Display(",")
	fmt.Println(q.EnQueue(2))
	q.Display(",")
	fmt.Println(q.EnQueue(8))
	q.Display(",")
	fmt.Println(q.EnQueue(12))
	q.Display(",")
	fmt.Println(q.EnQueue(10))
	q.Display(",")
}

func debugBinaryTree() {
	binaryTree := binary_tree.BinaryTree{}
	binaryTree.Create([]int{5, 8, 6, -1, 9, 3, -1, 4, 2}, -1)
	binaryTree.PreOrderUsingRecursion(",")
	binaryTree.PreOrderUsingLoop(",")
	binaryTree.PostOrderUsingRecursion(",")
	binaryTree.InOrderUsingRecursion(",")
	binaryTree.InOrderUsingLoop(",")
	binaryTree.LevelOrderUsingLoop(",")
}
