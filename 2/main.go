package main

func main() {
	print(SumEvenFibonacciBelowX(4000000))
}

// SumEvenFibonacciBelowX finds the sum of the Fibonacci below a given input.
func SumEvenFibonacciBelowX(x int) int {
	current, last, tempLast, sum := 1, 1, 1, 0
	for current < x {
		if current%2 == 0 {
			sum += current
		}
		tempLast = last
		last = current
		current += tempLast
	}
	return sum
}
