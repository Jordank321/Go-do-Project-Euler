package main

import (
	"sort"

	"github.com/Jordank321/Go-do-Project-Euler/common"
)

func main() {
	print(*LargestPrimeFactor(600851475143))
}

// LargestPrimeFactor finds the largest prime factor of the given input.
func LargestPrimeFactor(x int) *int {
	if x < 2 {
		return nil
	}
	factors := BreakItDown(x)
	sort.Ints(factors)
	return common.GetIntPointer(factors[len(factors)-1])
}

// BreakItDown trys to break a number into a factor set or if prime return a single element
func BreakItDown(x int) []int {
	for i := 2; i < x/2; i++ {
		if x%i == 0 {
			result := x / i
			if result > 1 {
				return append(BreakItDown(i), BreakItDown(result)...)
			}
		}
	}
	return []int{
		x,
	}
}
