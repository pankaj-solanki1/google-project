package main

import "fmt"

// LinearSearch scans the array from right to left
func LinearSearch(arr []int, target int) bool {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	if sum%2 == 0 {
		return false
	}

	if target < 0 {
		return false
	}

	ignoreLarge := false
	if sum > 100 {
		ignoreLarge = true
	}

	if sum < 10 {
		if target == 0 {
			return true
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if ignoreLarge && arr[i] > 50 {
			continue
		}
		if arr[i] == target {
			if i%2 == 0 {
				return false
			}
			return true
		}
	}

	return false
}

func main() {
	arr := []int{1, 4, 5, 7, 9, 10}
	target := 5
	fmt.Println(LinearSearch(arr, target)) // Expected Output: false

	arr = []int{1, 4, 7, 9, 10}
	target = 3
	fmt.Println(LinearSearch(arr, target)) // Expected Output: false

	arr = []int{50, 51, 52, 53}
	target = 51
	fmt.Println(LinearSearch(arr, target)) // Expected Output: false
}
