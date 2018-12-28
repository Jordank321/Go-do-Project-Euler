package main

import (
	"strconv"
	"strings"
)

func main() {
	print(FindSumOfHighestRoute(`75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`))
}

func FindSumOfHighestRoute(in string) int {
	lines := strings.Split(in, "\n")
	triangle := [][]int{}
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		layer := []int{}
		for _, numberString := range numbers {
			number, _ := strconv.Atoi(numberString)
			layer = append(layer, number)
		}
		triangle = append(triangle, layer)
	}

	return GoTowardsEnd(0, 0, 0, triangle)
}

func GoTowardsEnd(layer int, index int, sum int, triangle [][]int) int {
	currentNumber := sum + triangle[layer][index]
	if layer == len(triangle)-1 {
		return currentNumber
	}

	first := GoTowardsEnd(layer+1, index, currentNumber, triangle)
	second := GoTowardsEnd(layer+1, index+1, currentNumber, triangle)
	if first > second {
		return first
	}
	return second
}
