package main

import (
	"fmt"
	"slices"
)

func yieldAll(index int, value string) bool {
	return true
}

func main() {

	iterAll := slices.All([]string{"a", "b", "c"})
	iterAll(yieldAll)

	names := []string{"Alice", "Bob", "Vera"}
	names = slices.AppendSeq(names, slices.Values([]string{"Dave", "Eve"}))

	for range slices.Backward([]string{"Ahmet", "Serra", "Ömer"}) {
	}

	langs := []string{"Go", "Java", "Python", "TypeScript"}
	for i := range slices.Chunk(langs, 3) {
		fmt.Println(i)
	}
	fmt.Println()

	slices.Concat(
		[]int{0},
		[]int{1, 2, 3},
		[]int{4, 5},
		[]int{},
	)

}
