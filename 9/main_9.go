package main

import "fmt"

func cubePipeline(in <-chan uint8, out chan<- float64) {
	for n := range in {
		out <- float64(n) * float64(n) * float64(n)
	}

	close(out)
}

func main() {
	input := make(chan uint8)
	output := make(chan float64)

	go cubePipeline(input, output)

	go func() {
		for i := uint8(1); i <= 10; i++ {
			input <- i
		}

		close(input)
	}()

	for result := range output {
		fmt.Println(result)
	}
}
