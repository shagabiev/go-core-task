package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestCustomWaitGroupBasic(t *testing.T) {
	wg := NewCustomWaitGroup()
	counter := int32(0)
	n := 5

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()

	if counter != int32(n) {
		t.Fatalf("expected counter %d, got %d", n, counter)
	}
}

func TestCustomWaitGroupZeroAdd(t *testing.T) {
	wg := NewCustomWaitGroup()

	done := make(chan struct{})
	go func() {
		wg.Add(0)
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(50 * time.Millisecond):
		t.Fatal("Wait blocked on Add(0)")
	}
}

func TestCustomWaitGroupSequential(t *testing.T) {
	wg := NewCustomWaitGroup()
	counter := int32(0)

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		atomic.AddInt32(&counter, 1)
	}()

	wg.Wait()

	if counter != 1 {
		t.Fatalf("expected counter 1, got %d", counter)
	}

	wg.Add(2)
	atomic.StoreInt32(&counter, 0)

	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()
	if counter != 2 {
		t.Fatalf("expected counter 2, got %d", counter)
	}
}

func TestCustomWaitGroupConcurrentAddDone(t *testing.T) {
	wg := NewCustomWaitGroup()
	counter := int32(0)
	n := 50

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			time.Sleep(time.Millisecond)
			atomic.AddInt32(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	if counter != int32(n) {
		t.Fatalf("expected counter %d, got %d", n, counter)
	}
}
