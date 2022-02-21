package main

import (
	"fmt"
	"go_dsa/linkedlist/circular"
)

func main() {
	fmt.Println("Data Structures and Algorithms")

	list1 := circular.LinkedList{}
	list1.Create([]int{1, 2, 3, 4}...)
	list1.DisplayUsingLoop(",")
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

	fmt.Println(list1.InsertUsingLoop(0, 10))
	fmt.Println(list1.InsertUsingLoop(1, 11))
	fmt.Println(list1.InsertUsingLoop(2, 12))
	fmt.Println(list1.InsertUsingLoop(4, 15))
	fmt.Println(list1.InsertUsingLoop(9, 20))
	list1.DisplayUsingLoop(",")

	// fmt.Println(list1.InsertUsingRecursion(0, 10))
	// fmt.Println(list1.InsertUsingRecursion(1, 11))
	// fmt.Println(list1.InsertUsingRecursion(2, 12))
	// fmt.Println(list1.InsertUsingRecursion(4, 15))
	// fmt.Println(list1.InsertUsingRecursion(9, 20))
	// list1.DisplayUsingLoop(",")

	list1.AddFirst(20)
	list1.AddLastUsingLoop(21)
	list1.AddLastUsingRecursion(22)
	list1.AddLastUsingLoop(211)
	list1.AddFirst(200)
	list1.DisplayUsingLoop(",")

}
