package main

import "fmt"

var one = "One"

func main() {
	fmt.Println(one)
	myFunc()
}

func myFunc() {
	fmt.Println(one)
}
