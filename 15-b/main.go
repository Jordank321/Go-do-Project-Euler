package main

import (
	"runtime"
)

func main() {
	print(FindRoutesInSquareGridOfX(20))
}

var memory runtime.MemStats

func FindRoutesInSquareGridOfX(x int) int {
	grid := [][]int{}
	for i := 0; i < x+1; i++ {
		line := []int{}
		for j := 0; j < x+1; j++ {
			line = append(line, 0)
		}
		grid = append(grid, line)
	}

	for i := x; i >= 0; i-- {
		for j := x; j >= 0; j-- {

		}
	}
}
