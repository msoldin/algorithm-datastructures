package raddix_sort

import "slices"

func sort(arr []int) {
	maximum := slices.Max(arr)
	for exp := 1; maximum/exp > 0; exp *= 10 {
		sortExponent(arr, exp)
	}
}

func sortExponent(arr []int, exp int) {

	output := make([]int, len(arr))
	count := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		count[arr[i]/exp%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		output[count[arr[i]/exp%10]-1] = arr[i]
		count[arr[i]/exp%10]--
	}

	for i := 0; i < len(output); i++ {
		arr[i] = output[i]
	}
}
