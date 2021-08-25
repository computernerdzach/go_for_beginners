package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
	intMap := make(map[string]int)

	intMap["One"] = 1
	intMap["Two"] = 2
	intMap["Three"] = 3
	intMap["Four"] = 4
	intMap["Five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	// delete(intMap, "Four")

	el, ok := intMap["Four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")
	}

	intMap["Two"] = 4
}
