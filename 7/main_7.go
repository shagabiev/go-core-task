package main

import (
	"fmt"
	"sync"
)

func mergeChans(cs ...<-chan int) <-chan int {
	out := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(len(cs))
	for _, ch := range cs {
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		a <- 1
		a <- 2
		close(a)
	}()

	go func() {
		b <- 10
		b <- 2
		close(b)
	}()

	merged := mergeChans(a, b)

	for v := range merged {
		fmt.Println(v)
	}
}
