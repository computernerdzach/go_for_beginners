package main

import (
	"bufio"
	"fmt"
	"myapp/eliza_project/doctor"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		userInput, _ := reader.ReadString('\n')

	}

}
