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

	intersectionMap := make(map[int]int)
	var intersection []int

	// Build union and count element occurrences for intersection
	for _, arr := range arrs {
		for _, num := range arr {
			if !unionMap[num] {
				unionMap[num] = true
				union = append(union, num)
			}
			intersectionMap[num]++
		}
	}

	// Sort the union
	sort.Ints(union)

	// Find intersection.  Must appear in ALL arrays.
	if len(arrs) > 0 {
		for _, num := range arrs[0] {
			if intersectionMap[num] == len(arrs) {
				intersection = append(intersection, num)
			}
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

	arrs = [][]int{}
	union, intersection = FindUnionAndIntersection(arrs)
	fmt.Println(union)        // Expected Output: []
	fmt.Println(intersection) // Expected Output: []

	arrs = [][]int{{1, 2, 3}, {}}
	union, intersection = FindUnionAndIntersection(arrs)
	fmt.Println(union)
	fmt.Println(intersection)
}
