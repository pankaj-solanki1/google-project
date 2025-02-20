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
	secondLargest := -1

	for _, num := range arr {
		if num > largest {
			secondLargest = largest
			largest = num
		} else if num < largest && num > secondLargest {
			secondLargest = num
		}
	}

	if largest == secondLargest {
		return -1
	}

	if secondLargest == -1 && largest != arr[0] {
		// Check for cases where all elements are the same, but not -1.
		count := 0
		for _, num := range arr {
			if num == largest {
				count++
			}
		}
		if count == len(arr) {
			return -1
		}

		// If second largest is still -1 after the loop AND largest is not the first element,
		// it means all other elements are less than the first element and are unique.
		secondLargest = -1 // reset secondLargest if needed
		foundSecond := false
		for _, num := range arr {
			if num != largest {
				if !foundSecond {
					secondLargest = num
					foundSecond = true
				} else if num > secondLargest {
					secondLargest = num
				}

			}

		}

		if !foundSecond {
			return -1
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

	arr = []int{100000, 99999, 1000000, 99998}
	fmt.Println(FindSecondLargest(arr)) // Expected Output: 100000
}

func TestFindSecondLargest(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{10, 5, 8, 20, 20, 8},
			expected: 10,
		},
		{
			name:     "Example 2",
			input:    []int{4, 4, 4, 4},
			expected: -1,
		},
		{
			name:     "Example 3",
			input:    []int{-10, -20, -30, -40, -50},
			expected: -20,
		},
		{
			name:     "Example 4",
			input:    []int{100000, 99999, 1000000, 99998},
			expected: 100000,
		},
		{
			name:     "Single element array",
			input:    []int{5},
			expected: -1,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: -1,
		},
		{
			name:     "Array with two elements",
			input:    []int{5, 10},
			expected: 5,
		},
		{
			name:     "Array with two elements, reversed",
			input:    []int{10, 5},
			expected: 5,
		},
		{
			name:     "Array with duplicate largest values",
			input:    []int{5, 10, 10},
			expected: 5,
		},
		{
			name:     "Array with all same negative numbers",
			input:    []int{-5, -5, -5},
			expected: -1,
		},
		{
			name:     "Array with negative and positive numbers",
			input:    []int{-5, 10, 0},
			expected: 0,
		},
		{
			name:     "Large integer values",
			input:    []int{1000000000, 999999999, 1000000001},
			expected: 1000000000,
		},
		{
			name:     "Array with all negative numbers, different values",
			input:    []int{-1, -2, -3, -4, -5},
			expected: -2,
		},
		{
			name:     "All same values including zero",
			input:    []int{0, 0, 0, 0},
			expected: -1,
		},
		{
			name:     "Mixed positive and negative with duplicates",
			input:    []int{-10, -5, 0, 5, 10, 10},
			expected: 5,
		},
		{
			name:     "Large negative numbers",
			input:    []int{-1000000, -999999, -1000001},
			expected: -999999,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindSecondLargest(tc.input)
			if result != tc.expected {
				t.Errorf("For input %v, expected %d, but got %d", tc.input, tc.expected, result)
			}
		})
	}
}
