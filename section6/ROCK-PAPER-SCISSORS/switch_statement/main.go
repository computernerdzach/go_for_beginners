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

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	fmt.Println("Rock, Paper, & Scissors")
	fmt.Println("-----------------------")
	fmt.Println("Game is played for three rounds, and best of three wins the game. Good luck!")
	fmt.Println()
	fmt.Println("-------")

	for i := 0; i < 3; i++ {

		computerValue := rand.Intn(2)

		fmt.Println("Round", i+1)
		fmt.Println("-------")
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
				fmt.Println("--Invalid choice!--")
			}
			fmt.Println()
			if i < 2 {
				fmt.Println("-------")
			}

		}
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
