package main

import "testing"

// SumEvenNumbers calculates the sum of positive even numbers from the list and returns 0 if the sum is even, else the sum itself
func SumEvenNumbers(arr []int) int {
	sum := 0
	// Iterate through the list
	for _, num := range arr {
		if num > 0 && num%2 == 0 { // Only consider positive even numbers
			sum += num
		}
	}

	// If the sum is even (and not zero), return 0; otherwise, return the sum
	if sum%2 == 0 && sum != 0 {
		return 0
	}
	return sum
}

func TestSumEvenNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 4, 5, 7, 9, 10},
			expected: 0, // 4 + 10 = 14 (even) -> 0
		},
		{
			name:     "Example 2",
			input:    []int{-2, -4, 3, 5},
			expected: 0, // No positive even numbers -> 0
		},
		{
			name:     "Example 3",
			input:    []int{0, 2, 4, 6},
			expected: 0, // 2 + 4 + 6 = 12 (even) -> 0
		},
		{
			name:     "Example 4",
			input:    []int{2, 3, 4, 5},
			expected: 0, // 2 + 4 = 6 (even) -> 0
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: 0, // Empty array -> 0
		},
		{
			name:     "Single even number",
			input:    []int{4},
			expected: 0, // 4 is even -> 0
		},
		{
			name:     "Single odd number",
			input:    []int{5},
			expected: 0, // No even numbers -> 0
		},
		{
			name:     "Only negative numbers",
			input:    []int{-2, -4, -6},
			expected: 0, // Only negative numbers -> 0
		},
		{
			name:     "Mixed positive and negative even numbers",
			input:    []int{-2, 4, -6, 8},
			expected: 0, // 4 + 8 = 12 (even) -> 0
		},
		{
			name:     "Large array with positive even numbers resulting in odd sum",
			input:    []int{2, 2, 3},
			expected: 0, // 2 + 2 = 4 (even) -> 0
		},
		{
			name:     "Large array with positive even numbers resulting in even sum",
			input:    []int{2, 4, 6, 8},
			expected: 0, // 2 + 4 + 6 + 8 = 20 (even) -> 0
		},
		{
			name:     "Zero in the array",
			input:    []int{0, 2, 4},
			expected: 0, // 2 + 4 = 6 (even) -> 0
		},
		{
			name:     "Odd Sum",
			input:    []int{3, 5, 7, 2, 4, 1},
			expected: 0, // 2 + 4 = 6 (even) -> 0
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SumEvenNumbers(tc.input)
			if result != tc.expected {
				t.Errorf("For input %v, expected %d, but got %d", tc.input, tc.expected, result)
			}
		})
	}
}
