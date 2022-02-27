package bubble

func Sort(values []int) {
	n, flag := len(values), false
	for i := 0; i < n-1; i++ {
		flag = false
		for j := 0; j < n-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
