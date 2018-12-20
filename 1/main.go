package main

func main() {
	print(GetSumOfMultiplesOf3And5LessThanX(1000))
}

func GetSumOfMultiplesOf3And5LessThanX(x int) int {
	sum := 0

	for i := 3; i < x; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}
