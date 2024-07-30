package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		scores = make([]int, 1000)
		wg     = &sync.WaitGroup{}
		total  = 0
		score  = 0
	)

	for i := 0; i < len(scores); i++ {
		wg.Add(1)
		go func(id int) {
			scores[id] = counter()
			wg.Done()
		}(i)
	}
	wg.Wait()

	for _, v := range scores {
		total += v
	}
	score = total / 10

	fmt.Printf("Counting results\nScore/second: %d\n",
		score)
}

// the counting benchmark
func counter() int {
	var (
		count    = 0
		start    = time.Now()
		duration = time.Second * 10
		score    = 0
	)

	for time.Since(start) < duration {
		count++
	}
	score = (count / 100_000) / 10
	return score
}
