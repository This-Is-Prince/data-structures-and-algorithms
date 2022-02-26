package heap

import "fmt"

type Heap struct {
	data []int
}

func (heap *Heap) Create(values ...int) {
	for _, value := range values {
		heap.Insert(value)
	}
}

func (heap *Heap) Insert(value int) {
	if heap.data == nil {
		heap.data = []int{0}
	}
	heap.data = append(heap.data, value)
	index := len(heap.data) - 1
	for index > 1 && value > heap.data[index/2] {
		heap.data[index] = heap.data[index/2]
		index = index / 2
	}
	heap.data[index] = value
}

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
