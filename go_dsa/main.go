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
	fmt.Println(list1.CountUsingLoop())
	fmt.Println(list1.CountUsingRecursion())
	list1.DisplayUsingLoop(",")

	fmt.Println(list1.SumUsingLoop())
	fmt.Println(list1.SumUsingRecursion())
	list1.DisplayUsingLoop(",")

	fmt.Println(list1.AverageUsingLoop())
	fmt.Println(list1.AverageUsingRecursion())
	list1.DisplayUsingLoop(",")

	fmt.Println(list1.MaxUsingLoop())
	fmt.Println(list1.MaxUsingRecursion())
	list1.DisplayUsingLoop(",")

	fmt.Println(list1.MinUsingLoop())
	fmt.Println(list1.MinUsingRecursion())
	list1.DisplayUsingLoop(",")
}
