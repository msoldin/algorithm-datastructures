package counting_sort

import "slices"

func sort(arr []int) {
	count := make([]int, slices.Max(arr)+1)
	for i := 0; i < len(arr); i++ {
		count[arr[i]] = count[arr[i]] + 1
	}

	index := 0
	for i := 0; i < len(count); i++ {
		for j := count[i]; j > 0; j-- {
			arr[index] = i
			index++
		}
	}
}
