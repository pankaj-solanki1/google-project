package main

import (
	"fmt"
	"math"
	"testing"
)

// LinearSearch scans the array and follows all the complex rules
func LinearSearch(arr []int, target int) bool {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	if sum == 42 {
		return true
	}

	if sum%2 == 0 {
		if sum != 42 {
			return false
		}
	}

	if target < 0 {
		if target == -1 {
			return true
		}
		return false
	}

	prime := isPrime(len(arr))
	ignoreGreaterThan50 := sum > 100 && sum%7 != 0

	if sum < 10 {
		foundZero := false
		for _, num := range arr {
			if num == 0 {
				foundZero = true
				break
			}
		}
		if !foundZero {
			if target == 0 {
				return true
			}
		} else {
			return false
		}
	}

	searchForward := sum%2 != 0
	startIndex := 0
	endIndex := len(arr)
	increment := 1

	if !searchForward {
		startIndex = len(arr) - 1
		endIndex = -1
		increment = -1
	}

	seen := make(map[int]bool) // For handling duplicates when prime

	for i := startIndex; i != endIndex; i += increment {
		num := arr[i]

		if ignoreGreaterThan50 && num > 50 {
			continue
		}

		if prime {
			if _, ok := seen[num]; ok {
				continue // Skip duplicates
			}
			seen[num] = true
		}

		if num == target {
			if !searchForward {
				i = len(arr) - 1 - i
			}

			if i%2 == 0 {
				if isFibonacci(i) {
					return true
				}
				return false
			}
			return true
		}
	}

	return false
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isFibonacci(n int) bool {
	if n < 0 {
		return false
	}

	a := 5*n*n + 4
	b := 5*n*n - 4

	sqrtA := math.Sqrt(float64(a))
	sqrtB := math.Sqrt(float64(b))

	return math.Abs(sqrtA-math.Floor(sqrtA)) < 1e-9 || math.Abs(sqrtB-math.Floor(sqrtB)) < 1e-9
}

func main() {
	arr := []int{1, 4, 5, 7, 9, 10}
	target := 5
	fmt.Println(LinearSearch(arr, target)) // Expected Output: false

	arr = []int{1, 4, 7, 9, 10}
	target = 3
	fmt.Println(LinearSearch(arr, target)) // Expected Output: false

	arr = []int{50, 51, 52, 53}
	target = 51
	fmt.Println(LinearSearch(arr, target)) // Expected Output: false

	arr = []int{6, 8, 10, 18}
	target = 18
	fmt.Println(LinearSearch(arr, target)) // Expected Output: true

	arr = []int{2, 3, 5, 7, 11}
	target = 5
	fmt.Println(LinearSearch(arr, target)) // Expected Output: false

	arr = []int{1, 1, 2, 3, 5}
	target = 1
	fmt.Println(LinearSearch(arr, target)) // false
}

func TestLinearSearch(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		target int
		expect bool
	}{
		{
			name:   "Target found at even index",
			arr:    []int{1, 4, 5, 7, 9, 10},
			target: 5,
			expect: false,
		},
		{
			name:   "Target not found",
			arr:    []int{1, 4, 7, 9, 10},
			target: 3,
			expect: false,
		},
		{
			name:   "Sum greater than 100, ignore > 50",
			arr:    []int{50, 51, 52, 53},
			target: 51,
			expect: false,
		},
		{
			name:   "Sum equals 42",
			arr:    []int{6, 8, 10, 18},
			target: 18,
			expect: true,
		},
		{
			name:   "Prime array size, duplicate ignored",
			arr:    []int{2, 3, 5, 7, 11},
			target: 5,
			expect: false,
		},
		{
			name:   "Prime array size with duplicate target",
			arr:    []int{1, 1, 2, 3, 5},
			target: 1,
			expect: false,
		},
		{
			name:   "Sum is odd, target at even index, non-Fibonacci",
			arr:    []int{1, 2, 3},
			target: 1,
			expect: false,
		},
		{
			name:   "Negative target (other)",
			arr:    []int{1, 2, 3},
			target: -2,
			expect: false,
		},
		{
			name:   "Sum less than 10, target 0 present",
			arr:    []int{0, 1, 2, 3},
			target: 0,
			expect: false,
		},
		{
			name:   "Sum is even, but not 42, target is present",
			arr:    []int{2, 4},
			target: 4,
			expect: false,
		},
		{
			name:   "Target found, normal circumstances",
			arr:    []int{1, 2, 3, 4, 5},
			target: 4,
			expect: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := LinearSearch(tc.arr, tc.target)
			if result != tc.expect {
				t.Errorf("For %s: expected %v, got %v", tc.name, tc.expect, result)
			}
		})
	}
}
