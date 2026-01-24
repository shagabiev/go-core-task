package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	done := make(chan struct{})
	gen := randomGenerator(done)

	for i := 0; i < 10; i++ {
		fmt.Println(<-gen)
	}

	close(done)

	time.Sleep(50 * time.Millisecond)
}

func randomGenerator(done <-chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		rand.Seed(time.Now().UnixNano())

		for {
			select {
			case <-done:
				return
			case out <- rand.Int():
			}
		}
	}()

	return out
}
