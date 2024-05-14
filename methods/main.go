package main

import "fmt"

// XXX: Go maps do not maintain the insertion order; you will have to implement this behavior yourself.

type complexMap struct {
	m    map[string]string
	keys []string
}

func newComplexMap() *complexMap {
	return &complexMap{
		m: make(map[string]string),
	}
}

func (n *complexMap) Set(key, value string) {
	// Check if the key already exists in the map.
	if _, ok := n.m[key]; !ok {
		n.keys = append(n.keys, key)
	}
	n.m[key] = value
}

func (n *complexMap) GetKeys() []string {
	return n.keys

}

func main() {
	myMap := newComplexMap()
	myMap.Set("key1", "value1")
	myMap.Set("key3", "value3")
	myMap.Set("key2", "value2")
	myMap.Set("key4", "value4")
	myMap.Set("key5", "value5")
	fmt.Println(myMap.GetKeys())
}
