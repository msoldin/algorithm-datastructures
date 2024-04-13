package fibbonacci

func Fibonnacci(number int) int {
	prev1 := 0
	prev2 := 1
	for i := 1; i < number; i++ {
		current := prev1 + prev2
		prev1 = prev2
		prev2 = current
	}
	return prev2
}

func RecFibonnacci(number int) int {
	if number <= 1 {
		return number
	}
	return RecFibonnacci(number-1) + RecFibonnacci(number-2)
}
