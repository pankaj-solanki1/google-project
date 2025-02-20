package main

import "fmt"

// HasDuplicates checks if an array contains duplicate elements.
func HasDuplicates(arr []int) bool {
	if len(arr) <= 1 {
		return false
	}

	seen := make(map[int]bool)
	for _, num := range arr {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 2}
	fmt.Println(HasDuplicates(arr)) // Expected Output: true

	arr = []int{1, 3, 5, 7, 9}
	fmt.Println(HasDuplicates(arr)) // Expected Output: false

	arr = []int{10, 20, 30, 40, 50}
	fmt.Println(HasDuplicates(arr)) // Expected Output: false

	arr = []int{5, 5, 5, 5, 5}
	fmt.Println(HasDuplicates(arr)) // Expected Output: true

	arr = []int{5}
	fmt.Println(HasDuplicates(arr)) // Expected Output: false
}
