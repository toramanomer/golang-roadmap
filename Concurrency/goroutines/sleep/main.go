package main

import (
	"fmt"
	"time"
)

func routine() {
	fmt.Println("From Goroutine!")
}

func main() {
	go routine()            // Start a goroutine
	time.Sleep(time.Second) // Wait for the goroutine to finish
	fmt.Println("Main function done!")
}
