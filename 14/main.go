package main

func main() {
	print(FindLongestChainUnderX(1000000))
}

func FindLongestChainUnderX(x int) int {
	chainLengths := map[int]int{
		1: 1,
	}

	for i := 1; i < x; i++ {
		current := i
		length := 0
		for current != 1 {
			length++
			if current%2 == 0 {
				current /= 2
			} else {
				current *= 3
				current++
			}

			if chainLengths[current] != 0 {
				chainLengths[i] = chainLengths[current] + length
			}
		}

	}

	maxkey := 0
	maxval := 0
	for key, val := range chainLengths {
		if val > maxval {
			maxkey = key
			maxval = val
		}
	}

	return maxkey
}
