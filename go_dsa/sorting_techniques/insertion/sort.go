package insertion

/*
	Critera for Analysis
	1. Number of Comparisons
	2. Number of Swaps
	3. Adaptive (if complexity changes based on pre-sorted input i.e. pre-sorted array of the input affects the running time.)
	4. Stable (Sorting algorithm is stable if two elements with equal values appear in the same order in output as it was in the input.)
	5. Extra Memory

	1. min - O(n), max - O(n^2)
	2. min - O(1), max - O(n^2)
	3. insertion sort is by default adaptive
	4. insertion sort is stable
	5. no extra memory O(1)

	TIME
	min - O(n), max - O(n^2)

	insertion sort is more usefull in linkedlist than array
*/

func Sort(values []int) {
	n, j, value := len(values), 0, 0
	for i := 1; i < n; i++ {
		j, value = i-1, values[i]
		for j > -1 && values[j] > value {
			values[j+1] = values[j]
			j--
		}
		values[j+1] = value
	}
}
