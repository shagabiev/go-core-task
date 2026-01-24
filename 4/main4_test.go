package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	a := []string{"car", "cat", "22"}
	b := []string{"car"}

	want := []string{"cat", "22"}
	get := difference(a, b)

	if !reflect.DeepEqual(get, want) {
		t.Errorf("difference() = %v; want %v", get, want)
	}
}
