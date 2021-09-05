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
	playerWins := 0
	computerWins := 0
	roundOver := false

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	fmt.Println("Rock, Paper, & Scissors")
	fmt.Println("-----------------------")
	fmt.Println("Game is played for three rounds, and best of three wins the game. Good luck!")
	fmt.Println()
	fmt.Println("-------")

	for i := 0; i < 3; i++ {

		computerValue := rand.Intn(2)

		for !roundOver {

			fmt.Println("Round", i+1)
			fmt.Println("-------")
			fmt.Print("Please enter rock, paper, or scissors ->")

			playerChoice, _ = reader.ReadString('\n')
			playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
			playerChoice = strings.Replace(playerChoice, "\n", "", -1)

			fmt.Println()

			if playerChoice == "rock" {
				playerValue = ROCK
				fmt.Println("Player chose ROCK")
				roundOver = true
			} else if playerChoice == "paper" {
				playerValue = PAPER
				fmt.Println("Player chose PAPER")
				roundOver = true
			} else if playerChoice == "scissors" {
				playerValue = SCISSORS
				fmt.Println("Player chose SCISSORS")
				roundOver = true
			} else {
				fmt.Println("Invalid entry.")
				if i < 2 {
					fmt.Println("-------")
				}
			}
		}

		switch computerValue {
		case ROCK:
			fmt.Println("Computer chose ROCK")
		case PAPER:
			fmt.Println("Computer chose PAPER")
		case SCISSORS:
			fmt.Println("Computer chose SCISSORS")
		default:
		}

		fmt.Println()

		if playerValue == computerValue {
			fmt.Println("It's a draw")
		} else {
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					fmt.Println("--Computer wins!--")
					computerWins++
				} else {
					fmt.Println("--Player wins!--")
					playerWins++
				}
			case PAPER:
				if computerValue == SCISSORS {
					fmt.Println("--Computer wins!--")
					computerWins++
				} else {
					fmt.Println("--Player wins!--")
					playerWins++
				}
			case SCISSORS:
				if computerValue == ROCK {
					fmt.Println("--Computer wins!--")
					computerWins++
				} else {
					fmt.Println("--Player wins!--")
					playerWins++
				}
			default:
			}

		}
		fmt.Println()
		if i < 2 {
			fmt.Println("-------")
		}

		playerChoice = ""
		playerValue = -1
		roundOver = false

	}

	if playerWins < computerWins {
		fmt.Println("*********************")
		fmt.Println("*Winner is Computer!*")
		fmt.Println("*********************")
	} else if playerWins > computerWins {
		fmt.Println("*******************")
		fmt.Println("*Winner is Player!*")
		fmt.Println("*******************")
	} else {
		fmt.Println("*****************")
		fmt.Println("*Game is a draw!*")
		fmt.Println("*****************")
	}

	fmt.Println()
	fmt.Println("Computer:", computerWins, "/ 3.")
	fmt.Println("Player  :", playerWins, "/ 3.")
}

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
