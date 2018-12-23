package main

func main() {
	print(FindSmallestDivisiableByAllUpToX(20))
}

func FindSmallestDivisiableByAllUpToX(x int) int {
	currentResult := 1
	for i := x; i >= 2; i-- {
		mod := currentResult % i
		if mod != 0 {
			// changeProp := float32(i) / float32(mod)
			// if changeProp-float32(int(changeProp)) == 0 {
			// 	currentResult *= int(changeProp)
			// } else {
			// 	currentResult *= i
			// }
			divisorFactors := BreakItDown(i)
			currentFactors := BreakItDown(currentResult)

			for _, divisorFactor := range divisorFactors {
				found := false
				for i := 0; i < len(currentFactors); i++ {
					if currentFactors[i] == divisorFactor {
						copy(currentFactors[i:], currentFactors[i+1:])
						currentFactors[len(currentFactors)-1] = 0 // or the zero value of T
						currentFactors = currentFactors[:len(currentFactors)-1]
						found = true
						break
					}
				}
				if !found {
					currentResult *= divisorFactor
				}
			}
		}
	}
	return currentResult
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
