package main

import "fmt"

// FindDuplicates scans the array in a single pass to determine duplicate numbers
func FindDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	duplicates := []int{}
	for _, num := range arr {
		if _, ok := seen[num]; ok {
			// Check if the duplicate has already been added
			alreadyAdded := false
			for _, dup := range duplicates {
				if dup == num {
					alreadyAdded = true
					break
				}
			}
			if !alreadyAdded {
				duplicates = append(duplicates, num)
			}
		} else {
			seen[num] = true
		}
	}
	return duplicates
}

func main() {
	arr := []int{4, 5, 6, 7, 4, 6, 8, 5}
	fmt.Println(FindDuplicates(arr)) // Expected Output: [4, 5, 6]

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(FindDuplicates(arr)) // Expected Output: []

	arr = []int{-1, -2, -3, -1, -2, -4}
	fmt.Println(FindDuplicates(arr)) // Expected Output: [-1, -2]
}
