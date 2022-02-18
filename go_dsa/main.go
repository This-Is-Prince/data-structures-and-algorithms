package main

import (
	"fmt"
	"go_dsa/linkedlist"
)

func main() {
	fmt.Println("Data Structures and Algorithms")

	linkedList := linkedlist.SinglyLinkedList{}
	// fmt.Println(linkedList)
	linkedList.Create([]int{1, 3, 6, 7, 10, 15, 21, 44, 59, 66, 68, 100, 150}...)
	// fmt.Println(linkedList)
	// linkedList.DisplayUsingRecursion("-")
	linkedList.DisplayUsingLoop(",")
	// linkedList.DisplayUsingRecursion("-")
	// linkedList.DisplayUsingLoop(",")
	// linkedList.DisplayUsingLoopInReverseOrder(" , ")
	// linkedList.DisplayUsingRecursionInReverseOrder(" : ")
	// fmt.Println(linkedList.CountUsingRecursion())
	// fmt.Println(linkedList.CountUsingLoop())
	// fmt.Println(linkedList.SumUsingLoop())
	// fmt.Println(linkedList.SumUsingRecursion())
	// fmt.Println(linkedList.MaxUsingLoop())
	// fmt.Println(linkedList.MaxUsingRecursion())
	// fmt.Println(linkedList.MinUsingLoop())
	// fmt.Println(linkedList.MinUsingRecursion())

	// if node, err := linkedList.LinearSearchUsingLoop(3); err == nil {
	// 	fmt.Println(node.Data)
	// 	node.Data = 55
	// } else {
	// 	fmt.Println(err)
	// }

	// linkedList.DisplayUsingLoop(",")
	// if node, err := linkedList.LinearSearchUsingRecursion(55); err == nil {
	// 	fmt.Println(node.Data)
	// 	node.Data = 77
	// } else {
	// 	fmt.Println(err)
	// }

	// fmt.Println(linkedList.InsertUsingLoop(9, 200))
	// linkedList.DisplayUsingLoop(",")
	// fmt.Println(linkedList.InsertUsingRecursion(5, 300))
	// linkedList.DisplayUsingLoop("-")

	// linkedList.InsertInSortedListUsingLoop(8)
	// linkedList.DisplayUsingLoop(",")

	// fmt.Println(linkedList.DeleteFirst())
	// fmt.Println(linkedList.DeleteLastUsingLoop())
	// fmt.Println(linkedList.DeleteLastUsingRecursion())
	// linkedList.DisplayUsingLoop(",")
	// fmt.Println(linkedList.DeleteUsingLoop(0))
	// linkedList.DisplayUsingLoop(",")
	// fmt.Println(linkedList.DeleteUsingRecursion(8))
	// linkedList.DisplayUsingLoop(",")

	fmt.Println(linkedList.IsSortedUsingLoop())
	fmt.Println(linkedList.IsSortedUsingRecursion())
}
