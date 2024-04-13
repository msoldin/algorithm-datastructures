package binary_search

func search(arr []int, elem int) int {
	return binarySearch(arr, elem, 0, len(arr)-1)
}

func binarySearch(arr []int, elem int, left int, right int) int {
	if right < left {
		return -1
	}

	center := (left + right) / 2
	if elem < arr[center] {
		return binarySearch(arr, elem, left, center-1)
	} else if elem > arr[center] {
		return binarySearch(arr, elem, center+1, right)
	}
	return center
}
