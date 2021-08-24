package main

import (
	"fmt"
	"myapp/section2/packageone"
)

var one = "One"

func main() {
	var somethingElse = "this is a block level variable"
	fmt.Println(somethingElse)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	packageone.Exported()
}

func myFunc() {
	fmt.Println(one)
}
