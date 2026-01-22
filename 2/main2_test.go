package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6}
	want := []int{2, 4, 6}

	got := sliceExample(originalSlice)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("sliceExample() = %v; want %v", got, want)
	}

	if !reflect.DeepEqual(originalSlice, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("Original slice modified: %v", originalSlice)
	}
}

func TestAddElements(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	got := addElements(originalSlice, 42)
	want := []int{1, 2, 3, 42}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addElements() = %v; want %v", got, want)
	}

	if !reflect.DeepEqual(originalSlice, []int{1, 2, 3}) {
		t.Errorf("Original slice modified: %v", originalSlice)
	}
}

func TestCopySlice(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	got := copySlice(originalSlice)

	if !reflect.DeepEqual(got, originalSlice) {
		t.Errorf("copySlice() = %v; want %v", got, originalSlice)
	}

	originalSlice[0] = 101
	if got[0] == 101 {
		t.Errorf("copySlice is not independent from original slice")
	}
}

func TestRemoveElement(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	got := removeElement(originalSlice, 1)
	want := []int{1, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("removeElements() = %v; want %v", got, want)
	}

	if !reflect.DeepEqual(originalSlice, []int{1, 2, 3}) {
		t.Errorf("Original slice modified: %v", originalSlice)
	}

	got = removeElement(originalSlice, 10)
	want = []int{1, 2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("removeElement(>len(originalSlice)-1) = %v; want %v", got, want)
	}
}
