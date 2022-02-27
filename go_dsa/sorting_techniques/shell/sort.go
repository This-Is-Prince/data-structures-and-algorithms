package shell

import "errors"

func Sort(values []int) error {
	if n := len(values); n > 0 {
		for gap := n / 2; gap >= 1; gap /= 2 {
			for i := gap; i < n; i++ {
				tmp := values[i]
				j := i - gap
				for j >= 0 && values[j] > tmp {
					values[j+gap] = values[j]
					j = j - gap
				}
				values[j+gap] = tmp
			}
		}
		return nil
	}
	return errors.New("there is no values to sort")
}
