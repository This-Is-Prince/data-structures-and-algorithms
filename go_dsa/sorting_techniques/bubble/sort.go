package bubble

/*
	Critera for Analysis
	1. Number of Comparisons
	2. Number of Swaps
	3. Adaptive (if complexity changes based on pre-sorted input i.e. pre-sorted array of the input affects the running time.)
	4. Stable (Sorting algorithm is stable if two elements with equal values appear in the same order in output as it was in the input.)
	5. Extra Memory

	1. min - O(n), max - O(n^2)
	2. min - O(1), max - O(n^2)
	3. we make bubble sort adaptive
	4. bubble sort is stable
	5. no extra memory required

	TIME
	min - O(n), max - O(n^2)
*/
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

/*
           | Bubble Sort | Insertion Sort |
Min Comp   | 	O(n)     |		O(n)	  | Already In Ascending
Max Comp   |    O(n^2)   |      O(n^2)    | Already In Descending
Min Swap   |    O(1)     |      O(1)      | Ascending
Max Swap   |    O(n^2)   |      O(n^2)    | Descending
Adaptive   |     True    |      True      |
Stable     |     True    |      True      |
LinkedList |     NO      |      YES       |
K Passes   |     YES     |      NO        |
*/
