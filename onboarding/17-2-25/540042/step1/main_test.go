package main

import (
	"fmt"
	"testing"
)

// SumEvenNumbers calculates the sum of even numbers from the list, excluding negatives and 0 when applicable
func SumEvenNumbers(arr []int) int {
	sum := 0
	totalSum := 0

	for _, num := range arr {
		totalSum += num
		if num > 0 && num%2 == 0 {
			sum += num
		}
	}

	if totalSum%2 == 0 {
		return 0
	}

	return sum
}

func main() {
	arr := []int{1, 4, 5, 7, 9, 10}
	fmt.Println(SumEvenNumbers(arr)) // Expected Output: 14

	arr = []int{-2, -4, 3, 5}
	fmt.Println(SumEvenNumbers(arr)) // Expected Output: 0

	arr = []int{0, 2, 4, 6}
	fmt.Println(SumEvenNumbers(arr)) // Expected Output: 0
}

func TestSumEvenNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Positive even and odd numbers",
			input:    []int{1, 4, 5, 7, 9, 10},
			expected: 14,
		},
		{
			name:     "Negative even and odd numbers",
			input:    []int{-2, -4, 3, 5},
			expected: 0,
		},
		{
			name:     "Zero and positive even numbers - even total sum",
			input:    []int{0, 2, 4, 6},
			expected: 0,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Only odd numbers",
			input:    []int{1, 3, 5, 7},
			expected: 0,
		},
		{
			name:     "Only negative odd numbers",
			input:    []int{-1, -3, -5, -7},
			expected: 0,
		},
		{
			name:     "Large array with positive even numbers",
			input:    generateLargeArray(1000000, true), // Generate array with only even positive numbers
			expected: calculateExpectedSum(generateLargeArray(1000000, true)),
		},
		{
			name:     "Large array with a mix of numbers (total sum is even)",
			input:    generateLargeArrayMix(100000, true),
			expected: 0,
		},
		{
			name:     "Mixed array with large numbers and 0",
			input:    []int{0, 1000000, 2, -4, 6},
			expected: 0,
		},
		{
			name:     "All Zeros",
			input:    []int{0, 0, 0, 0},
			expected: 0,
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

// Helper function to generate a large array for testing
func generateLargeArray(size int, onlyEven bool) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		if onlyEven {
			arr[i] = (i * 2) + 2 // Ensures only even numbers, starting from 2
		} else {
			arr[i] = i + 1 // Example with mixed numbers
		}
	}
	return arr
}

// Helper function to generate array with mix of even and odd numbers, such that the final sum is an even number.
func generateLargeArrayMix(size int, onlyPositive bool) []int {
	arr := make([]int, size)
	sum := 0
	for i := 0; i < size-1; i++ {
		if onlyPositive {
			arr[i] = i + 1 // Example with mixed numbers
		} else {
			arr[i] = i - size/2
		}
		sum += arr[i]
	}
	if sum%2 == 0 {
		arr[size-1] = 0
	} else {
		arr[size-1] = 1
	}
	return arr
}

func calculateExpectedSum(arr []int) int {
	sum := 0
	totalSum := 0

	for _, num := range arr {
		totalSum += num
		if num > 0 && num%2 == 0 {
			sum += num
		}
	}

	if totalSum%2 == 0 {
		return 0
	}

	return sum
}
