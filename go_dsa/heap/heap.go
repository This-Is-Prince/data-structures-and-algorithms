package heap

import (
	"errors"
	"fmt"
)

const (
	MAX = "max"
	MIN = "min"
)

type Heap struct {
	Which   string
	data    []int
	cmpFunc func(int, int) bool
}

// Create function for creating heap with given values
func (heap *Heap) Create(values ...int) {
	for _, value := range values {
		heap.Insert(value)
	}
}

// insert function for inserting given value in heap
func (heap *Heap) Insert(value int) {
	if heap.data == nil {
		if heap.Which == MIN {
			heap.cmpFunc = func(child, parent int) bool {
				return child < parent
			}
		} else {
			heap.Which = MAX
			heap.cmpFunc = func(child, parent int) bool {
				return child > parent
			}
		}
		heap.data = []int{0}
	}
	heap.data = append(heap.data, value)
	index := len(heap.data) - 1
	for index > 1 && heap.cmpFunc(value, heap.data[index/2]) {
		heap.data[index] = heap.data[index/2]
		index = index / 2
	}
	heap.data[index] = value
}

// delete max priority heap value
func (heap *Heap) Delete() (int, error) {
	if length := len(heap.data); length <= 1 {
		return 0, errors.New("heap is empty")
	} else {
		value, i, j := heap.data[1], 1, 2
		length--
		heap.data[1] = heap.data[length]
		heap.data = heap.data[:length]
		for j < length {
			if j < length-1 {
				if heap.cmpFunc(heap.data[j+1], heap.data[j]) {
					j = j + 1
				}
			}
			if heap.cmpFunc(heap.data[j], heap.data[i]) {
				heap.data[i], heap.data[j] = heap.data[j], heap.data[i]
				i = j
				j = 2 * i
			} else {
				break
			}
		}
		return value, nil
	}
}

// display all heap values
func (heap *Heap) Display(sep string) {
	if length := len(heap.data); length > 1 {
		index := 1
		for index < length-1 {
			fmt.Print(heap.data[index], sep)
			index++
		}
		fmt.Println(heap.data[index])
	}
}
