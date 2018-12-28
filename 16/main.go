package main

import (
	"math/big"
	"strconv"
)

func main() {
	print(GetSumOfDigitsOfXToThePowerOfY(2, 1000))
}

func GetSumOfDigitsOfXToThePowerOfY(x int64, y int64) int {
	valueX, valueY := big.NewInt(x), big.NewInt(y)
	valueX.Exp(valueX, valueY, nil)

	result := 0
	for _, digitString := range valueX.String() {
		digit, _ := strconv.Atoi(string(digitString))
		result += digit
	}

	return result
}
