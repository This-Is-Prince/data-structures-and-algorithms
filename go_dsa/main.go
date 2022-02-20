package main

import (
	"fmt"
	"go_dsa/linkedlist"
)

func main() {
	fmt.Println("Data Structures and Algorithms")

	// linkedList1 := linkedlist.SinglyLinkedList{}
	// linkedList2 := linkedlist.SinglyLinkedList{}

	// linkedList1.Create([]int{2, 6, 8, 9, 10, 15}...)
	// linkedList2.Create([]int{4, 7, 12, 14}...)

	// linkedList1.ConcatUsingRecursion(&linkedList2)
	// linkedList1.DisplayUsingLoop(",")
	// fmt.Println(linkedList2.ReverseListUsingLoop())
	// linkedList2.DisplayUsingLoop(",")
	// linkedList1.DisplayUsingLoop(",")

	// linkedList1.DisplayUsingLoop(",")
	// linkedList2.DisplayUsingLoop(",")
	// linkedList1.MergeUsingLoop(&linkedList2)
	// linkedList1.DisplayUsingLoop(",")
	// linkedList2.DisplayUsingLoop(",")

	// linkedList1.DisplayUsingLoop(",")
	// linkedList2.DisplayUsingLoop(",")
	// linkedList1.MergeUsingRecursion(&linkedList2)
	// linkedList1.DisplayUsingLoop(",")
	// linkedList2.DisplayUsingLoop(",")

	linkedList3 := linkedlist.SinglyLinkedList{}
	linkedList3.Create([]int{10, 20, 30, 40, 50}...)
	root := linkedList3.GetRoot()
	first := root.Next
	second := root.Next.Next.Next.Next
	second.Next = first
	fmt.Println(linkedList3.CheckLoopUsingRecursion())
	fmt.Println(linkedList3.CheckLoopUsingLoop())
}
