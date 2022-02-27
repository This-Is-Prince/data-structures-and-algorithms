package radix

import (
	"errors"
	"fmt"
	"math"
)

type node struct {
	data int
	next *node
}

func Sort(values []int) error {
	if n := len(values); n > 0 {
		pass, err := maxDigits(values)
		if err != nil {
			return err
		}
		bins := make([]*node, 10)
		for i := 0; i < pass; i++ {
			for _, value := range values {
				insert(bins, getIndex(value, i), value)
			}
			k := 0
			for j := 0; j < 10; j++ {
				for bins[j] != nil {
					values[k], err = delete(bins, j)
					if err != nil {
						return err
					}
					k++
				}
			}
		}
		return nil
	}
	return errors.New("there is no value to sort")
}

func insert(bins []*node, index, value int) error {
	if n := len(bins); n == 0 {
		return errors.New("empty bins")
	} else if n <= index || index < 0 {
		return errors.New("invalid index")
	} else {
		newNode := &node{data: value}
		if bins[index] == nil {
			bins[index] = newNode
		} else {
			head := bins[index]
			for head.next != nil {
				head = head.next
			}
			head.next = newNode
		}
		return nil
	}
}

func delete(bins []*node, index int) (int, error) {
	if n := len(bins); n == 0 {
		return 0, errors.New("empty bins")
	} else if n <= index || index < 0 {
		return 0, errors.New("invalid index")
	} else {
		value := bins[index]
		bins[index] = bins[index].next
		return value.data, nil
	}
}

func maxDigits(values []int) (int, error) {
	max := 0
	for _, value := range values {
		if value < 0 {
			return 0, errors.New("value can't be negative")
		}
		value = len(fmt.Sprint(value))
		if max < value {
			max = value
		}
	}
	return max, nil
}

func getIndex(value, index int) int {
	return (value / int(math.Pow(10, float64(index)))) % 10
}
