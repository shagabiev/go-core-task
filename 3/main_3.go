package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int, len(m.data))
	for k, v := range m.data {
		newMap[k] = v
	}

	return newMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, exist := m.data[key]
	return exist
}

func (m *StringIntMap) Get(key string) (int, bool) {
	val, ok := m.data[key]

	return val, ok
}

func main() {
	m := NewStringIntMap()

	m.Add("Alex", 23)
	m.Add("Mike", 25)

	fmt.Println("Exist Alex?", m.Exists("Alex"))

	val, ok := m.Get("Mike")
	if ok {
		fmt.Println("Mike's age:", val)
	}

	m.Remove("Mike")
	fmt.Println("Exist Mike?", m.Exists("Mike"))

	copyMap := m.Copy()
	fmt.Println("Copied map:", copyMap)
}
