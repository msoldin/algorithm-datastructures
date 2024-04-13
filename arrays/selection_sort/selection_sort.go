package selection_sort

// O(n^2) time complexity
func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		lowestIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[lowestIndex] {
				lowestIndex = j
			}
		}
		arr[i], arr[lowestIndex] = arr[lowestIndex], arr[i]
	}
}
