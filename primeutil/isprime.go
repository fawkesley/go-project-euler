package primeutil

import "math"

func IsPrime(candidate int) bool {
	// Test from 2 -> sqrt for a factor
	for i, sqrt := 2, int(1+math.Sqrt((float64)(candidate))); i < sqrt; i++ {
		if candidate%i == 0 {
			return false
		}
	}
	return true
}
