package heap

import "fmt"

const (
	MAX = "max"
	MIN = "min"
)

type Heap struct {
	Which   string
	data    []int
	cmpFunc func(int, int) bool
}

func (heap *Heap) Create(values ...int) {
	for _, value := range values {
		heap.Insert(value)
	}
}

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
