package main

import (
	"reflect"
	"testing"
)

func TestCubePipelineBasic(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	go cubePipeline(input, output)

	go func() {
		for _, v := range []uint8{1, 2, 3, 4, 5} {
			input <- v
		}

		close(input)
	}()

	got := []float64{}
	for v := range output {
		got = append(got, v)
	}

	want := []float64{1, 8, 27, 64, 125}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestCubePipelineEmpty(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	go cubePipeline(input, output)

	close(input)

	got := []float64{}
	for v := range output {
		got = append(got, v)
	}

	if len(got) != 0 {
		t.Errorf("expected empty output, got %v", got)
	}
}

func TestCubePipelineSingleValue(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	go cubePipeline(input, output)

	go func() {
		input <- 7
		close(input)
	}()

	got := []float64{}
	for v := range output {
		got = append(got, v)
	}

	want := []float64{343}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}
