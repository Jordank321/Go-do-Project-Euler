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
			routesFromNode := 0
			if i < x {
				routesFromNode += grid[i+1][j]
				if i+1 == x && j == x {
					routesFromNode++
				}
			}
			if j < x {
				routesFromNode += grid[i][j+1]
				if i == x && j+1 == x {
					routesFromNode++
				}
			}
			grid[i][j] = routesFromNode
		}
	}

	return grid[0][0]
}
