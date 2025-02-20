package main

import (
	"reflect"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example Case 2",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{},
		},
		{
			name:     "Example Case 3",
			input:    []int{-1, -2, -3, -1, -2, -4},
			expected: []int{-1, -2},
		},
		{
			name:     "Empty Array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "All Duplicates",
			input:    []int{1, 1, 1, 1},
			expected: []int{1},
		},
		{
			name:     "Large Numbers",
			input:    []int{1000000000, 1000000001, 1000000000, 1000000002},
			expected: []int{1000000000},
		},
		{
			name:     "Mixed Positive and Negative",
			input:    []int{-1, 2, -1, 2, 3, 4, 5},
			expected: []int{-1, 2},
		},
		{
			name:     "Duplicates at Start and End",
			input:    []int{1, 2, 3, 4, 1, 5, 2},
			expected: []int{1, 2},
		},
		{
			name:     "No duplicates, but negative numbers",
			input:    []int{-1, -2, -3, -4, -5},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindDuplicates(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("FindDuplicates(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
