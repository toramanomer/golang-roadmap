package main

import "fmt"

func routine() {
	fmt.Println("From Goroutine!")
}

func main() {
	go routine()                       // Start a goroutine
	fmt.Println("Main function done!") // Main function exits immediately without waiting for routine to end
}
