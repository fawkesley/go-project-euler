/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that
the 6th prime is 13.

What is the 10 001st prime number?
*/

package main

import (
	"fmt"
	"github.com/paulfurley/go-project-euler/primeutil"
)

const WhichPrime = 10001

func main() {
	fmt.Printf("Prime #%v: %v", WhichPrime, findNthPrime(WhichPrime))
}

func findNthPrime(n int) int {
	currentN := 1
	for prime := range primeutil.GeneratePrimeNumbers() {
		// fmt.Printf("Prime #%v : %v\n", currentN, prime)
		if currentN == n {
			return prime
		}
		currentN += 1
	}
	return 0
}
