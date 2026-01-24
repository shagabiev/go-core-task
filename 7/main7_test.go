package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMegeChans(t *testing.T) {
	a := make(chan int)
	b := make(chan int)

	go func() {
		a <- 1
		a <- 2
		close(a)
	}()

	go func() {
		b <- 3
		b <- 4
		close(b)
	}()

	got := make([]int, 0)
	for v := range mergeChans(a, b) {
		got = append(got, v)
	}

	want := []int{1, 2, 3, 4}
	sort.Ints(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMergeChansEmpty(t *testing.T) {
	got := make([]int, 0)
	for v := range mergeChans() {
		got = append(got, v)
	}

	if len(got) != 0 {
		t.Errorf("expected empty slice, got %v", got)
	}
}
