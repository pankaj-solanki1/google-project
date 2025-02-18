package main

import "fmt"

// SumEvenNumbers calculates the sum of even numbers from the list, excluding negatives and 0 when applicable
func SumEvenNumbers(arr []int) int {
	totalSum := 0
	evenSum := 0

	for _, num := range arr {
		totalSum += num
		if num > 0 && num%2 == 0 {
			evenSum += num
		}
	}

	if totalSum%2 == 0 {
		return 0
	}

	return evenSum
}

func main() {
	arr := []int{1, 4, 5, 7, 9, 10}
	fmt.Println(SumEvenNumbers(arr)) // Expected Output: 14

	arr = []int{-2, -4, 3, 5}
	fmt.Println(SumEvenNumbers(arr)) // Expected Output: 0

	arr = []int{0, 2, 4, 6}
	fmt.Println(SumEvenNumbers(arr)) // Expected Output: 0

	arr = []int{1, 2, 3}
	fmt.Println(SumEvenNumbers(arr)) // Expected Output: 2

	arr = []int{2, 4, 6}
	fmt.Println(SumEvenNumbers(arr)) // Expected Output: 0

	arr = []int{}
	fmt.Println(SumEvenNumbers(arr)) // Expected Output: 0

	arr = []int{0}
	fmt.Println(SumEvenNumbers(arr)) // Expected Output: 0

	arr = []int{-100, -200, 1, 2}
	fmt.Println(SumEvenNumbers(arr)) // Expected output: 2

	arr = []int{2}
	fmt.Println(SumEvenNumbers(arr)) // Expected output: 0

	arr = []int{1}
	fmt.Println(SumEvenNumbers(arr)) // Expected output: 0
}
