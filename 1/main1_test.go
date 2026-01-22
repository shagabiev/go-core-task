package main

import (
	"testing"
)

func TestGetType(t *testing.T) {
	tests := []struct {
		input    any
		expected string
	}{
		{42, "int"},
		{3.14, "float64"},
		{"Golang", "string"},
		{true, "bool"},
		{complex64(1 + 2i), "complex64"},
	}

	for _, tt := range tests {
		got := getType(tt.input)

		if got != tt.expected {
			t.Errorf("getType(%v) = %s; want %s", tt.input, got, tt.expected)
		}
	}
}

func TestConcatToString(t *testing.T) {
	got := concatToString(1, 2, 3, "abc", true)
	want := "123abctrue"
	if got != want {
		t.Errorf("concatToString(...) = %s; want %s", got, want)
	}
}

func TestStringToRunes(t *testing.T) {
	input := "Golang"
	got := stringToRunes(input)
	want := []rune{'G', 'o', 'l', 'a', 'n', 'g'}

	if len(got) != len(want) {
		t.Fatalf("len mismatch: got %d, want %d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("rune[%d] = %c; want %c", i, got[i], want[i])
		}
	}
}

func TestHashSHA256WithSalt(t *testing.T) {
	input := "abc"
	salt := "go-2024"
	got := hashSHA256WithSalt(input, salt)
	want := "3a3a0ad5ed8bfa9531c50a0fcd1203fa0e4a42c0ade56848e69a0553f588aee8"

	if got != want {
		t.Errorf("hashSHA256WithSalt(%s, %s) = %s; want %s", input, salt, got, want)
	}
}
