package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	const numGs = 500

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(numGs)

	for range numGs {
		go func() {
			defer waitGroup.Done()
			count++
		}()
	}

	waitGroup.Wait()
	fmt.Println("Main function done with count equal to", count)
}
