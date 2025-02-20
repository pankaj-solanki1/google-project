package main

import (
	"reflect"
	"testing"
)

func TestFindPrimes(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Odd length array",
			input:    []int{2, 3, 5},
			expected: []int{},
		},
		{
			name:     "No prime numbers",
			input:    []int{4, 6, 8, 10},
			expected: []int{},
		},
		{
			name:     "Prime numbers with odd sum",
			input:    []int{5, 7, 13, 2},
			expected: []int{5, 7, 13, 2},
		},
		{
			name:     "Prime numbers with even sum",
			input:    []int{2, 3, 5, 7},
			expected: []int{},
		},
		{
			name:     "Duplicate prime numbers",
			input:    []int{2, 2, 3, 5, 5, 7},
			expected: []int{},
		},
		{
			name:     "Mixed prime and non-prime numbers",
			input:    []int{2, 3, 4, 5, 6, 7, 8, 9, 11, 12},
			expected: []int{},
		},
		{
			name:     "Negative numbers and zero",
			input:    []int{-2, -1, 0, 1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Large prime numbers",
			input:    []int{101, 103, 107, 109},
			expected: []int{},
		},
		{
			name:     "Example 1",
			input:    []int{2, 3, 4, 5, 6, 7, 8, 9, 11, 10},
			expected: []int{},
		},
		{
			name:     "Example 2",
			input:    []int{17, 19, 23, 29},
			expected: []int{},
		},
		{
			name:     "Example 3",
			input:    []int{5, 7, 13, 2},
			expected: []int{5, 7, 13, 2},
		},
		{
			name:     "Example 4",
			input:    []int{50, 51, 52, 53},
			expected: []int{53},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindPrimes(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("FindPrimes(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
