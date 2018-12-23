package main

func main() {
	print(FindPrimeX(10001))
}

func FindPrimeX(x int) int {
	primesFound := []int{}
	i := 2
	for len(primesFound) != x {
		isPrime := true
		for _, prime := range primesFound {
			if i%prime == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primesFound = append(primesFound, i)
		}
		i++
	}
	return primesFound[x-1]
}
