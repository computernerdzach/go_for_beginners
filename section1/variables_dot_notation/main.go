package main

import (
	"fmt"
	"myapp/section1/variables_dot_notation/doctor"
)

func main() {
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)
}
