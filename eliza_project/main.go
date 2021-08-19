package main

import (
	"bufio"
	"fmt"
	"myapp/variables_dot_notation/doctor"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

}
