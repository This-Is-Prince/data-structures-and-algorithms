package quick

// Selection Exchange Sort
// Partition Exchange Sort
// Quick Sort

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
	sortUtil(values, 0, n-1)
}

func sortUtil(values []int, l, h int) {
	if l < h {
		j := partition(values, l, h)
		sortUtil(values, l, j-1)
		sortUtil(values, j+1, h)
	}
}

func partition(values []int, l, h int) int {
	pivot, i, j := values[l], l, h
	for i <= j {
		for i <= j && values[i] <= pivot {
			i++
		}
		for j >= i && values[j] > pivot {
			j--
		}
		if i < j {
			values[i], values[j] = values[j], values[i]
		}
	}
	values[l], values[j] = values[j], values[l]
	return j
}
