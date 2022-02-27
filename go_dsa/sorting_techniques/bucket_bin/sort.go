package bucket_bin

import "errors"

type node struct {
	data int
	next *node
}

func Sort(values []int) error {
	max, err := max(values)
	if err != nil {
		return err
	}
	bins := make([]*node, max+1)
	for _, value := range values {
		insert(bins, value, value)
	}

	for i, j := 0, 0; i < max+1; i++ {
		for bins[i] != nil {
			values[j], err = delete(bins, i)
			if err != nil {
				return err
			}
			j++
		}
	}
	return nil
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

func max(values []int) (int, error) {
	if n := len(values); n > 0 {
		max := values[0]
		for i := 1; i < n; i++ {
			if max < values[i] {
				max = values[i]
			}
		}
		return max, nil
	} else {
		return 0, errors.New("there is no values to sort")
	}
}
