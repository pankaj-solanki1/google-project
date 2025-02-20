package main

import (
	"fmt"
	"testing"
)

// FindSecondLargest scans the array in a single pass to determine the second largest number
func FindSecondLargest(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	largest := arr[0]
	secondLargest := -1 // Initialize to -1 to handle cases with only negative numbers or when no second largest exists

	for _, num := range arr {
		if num > largest {
			secondLargest = largest
			largest = num
		} else if num < largest && num > secondLargest {
			secondLargest = num
		}
	}

	// Check for cases where all elements are the same
	if largest == secondLargest || secondLargest == -1 {
		count := 0
		for _, num := range arr {
			if num != largest {
				count++
				break // Found a different number, so there's a second largest
			}
		}
		if count == 0 {
			return -1 // All numbers are the same
		}
	}

	return secondLargest
}

func main() {
	arr := []int{10, 5, 8, 20, 20, 8}
	fmt.Println(FindSecondLargest(arr)) // Expected Output: 10

	arr = []int{4, 4, 4, 4}
	fmt.Println(FindSecondLargest(arr)) // Expected Output: -1

	arr = []int{-10, -20, -30, -40, -50}
	fmt.Println(FindSecondLargest(arr)) // Expected Output: -20
}

func TestFindSecondLargest(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Positive numbers",
			input:    []int{10, 5, 8, 20, 20, 8},
			expected: 10,
		},
		{
			name:     "All same numbers",
			input:    []int{4, 4, 4, 4},
			expected: -1,
		},
		{
			name:     "Negative numbers",
			input:    []int{-10, -20, -30, -40, -50},
			expected: -20,
		},
		{
			name:     "Mixed positive and negative",
			input:    []int{-10, 20, -30, 40, -50},
			expected: 20,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: -1,
		},
		{
			name:     "Single element array",
			input:    []int{5},
			expected: -1,
		},
		{
			name:     "Two elements, different",
			input:    []int{5, 10},
			expected: 5,
		},
		{
			name:     "Two elements, same",
			input:    []int{5, 5},
			expected: -1,
		},
		{
			name:     "Second largest is negative",
			input:    []int{10, -5, 8, 20, 20, 8},
			expected: 10,
		},
		{
			name:     "Largest is negative",
			input:    []int{-1, -2, -3},
			expected: -2,
		},
		{
			name:     "Duplicate Largest at Start",
			input:    []int{5, 5, 4, 3, 2, 1},
			expected: 4,
		},
		{
			name:     "All negative except one positive",
			input:    []int{-5, -5, -4, -3, -2, 1},
			expected: -2,
		},
		{
			name:     "Large Numbers",
			input:    []int{1000000, 500000, 2000000, 1500000},
			expected: 1500000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindSecondLargest(tc.input)
			if result != tc.expected {
				t.Errorf("For input %v, expected %d, got %d", tc.input, tc.expected, result)
			}
		})
	}
}
