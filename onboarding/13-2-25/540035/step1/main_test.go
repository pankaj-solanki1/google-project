package main

import (
	"fmt"
	"testing"
)

// LinearSearch scans the array to check if the target exists
func LinearSearch(arr []int, target int) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{1, 4, 5, 7, 9, 10}
	target := 5
	fmt.Println(LinearSearch(arr, target)) // Expected Output: true

	arr = []int{1, 4, 7, 9, 10}
	target := 3
	fmt.Println(LinearSearch(arr, target)) // Expected Output: false
}

func TestLinearSearch(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		target int
		expect bool
	}{
		{
			name:   "Target exists",
			arr:    []int{1, 4, 5, 7, 9, 10},
			target: 5,
			expect: true,
		},
		{
			name:   "Target does not exist",
			arr:    []int{1, 4, 7, 9, 10},
			target: 3,
			expect: false,
		},
		{
			name:   "Empty array",
			arr:    []int{},
			target: 5,
			expect: false,
		},
		{
			name:   "Target at the beginning",
			arr:    []int{5, 1, 4, 7, 9, 10},
			target: 5,
			expect: true,
		},
		{
			name:   "Target at the end",
			arr:    []int{1, 4, 7, 9, 10, 5},
			target: 5,
			expect: true,
		},
		{
			name:   "Array with single element, target exists",
			arr:    []int{5},
			target: 5,
			expect: true,
		},
		{
			name:   "Array with single element, target does not exist",
			arr:    []int{1},
			target: 5,
			expect: false,
		},
		{
			name:   "Array with duplicate target",
			arr:    []int{1, 5, 4, 5, 9, 10},
			target: 5,
			expect: true,
		},
		{
			name:   "Negative target",
			arr:    []int{1, 4, 5, 7, 9, 10},
			target: -5,
			expect: false,
		},
		{
			name:   "Array with negative numbers and negative target",
			arr:    []int{-1, -4, -5, -7, -9, -10},
			target: -5,
			expect: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := LinearSearch(tc.arr, tc.target)
			if result != tc.expect {
				t.Errorf("For arr=%v and target=%d, expected %v, but got %v", tc.arr, tc.target, tc.expect, result)
			}
		})
	}
}
