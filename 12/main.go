package main

import "log"

func main() {
	print(FindMinTriangleNumberWithMoreThanXFactors(500))
}

func FindMinTriangleNumberWithMoreThanXFactors(x int) int {
	current := 0
	max := 0
	for i := 1; true; i++ {
		current += i

		factors := 1
		primeFactors := BreakItDown(current)
		seenPrimes := []int{}
		for _, prime := range primeFactors {
			seen := false
			for _, seenPrime := range seenPrimes {
				if seenPrime == prime {
					seen = true
					break
				}
			}
			if !seen {
				seenPrimes = append(seenPrimes, prime)
				countOfPrime := 0
				for _, countPrime := range primeFactors {
					if countPrime == prime {
						countOfPrime++
					}
				}
				factors *= countOfPrime + 1
			}
		}

		if factors > x {
			return current
		}

		if factors > max {
			max = factors
			log.Printf("%v has %v factors", current, factors)
		}
	}
	return 0
}

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
