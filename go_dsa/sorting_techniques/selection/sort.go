package selection

/*
	Critera for Analysis
	1. Number of Comparisons
	2. Number of Swaps
	3. Adaptive (if complexity changes based on pre-sorted input i.e. pre-sorted array of the input affects the running time.)
	4. Stable (Sorting algorithm is stable if two elements with equal values appear in the same order in output as it was in the input.)
	5. Extra Memory

	1. min - O(n^2), max - O(n^2)
	2. min - O(1), max - O(n)
	3. selection sort is not adaptive
	4. insertion sort is not stable
	5. no extra memory O(1)

	TIME
	min - O(n^2), max - O(n^2)

	if i performed 1 pass i can get 1 sorted element (1 smallest element)
	if i performed k pass i can get k sorted element (k smallest element)
*/

func Sort(values []int) {
	n := len(values)
	for i, k := 0, 0; i < n-1; i++ {
		k = i
		for j := i; j < n; j++ {
			if values[j] < values[k] {
				k = j
			}
		}
		values[i], values[k] = values[k], values[i]
	}
}
