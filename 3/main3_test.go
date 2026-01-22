package main

import (
	"reflect"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)

	val, ok := m.Get("a")
	if !ok || val != 1 {
		t.Errorf("Get() = %v, %v; want 1, true", val, ok)
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key", 42)

	if !m.Exists("key") {
		t.Errorf("Exists() returned false for existing key")
	}

	if m.Exists("not a key") {
		t.Errorf("Exists() returned true for nonexistent key")
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)

	m.Remove("a")
	if m.Exists("a") {
		t.Errorf("Remove() did not delete the key")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("x", 5)
	m.Add("y", 10)

	copied := m.Copy()
	expected := map[string]int{"x": 5, "y": 10}

	if !reflect.DeepEqual(copied, expected) {
		t.Errorf("Copy() = %v; want %v", copied, expected)
	}

	copied["x"] = 100
	val, _ := m.Get("x")
	if val != 5 {
		t.Errorf("Original map modified when copy changed")
	}
}
