package main

import (
	"fmt"
	"myapp/variables_dot_notation/doctor"
)

func main() {
	var whatToSay string

	whatToSay = doctor.Intro()

	fmt.Println(whatToSay)
}
