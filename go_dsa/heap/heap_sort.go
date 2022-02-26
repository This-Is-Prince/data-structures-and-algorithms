package heap

const (
	INCREASING = "inc"
	DECREASING = "dec"
)

func Sort(order string, values ...int) []int {
	h := Heap{Which: MAX}
	if order == DECREASING {
		h.Which = MIN
	}
	h.Create(values...)
	_, err := h.Delete()
	for err == nil {
		_, err = h.Delete()
	}
	return h.data[1:]
}
