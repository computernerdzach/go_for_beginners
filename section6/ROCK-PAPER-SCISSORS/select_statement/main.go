package main

import (
	"myapp/select_statement/game"
)

func main() {
	displayChan := make(chan string)
	roundChan := make(chan int)

	game := game.Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}

	go game.Rounds()
	game.ClearScreen()
	game.PrintIntro()

	for {

		game.RoundChan <- 1
		<-game.RoundChan
		if game.Round.RoundNumber > 3 {
			break
		}

		if !game.PlayRound() {
			game.RoundChan <- -1
			<-game.RoundChan
		}
	}
	// if playerWins < computerWins {
	// 	fmt.Println("*********************")
	// 	fmt.Println("*Winner is Computer!*")
	// 	fmt.Println("*********************")
	// } else if playerWins > computerWins {
	// 	fmt.Println("*******************")
	// 	fmt.Println("*Winner is Player!*")
	// 	fmt.Println("*******************")
	// } else {
	// 	fmt.Println("*****************")
	// 	fmt.Println("*Game is a draw!*")
	// 	fmt.Println("*****************")
	// }
	// 	fmt.Println()
	// 	fmt.Println("Computer:", computerWins, "/ 3.")
	// 	fmt.Println("Player  :", playerWins, "/ 3.")
}