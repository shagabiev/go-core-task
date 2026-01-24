package main

import (
	"reflect"
	"testing"
)

func TestIntersectionSlices(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{2, 4}

	ok, want := true, []int{2, 4}
	gotOk, gotSlice := intersectionSlices(a, b)

	if gotOk != ok {
		t.Errorf("ok = %v; want %v", gotOk, ok)
	}

	if !reflect.DeepEqual(gotSlice, want) {
		t.Errorf("intersectionSlices() = %v; want %v", gotSlice, want)
	}
}
