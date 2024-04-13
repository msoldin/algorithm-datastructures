package linear_search

func search(arr []int, elem int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			return i
		}
	}
	return -1
}
