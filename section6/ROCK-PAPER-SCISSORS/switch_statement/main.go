package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(2)

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	for i := 0; i < 3; i++ {
		fmt.Print("Please enter rock, paper, or scissors ->")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)

		if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		}

		switch computerValue {
		case ROCK:
			fmt.Println("computer chose ROCK")
		case PAPER:
			fmt.Println("computer chose PAPER")
		case SCISSORS:
			fmt.Println("computer chose SCISSORS")
		default:
		}

		fmt.Println()

		if playerValue == computerValue {
			fmt.Println("It's a draw")
		} else {
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					fmt.Println("Computer wins!")
				} else {
					fmt.Println("Player wins!")
				}
			case PAPER:
				if computerValue == SCISSORS {
					fmt.Println("Computer wins!")
				} else {
					fmt.Println("Player wins!")
				}
			case SCISSORS:
				if computerValue == ROCK {
					fmt.Println("Computer wins!")
				} else {
					fmt.Println("Player wins!")
				}
			default:
				fmt.Println("Invalid choice!")
			}
		}
	}
}

// 	fmt.Print("Please enter rock, paper, or scissors ->")
// 	playerChoice, _ = reader.ReadString('\n')
// 	playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
// 	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

// 	if playerChoice == "rock" {
// 		playerValue = ROCK
// 	} else if playerChoice == "paper" {
// 		playerValue = PAPER
// 	} else if playerChoice == "scissors" {
// 		playerValue = SCISSORS
// 	}

// 	switch computerValue {
// 	case ROCK:
// 		fmt.Println("computer chose ROCK")
// 	case PAPER:
// 		fmt.Println("computer chose PAPER")
// 	case SCISSORS:
// 		fmt.Println("computer chose SCISSORS")
// 	default:
// 	}

// 	fmt.Println()

// 	if playerValue == computerValue {
// 		fmt.Println("It's a draw")
// 	} else {
// 		switch playerValue {
// 		case ROCK:
// 			if computerValue == PAPER {
// 				fmt.Println("Computer wins!")
// 			} else {
// 				fmt.Println("Player wins!")
// 			}
// 		case PAPER:
// 			if computerValue == SCISSORS {
// 				fmt.Println("Computer wins!")
// 			} else {
// 				fmt.Println("Player wins!")
// 			}
// 		case SCISSORS:
// 			if computerValue == ROCK {
// 				fmt.Println("Computer wins!")
// 			} else {
// 				fmt.Println("Player wins!")
// 			}
// 		default:
// 			fmt.Println("Invalid choice!")
// 		}
// 	}
// }

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
