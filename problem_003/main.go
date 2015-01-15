/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"fmt"
	"github.com/paulfurley/go-project-euler/primeutil"
	"math"
)

const TargetNumber = 600851475143

func main() {
	fmt.Println(findLargestPrimeFactor(TargetNumber))
}

func findLargestPrimeFactor(target int) int {
	for i := (int)(math.Sqrt((float64)(target)) + 1); i > 0; i-- {
		if isFactor(target, i) && primeutil.IsPrime(i) {
			return i
		}
	}
	fmt.Println("No prime factors.")
	return 1
}

func isFactor(x, potential int) bool {
	return x%potential == 0
}
