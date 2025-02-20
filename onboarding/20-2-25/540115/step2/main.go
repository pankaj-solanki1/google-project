package main

import (
	"fmt"
	"sort"
)

// FindUnionAndIntersection returns the union and intersection of multiple integer arrays.
func FindUnionAndIntersection(arrs [][]int) ([]int, []int) {
	if len(arrs) == 0 {
		return []int{}, []int{}
	}

	unionMap := make(map[int]bool)
	var union []int

	// Build the union and use a map to track duplicates efficiently.
	for _, arr := range arrs {
		for _, num := range arr {
			if !unionMap[num] {
				unionMap[num] = true
				union = append(union, num)
			}
		}
	}

	// Sort the union.
	sort.Ints(union)

	// Find the intersection.
	if len(arrs) == 0 {
		return union, []int{} // Return empty intersection if there are no arrays.
	}

	intersectionMap := make(map[int]int) // Element -> count map.
	for _, num := range arrs[0] {
		intersectionMap[num] = 1 // Initialize counts based on the first array.
	}

	for i := 1; i < len(arrs); i++ {
		tempMap := make(map[int]bool) //Use bool to mark presence in current array
		for _, num := range arrs[i] {
			tempMap[num] = true
		}

		// Iterate through existing intersectionMap, check in tempMap, update.
		for num, count := range intersectionMap {
			if !tempMap[num] {
				delete(intersectionMap, num) // Remove if not in current array
			} else {
				intersectionMap[num] = count + 1 // increment count. Unused. but could be useful
			}
		}

		//If map is now empty, no intersection.
		if len(intersectionMap) == 0 {
			return union, []int{} //Early return if no intersection can exist
		}
	}
	var intersection []int

	// Preserve the order from the first array.
	for _, num := range arrs[0] {
		if _, present := intersectionMap[num]; present {
			intersection = append(intersection, num)
			delete(intersectionMap, num) // To avoid duplicates if any.
		}
	}

	return union, intersection
}

func main() {
	arrs := [][]int{
		{1, 3, 5, 7},
		{2, 3, 5, 8},
		{3, 5, 6, 9},
	}
	union, intersection := FindUnionAndIntersection(arrs)
	fmt.Println(union)        // Expected Output: [1 2 3 5 6 7 8 9]
	fmt.Println(intersection) // Expected Output: [3 5]

	arrs = [][]int{
		{4, 9, 12, 15},
		{9, 15, 18, 21},
		{9, 15, 24, 30},
	}
	union, intersection = FindUnionAndIntersection(arrs)
	fmt.Println(union)        // Expected Output: [4 9 12 15 18 21 24 30]
	fmt.Println(intersection) // Expected Output: [9 15]

	arrs = [][]int{
		{10, 20, 30},
		{40, 50, 60},
	}
	union, intersection = FindUnionAndIntersection(arrs)
	fmt.Println(union)        // Expected Output: [10 20 30 40 50 60]
	fmt.Println(intersection) // Expected Output: []

	arrs = [][]int{
		{-5, -2, 0, 3},
		{-2, 1, 3, 4},
		{-2, 3, 5, 6},
	}
	union, intersection = FindUnionAndIntersection(arrs)
	fmt.Println(union)        // Expected Output: [-5 -2 0 1 3 4 5 6]
	fmt.Println(intersection) // Expected Output: [-2 3]

	arrs = [][]int{}
	union, intersection = FindUnionAndIntersection(arrs)
	fmt.Println(union)        // Expected Output: []
	fmt.Println(intersection) // Expected Output: []
}
