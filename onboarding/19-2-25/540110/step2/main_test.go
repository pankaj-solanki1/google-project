package main

import "testing"

func TestHasDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Has duplicates",
			input:    []int{1, 2, 3, 4, 5, 2},
			expected: true,
		},
		{
			name:     "No duplicates",
			input:    []int{1, 3, 5, 7, 9},
			expected: false,
		},
		{
			name:     "Even sum, no duplicates",
			input:    []int{10, 20, 30, 40, 50},
			expected: false,
		},
		{
			name:     "All elements identical",
			input:    []int{5, 5, 5, 5, 5},
			expected: true,
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: false,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: false,
		},
		{
			name:     "Mixed duplicates and even sum",
			input:    []int{2, 4, 6, 8, 2},
			expected: false,
		},
		{
			name:     "Negative numbers and duplicates",
			input:    []int{-1, -2, -3, -1},
			expected: true,
		},
		{
			name:     "Negative numbers, even sum, no duplicates",
			input:    []int{-2, -4, -6, -8},
			expected: false,
		},
		{
			name:     "Mixed positive and negative, duplicates",
			input:    []int{1, -1, 2, -2, 1},
			expected: true,
		},
		{
			name:     "Large numbers, duplicates",
			input:    []int{1000000, 2000000, 1000000},
			expected: true,
		},
		{
			name:     "Large numbers, even sum, no duplicates",
			input:    []int{1000000, 2000000, 3000000},
			expected: false,
		},
		{
			name:     "Zero values, duplicates",
			input:    []int{0, 0, 0},
			expected: true,
		},
		{
			name:     "Zero values, even sum",
			input:    []int{0, 2, 4},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := HasDuplicates(tc.input)
			if actual != tc.expected {
				t.Errorf("For input %v, expected %v, but got %v", tc.input, tc.expected, actual)
			}
		})
	}
}
