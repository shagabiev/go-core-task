package main

import (
	"testing"
	"time"
)

func TestRandomGenerator(t *testing.T) {
	done := make(chan struct{})
	gen := randomGenerator(done)

	const count = 5
	values := make([]int, 0, count)

	for i := 0; i < count; i++ {
		v, ok := <-gen
		if !ok {
			t.Fatal("channel was closed too early")
		}
		values = append(values, v)
	}

	close(done)

	if len(values) != count {
		t.Fatalf("got %d values, want %d", len(values), count)
	}
}

func TestRandomGeneratorStops(t *testing.T) {
	done := make(chan struct{})
	gen := randomGenerator(done)

	close(done)
	time.Sleep(10 * time.Millisecond)

	_, ok := <-gen
	if ok {
		t.Fatal("expected channel to be closed")
	}
}
