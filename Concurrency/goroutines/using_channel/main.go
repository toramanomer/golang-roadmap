package main

import "fmt"

func routine(done chan<- bool) {
	fmt.Println("From Goroutine!")
	done <- true
}

func main() {
	done := make(chan bool) // Create a channel of type boolean
	go routine(done)        // Start the goroutine
	<-done                  // Block until the goroutine is complete
	fmt.Println("Main function done!")
}
