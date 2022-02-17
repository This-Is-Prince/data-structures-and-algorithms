package main

import (
	"fmt"
	"go_dsa/linkedlist"
)

func main() {
	fmt.Println("Data Structures and Algorithms")

	linkedList := linkedlist.SinglyLinkedList{}
	fmt.Println(linkedList)
	linkedList.Create([]int{1, 2, 3, 4, 10, 9, 8, 5, 6}...)
	fmt.Println(linkedList)
	linkedList.DisplayUsingRecursion("-")
	linkedList.DisplayUsingLoop(",")
	linkedList.DisplayUsingRecursion("-")
	linkedList.DisplayUsingLoop(",")
	linkedList.DisplayUsingLoopInReverseOrder(" , ")
	linkedList.DisplayUsingRecursionInReverseOrder(" : ")
	fmt.Println(linkedList.CountUsingRecursion())
	fmt.Println(linkedList.CountUsingLoop())
	fmt.Println(linkedList.SumUsingLoop())
	fmt.Println(linkedList.SumUsingRecursion())
	fmt.Println(linkedList.Max())
}
