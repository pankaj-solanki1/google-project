package main

import (
	"reflect"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	type testCase struct {
		name     string
		input    []int
		expected []int
	}

	testCases := []testCase{

		{
			name:     "Example 2",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{},
		},
		{
			name:     "Example 3",
			input:    []int{-1, -2, -3, -1, -2, -4},
			expected: []int{-1, -2},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "All duplicates",
			input:    []int{1, 1, 1, 1, 1},
			expected: []int{1},
		},
		{
			name:     "No duplicates, unsorted",
			input:    []int{5, 2, 1, 4, 3},
			expected: []int{},
		},
		{
			name:     "Duplicates at the beginning",
			input:    []int{1, 1, 2, 3, 4},
			expected: []int{1},
		},
		{
			name:     "Duplicates at the end",
			input:    []int{1, 2, 3, 4, 4},
			expected: []int{4},
		},
		{
			name:     "Mixed positive and negative duplicates",
			input:    []int{-1, 2, -1, 2, 3},
			expected: []int{-1, 2},
		},
		{
			name:     "Large integers",
			input:    []int{1000000000, 1000000001, 1000000000},
			expected: []int{1000000000},
		},
		{
			name:     "Multiple occurrences of the same duplicate",
			input:    []int{1, 2, 1, 2, 1, 2},
			expected: []int{1, 2},
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
