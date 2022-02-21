package main

import (
	"fmt"
	"go_dsa/linkedlist/circular"
)

func main() {
	fmt.Println("Data Structures and Algorithms")

	// list1 := circular.LinkedList{}
	// list1.Create([]int{10, 20, 30, 40}...)
	// list1.DisplayUsingLoop(",")
	// fmt.Println(list1.CountUsingLoop())
	// fmt.Println(list1.CountUsingRecursion())
	// list1.DisplayUsingLoop(",")

	// fmt.Println(list1.SumUsingLoop())
	// fmt.Println(list1.SumUsingRecursion())
	// list1.DisplayUsingLoop(",")

	// fmt.Println(list1.AverageUsingLoop())
	// fmt.Println(list1.AverageUsingRecursion())
	// list1.DisplayUsingLoop(",")

	// fmt.Println(list1.MaxUsingLoop())
	// fmt.Println(list1.MaxUsingRecursion())
	// list1.DisplayUsingLoop(",")

	// fmt.Println(list1.MinUsingLoop())
	// fmt.Println(list1.MinUsingRecursion())
	// list1.DisplayUsingLoop(",")

	// node2, err2 := list1.LinearSearchUsingLoop(2)
	// if err2 == nil {
	// 	node2.Data = -node2.Data
	// 	fmt.Println(node2.Data)
	// } else {
	// 	fmt.Println(err2)
	// }
	// node40, err40 := list1.LinearSearchUsingRecursion(40)
	// if err40 == nil {
	// 	node40.Data = -node40.Data
	// 	fmt.Println(node40.Data)
	// } else {
	// 	fmt.Println(err40)
	// }
	// list1.DisplayUsingLoop(",")

	// fmt.Println(list1.InsertUsingLoop(0, 10))
	// fmt.Println(list1.InsertUsingLoop(1, 11))
	// fmt.Println(list1.InsertUsingLoop(2, 12))
	// fmt.Println(list1.InsertUsingLoop(4, 15))
	// fmt.Println(list1.InsertUsingLoop(9, 20))
	// list1.DisplayUsingLoop(",")

	// fmt.Println(list1.InsertUsingRecursion(0, 10))
	// fmt.Println(list1.InsertUsingRecursion(1, 11))
	// fmt.Println(list1.InsertUsingRecursion(2, 12))
	// fmt.Println(list1.InsertUsingRecursion(4, 15))
	// fmt.Println(list1.InsertUsingRecursion(9, 20))
	// list1.DisplayUsingLoop(",")

	// list1.AddFirst(20)
	// list1.AddLastUsingLoop(21)
	// list1.AddLastUsingRecursion(22)
	// list1.AddLastUsingLoop(211)
	// list1.AddFirst(200)
	// list1.DisplayUsingLoop(",")

	// list1.InsertInSortedListUsingLoop(5)
	// list1.InsertInSortedListUsingLoop(50)
	// list1.InsertInSortedListUsingLoop(15)
	// list1.InsertInSortedListUsingLoop(35)
	// list1.DisplayUsingLoop(",")

	// list1.InsertInSortedListUsingRecursion(5)
	// list1.InsertInSortedListUsingRecursion(50)
	// list1.InsertInSortedListUsingRecursion(15)
	// list1.InsertInSortedListUsingRecursion(35)
	// list1.DisplayUsingLoop(",")

	// list1.DeleteFirst()
	// list1.DeleteLastUsingLoop()
	// list1.DeleteLastUsingRecursion()
	// list1.DisplayUsingLoop(",")

	// fmt.Println(list1.DeleteUsingLoop(-1))
	// fmt.Println(list1.DeleteUsingLoop(0))
	// fmt.Println(list1.DeleteUsingLoop(2))
	// fmt.Println(list1.DeleteUsingLoop(1))
	// fmt.Println(list1.DeleteUsingLoop(1))
	// list1.DisplayUsingLoop(",")

	// fmt.Println(list1.DeleteUsingRecursion(-1))
	// fmt.Println(list1.DeleteUsingRecursion(0))
	// fmt.Println(list1.DeleteUsingRecursion(2))
	// fmt.Println(list1.DeleteUsingRecursion(1))
	// fmt.Println(list1.DeleteUsingRecursion(1))
	// list1.DisplayUsingLoop(",")

	// fmt.Println(list1.IsSortedUsingLoop())
	// fmt.Println(list1.IsSortedUsingRecursion())
	// list1.DisplayUsingLoop(",")

	// fmt.Println(list1.RemoveDuplicateFromSortedListUsingLoop())
	// fmt.Println(list1.RemoveDuplicateFromSortedListUsingRecursion())
	// list1.DisplayUsingLoop(",")

	// fmt.Println(list1.ReverseListUsingLoop())
	// list1.DisplayUsingLoop(",")
	// fmt.Println(list1.ReverseListUsingRecursion())
	// list1.DisplayUsingLoop(",")

	list3 := circular.LinkedList{}
	list3.Create([]int{10, 20, 30, 30, 40}...)
	fmt.Println("list3")
	list3.DisplayUsingLoop(",")
	list4 := circular.LinkedList{}
	list4.Create([]int{15, 45}...)
	fmt.Println("list4")
	list4.DisplayUsingLoop(",")
	// // list3.ConcatUsingLoop(&list4)
	// list3.ConcatUsingRecursion(&list4)
	// list3.DisplayUsingLoop(",")

	fmt.Println("mergelist")
	list3.MergeUsingLoop(&list4)
	list3.DisplayUsingLoop(",")

}
