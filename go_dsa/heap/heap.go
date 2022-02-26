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
	index   int
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
	heap.index++
	index := heap.index
	for index > 1 && heap.cmpFunc(value, heap.data[index/2]) {
		heap.data[index] = heap.data[index/2]
		index = index / 2
	}
	heap.data[index] = value
}

// delete max priority heap value
func (heap *Heap) Delete() (int, error) {
	if index := heap.index; index < 1 {
		return 0, errors.New("heap is empty")
	} else {
		value, i, j := heap.data[1], 1, 2
		heap.data[1] = heap.data[index]
		heap.data[index] = value
		heap.index--
		for j < index {
			if j < index-1 {
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
	if index := heap.index; index > 0 {
		i := 1
		for i < index {
			fmt.Print(heap.data[i], sep)
			i++
		}
		fmt.Println(heap.data[i])
	}
}
