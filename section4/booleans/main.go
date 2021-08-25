package main

import "fmt"

func main() {
	apples := 18
	oranges := 9

	// boolean expression test
	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	// > < >= <=
	fmt.Printf("%d > %d: %t", apples, oranges, apples > oranges)
	fmt.Println()
	fmt.Printf("%d < %d: %t", apples, oranges, apples < oranges)
	fmt.Println()
	fmt.Printf("%d >= %d: %t", apples, oranges, apples >= oranges)
	fmt.Println()
	fmt.Printf("%d <= %d: %t", apples, oranges, apples <= oranges)
	fmt.Println()
}
