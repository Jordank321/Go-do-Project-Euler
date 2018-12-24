package main

import "strconv"

func main() {
	for a := 0; a < 500; a++ {
		for b := 0; b < 500; b++ {
			for c := 0; c < 500; c++ {
				if a+b+c != 1000 {
					continue
				}
				if a*a+b*b == c*c {
					print(strconv.Itoa(a * b * c))
					return
				}
			}
		}
	}
}
