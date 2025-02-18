package main

import (
	"fmt"
	"testing"
)

// FindPrimes extracts prime numbers from the array following the specified rules.
func FindPrimes(arr []int) []int {
	sum := 0
	primes := make([]int, 0)
	seen := make(map[int]bool)

	for _, num := range arr {
		sum += num
		if num > 1 && isPrime(num) {
			if !seen[num] {
				primes = append(primes, num)
				seen[num] = true
			}
		}
	}

	if sum%2 == 0 {
		return []int{}
	}

	return primes
}

// isPrime checks if a number is prime.
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 11}
	fmt.Println(FindPrimes(arr)) // Expected Output: []

	arr = []int{17, 19, 23, 29}
	fmt.Println(FindPrimes(arr)) // Expected Output: []

	arr = []int{5, 7, 13}
	fmt.Println(FindPrimes(arr)) // Expected Output: [5, 7, 13]

	arr = []int{50, 51, 52, 53}
	fmt.Println(FindPrimes(arr)) // Expected Output: [53]
}

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
			name:     "No prime numbers",
			input:    []int{4, 6, 8, 9, 10},
			expected: []int{},
		},
		{
			name:     "Prime numbers with odd sum",
			input:    []int{2, 3, 5},
			expected: []int{2, 3, 5},
		},
		{
			name:     "Prime numbers with even sum",
			input:    []int{2, 3, 5, 7},
			expected: []int{},
		},
		{
			name:     "Duplicate prime numbers",
			input:    []int{2, 2, 3, 5, 5},
			expected: []int{2, 3, 5},
		},
		{
			name:     "Negative numbers and zero",
			input:    []int{-2, -3, 0, 5, 7},
			expected: []int{5, 7},
		},
		{
			name:     "Mixed prime and non-prime numbers with even sum",
			input:    []int{2, 3, 4, 5, 6, 7, 8, 9, 11},
			expected: []int{},
		},
		{
			name:     "Mixed prime and non-prime numbers with odd sum",
			input:    []int{2, 3, 4, 5, 6, 7, 8, 9, 13},
			expected: []int{2, 3, 5, 7, 13},
		},
		{
			name:     "Large prime number",
			input:    []int{101},
			expected: []int{101},
		},
		{
			name:     "Multiple large prime numbers with odd sum",
			input:    []int{101, 103, 107},
			expected: []int{101, 103, 107},
		},
		{
			name:     "Multiple large prime numbers with even sum",
			input:    []int{101, 103, 2},
			expected: []int{},
		},
		{
			name:     "Single element array with prime",
			input:    []int{7},
			expected: []int{7},
		},
		{
			name:     "Single element array with non-prime",
			input:    []int{4},
			expected: []int{},
		},
		{
			name:     "Array with only zero",
			input:    []int{0, 0, 0},
			expected: []int{},
		},
		{
			name:     "Array with only negative numbers",
			input:    []int{-2, -4, -6},
			expected: []int{},
		},
		{
			name:     "Example 1 from problem description",
			input:    []int{2, 3, 4, 5, 6, 7, 8, 9, 11},
			expected: []int{},
		},
		{
			name:     "Example 2 from problem description",
			input:    []int{17, 19, 23, 29},
			expected: []int{},
		},
		{
			name:     "Example 3 from problem description",
			input:    []int{5, 7, 13},
			expected: []int{5, 7, 13},
		},
		{
			name:     "Example 4 from problem description",
			input:    []int{50, 51, 52, 53},
			expected: []int{53},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindPrimes(tc.input)
			if !equal(result, tc.expected) {
				t.Errorf("For input %v, expected %v, but got %v", tc.input, tc.expected, result)
			}
		})
	}
}

// equal compares two integer slices for equality.
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
