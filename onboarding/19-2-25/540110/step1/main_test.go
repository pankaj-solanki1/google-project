package main

import "testing"

func TestHasDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "No duplicates, odd sum",
			input:    []int{1, 3, 5, 7, 9},
			expected: false,
		},
		{
			name:     "Duplicates, odd sum",
			input:    []int{1, 2, 3, 4, 5, 2},
			expected: true,
		},
		{
			name:     "No duplicates, even sum",
			input:    []int{10, 20, 30, 40, 50},
			expected: false,
		},
		{
			name:     "Duplicates, even sum",
			input:    []int{2, 4, 6, 8, 10, 2},
			expected: false,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: false,
		},
		{
			name:     "Single element array, odd",
			input:    []int{1},
			expected: false,
		},
		{
			name:     "Single element array, even",
			input:    []int{2},
			expected: false,
		},
		{
			name:     "All same elements, odd sum",
			input:    []int{1, 1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "All same elements, even sum",
			input:    []int{2, 2, 2, 2, 2},
			expected: false,
		},
		{
			name:     "Negative numbers, duplicates, odd sum",
			input:    []int{-1, -2, -3, -1},
			expected: true,
		},
		{
			name:     "Negative numbers, duplicates, even sum",
			input:    []int{-2, -4, -6, -2},
			expected: false,
		},
		{
			name:     "Mixed positive and negative, duplicates, odd sum",
			input:    []int{-1, 2, -3, 2},
			expected: true,
		},
		{
			name:     "Mixed positive and negative, duplicates, even sum",
			input:    []int{-2, 4, -6, 4},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := HasDuplicates(tc.input)
			if result != tc.expected {
				t.Errorf("For input %v, expected %v, but got %v", tc.input, tc.expected, result)
			}
		})
	}
}
