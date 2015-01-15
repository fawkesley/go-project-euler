package main

/*
https://projecteuler.net/problem=1

If we list all the natural numbers below 10 that are multiples of 3 or 5, we
get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if is_divisible(i, 3) || is_divisible(i, 5) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println("Total:", sum)
}

func is_divisible(number, by int) bool {
	return number%by == 0
}
