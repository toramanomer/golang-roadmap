package main

func main() {
	array1 := [3]int{5, 7, 8}
	array2 := array1 // It is copied by value
	array2[0] = 10   // Will not change array1[0]

}
