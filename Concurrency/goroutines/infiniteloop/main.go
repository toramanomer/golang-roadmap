package main

import (
	"fmt"
	"time"
)

var done = false

func routine() {
	fmt.Println("From Goroutine!")
	done = true
}

func main() {
	go routine()

	for !done {
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("Main function done!")

}
