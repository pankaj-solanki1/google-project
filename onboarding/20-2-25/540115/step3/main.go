package main

import "fmt"

// FindDuplicates finds duplicate numbers while maintaining order
func FindDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	duplicates := make(map[int]bool)
	result := []int{}

	for _, num := range arr {
		if seen[num] {
			// Append only if it's the first time we detect a duplicate
			if !duplicates[num] {
				result = append(result, num)
				duplicates[num] = true
			}
		} else {
			seen[num] = true
		}
	}

	return result
}

func main() {
	arr := []int{4, 5, 6, 7, 4, 6, 8, 5}
	fmt.Println(FindDuplicates(arr)) // Expected Output: [4, 5, 6]

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(FindDuplicates(arr)) // Expected Output: []

	arr = []int{-1, -2, -3, -1, -2, -4}
	fmt.Println(FindDuplicates(arr)) // Expected Output: [-1, -2]
}
