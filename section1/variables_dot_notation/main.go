package main

import (
	"fmt"
	"myapp/variables_dot_notation/doctor"
)

func main() {
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)
}
