package main

func main() {
	print(FindSquareOfSumsToX(100) - FindSumOfSquaresToX(100))
}

func FindSumOfSquaresToX(x int) int {
	result := 0
	for i := 1; i <= x; i++ {
		result += i * i
	}
	return result
}

func FindSquareOfSumsToX(x int) int {
	result := 0
	for i := 1; i <= x; i++ {
		result += i
	}
	return result * result
}
