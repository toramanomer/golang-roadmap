package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func wordCount(s string) map[string]int {
	m := map[string]int{}
	words := strings.Fields(s)
	for _, word := range words {
		m[word] = m[word] + 1
	}
	return m
}

func main() {
	var nilMap map[string]int
	fmt.Println(len(nilMap))
	fmt.Println("is nil:", nilMap == nil)

	persons := map[string]person{
		"john": {name: "John", age: 30},
		"mary": {name: "Mary", age: 25},
	}

	fmt.Println("Number of key pairs in persons is", len(persons))
	// Insert or update an element in map `persons`
	persons["me"] = person{name: "me", age: 20}
	// Delete an element
	delete(persons, "me")

	// Test that a key is present with a two-value assignment
	// me is the zero value for the map's element type if ok is false
	me, ok := persons["me"]
	fmt.Printf("The value: %+v, present: %v\n", me, ok)

	fmt.Println(wordCount("I am learning Go!"))
}
