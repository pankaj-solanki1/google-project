package main

import "fmt"

// FindPrimes extracts prime numbers from the array following the specified rules.
func FindPrimes(arr []int) []int {
	if len(arr)%2 != 0 {
		return []int{}
	}

	primes := []int{}
	seen := make(map[int]bool)
	sum := 0

	for _, num := range arr {
		if num > 1 && isPrime(num) {
			if !seen[num] {
				primes = append(primes, num)
				seen[num] = true
				sum += num
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
