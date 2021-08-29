package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := map[int]string{
		1: "Cappucino",
		2: "Latte",
		3: "Americano",
		4: "Mocha",
		5: "Macchiato",
		6: "Espresso",
	}

	fmt.Println("MENU")
	fmt.Println("----")

	keySlice := make([]int, 0, len(coffees))
	for num := range coffees {
		keySlice = append(keySlice, num)
	}
	sort.Ints(keySlice)

	for _, item := range keySlice {
		fmt.Println(item, "-", coffees[item])
	}
	fmt.Println("Q - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Printf("You chose %s.\n", coffees[i])
	}

	fmt.Println("Program exiting...")
}
