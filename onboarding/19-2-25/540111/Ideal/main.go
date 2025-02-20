package main

import "fmt"

// FindDuplicates finds all duplicates in an array while preserving their first appearance order.
func FindDuplicates(arr []int) []int {
	seen := make(map[int]bool)       // Tracks numbers seen before
	result := []int{}                // Stores the first appearance of duplicates
	duplicates := make(map[int]bool) // Ensures duplicates are added only once

	for _, num := range arr {
		if seen[num] {
			if !duplicates[num] { // Add only if it's the first duplicate
				duplicates[num] = true
				result = append(result, num) // Maintain the order
			}
		} else {
			seen[num] = true
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(FindDuplicates([]int{4, 5, 6, 7, 4, 6, 8, 5})) // Expected: [4, 5, 6]
	fmt.Println(FindDuplicates([]int{1, 2, 3, 4, 5}))          // Expected: []
	fmt.Println(FindDuplicates([]int{-1, -2, -3, -1, -2, -4})) // Expected: [-1, -2]
	fmt.Println(FindDuplicates([]int{5, 2, 1, 4, 3}))          // Expected: []
}
