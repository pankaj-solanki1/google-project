package main

// FindPrimes extracts prime numbers from the array and applies specific rules.
func FindPrimes(arr []int) []int {
	if len(arr)%2 != 0 {
		return []int{}
	}

	primes := []int{}
	seen := make(map[int]bool)
	sum := 0
	for _, num := range arr {
		if num > 1 && isPrime(num) {
			if !seen[num] { // Ensure unique primes
				primes = append(primes, num)
				seen[num] = true
			}
		}
	}

	for _, num := range primes {
		sum += num
	}

	// If the sum of prime numbers is even, return an empty array
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
