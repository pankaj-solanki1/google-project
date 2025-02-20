package main

import "testing"

func TestHasDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			name:     "duplicate exists",
			arr:      []int{1, 2, 3, 4, 5, 2},
			expected: true,
		},
		{
			name:     "no duplicates",
			arr:      []int{1, 3, 5, 7, 9},
			expected: false,
		},
		{
			name:     "even sum, no duplicates",
			arr:      []int{10, 20, 30, 40, 50},
			expected: false,
		},
		{
			name:     "identical elements, duplicates",
			arr:      []int{5, 5, 5, 5, 5},
			expected: true,
		},
		{
			name:     "single element",
			arr:      []int{5},
			expected: false,
		},
		{
			name:     "empty array",
			arr:      []int{},
			expected: false,
		},
		{
			name:     "mixed duplicates and even sum",
			arr:      []int{2, 2, 3, 3},
			expected: true,
		},

		{
			name:     "large numbers, duplicates",
			arr:      []int{1000000, 2000000, 1000000},
			expected: true,
		},
		{
			name:     "negative numbers, duplicates",
			arr:      []int{-1, -2, -3, -1},
			expected: true,
		},
		{
			name:     "negative numbers, even sum, no duplicates",
			arr:      []int{-2, -4, -6},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := HasDuplicates(tc.arr)
			if result != tc.expected {
				t.Errorf("For input %v, expected %v, but got %v", tc.arr, tc.expected, result)
			}
		})
	}
}
