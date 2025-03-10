package main

import "fmt"

func sumInts(m map[string]int64) int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}

func sumFloats(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}

type Number interface {
	int64 | float64
}

func sum[V Number](m map[string]V) V {
	var total V
	for _, value := range m {
		total += value
	}
	return total
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		sumInts(ints),
		sumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		sum(ints),
		sum(floats))
}
