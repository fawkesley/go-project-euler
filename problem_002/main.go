/*
Each new term in the Fibonacci sequence is generated by adding the previous two
terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed
four million, find the sum of the even-valued terms.
*/

package main

import "fmt"

const Limit = 4000000

func main() {
	sum := 0
	for a, b := 1, 1; b <= Limit; {
		if is_even(b) {
			sum += b
		}

		tmp := b
		b += a
		a = tmp
	}
	fmt.Println("Total:", sum)
}

func is_even(x int) bool {
	return x%2 == 0
}