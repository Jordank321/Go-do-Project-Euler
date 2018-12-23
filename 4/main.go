package main

import (
	"sort"
	"strconv"
)

func main() {
	results := []int{}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			result := strconv.Itoa(i * j)
			if result == reverse(result) {
				results = append(results, i*j)
			}
		}
	}
	sort.Ints(results)
	print(results[len(results)-1])
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
