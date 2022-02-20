package main

import (
	"fmt"
	"go_dsa/linkedlist/circular"
)

func main() {
	fmt.Println("Data Structures and Algorithms")

	list1 := circular.LinkedList{}
	list1.DisplayUsingLoop(",")
	list1.Create([]int{1, 2, 3, 4}...)
	list1.DisplayUsingRecursion("-")
	list1.DisplayUsingLoop(",")
}
