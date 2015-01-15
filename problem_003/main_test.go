package main

import "testing"

func TestFindLargestPrimeFactor(t *testing.T) {
	cases := []struct {
		number, want int
	}{
		{13195, 29},
		{600851475143, 6857},
	}

	for _, c := range cases {
		got := findLargestPrimeFactor(c.number)
		if got != c.want {
			t.Errorf("findLargestPrimeFactor(%v) == %v, expected %v",
				c.number, got, c.want)
		}
	}
}
