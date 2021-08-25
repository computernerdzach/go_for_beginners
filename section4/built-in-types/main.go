package main

import "fmt"

// aggregate types (array, struct)

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
	var myStrings [3]int

	myStrings[0] = 1
	myStrings[1] = 5
	myStrings[2] = 7

	fmt.Println("First element in array is", myStrings[0])

}
