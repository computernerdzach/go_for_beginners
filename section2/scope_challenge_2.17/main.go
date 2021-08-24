package main

import "myapp/section2/packagetwo"

var myVar = "[package level variable in package main]"

func main() {
	blockVar := "[block level variable in package main, func main]"
	packagetwo.PrintMe(myVar, blockVar)
}
