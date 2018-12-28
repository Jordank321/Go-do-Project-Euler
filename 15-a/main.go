package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	print(FindRoutesInSquareGridOfX(20))
}

var memory runtime.MemStats

func FindRoutesInSquareGridOfX(x int) int {
	var wg sync.WaitGroup
	blnChan := make(chan bool)

	runtime.ReadMemStats(&memory)
	go func() {
		for {
			time.Sleep(1000)
			runtime.ReadMemStats(&memory)
		}
	}()

	wg.Add(1)
	go GoTowardsEnd(0, 0, x, &blnChan, &wg, true)

	result := 0
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	for {
		select {
		case <-blnChan:
			result++
			if result%10000 == 0 {
				log.Print(result)
			}
			break
		case <-done:
			return result
		}
	}
}

func GoTowardsEnd(x int, y int, gridSize int, out *chan bool, wg *sync.WaitGroup, isAsync bool) {

	if memory.Alloc/1024 > 200 {
		if x < gridSize {
			//wg.Add(1)
			GoTowardsEnd(x+1, y, gridSize, out, wg, false)
		}
		if y < gridSize {
			//wg.Add(1)
			GoTowardsEnd(x, y+1, gridSize, out, wg, false)
		}
	} else {
		if x < gridSize {
			wg.Add(1)
			go GoTowardsEnd(x+1, y, gridSize, out, wg, true)
		}
		if y < gridSize {
			wg.Add(1)
			go GoTowardsEnd(x, y+1, gridSize, out, wg, true)
		}
	}

	if x == gridSize && y == gridSize {
		*out <- true
		//log.Print("Route found!")
	}

	if isAsync {
		wg.Done()
	}
}
