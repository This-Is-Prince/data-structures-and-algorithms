package count

import "errors"

func Sort(values []int) error {
	auxLen, err := max(values)
	if err != nil {
		return err
	}
	aux := make([]int, auxLen+1)
	for _, value := range values {
		if value < 0 {
			return errors.New("value can't be negative")
		}
		aux[value]++
	}
	i, j := 0, 0
	for i < auxLen+1 {
		if aux[i] > 0 {
			values[j] = i
			j++
			aux[i]--
		} else {
			i++
		}
	}
	return nil
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
