package merge_sort

// O(n^2) time complexity
func sort(arr []int) {
	sorted := mergeSort(arr)
	for i := 0; i < len(sorted); i++ {
		arr[i] = sorted[i]
	}
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left := mergeSort(arr[:len(arr)/2])
	right := mergeSort(arr[len(arr)/2:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	i := 0
	j := 0
	result := make([]int, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[i+j] = left[i]
			i++
		} else {
			result[i+j] = right[j]
			j++
		}
	}
	for ; i < len(left); i++ {
		result[i+j] = left[i]
	}
	for ; j < len(right); j++ {
		result[i+j] = right[j]
	}
	return result
}
