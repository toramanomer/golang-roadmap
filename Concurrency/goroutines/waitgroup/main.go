package main

import (
	"fmt"
	"sync"
)

func routine(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println("From Goroutine!")
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3) // Add one goroutine to wait for

	go routine(&waitGroup)
	go routine(&waitGroup)
	go routine(&waitGroup)

	waitGroup.Wait() // Block until all goroutines call waitGroup.Done()
	fmt.Println("Main function done!")
}
