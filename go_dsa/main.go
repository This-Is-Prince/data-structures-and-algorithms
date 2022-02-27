package main

import (
	"fmt"

	// Heap
	heap "go_dsa/heap"
	// Queue
	circular_queue_using_array "go_dsa/queues/circular/using_array"
	circular_queue_using_linkedlist "go_dsa/queues/circular/using_linkedlist"
	dequeue_using_array "go_dsa/queues/double_ended/using_array"
	dequeue_using_linkedlist "go_dsa/queues/double_ended/using_linkedlist"
	linear_queue_using_array "go_dsa/queues/linear/using_array"
	linear_queue_using_linkedlist "go_dsa/queues/linear/using_linkedlist"
	queue_using_2_stack "go_dsa/queues/queue_using_2_stack"

	// Sorting Techniques
	"go_dsa/sorting_techniques/bubble"
	"go_dsa/sorting_techniques/count"
	"go_dsa/sorting_techniques/insertion"
	"go_dsa/sorting_techniques/merge"
	"go_dsa/sorting_techniques/quick"
	"go_dsa/sorting_techniques/selection"

	// Stack
	stk_questions "go_dsa/stack/questions/parenthesis_matching"
	stk_using_array "go_dsa/stack/using_array"
	stk_using_linkedlist "go_dsa/stack/using_linkedlist"

	// Tree
	binary_search_tree "go_dsa/trees/binary_search"
	binary_tree "go_dsa/trees/binary_tree"
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
	// debugBinaryTree()
	// debugBinarySearchTree()
	// debugHeap()
	debugSortingTechniques()
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
	// binaryTree.Create([]int{8, 3, 9, 7, -1, 6, 4, 5, -1, -1, 2}, -1)
	// fmt.Println(binaryTree.GenerateUsingPreAndInOrder([]int{4, 7, 9, 6, 3, 2, 5, 8, 1}, []int{7, 6, 9, 3, 4, 5, 8, 2, 1}))
	// fmt.Println(binaryTree.GenerateUsingPostAndInOrder([]int{6, 3, 9, 7, 8, 5, 1, 2, 4}, []int{7, 6, 9, 3, 4, 5, 8, 2, 1}))
	fmt.Println(binaryTree.GenerateUsingPostAndInOrder([]int{20, 130, 140, 60, 80, 70, 50, 160, 150, 120, 110, 40, 30, 10}, []int{20, 10, 50, 80, 130, 60, 140, 70, 30, 40, 160, 120, 150, 110}))
	fmt.Println("PreOrder")
	binaryTree.PreOrderUsingRecursion(",")
	binaryTree.PreOrderUsingLoop(",")
	fmt.Println("PostOrder")
	binaryTree.PostOrderUsingRecursion(",")
	binaryTree.PostOrderUsingLoop(",")
	fmt.Println("InOrder")
	binaryTree.InOrderUsingRecursion(",")
	binaryTree.InOrderUsingLoop(",")
	fmt.Println("LevelOrder")
	binaryTree.LevelOrderUsingLoop(",")

	fmt.Println(binaryTree.Count())
	fmt.Println(binaryTree.Height())
	fmt.Println(binaryTree.CountLeafNode())
	fmt.Println(binaryTree.CountNonLeafNode())
}

func debugBinarySearchTree() {
	binarySearchTree := binary_search_tree.BinarySearchTree{}
	// fmt.Println(binarySearchTree.CreateUsingLoop(10, 20, 30, 40, 50))
	// fmt.Println(binarySearchTree.CreateUsingRecursion(10, 20, 30, 40, 50))

	// fmt.Println(binarySearchTree.CreateUsingLoop(30, 50, 40, 20, 10, 5, 15, 45))
	// fmt.Println(binarySearchTree.CreateUsingRecursion(30, 50, 40, 20, 10, 5, 15, 45))

	// fmt.Println(binarySearchTree.CreateUsingLoop(30, 50, 40, 20, 10, 5, 5, 15, 45))
	// fmt.Println(binarySearchTree.CreateUsingRecursion(30, 50, 40, 20, 10, 5,5, 15, 45))

	// binarySearchTree.PreOrderUsingRecursion(",")
	// binarySearchTree.InOrderUsingRecursion(",")
	// binarySearchTree.PostOrderUsingRecursion(",")

	// fmt.Println("Search Using Loop")
	// fmt.Println(binarySearchTree.SearchUsingLoop(10))
	// fmt.Println(binarySearchTree.SearchUsingLoop(11))
	// fmt.Println(binarySearchTree.SearchUsingLoop(15))
	// fmt.Println(binarySearchTree.SearchUsingLoop(-1))
	// fmt.Println("Search Using Recursion")
	// fmt.Println(binarySearchTree.SearchUsingRecursion(10))
	// fmt.Println(binarySearchTree.SearchUsingRecursion(11))
	// fmt.Println(binarySearchTree.SearchUsingRecursion(15))
	// fmt.Println(binarySearchTree.SearchUsingRecursion(-1))

	// fmt.Println(binarySearchTree.CreateUsingLoop(9, 15, 5, 20, 16, 8, 12, 3, 6))
	// fmt.Println(binarySearchTree.CreateUsingRecursion(9, 15, 5, 20, 16, 8, 12, 3, 6))

	// binarySearchTree.PreOrderUsingRecursion(",")
	// binarySearchTree.InOrderUsingRecursion(",")
	// binarySearchTree.PostOrderUsingRecursion(",")

	// fmt.Println("--------Search Using Loop--------")
	// fmt.Println(binarySearchTree.SearchUsingLoop(9))
	// fmt.Println(binarySearchTree.SearchUsingLoop(1))
	// fmt.Println(binarySearchTree.SearchUsingLoop(16))
	// fmt.Println(binarySearchTree.SearchUsingLoop(-1))
	// fmt.Println("--------Search Using Recursion--------")
	// fmt.Println(binarySearchTree.SearchUsingRecursion(9))
	// fmt.Println(binarySearchTree.SearchUsingRecursion(1))
	// fmt.Println(binarySearchTree.SearchUsingRecursion(16))
	// fmt.Println(binarySearchTree.SearchUsingRecursion(-1))

	// fmt.Println("--------Delete--------")
	// binarySearchTree.PreOrderUsingRecursion(",")
	// binarySearchTree.InOrderUsingRecursion(",")
	// binarySearchTree.PostOrderUsingRecursion(",")

	// fmt.Println("Deleting 9:- ", binarySearchTree.Delete(9))
	// binarySearchTree.PreOrderUsingRecursion(",")
	// binarySearchTree.InOrderUsingRecursion(",")
	// binarySearchTree.PostOrderUsingRecursion(",")

	// fmt.Println("Deleting 5:- ", binarySearchTree.Delete(5))
	// binarySearchTree.PreOrderUsingRecursion(",")
	// binarySearchTree.InOrderUsingRecursion(",")
	// binarySearchTree.PostOrderUsingRecursion(",")

	// fmt.Println("Deleting 16:- ", binarySearchTree.Delete(16))
	// binarySearchTree.PreOrderUsingRecursion(",")
	// binarySearchTree.InOrderUsingRecursion(",")
	// binarySearchTree.PostOrderUsingRecursion(",")

	// fmt.Println("Deleting -1:- ", binarySearchTree.Delete(-1))
	// binarySearchTree.PreOrderUsingRecursion(",")
	// binarySearchTree.InOrderUsingRecursion(",")
	// binarySearchTree.PostOrderUsingRecursion(",")

	fmt.Println(binarySearchTree.GenerateUsingPreOrder(30, 20, 10, 15, 25, 40, 50, 45))
	binarySearchTree.PreOrderUsingRecursion(",")
	binarySearchTree.InOrderUsingRecursion(",")
	binarySearchTree.PostOrderUsingRecursion(",")
}

func debugHeap() {
	// h := heap.Heap{}
	// h.Display(",")
	// h.Insert(10)
	// h.Display(",")
	// h.Insert(20)
	// h.Display(",")
	// h.Insert(30)
	// h.Display(",")
	// h.Insert(25)
	// h.Display(",")
	// h.Insert(5)
	// h.Display(",")
	// h.Insert(40)
	// h.Display(",")
	// h.Insert(35)
	// h.Display(",")

	// h := heap.Heap{Which: heap.MIN}
	// h := heap.Heap{Which: heap.MAX}
	// h.Display(",")
	// h.Create(10, 20, 30, 25, 5, 40, 35)
	// h.Display(",")
	// fmt.Println(h.Delete())
	// h.Display(",")
	// fmt.Println(h.Delete())
	// h.Display(",")
	// fmt.Println(h.Delete())
	// h.Display(",")
	// fmt.Println(h.Delete())
	// h.Display(",")
	// fmt.Println(h.Delete())
	// h.Display(",")
	// fmt.Println(h.Delete())
	// h.Display(",")
	// fmt.Println(h.Delete())
	// h.Display(",")
	// fmt.Println(h.Delete())

	values := heap.Sort(heap.INCREASING, 10, 20, 30, 25, 5, 40, 35)
	fmt.Println(values)
}

func debugSortingTechniques() {
	// Bubble Sort
	debugBubbleSort([]int{3, 7, 9, 10, 6, 5, 12, 4, 11, 2})
	// Insertion Sort
	debugInsertionSort([]int{8, 5, 7, 3, 2})
	// Selection Sort
	debugSelectionSort([]int{11, 13, 7, 2, 6, 9, 4, 5, 10, 3})
	// Quick Sort
	debugQuickSort([]int{11, 13, 7, 2, 6, 9, 4, 5, 10, 3})
	// Merge Sort
	debugMergeSort([]int{11, 13, 7, 2, 6, 9, 4, 5, 10, 3})
	// Count Sort
	debugCountSort([]int{11, 13, 7, 2, 6, 9, 4, 5, 10, 3})
}

func debugBubbleSort(values []int) {
	fmt.Println("Before Bubble Sort: ", values)
	bubble.Sort(values)
	fmt.Println("After Bubble Sort: ", values)
}

func debugInsertionSort(values []int) {
	fmt.Println("Before Insertion Sort: ", values)
	insertion.Sort(values)
	fmt.Println("After Insertion Sort: ", values)
}

func debugSelectionSort(values []int) {
	fmt.Println("Before Selection Sort: ", values)
	selection.Sort(values)
	fmt.Println("After Selection Sort: ", values)
}

func debugQuickSort(values []int) {
	fmt.Println("Before Quick Sort: ", values)
	quick.Sort(values)
	fmt.Println("After Quick Sort: ", values)
}

func debugMergeSort(values []int) {
	fmt.Println("Before Merge Sort: ", values)
	merge.Sort(values)
	fmt.Println("After Merge Sort: ", values)
}

func debugCountSort(values []int) {
	fmt.Println("Before Count Sort: ", values)
	err := count.Sort(values)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After Count Sort: ", values)
}
