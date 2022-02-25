package main

import (
	"fmt"
	// "go_dsa/trees/binary_tree"
	// questions "go_dsa/stack/questions/parenthesis_matching"
	// "go_dsa/stack/using_array"
	// "go_dsa/stack/using_linkedlist"
	"go_dsa/queues/linear/using_array"
)

func main() {
	fmt.Println("Data Structures and Algorithms")

	// binaryTree := binary_tree.BinaryTree{}
	// binaryTree.Create([]int{5, 8, 6, -1, 9, 3, -1, 4, 2}, -1)
	// binaryTree.PreOrderUsingRecursion(",")
	// binaryTree.PostOrderUsingRecursion(",")
	// binaryTree.InOrderUsingRecursion(",")
	// binaryTree.InOrderUsingLoop(",")
	// binaryTree.LevelOrderUsingLoop(",")

	// stk, err := using_array.Create(10)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	stk.Display(",")
	// 	stk.Push(1)
	// 	stk.Push(2)
	// 	stk.Push(3)
	// 	stk.Push(4)
	// 	stk.Push(5)
	// 	stk.Display(",")
	// }

	// stk1 := using_linkedlist.Stack{}
	// stk1.Display(",")
	// stk1.Push(1)
	// stk1.Push(2)
	// stk1.Push(3)
	// stk1.Push(4)
	// stk1.Push(5)
	// stk1.Display(",")

	// fmt.Println("Is '((a+b)*(c-d))' Balance", questions.IsBalance("((a+b)*(c-d))"))
	// fmt.Println("Is '((a+b)*(c-d)' Balance", questions.IsBalance("((a+b)*(c-d)"))
	// fmt.Println("Is '((a+b)*(c-d)))' Balance", questions.IsBalance("((a+b)*(c-d)))"))
	// fmt.Println("Is '{([a+b]*[c-d])/e}' Balance", questions.IsBalance("{([a+b]*[c-d])/e}"))
	// fmt.Println("Is '{([a+b]*[c-d]/e}' Balance", questions.IsBalance("{([a+b]*[c-d]/e}"))
	// fmt.Println("Is '{([a+b]*[c-d)/e}' Balance", questions.IsBalance("{([a+b]*[c-d)/e}"))
	// fmt.Println("Is '{([a+b)*[c-d]]/e}' Balance", questions.IsBalance("{([a+b)*[c-d]]/e}"))

	q, err := using_array.Create(5)
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
}
