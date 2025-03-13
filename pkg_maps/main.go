package main

import (
	"fmt"
	"maps"
	"strings"
)

func yieldAll[K comparable, V any](name K, value V) bool {
	return true
}

func yieldKeys[K comparable](key K) bool {
	return true
}

func yieldValues[V any](value V) bool {
	return true
}

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	iterAll := maps.All(m)
	iterAll(yieldAll)

	iterKeys := maps.Keys(m)
	iterKeys(yieldKeys)

	iterValues := maps.Values(m)
	iterValues(yieldValues)

	v := map[string]int{
		"d": 4,
		"e": 5,
		"f": 6,
	}

	maps.Insert(m, maps.All(v))
	fmt.Printf("%v\n", m)

	maps.Equal(m, v)

	maps.EqualFunc(
		map[string]string{
			"city": "ist",
		},
		map[string]string{
			"city": "Ist",
		},
		func(m1, m2 string) bool {
			return strings.EqualFold(m1, m2)
		},
	)
}
