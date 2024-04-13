package quicksort

func sort(arr []int) {
	quicksort(arr, 0, len(arr)-1)
}

func quicksort(arr []int, low int, high int) {
	if low < high {
		pivotElement := partition(arr, low, high)
		quicksort(arr, low, pivotElement-1)
		quicksort(arr, pivotElement+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	//Set pivot element to the right
	pivot := arr[high]

	//Set pointer of the greatest element to the left
	pointer := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {

			pointer++
			//Swap smaller element with pointer to the greatest element
			temp := arr[pointer]
			arr[pointer] = arr[i]
			arr[i] = temp
		}
	}

	temp := arr[pointer+1]
	arr[pointer+1] = pivot
	arr[high] = temp

	return pointer + 1
}
