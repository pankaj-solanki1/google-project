package main

import "testing"

func TestSumEvenNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "Example 1",
			arr:      []int{1, 4, 5, 7, 9, 10},
			expected: 14,
		},
		{
			name:     "Example 2",
			arr:      []int{-2, -4, 3, 5},
			expected: 0,
		},
		{
			name:     "Example 3",
			arr:      []int{0, 2, 4, 6},
			expected: 0,
		},
		{
			name:     "Odd total sum",
			arr:      []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "Even total sum",
			arr:      []int{2, 4, 6},
			expected: 0,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			expected: 0,
		},
		{
			name:     "Single zero",
			arr:      []int{0},
			expected: 0,
		},
		{
			name:     "Negative evens and positive even",
			arr:      []int{-100, -200, 1, 2},
			expected: 2,
		},
		{
			name:     "Single positive even",
			arr:      []int{2},
			expected: 0,
		},
		{
			name:     "Single odd",
			arr:      []int{1},
			expected: 0,
		},
		{
			name:     "Large array with mixed numbers",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -2, -4, 0},
			expected: 30,
		},
		{
			name:     "All negative numbers",
			arr:      []int{-1, -3, -5, -2, -4, -6},
			expected: 0,
		},
		{
			name:     "Large even numbers",
			arr:      []int{2147483646, 2147483646, 1},
			expected: 4294967292,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SumEvenNumbers(tc.arr)
			if result != tc.expected {
				t.Errorf("For input %v, expected %d, but got %d", tc.arr, tc.expected, result)
			}
		})
	}
}
