package merge

/*
	Critera for Analysis
	1. Number of Comparisons
	2. Number of Swaps
	3. Adaptive (if complexity changes based on pre-sorted input i.e. pre-sorted array of the input affects the running time.)
	4. Stable (Sorting algorithm is stable if two elements with equal values appear in the same order in output as it was in the input.)
	5. Extra Memory

	1. min - O(nlogn), max - O(n^2)
	2. min - O(1), max - O(n)
	3. quick sort is not adaptive
	4. quick sort is not stable
	5. no extra memory O(1)

	### If First Element is pivot
	Best Case - if partitioning is in middle
	Time - O(nlogn)
	Worst Case -if partitioning is on any end
	Time - O(n^2) (already sorted)
	Average Case
	Time - O(nlogn)

	### If Middle Element is pivot
	Best Case - if partitioning is on any end
	Time - O(nlogn) (already sorted)
	Worst Case - if partitioning is in middle
	Time - O(n^2)
	Average Case
	Time - O(nlogn)
*/

func Sort(values []int) {
	n := len(values)
	aux := make([]int, n)
	copy(aux, values)
	sort(values, aux, 0, n-1)
}

func sort(values, aux []int, l, h int) {
	if l < h {
		mid := (l + h) / 2
		sort(values, aux, l, mid)
		copy(aux, values)
		sort(values, aux, mid+1, h)
		copy(aux, values)
		merge(values, aux, l, mid, h)
	}
}

func merge(values, aux []int, l, mid, h int) {
	i, j, k := l, mid+1, l
	for i <= mid && j <= h {
		if aux[i] < aux[j] {
			values[k] = aux[i]
			i++
		} else {
			values[k] = aux[j]
			j++
		}
		k++
	}
	for ; i <= mid; i++ {
		values[k] = aux[i]
		k++
	}
	for ; j <= h; j++ {
		values[k] = aux[j]
		k++
	}
}
