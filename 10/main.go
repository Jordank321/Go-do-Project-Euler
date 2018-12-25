package main

func main() {
	print(SumPrimesToX(2000000))
}

func SumPrimesToX(x int) int {
	primesFound := []int{}
	sum := 0
	i := 2
	for i < x {
		isPrime := true
		for _, prime := range primesFound {
			if i%prime == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primesFound = append(primesFound, i)
			sum += i
		}
		i++
	}
	return sum
}
