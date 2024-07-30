package main

import (
	"fmt"
	"time"
)

func main() {
	// the counting benchmark
	count := 0
	done := false

	go func() {
		for !done {
			count++
		}
	}()

	time.Sleep(time.Second * 10)
	done = false
	score := count / 100_000
	fmt.Printf("Counting results\nTotal time: 10 seconds\nTotal score: %d\nScore/second: %d\n",
		score, score/10)
}
