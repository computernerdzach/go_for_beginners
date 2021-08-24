package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and don't type your number in, just press ENTER when ready."

func main() {
	var firstNumber = 2
	var secondNumber = 5
	var subtraction = 7
	var answer int

	reader := bufio.NewReader(os.Stdin)

	// display welcome / instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them throught the game
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	answer = firstNumber*secondNumber - subtraction
	fmt.Println("The answer is", answer)
}
